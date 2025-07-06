create table if not exists "user" (
    id bigserial primary key,
    pid text NOT NULL UNIQUE CHECK (pid LIKE 'usr_%'),
    phone varchar(25) not null unique,
    country_code varchar(2) not null,
    password_hash text not null,
    validated bool not null default false,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);