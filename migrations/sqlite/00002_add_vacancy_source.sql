-- +goose Up
-- +goose StatementBegin
ALTER TABLE vacancies ADD COLUMN source TEXT NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE vacancies ADD COLUMN responsibility TEXT;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE vacancies ADD COLUMN remote INTEGER NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_vacancies_created_at ON vacancies (created_at DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_vacancies_created_at;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE vacancies DROP COLUMN remote;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE vacancies DROP COLUMN responsibility;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE vacancies DROP COLUMN source;
-- +goose StatementEnd
