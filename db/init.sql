-- users table is used for authentication (members)
CREATE TABLE users (
  id serial PRIMARY KEY,
  username VARCHAR (255) NOT NULL,
  email VARCHAR (255) NOT NULL UNIQUE,
  password VARCHAR (255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_login TIMESTAMP
);

-- teams table is used to store teams
CREATE TABLE teams (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  owner_id INTEGER REFERENCES users(id)
);

-- players table is used to store quiddditch players (npc)
CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  nationality TEXT NOT NULL,
  power INTEGER NOT NULL,
  stamina INTEGER NOT NULL,
  position TEXT NOT NULL,
  starting BOOLEAN NOT NULL DEFAULT false,
  team_id INTEGER REFERENCES teams(id)
);

-- seasons table is used to store seasons
CREATE TABLE seasons (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
    start_date TIMESTAMP NOT NULL
);

-- matches table is used to store matches results
CREATE TABLE matches (
  id SERIAL PRIMARY KEY,
  home_team_id INTEGER REFERENCES teams(id),
  away_team_id INTEGER REFERENCES teams(id),
  home_team_score INTEGER NOT NULL,
  away_team_score INTEGER NOT NULL,
  season_id INTEGER REFERENCES seasons(id)
);

-- season_standings table is used to store teams points for each season
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

INSERT INTO settings (key, value) VALUES ('schema_version', '1');

-- init first season
INSERT INTO seasons (name, start_date) VALUES (TO_CHAR(NOW(), 'MonYYYY'), NOW());

-- init bot teams
INSERT INTO teams ("id", "name", "owner_id") VALUES
(1,	'Appleby Arrows',	NULL),
(2,	'Ballycastle Bats',	NULL),
(3,	'Caerphilly Catapults',	NULL),
(4,	'Chudley Cannons',	NULL),
(5,	'Falmouth Falcons',	NULL),
(6,	'Holyhead Harpies',	NULL),
(7,	'Kenmare Kestrels',	NULL),
(8,	'Montrose Magpies',	NULL),
(9,	'Puddlemere United',	NULL),
(10,	'Tutshill Tornados',	NULL),
(11,	'Wigtown Wanderers',	NULL);

INSERT INTO players ("id", "first_name", "last_name", "nationality", "power", "stamina", "position", "starting", "team_id") VALUES
(1,	'Margarita',	'Petkova',	'bulgaria',	35,	100,	'seeker',	't',	1),
(2,	'Oscar',	'Jackson',	'england',	29,	100,	'keeper',	't',	1),
(3,	'Liam',	'O''Reilly',	'ireland',	20,	100,	'beater',	't',	1),
(4,	'Jorge',	'López',	'spain',	21,	100,	'beater',	't',	1),
(5,	'Dimitar',	'Petrov',	'bulgaria',	29,	100,	'chaser',	't',	1),
(6,	'Ella',	'Anderson',	'england',	33,	100,	'chaser',	't',	1),
(7,	'Katerina',	'Angelova',	'bulgaria',	32,	100,	'chaser',	't',	1),
(8,	'Julian',	'Moreau',	'france',	20,	100,	'seeker',	't',	2),
(9,	'Maria',	'Jiménez',	'spain',	39,	100,	'keeper',	't',	2),
(10,	'David',	'Martínez',	'spain',	32,	100,	'beater',	't',	2),
(11,	'Harper',	'Harris',	'england',	29,	100,	'beater',	't',	2),
(12,	'Owen',	'O''Neill',	'ireland',	27,	100,	'chaser',	't',	2),
(13,	'Sophia',	'Lang',	'germany',	35,	100,	'chaser',	't',	2),
(14,	'Ava',	'Byrne',	'ireland',	28,	100,	'chaser',	't',	2),
(15,	'Noah',	'Martin',	'us',	35,	100,	'seeker',	't',	3),
(16,	'Harry',	'Jones',	'england',	24,	100,	'keeper',	't',	3),
(17,	'Harry',	'Taylor',	'england',	25,	100,	'beater',	't',	3),
(18,	'Sadie',	'Kennedy',	'ireland',	23,	100,	'beater',	't',	3),
(19,	'Noah',	'Jones',	'us',	37,	100,	'chaser',	't',	3),
(20,	'Harry',	'Anderson',	'england',	27,	100,	'chaser',	't',	3),
(21,	'Emma',	'Wagner',	'germany',	27,	100,	'chaser',	't',	3),
(22,	'Amelia',	'White',	'england',	20,	100,	'seeker',	't',	4),
(23,	'Léa',	'Boulanger',	'france',	33,	100,	'keeper',	't',	4),
(24,	'Sophia',	'Jones',	'us',	34,	100,	'beater',	't',	4),
(25,	'Liam',	'Sheehan',	'ireland',	23,	100,	'beater',	't',	4),
(26,	'Harper',	'Brown',	'england',	37,	100,	'chaser',	't',	4),
(27,	'Anne',	'Marceau',	'france',	37,	100,	'chaser',	't',	4),
(28,	'Vasilka',	'Dimitrov',	'bulgaria',	40,	100,	'chaser',	't',	4),
(29,	'Jack',	'Kelly',	'ireland',	24,	100,	'seeker',	't',	5),
(30,	'Benjamin',	'Boulanger',	'france',	36,	100,	'keeper',	't',	5),
(31,	'James',	'Johnson',	'england',	22,	100,	'beater',	't',	5),
(32,	'Sophie',	'Boyer',	'france',	32,	100,	'beater',	't',	5),
(33,	'Sadie',	'Walsh',	'ireland',	26,	100,	'chaser',	't',	5),
(34,	'Léa',	'Blanc',	'france',	38,	100,	'chaser',	't',	5),
(35,	'Lucia',	'López',	'spain',	20,	100,	'chaser',	't',	5),
(36,	'Finn',	'O''Brien',	'ireland',	34,	100,	'seeker',	't',	6),
(37,	'Connor',	'Byrne',	'ireland',	27,	100,	'keeper',	't',	6),
(38,	'Sebastian',	'Günther',	'germany',	39,	100,	'beater',	't',	6),
(39,	'Isabella',	'Jones',	'us',	33,	100,	'beater',	't',	6),
(40,	'Ashling',	'Byrne',	'ireland',	39,	100,	'chaser',	't',	6),
(41,	'Elena',	'Todorov',	'bulgaria',	40,	100,	'chaser',	't',	6),
(42,	'Jorge',	'Gómez',	'spain',	23,	100,	'chaser',	't',	6),
(43,	'Carmen',	'González',	'spain',	23,	100,	'seeker',	't',	7),
(44,	'Cristina',	'Vázquez',	'spain',	26,	100,	'keeper',	't',	7),
(45,	'Noah',	'Johnson',	'us',	24,	100,	'beater',	't',	7),
(46,	'Léa',	'Durand',	'france',	36,	100,	'beater',	't',	7),
(47,	'Martina',	'Todorova',	'bulgaria',	38,	100,	'chaser',	't',	7),
(48,	'Harper',	'Johnson',	'england',	24,	100,	'chaser',	't',	7),
(49,	'Johannes',	'Kaiser',	'germany',	20,	100,	'chaser',	't',	7),
(50,	'Ella',	'Miller',	'england',	33,	100,	'seeker',	't',	8),
(51,	'Ana',	'Sánchez',	'spain',	38,	100,	'keeper',	't',	8),
(52,	'Daniel',	'Fernández',	'spain',	35,	100,	'beater',	't',	8),
(53,	'Anne',	'Boucher',	'france',	26,	100,	'beater',	't',	8),
(54,	'James',	'O''Reilly',	'ireland',	32,	100,	'chaser',	't',	8),
(55,	'Rosa',	'López',	'spain',	26,	100,	'chaser',	't',	8),
(56,	'Sarah',	'Walsh',	'ireland',	28,	100,	'chaser',	't',	8),
(57,	'Chloé',	'Dubois',	'france',	20,	100,	'seeker',	't',	9),
(58,	'Todor',	'Atanasov',	'bulgaria',	30,	100,	'keeper',	't',	9),
(59,	'Charlotte',	'Taylor',	'england',	33,	100,	'beater',	't',	9),
(60,	'Sebastian',	'Stein',	'germany',	39,	100,	'beater',	't',	9),
(61,	'Niklas',	'Pfeiffer',	'germany',	20,	100,	'chaser',	't',	9),
(62,	'Jacques',	'Simon',	'france',	34,	100,	'chaser',	't',	9),
(63,	'Benjamin',	'Durand',	'france',	23,	100,	'chaser',	't',	9),
(64,	'Benjamin',	'Da Silva',	'france',	39,	100,	'seeker',	't',	10),
(65,	'Niamh',	'Walsh',	'ireland',	22,	100,	'keeper',	't',	10),
(66,	'Ella',	'Miller',	'england',	26,	100,	'beater',	't',	10),
(67,	'Isabella',	'Robinson',	'england',	31,	100,	'beater',	't',	10),
(68,	'Benjamin',	'Robert',	'france',	38,	100,	'chaser',	't',	10),
(69,	'Vasil',	'Kovachev',	'bulgaria',	34,	100,	'chaser',	't',	10),
(70,	'Paul',	'Menard',	'france',	23,	100,	'chaser',	't',	10),
(71,	'Tobias',	'Pfeiffer',	'germany',	21,	100,	'seeker',	't',	11),
(72,	'Marie',	'Richter',	'germany',	31,	100,	'keeper',	't',	11),
(73,	'Rositsa',	'Stoyanova',	'bulgaria',	29,	100,	'beater',	't',	11),
(74,	'Simeon',	'Angelova',	'bulgaria',	20,	100,	'beater',	't',	11),
(75,	'Estelle',	'Simon',	'france',	31,	100,	'chaser',	't',	11),
(76,	'Martina',	'Nikolova',	'bulgaria',	30,	100,	'chaser',	't',	11),
(77,	'Luis',	'González',	'spain',	35,	100,	'chaser',	't',	11);

INSERT INTO season_standings ("id", "team_id", "season_id", "points") VALUES
( 1,  1, 1,	0),
( 2,  2, 1,	0),
( 3,  3, 1,	0),
( 4,  4, 1,	0),
( 5,  5, 1,	0),
( 6,  6, 1,	0),
( 7,  7, 1,	0),
( 8,  8, 1,	0),
( 9,  9, 1,	0),
(10, 10, 1,	0),
(11, 11, 1,	0);