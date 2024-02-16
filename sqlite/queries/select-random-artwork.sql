-- name: SelectRandArtwork :one
SELECT artwork
FROM art_museum
ORDER BY RANDOM()
LIMIT 1;