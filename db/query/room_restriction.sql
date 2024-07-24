-- name: CreateRoomRestriction :one
INSERT INTO room_restrictions (
	start_date,
	end_date,
	room_id,
	reservation_id,
	restriction_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetRoomRestriction :one
SELECT * FROM room_restrictions
WHERE id = $1 LIMIT 1;

-- name: GetRoomRestrictionForUpdate :one
SELECT * FROM room_restrictions
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListRoomRestrictions :many
SELECT * FROM room_restrictions
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetRestrictionsForRoomByDate :many
select *
from room_restrictions where $1 < end_date and $2 >= start_date
and room_id = $3
ORDER BY id
LIMIT $4
OFFSET $5;

-- name: SearchAvailabilityByDatesByRoomID :one
select count(id)
from room_restrictions
where room_id = $1
and $2 < end_date and $3 >= start_date;

-- name: UpdateRoomRestriction :one
UPDATE room_restrictions
SET
  updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
  start_date = COALESCE(sqlc.narg(start_date), start_date),
  end_date = COALESCE(sqlc.narg(end_date), end_date),
  room_id = COALESCE(sqlc.narg(room_id), room_id),
  reservation_id = COALESCE(sqlc.narg(reservation_id), reservation_id),
  restriction_id = COALESCE(sqlc.narg(restriction_id), restriction_id)
WHERE
id = sqlc.arg(id)
RETURNING *;

-- name: DeleteRoomRestriction :exec
DELETE FROM room_restrictions
WHERE id = $1;