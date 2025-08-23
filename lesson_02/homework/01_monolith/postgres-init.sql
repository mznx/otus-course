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

CREATE TABLE "dialogs" (
    id text primary key,
    created_at timestamp NOT NULL default current_timestamp,
    updated_at timestamp NOT NULL default current_timestamp
);

CREATE TABLE "dialog_members" (
    user_id text NOT NULL,
    dialog_id text NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    updated_at timestamp NOT NULL default current_timestamp,

    primary key (user_id, dialog_id)
);

CREATE TABLE "dialog_messages" (
    id text primary key,
    dialog_id text NOT NULL,
    sender_id text NOT NULL,
    message text NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    updated_at timestamp NOT NULL default current_timestamp
);

CREATE TABLE "dialog_private" (
    user_pair_hash text primary key,
    dialog_id text NOT NULL,
    user_id_1 text NOT NULL,
    user_id_2 text NOT NULL,
    created_at timestamp NOT NULL default current_timestamp,
    updated_at timestamp NOT NULL default current_timestamp
);
