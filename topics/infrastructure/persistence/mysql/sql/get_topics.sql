SELECT id         AS topic_id,
       title      AS topic_title,
       created_at AS topic_created_at
FROM scraped_topics
ORDER BY created_at DESC;