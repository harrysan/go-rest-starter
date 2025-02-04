CREATE TABLE users (
    id serial4 NOT NULL,
    username varchar(64) NOT NULL UNIQUE,
    "name" varchar(64) NOT NULL,
    "password" varchar(64) NOT NULL,
    phone varchar(32) NULL,
    email varchar(128) NOT NULL UNIQUE,
    remark varchar(1024) NULL,
    status varchar(20) NOT NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_users_created_at ON users USING btree (created_at);

CREATE INDEX idx_users_updated_at ON users USING btree (updated_at);

CREATE INDEX idx_users_username ON users USING btree (username);