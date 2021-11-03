package schema

var Schema = `
	CREATE TABLE rating (
		id BIGSERIAL PRIMARY KEY,
		paraphrase_id BIGINT NOT NULL,
		timestamp TIMESTAMPTZ NOT NULL,
		comment TEXT,
		value SMALLINT
	);
`
