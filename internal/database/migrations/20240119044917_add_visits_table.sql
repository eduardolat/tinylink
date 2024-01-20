-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS visits (
  id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  link_id UUID NOT NULL REFERENCES links (id) ON DELETE CASCADE,
  ip TEXT NOT NULL,
  user_agent TEXT NOT NULL,
  referer TEXT,
  is_redirected BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS visits_link_id_idx ON visits (link_id);
CREATE INDEX IF NOT EXISTS visits_ip_idx ON visits (ip);
CREATE INDEX IF NOT EXISTS visits_is_redirected_idx ON visits (is_redirected);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS visits_link_id_idx;
DROP INDEX IF EXISTS visits_ip_idx;
DROP INDEX IF EXISTS visits_is_redirected_idx;
DROP TABLE IF EXISTS visits;
-- +goose StatementEnd
