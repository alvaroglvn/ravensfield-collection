-- name: SelectRandStyle :one
SELECT style
FROM art_museum
ORDER BY RANDOM()
LIMIT 1;