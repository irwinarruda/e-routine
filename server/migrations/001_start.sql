-- +goose Up

CREATE TABLE todos(
	id SERIAL PRIMARY KEY,
	title VARCHAR(50) UNIQUE NOT NULL,
	description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE todos;
