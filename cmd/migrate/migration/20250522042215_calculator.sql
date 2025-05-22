-- +goose Up
-- +goose StatementBegin
CREATE TABLE calculations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    principal FLOAT,
    rate FLOAT,
    compounds_per_year INT,
    years FLOAT,
    result FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd




-- +goose Down
-- +goose StatementBegin
DROP TABLE calculations;
DROP TABLE users
-- +goose StatementEnd
