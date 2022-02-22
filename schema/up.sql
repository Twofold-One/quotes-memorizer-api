create table users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null
);

create table quotes
(
    id serial not null unique,
    author varchar(255) not null,
    quote varchar(255) not null,
    source varchar(255) not null,
    sourcetype varchar(255) not null,
    created date not null
);

create table users_quotes
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    quote_id int references quotes(id) on delete cascade not null
);


