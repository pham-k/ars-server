-- name: RegisterUserWithEmail :one
insert into "user" (pid, authn_type, email, password_hash)
values (@pid, @authn_type, @email, @password_hash)
returning *;

-- name: ValidateUser :one
update "user"
set validated = true
where email = @email
returning validated;

-- name: GetUserFromEmail :one
select id, pid, authn_type, email, password_hash, validated
from "user"
where email = @email and authn_type = @authn_type;