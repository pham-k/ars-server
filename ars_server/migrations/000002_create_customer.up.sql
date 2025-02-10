begin;

create table if not exists "customer" (
    id serial primary key,
    pid text NOT NULL UNIQUE CHECK (pid LIKE 'cus_%'),
    email citext unique,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

create table authn_type (
    id serial primary key,
    type text unique,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

create table if not exists "authn" (
    id serial primary key,
    customer_id serial references customer(id) on delete restrict on update cascade,
    type text references authn_type (type) on delete restrict on update cascade,
    ref_id serial,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

create table if not exists "authn_email" (
    id serial primary key,
    email citext unique,
    password_hash text not null,
    activated bool not null default false,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

INSERT INTO authn_type(type)
VALUES ('email');

commit;