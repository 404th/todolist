CREATE TABLE users (
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(250) NOT NULL,
    username VARCHAR(250) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL 
);

CREATE TABLE todo_items (
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE todo_lists (
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE users_lists (
    id SERIAL,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE lists_items (
    id SERIAL,
    item_id INT REFERENCES todo_items(id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL
);
