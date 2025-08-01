UPDATE users
SET
    total_requests        = $2,
    created_at            = $3,
    generations_per_month = $4
WHERE id = $1