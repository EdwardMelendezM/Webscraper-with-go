SELECT EXISTS(
  SELECT 1
  FROM scraped_results
  WHERE project_id = $1
    AND url = $2
);

