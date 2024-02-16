-- name: SelectRandMood :one
SELECT mood
FROM art_museum
ORDER BY RANDOM()
LIMIT 1;