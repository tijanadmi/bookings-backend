-- name: CreateRoom :one
INSERT INTO rooms (
	room_name_sr,
	room_name_en,
	room_name_bg,
	room_short_des_sr,
	room_short_des_en,
	room_short_des_bg,
	room_description_sr,
	room_description_en,
	room_description_bg,
	room_pictures_folder,
	room_guest_number,
	room_price_en
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: GetRoom :one
SELECT * FROM rooms
WHERE id = $1 LIMIT 1;

-- name: GetRoomForUpdate :one
SELECT * FROM rooms
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListRooms :many
SELECT * FROM rooms
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: SearchAvailabilityForAllRooms :many
select r.* from rooms r
where r.id not in 
		(select room_id from room_restrictions rr where $1 < rr.end_date and $2 >= rr.start_date)
ORDER BY r.id
LIMIT $3
OFFSET $4;

-- name: UpdateRoom :one
UPDATE rooms
SET
  updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
  room_name_sr = COALESCE(sqlc.narg(room_name_sr), room_name_sr),
  room_name_en = COALESCE(sqlc.narg(room_name_en), room_name_en),
  room_name_bg = COALESCE(sqlc.narg(room_name_bg), room_name_bg),
  room_short_des_sr = COALESCE(sqlc.narg(room_short_des_sr), room_short_des_sr),
  room_short_des_en = COALESCE(sqlc.narg(room_short_des_en), room_short_des_en),
  room_short_des_bg = COALESCE(sqlc.narg(room_short_des_bg), room_short_des_bg),
  room_description_sr = COALESCE(sqlc.narg(room_description_sr), room_description_sr),
  room_description_en = COALESCE(sqlc.narg(room_description_en), room_description_en),
  room_description_bg = COALESCE(sqlc.narg(room_description_bg), room_description_bg),
  room_pictures_folder = COALESCE(sqlc.narg(room_pictures_folder), room_pictures_folder),
  room_guest_number = COALESCE(sqlc.narg(room_guest_number), room_guest_number),
  room_price_en = COALESCE(sqlc.narg(room_price_en), room_price_en)
WHERE
id = sqlc.arg(id)
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id = $1;