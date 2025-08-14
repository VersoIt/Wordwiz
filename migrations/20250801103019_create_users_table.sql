-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id                    BIGINT PRIMARY KEY,
    total_requests        BIGINT,
    created_at            TIMESTAMP,
    generations_per_month BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
