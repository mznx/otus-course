CREATE TABLE "users" (
    id text primary key,
    first_name text NOT NULL,
    second_name text NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    updated_at timestamp NOT NULL default current_timestamp
);

CREATE TABLE "user_auth" (
    user_id text NOT NULL,
    pass_hash text NOT NULL,
    token text NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    updated_at timestamp NOT NULL default current_timestamp
);

CREATE TABLE "friends" (
    user_id text NOT NULL,
    friend_id text NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    updated_at timestamp NOT NULL default current_timestamp,

    primary key (user_id, friend_id)
);

CREATE TABLE "posts" (
    id text primary key,
    author_id text NOT NULL,
    text text NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    updated_at timestamp NOT NULL default current_timestamp
);
