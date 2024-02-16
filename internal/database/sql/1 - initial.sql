BEGIN TRANSACTION;

CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL
);

CREATE TABLE sessions (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE groups (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    owner_id INTEGER NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

CREATE TABLE retro ( 
    id INTEGER PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    owner_id INTEGER NOT NULL,
    hosting_group_id INTEGER NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users(id),
    FOREIGN KEY (hosting_group_id) REFERENCES groups(id)
);

create TABLE retro_column (
    id INTEGER PRIMARY KEY,
    retro_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    FOREIGN KEY (retro_id) REFERENCES retro(id)
);

CREATE TABLE retro_idea (
    id INTEGER PRIMARY KEY,
    retro_column_id INTEGER NOT NULL,
    description VARCHAR(255) NOT NULL,
    votes INTEGER NOT NULL,
    discussion_start_time TIMESTAMP NOT NULL,
    discussion_end_time TIMESTAMP NOT NULL,
    FOREIGN KEY (retro_column_id) REFERENCES retro_column(id)
);

create TABLE todo (
    id INTEGER PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    retro_id INTEGER NOT NULL,
    FOREIGN KEY (retro_id) REFERENCES retro(id)
); 

COMMIT;