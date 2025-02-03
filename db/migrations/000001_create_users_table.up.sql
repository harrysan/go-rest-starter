CREATE TABLE users (
    id serial4 NOT NULL,
    username varchar(64) NULL,
    "name" varchar(64) NULL,
    "password" varchar(64) NULL,
    phone varchar(32) NULL,
    email varchar(128) NULL,
    remark varchar(1024) NULL,
    status varchar(20) NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_users_created_at ON public.users USING btree (created_at);

CREATE INDEX idx_users_name ON public.users USING btree (name);

CREATE INDEX idx_users_status ON public.users USING btree (status);

CREATE INDEX idx_users_updated_at ON public.users USING btree (updated_at);

CREATE INDEX idx_users_username ON public.users USING btree (username);