-- name: CreateCustomer :one
insert into "customer" (pid, email)
values (@pid, @email)
returning *;

-- name: CreateAuthnEmail :one
insert into "authn_email" (email, password_hash)
values (@email, @passwordHash)
returning *;

-- name: CreateAuthn :one
insert into "authn" (customer_id, "type", ref_id)
values (@customerId, @type, @refID)
returning *;