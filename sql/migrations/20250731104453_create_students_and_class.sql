-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS classes (
    name text NOT NULL,
    program_code text NOT NULL,
    year integer NOT NULL,
    study_type text NOT NULL,
    active boolean NOT NULL,
    PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS students (
    uid text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    active boolean NOT NULL,
    class_name text NOT NULL REFERENCES classes(name),
    PRIMARY KEY (uid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS classes;
-- +goose StatementEnd
