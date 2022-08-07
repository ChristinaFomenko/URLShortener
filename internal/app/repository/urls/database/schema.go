package database

const query = `create table if not exists urls 
(
    id varchar(10) not null,
    url varchar(500) not null unique,
    user_id varchar(10) not null,
    created_at timestamp with time zone default now() not null,
	deleted_at  int DEFAULT (0)
);`
