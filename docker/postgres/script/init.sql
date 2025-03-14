create database ars_dev;
create user ars_dev with login password 'ars_dev';
alter database ars_dev owner to ars_dev;;
grant all privileges on database ars_dev to ars_dev;
grant all on schema public to ars_dev;

\c ars_dev
create extension citext;