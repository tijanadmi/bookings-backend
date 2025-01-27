-- name: CreateReservation :one
INSERT INTO reservations (
	room_id,
	first_name,
	last_name,
	email,
	phone,
	start_date,
	end_date,
	processed
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetReservation :one
SELECT * FROM reservations
WHERE id = $1 LIMIT 1;

-- name: GetReservationForUpdate :one
SELECT * FROM reservations
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListReservations :many
SELECT * FROM reservations
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateReservation :one
UPDATE reservations
SET
  updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
  room_id = COALESCE(sqlc.narg(room_id), room_id),
  first_name = COALESCE(sqlc.narg(first_name), first_name),
  last_name = COALESCE(sqlc.narg(last_name), last_name),
  email = COALESCE(sqlc.narg(email), email),
  phone = COALESCE(sqlc.narg(phone), phone),
  start_date = COALESCE(sqlc.narg(start_date), start_date),
  end_date = COALESCE(sqlc.narg(end_date), end_date),
  processed = COALESCE(sqlc.narg(processed), processed)
WHERE
id = sqlc.arg(id)
RETURNING *;

-- name: UpdateProcessedForReservation :one
UPDATE reservations
SET
  processed = $1
WHERE
id = $2
RETURNING *;

-- name: DeleteReservation :exec
DELETE FROM reservations
WHERE id = $1;

-- name: AllReservations :many
select r.id as reservation_id, rm.room_guest_number, rm.room_price_en, r.first_name, r.last_name, r.email, r.phone, r.start_date, 
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, rm.room_name_sr,
rm.room_name_en,
rm.room_name_bg
from reservations r
left join rooms rm on (r.room_id = rm.id)
order by r.start_date asc
LIMIT $1
OFFSET $2;

-- name: AllNewReservations :many
select r.id as reservation_id,rm.room_guest_number, rm.room_price_en, r.first_name, r.last_name, r.email, r.phone, r.start_date, 
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, rm.room_name_sr,
rm.room_name_en,
rm.room_name_bg
from reservations r
left join rooms rm on (r.room_id = rm.id)
where processed = 0
order by r.start_date asc
LIMIT $1
OFFSET $2;

-- name: AllProcessedReservations :many
select r.id as reservation_id,rm.room_guest_number, rm.room_price_en, r.first_name, r.last_name, r.email, r.phone, r.start_date, 
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, rm.room_name_sr,
rm.room_name_en,
rm.room_name_bg
from reservations r
left join rooms rm on (r.room_id = rm.id)
where processed = 1
order by r.start_date asc
LIMIT $1
OFFSET $2;

-- name: GetReservationByID :one
select r.id as reservation_id,rm.room_guest_number, rm.room_price_en, r.first_name, r.last_name, r.email, r.phone, r.start_date, 
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, rm.room_name_sr,
rm.room_name_en,
rm.room_name_bg
from reservations r
left join rooms rm on (r.room_id = rm.id)
where r.id = $1
order by r.start_date asc
LIMIT $2
OFFSET $3;