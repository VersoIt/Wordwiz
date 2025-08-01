UPDATE users
SET id                    = $1,
    total_requests        = $2,
    created_at            = $3,
    generations_per_month = $4