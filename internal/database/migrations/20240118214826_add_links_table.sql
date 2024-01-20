-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS links (
  id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  short_code TEXT NOT NULL,
  original_url TEXT NOT NULL,
  http_redirect_code SMALLINT NOT NULL,
  is_active BOOLEAN NOT NULL,
  description TEXT,
  tags TEXT[],
  password TEXT,
  expires_at TIMESTAMPTZ,
  created_by_ip TEXT,
  created_by_user_agent TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE links ADD CONSTRAINT links_short_code_unique UNIQUE (short_code);
CREATE INDEX IF NOT EXISTS links_short_code_idx ON links (short_code);
CREATE INDEX IF NOT EXISTS links_original_url_idx ON links (original_url);
CREATE INDEX IF NOT EXISTS links_is_active_idx ON links (is_active);
CREATE INDEX IF NOT EXISTS links_tags_idx ON links USING GIN (tags);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE links DROP CONSTRAINT IF EXISTS links_short_code_unique;
DROP INDEX IF EXISTS links_short_code_idx;
DROP INDEX IF EXISTS links_original_url_idx;
DROP INDEX IF EXISTS links_is_active_idx;
DROP INDEX IF EXISTS links_tags_idx;
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
