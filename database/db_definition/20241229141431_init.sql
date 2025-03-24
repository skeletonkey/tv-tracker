-- +goose Up
CREATE TABLE user (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT NOT NULL UNIQUE DEFAULT (lower(hex(randomblob(16)))),
  username TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  active BOOLEAN NOT NULL DEFAULT 1,
  password TEXT NOT NULL
);

CREATE TABLE show (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT NOT NULL UNIQUE DEFAULT (lower(hex(randomblob(16)))),
  name TEXT NOT NULL,
  type TEXT NOT NULL CHECK(type IN ('move', 'tv')),
  description TEXT
);

CREATE TABLE episode (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT NOT NULL UNIQUE DEFAULT (lower(hex(randomblob(16)))),
  show_id INTEGER NOT NULL,
  title TEXT NOT NULL,
  season INTEGER NOT NULL,
  episode_number INTEGER NOT NULL,
  air_date DATE,
  FOREIGN KEY (show_id) REFERENCES show (show_id)
);

CREATE TABLE external_source (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  default_url TEXT NOT NULL
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

CREATE TABLE user_X_show (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  show_id INTEGER NOT NULL,
  rank INT NOT NULL DEFAULT 1,
  visible BOOLEAN NOT NULL DEFAULT 1,
  FOREIGN KEY (user_id) REFERENCES user (id),
  FOREIGN KEY (show_id) REFERENCES show (id)
);

CREATE TABLE user_history (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  show_id INTEGER NOT NULL,
  episode_id INTEGER NOT NULL,
  status_id INTEGER NOT NULL,
  timestamp DATETIME NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user (id),
  FOREIGN KEY (show_id) REFERENCES show (id),
  FOREIGN KEY (episode_id) REFERENCES episode (id)
  FOREIGN KEY (status_id) REFERENCES status (id)
  CHECK (
    (show_id IS NOT NULL AND episode_id IS NULL) 
    OR 
    (show_id IS NULL AND episode_id IS NOT NULL)
  )
);

CREATE TABLE status (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  type TEXT NOT NULL
);

INSERT INTO status (name, type) values ("watched", "view"), ("started", "view"), ("new", "view"), ("unreleased", "scan"), ("released", "scan"), ("ended", "scan");

CREATE TABLE art (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  url TEXT NOT NULL,
  art_type_id INTEGER NOT NULL,
  height INT NOT NULL DEFAULT 0,
  width INT NOT NULL DEFAULT 0,
  FOREIGN KEY (art_type_id) REFERENCES art_type (id)
);

CREATE TABLE art_type (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

INSERT INTO art_type (name) values ("thumbnail"), ("image"); 

CREATE TABLE entity_art (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  show_id INTEGER NOT NULL,
  episode_id INTEGER NOT NULL,
  art_id INTEGER NOT NULL,
  default_art BOOLEAN NOT NULL DEFAULT 1,
  FOREIGN KEY (show_id) REFERENCES show (id),
  FOREIGN KEY (episode_id) REFERENCES episode (id),
  FOREIGN KEY (art_id) REFERENCES art (id)
);

CREATE TABLE user_X_entity_art (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  entity_art_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user (id),
  FOREIGN KEY (entity_art_id) REFERENCES entity_art (id)
);

CREATE TABLE source (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  url TEXT NOT NULL
);

CREATE TABLE user_X_source (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  source_id INTEGER NOT NULL,
  url TEXT NOT NULL,
  api_token TEXT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user (id),
  FOREIGN KEY (source_id) REFERENCES source (id)
);