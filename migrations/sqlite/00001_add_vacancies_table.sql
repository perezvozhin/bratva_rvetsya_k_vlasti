-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS vacancies (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    company TEXT NOT NULL,
    url TEXT NOT NULL,
    experience TEXT,
    salary_from INTEGER,
    salary_to INTEGER,
    currency TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vacancies;
-- +goose StatementEnd