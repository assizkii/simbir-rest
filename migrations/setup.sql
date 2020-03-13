create table if not exists users (
    id serial primary key,
    login text,
    password text
);