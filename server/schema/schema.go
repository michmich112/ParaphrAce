package schema

var Schema = `
	CREATE TABLE user (
		id BIGSERIAL PRIMARY KEY,
		session_token TEXT NOT NULL
	);

	CREATE TABLE paraphrase (
		id BIGSERIAL PRIMARY KEY,
		user_id BIGINT NOT NULL,
		rating_id BIGINT,
		timestamp TIMESTAMPTZ NOT NULL,
		start_time TIMESTAMPTZ,
		end_time TIMESTAMPTZ,
		original_file_uri TEXT NOT NULL,
		result_file_uri TEXT
	);

	CREATE TABLE rating (
		id BIGSERIAL PRIMARY KEY,
		paraphrase_id BIGINT NOT NULL,
		timestamp TIMESTAMPTZ NOT NULL,
		comment TEXT,
		value SMALLINT
	);
`
