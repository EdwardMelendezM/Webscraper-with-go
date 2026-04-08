-- +migrate Up
BEGIN;

CREATE TABLE IF NOT EXISTS scraped_topics (
  id uuid PRIMARY KEY,
  project_id uuid NOT NULL,
  title text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  deleted_at timestamptz NULL
);

CREATE INDEX IF NOT EXISTS idx_scraped_topics_project_id_created_at
  ON scraped_topics (project_id, created_at DESC)
  WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS scraped_results (
  id uuid PRIMARY KEY,
  project_id uuid NOT NULL,
  title text NOT NULL,
  url text NOT NULL,
  content text NOT NULL,
  number integer NOT NULL,
  title_corpus text NULL,
  content_corpus text NULL,
  work_key text NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  deleted_at timestamptz NULL
);

CREATE INDEX IF NOT EXISTS idx_scraped_results_project_id_created_at
  ON scraped_results (project_id, created_at)
  WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_scraped_results_project_id_number
  ON scraped_results (project_id, number DESC);

CREATE INDEX IF NOT EXISTS idx_scraped_results_project_id_url
  ON scraped_results (project_id, url);

COMMIT;

-- +migrate Down
BEGIN;

DROP TABLE IF EXISTS scraped_results;
DROP TABLE IF EXISTS scraped_topics;

COMMIT;

