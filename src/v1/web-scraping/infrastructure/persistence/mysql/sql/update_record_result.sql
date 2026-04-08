UPDATE scraped_results
SET content = $1
WHERE id = $2
  AND project_id = $3;
