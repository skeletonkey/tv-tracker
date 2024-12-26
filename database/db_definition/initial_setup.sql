CREATE TABLE user (
  user_id INTEGER PRIMARY KEY AUTOINCREMENT,
  username TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

CREATE TABLE show (
  show_id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  type TEXT NOT NULL CHECK(type IN ('move', 'tv'))
  description TEXT,
  tvdb_id INTEGER NOT NULL,
  thumbnail TEXT
);

CREATE TABLE episode (
  episode_id INTEGER PRIMARY KEY AUTOINCREMENT,
  show_id INTEGER NOT NULL,
  title TEXT NOT NULL,
  season INTEGER,
  episode_number INTEGER,
  air_date DATE,
  FOREIGN KEY (show_id) REFERENCES show (show_id)
);