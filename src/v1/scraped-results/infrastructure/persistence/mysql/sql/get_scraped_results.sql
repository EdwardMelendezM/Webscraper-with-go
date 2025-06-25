SELECT id             AS id,
       project_id     AS project_id,
       title          AS title,
       url            AS url,
       content        AS content,
       number         AS number,
       title_corpus   AS title_corpus,
       content_corpus AS content_corpus,
       work_key       AS work_key,
       created_at     AS created_at
FROM scraped_results
WHERE project_id = ?
  AND deleted_at IS NULL
ORDER BY number DESC;
