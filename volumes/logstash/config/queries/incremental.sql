SELECT
	j.journal_id, j.action_type,r.*
FROM restaurants_journal j
LEFT JOIN restaurants r ON r.id = j.id
WHERE j.journal_id > :sql_last_value
	AND j.action_time < NOW()
ORDER BY j.journal_id;