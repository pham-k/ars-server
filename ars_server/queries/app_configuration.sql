-- name: ListAppConfigsByScope :many
select c.pid as "pid", appconf.scope as "scope", appconf.name as "name", c.type as "type", appconf.text_value as "text_value"
from "app_configuration" appconf
join "config" c on appconf.name = c.name
where appconf.scope = @scope;