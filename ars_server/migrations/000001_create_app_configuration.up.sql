begin;

create table if not exists "config_type" (
    id serial primary key,
    type TEXT unique,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

create table if not exists "config_scope" (
    id serial primary key,
    scope TEXT unique,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

create table if not exists "config" (
    id serial primary key,
    pid text NOT NULL UNIQUE CHECK (pid LIKE 'conf_%'),
    name TEXT unique,
    type text REFERENCES config_type (type) on delete restrict on update cascade,
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

create table if not exists "app_configuration" (
    id bigserial primary key,
    name text references config (name) on delete restrict on update cascade,
    scope text references config_scope (scope) on delete restrict on update cascade,
    text_value text not null default '',
    version integer not null DEFAULT 1,
    updated_at timestamp(0) with time zone not null default now(),
    created_at timestamp(0) with time zone not null default now()
);

INSERT INTO config_scope(scope)
VALUES ('global');

INSERT INTO config_type(type)
VALUES ('bool');

INSERT INTO config(pid, name, type)
VALUES ('conf_mf1giXJ19PJ0Sg7lmvI59', 'feature_cash_flow_calculator', 'bool'),
       ('conf_NoOu2D1JGGUumIRafeQ6d', 'feature_date_reminder', 'bool');

INSERT INTO app_configuration(name, scope, text_value)
VALUES ('feature_cash_flow_calculator', 'global', 'true'),
       ('feature_date_reminder', 'global', 'true');

commit;