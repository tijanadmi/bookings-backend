-- name: CreateRestriction :one
INSERT INTO restrictions (
	restriction_name_sr,
	restriction_name_en,
	restriction_name_bg
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetRestriction :one
SELECT * FROM restrictions
WHERE id = $1 LIMIT 1;

-- name: GetRestrictionForUpdate :one
SELECT * FROM restrictions
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListRestrictions :many
SELECT * FROM restrictions
ORDER BY id
LIMIT $1
OFFSET $2;



-- name: UpdateRestriction :one
UPDATE restrictions
SET
  updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
  restriction_name_sr = COALESCE(sqlc.narg(restriction_name_sr), restriction_name_sr),
  restriction_name_en = COALESCE(sqlc.narg(restriction_name_en), restriction_name_en),
  restriction_name_bg = COALESCE(sqlc.narg(restriction_name_bg), restriction_name_bg)
WHERE
id = sqlc.arg(id)
RETURNING *;

-- name: DeleteRestriction :exec
DELETE FROM restrictions
WHERE id = $1;