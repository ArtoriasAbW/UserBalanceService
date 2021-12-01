
CREATE TABLE users
(
    id serial not null unique,
    balance integer default 0
);