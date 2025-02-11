-- name: CreateReservation :one
INSERT INTO reservations (
	room_id,
	first_name,
	last_name,
	email,
	phone,
	start_date,
	end_date,
	processed,
  num_nights,
  num_guests,
  status,
  total_price,
  extras_price,
  is_paid,
  has_breakfast
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
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
  processed = COALESCE(sqlc.narg(processed), processed),
  num_nights = COALESCE(sqlc.narg(num_nights), num_nights),
  num_guests = COALESCE(sqlc.narg(num_guests), num_guests),
  status = COALESCE(sqlc.narg(status), status),
  total_price = COALESCE(sqlc.narg(total_price), total_price),
  extras_price = COALESCE(sqlc.narg(extras_price), extras_price),
  is_paid = COALESCE(sqlc.narg(is_paid), is_paid),
  has_breakfast = COALESCE(sqlc.narg(has_breakfast), has_breakfast)
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
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, r.num_nights, r.num_guests, r.status, r.total_price, r.extras_price, r.is_paid, r.has_breakfast,
rm.room_name_sr,
rm.room_name_en,
rm.room_name_bg
from reservations r
left join rooms rm on (r.room_id = rm.id)
order by r.start_date asc
LIMIT $1
OFFSET $2;

-- name: AllNewReservations :many
select r.id as reservation_id,rm.room_guest_number, rm.room_price_en, r.first_name, r.last_name, r.email, r.phone, r.start_date, 
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, r.num_nights, r.num_guests, r.status, r.total_price, r.extras_price, r.is_paid, r.has_breakfast,
rm.room_name_sr,
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
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed,r.num_nights, r.num_guests, r.status, r.total_price, r.extras_price, r.is_paid, r.has_breakfast,
rm.room_name_sr,
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
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, r.num_nights, r.num_guests, r.status, r.total_price, r.extras_price, r.is_paid, r.has_breakfast,
rm.room_name_sr,
rm.room_name_en,
rm.room_name_bg
from reservations r
left join rooms rm on (r.room_id = rm.id)
where r.id = $1
order by r.start_date asc;

-- name: ListReservationsAfterDate :many
select r.id as reservation_id, rm.room_guest_number, rm.room_price_en, r.first_name, r.last_name, r.email, r.phone, r.start_date, 
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, r.num_nights, r.num_guests, r.status, r.total_price, r.extras_price, r.is_paid, r.has_breakfast,
rm.room_name_sr,
rm.room_name_en,
rm.room_name_bg
from reservations r
left join rooms rm on (r.room_id = rm.id)
where r.created_at::date BETWEEN $1 AND  $2
order by r.start_date asc;

-- name: ListStaysAfterDate :many
select r.id as reservation_id, rm.room_guest_number, rm.room_price_en, r.first_name, r.last_name, r.email, r.phone, r.start_date, 
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, r.num_nights, r.num_guests, r.status, r.total_price, r.extras_price, r.is_paid, r.has_breakfast,
rm.room_name_sr,
rm.room_name_en,
rm.room_name_bg
from reservations r
left join rooms rm on (r.room_id = rm.id)
where r.start_date::date BETWEEN $1 AND  $2
order by r.start_date asc;

-- name: ListTodayActivity :many
select r.id as reservation_id, rm.room_guest_number, rm.room_price_en, r.first_name, r.last_name, r.email, r.phone, r.start_date, 
r.end_date, r.room_id , r.created_at, r.updated_at, r.processed, r.num_nights, r.num_guests, r.status, r.total_price, r.extras_price, r.is_paid, r.has_breakfast,
rm.room_name_sr,
rm.room_name_en,
rm.room_name_bg
from reservations r
left join rooms rm on (r.room_id = rm.id)
 where (r.status = 'unconfirmed' AND r.start_date::date =$1)
    OR 
    (r.status = 'checked-in' AND r.end_date::date =$2)
order by r.start_date asc;
