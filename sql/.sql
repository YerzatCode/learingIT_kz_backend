CREATE TABLE User (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    age INTEGER DEFAULT 0,
    email TEXT UNIQUE NOT NULL,
    password TEXT UNIQUE NOT NULL ,
    role TEXT NOT NULL DEFAULT "user",
    cretated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);