SELECT id  AS scraped_result_id,
       url AS scraped_result_url
FROM scraped_results
ORDER BY created_at
LIMIT ?