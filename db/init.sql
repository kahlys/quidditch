CREATE TABLE users (
  id serial PRIMARY KEY,
  username VARCHAR (255) NOT NULL,
  email VARCHAR (255) NOT NULL,
  password VARCHAR (255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_login TIMESTAMP
);

CREATE TABLE teams (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  owner_id INTEGER REFERENCES users(id)
);

CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  nationality TEXT NOT NULL,
  power INTEGER NOT NULL,
  stamina INTEGER NOT NULL,
  position TEXT NOT NULL,
  team_id INTEGER REFERENCES teams(id)
);

CREATE TABLE seasons (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	first_place TEXT,
	second_place TEXT,
	third_place TEXT
);

CREATE TABLE matches (
  id SERIAL PRIMARY KEY,
  home_team_id INTEGER REFERENCES teams(id),
  away_team_id INTEGER REFERENCES teams(id),
  home_team_score INTEGER NOT NULL,
  away_team_score INTEGER NOT NULL,
  season_id INTEGER REFERENCES seasons(id)
);

CREATE TABLE season_standings (
  id SERIAL PRIMARY KEY,
  team_id INTEGER REFERENCES teams(id),
  season_id INTEGER REFERENCES seasons(id),
  points INTEGER NOT NULL
);

CREATE TABLE settings (
  key TEXT NOT NULL PRIMARY KEY,
  value TEXT NOT NULL
);

INSERT INTO settings (key, value) VALUES ('season_counter', '0');
INSERT INTO settings (key, value) VALUES ('schema_version', '1');