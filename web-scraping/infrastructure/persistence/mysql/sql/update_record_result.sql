UPDATE scraped_results
SET content = ?
WHERE id = ?
  AND project_id = ?;