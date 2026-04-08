INSERT INTO scraped_results (id, project_id, title, url, content, number, title_corpus, content_corpus, work_key,
                             created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
