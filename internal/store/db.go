package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/kahlys/quidditch/internal/game"
)

type Database struct {
	*pgxpool.Pool
}

func NewDatabase(connStr string) (Database, error) {
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return Database{}, err
	}
	return Database{
		Pool: pool,
	}, nil
}

func (db *Database) Team(teamid int) (game.Team, error) {
	team := game.Team{ID: teamid}

	tx, err := db.Begin(context.Background())
	if err != nil {
		return game.Team{}, err
	}
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(
		context.Background(),
		`SELECT name FROM teams WHERE id = $1`,
		teamid,
	).Scan(&team.Name)
	if err != nil {
		return game.Team{}, err
	}

	row, err := tx.Query(
		context.Background(),
		`SELECT id, first_name, last_name, nationality, power, stamina, position FROM players WHERE team_id=$1`,
		teamid,
	)
	if err != nil {
		return game.Team{}, err
	}
	defer row.Close()

	counter := 0
	counterChaser := 0
	counterBeater := 0
	for row.Next() {
		counter++
		var p game.Player
		err = row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Country, &p.Power, &p.Stamina, &p.Role)
		if err != nil {
			return game.Team{}, err
		}
		switch p.Role {
		case game.RoleSeeker:
			team.Squad.Seeker = p
		case game.RoleKeeper:
			team.Squad.Keeper = p
		case game.RoleBeater:
			counterBeater++
			switch counterBeater {
			case 1:
				team.Squad.Beater1 = p
			case 2:
				team.Squad.Beater2 = p
			default:
				return game.Team{}, fmt.Errorf("too many players for role %v", p.Role)
			}
		case game.RoleChaser:
			counterChaser++
			switch counterChaser {
			case 1:
				team.Squad.Chaser1 = p
			case 2:
				team.Squad.Chaser2 = p
			case 3:
				team.Squad.Chaser3 = p
			default:
				return game.Team{}, fmt.Errorf("too many players for role %v", p.Role)
			}
		default:
			return game.Team{}, fmt.Errorf("unknown role %v", p.Role)
		}
	}
	row.Close()

	if counter != 7 {
		return game.Team{}, fmt.Errorf("team must have 7 players but found %v", counter)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return game.Team{}, err
	}

	return team, nil
}

func (db *Database) Teams(ctx context.Context) ([]game.Team, error) {
	teams := []game.Team{}

	tx, err := db.Begin(ctx)
	if err != nil {
		return []game.Team{}, err
	}
	defer tx.Rollback(ctx)

	ids := []int{}
	row, err := tx.Query(ctx, `SELECT id FROM teams`)
	if err != nil {
		return []game.Team{}, err
	}
	for row.Next() {
		i := 0
		row.Scan(&i)
		ids = append(ids, i)
	}

	for _, id := range ids {
		team, err := db.Team(id)
		if err != nil {
			return []game.Team{}, err
		}
		teams = append(teams, team)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return []game.Team{}, err
	}

	return teams, nil
}

// TeamsCount return the number of teams
func (db *Database) TeamsCount(ctx context.Context) (int, error) {
	var count int
	err := db.QueryRow(ctx, `SELECT COUNT(*) FROM teams`).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *Database) CurrentSeason(ctx context.Context) (int, error) {
	var season int
	err := db.QueryRow(ctx, `SELECT id FROM seasons WHERE start_date = (SELECT MAX(start_date) FROM seasons)`).Scan(&season)
	if err != nil {
		return 0, err
	}
	return season, nil
}

func (db *Database) Matches(ctx context.Context) ([]game.Game, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		return []game.Game{}, err
	}
	defer tx.Rollback(ctx)

	season, err := db.CurrentSeason(ctx)
	if err != nil {
		return []game.Game{}, err
	}

	rows, err := tx.Query(
		ctx,
		`WITH ranked_teams AS (
			SELECT
				team_id,
				points,
				ROW_NUMBER() OVER (ORDER BY points ASC) as rank
			FROM season_standings
			WHERE season_id = $1
		)
		SELECT
			rt1.team_id as team1,
			rt2.team_id as team2
		FROM ranked_teams rt1
		JOIN ranked_teams rt2 ON rt1.rank = rt2.rank - 1
		WHERE rt1.rank % 2 = 1
		ORDER BY rt1.rank`,
		season,
	)
	if err != nil {
		return []game.Game{}, err
	}
	defer rows.Close()

	games := []game.Game{}
	for rows.Next() {
		var homeID, awayID int
		err = rows.Scan(&homeID, &awayID)
		if err != nil {
			return []game.Game{}, err
		}
		home, err := db.Team(homeID)
		if err != nil {
			return []game.Game{}, err
		}
		away, err := db.Team(awayID)
		if err != nil {
			return []game.Game{}, err
		}
		games = append(games, game.NewGame(season, home, away))
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return []game.Game{}, err
	}

	return games, nil
}

func (db *Database) SaveMatchResult(ctx context.Context, g game.Game) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(
		ctx,
		`INSERT INTO matches (season_id, home_team_id, away_team_id, home_team_score, away_team_score) VALUES ($1, $2, $3, $4, $5)`,
		g.SeasonID, g.Home.ID, g.Away.ID, g.Result.ScoreHome, g.Result.ScoreAway,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx,
		`UPDATE season_standings SET points = points + $1 WHERE season_id = $2 AND team_id = $3`,
		g.Result.ScoreHome, g.SeasonID, g.Home.ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx,
		`UPDATE season_standings SET points = points + $1 WHERE season_id = $2 AND team_id = $3`,
		g.Result.ScoreAway, g.SeasonID, g.Away.ID,
	)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// RecruitablePlayers return the list of recruitable players
func (db *Database) RecruitablePlayers(context.Context) ([]game.Player, error) {
	row, err := db.Query(
		context.Background(),
		`SELECT id, first_name, last_name, nationality, power, stamina, position FROM players WHERE team_id IS NULL`,
	)
	if err != nil {
		return []game.Player{}, err
	}
	defer row.Close()

	players := []game.Player{}
	for row.Next() {
		p := game.Player{}
		err = row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Country, &p.Power, &p.Stamina, &p.Role)
		if err != nil {
			return []game.Player{}, err
		}
		players = append(players, p)
	}

	return players, nil
}

// NewRecruitablePlayers replace recruitable players
func (db *Database) NewRecruitablePlayers(ctx context.Context, players []game.Player) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(
		context.Background(),
		`DELETE FROM players WHERE team_id IS NULL`,
	)
	if err != nil {
		return err
	}

	for _, p := range players {
		_, err = tx.Exec(
			context.Background(),
			`INSERT INTO players (first_name, last_name, nationality, power, stamina, position) VALUES ($1, $2, $3, $4, $5, $6)`,
			p.FirstName, p.LastName, p.Country, p.Power, p.Stamina, p.Role,
		)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
