-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id                    INT PRIMARY KEY,
    total_requests        bigint,
    created_at            TIMESTAMP,
    generations_per_month bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
