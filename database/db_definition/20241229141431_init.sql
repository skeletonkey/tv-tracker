-- +goose Up
CREATE TABLE user (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT NOT NULL UNIQUE DEFAULT (lower(hex(randomblob(16)))),
  username TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

CREATE TABLE show (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT NOT NULL UNIQUE DEFAULT (lower(hex(randomblob(16)))),
  name TEXT NOT NULL,
  type TEXT NOT NULL CHECK(type IN ('move', 'tv')),
  description TEXT,
  external_id INTEGER NOT NULL,
  thumbnail TEXT,
  FOREIGN KEY (external_id) REFERENCES external_X_internal (id)
);

CREATE TABLE episode (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT NOT NULL UNIQUE DEFAULT (lower(hex(randomblob(16)))),
  show_id INTEGER NOT NULL,
  title TEXT NOT NULL,
  season INTEGER NOT NULL,
  episode_number INTEGER NOT NULL,
  air_date DATE,
  external_id INTEGER NOT NULL,
  FOREIGN KEY (show_id) REFERENCES show (show_id),
  FOREIGN KEY (external_id) REFERENCES external_X_internal (id)
);

CREATE TABLE external_source (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

CREATE TABLE external_X_internal (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  show_id INTEGER,
  episode_id INTEGER,
  external_source_id INTEGER NOT NULL,
  external_id TEXT NOT NULL,
  CHECK (
    (show_id IS NOT NULL AND episode_id IS NULL) 
    OR 
    (show_id IS NULL AND episode_id IS NOT NULL)
  )
);