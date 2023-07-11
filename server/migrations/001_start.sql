-- +goose Up

CREATE TABLE todos(
	id SERIAL PRIMARY KEY,
	title VARCHAR(50) UNIQUE NOT NULL,
	description TEXT,
  created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE todos;
