begin;

create table if not exists "authn_type" (
    id bigserial primary key,
    type text unique,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

insert into authn_type(type)
values ('email'), ('google');

create table if not exists "user" (
    id bigserial primary key,
    pid text NOT NULL UNIQUE CHECK (pid LIKE 'usr_%'),
    authn_type text references authn_type (type) on delete restrict on update cascade,
    email citext unique,
    password_hash text not null,
    validated bool not null default false,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

commit;