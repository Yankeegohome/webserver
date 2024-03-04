CREATE TABLE IF NOT EXISTS users
(
    id        INTEGER PRIMARY KEY autoincrement,
    email     TEXT NOT NULL UNIQUE,
    password  TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS lab_sample
(
    id      INTEGER PRIMARY KEY autoincrement,
    ids     TEXT NOT NULL,
    text    TEXT NOT NULL,
    specimen_id INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS ID ON lab_sample (ID);

CREATE TABLE IF NOT EXISTS lab_research
(
    id                INTEGER PRIMARY KEY autoincrement,
    lab_sample_id     INTEGER NOT NULL,
    code    TEXT NOT NULL,
    text    TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS id ON lab_research (id);

CREATE TABLE IF NOT EXISTS specimen
(
    id       INTEGER PRIMARY KEY autoincrement,
    code     TEXT NOT NULL,
    text     TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS id ON specimen (id);
