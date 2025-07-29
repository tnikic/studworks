-- PostgreSQL

CREATE TABLE IF NOT EXISTS students (
    uid text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    active boolean NOT NULL,
    class_name text NOT NULL REFERENCES classes(name),
    PRIMARY KEY (uid)
);

CREATE TABLE IF NOT EXISTS teachers (
    uid text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    active boolean NOT NULL,
    PRIMARY KEY (uuid)
);

CREATE TABLE IF NOT EXISTS classes (
    name text NOT NULL,
    program_code text NOT NULL,
    year integer NOT NULL,
    study_type text NOT NULL,
    active boolean NOT NULL,
    PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS subjects (
    uuid uuid DEFAULT gen_random_uuid(),
    name text NOT NULL,
    program_code text NOT NULL,
    semester integer NOT NULL,
    PRIMARY KEY (uuid)
);

CREATE TABLE IF NOT EXISTS courses (
    uuid uuid DEFAULT gen_random_uuid(),
    subject_uuid uuid NOT NULL REFERENCES subjects(uuid),
    teacher_uid text NOT NULL REFERENCES teachers(uid),
    class_name text NOT NULL REFERENCES classes(name),
    PRIMARY KEY (uuid)
);

CREATE TABLE IF NOT EXISTS project_participants (
    uuid uuid DEFAULT gen_random_uuid(),
    points_earned integer NOT NULL,
    student_uid text NOT NULL REFERENCES students(uid),
    project_uuid uuid NOT NULL REFERENCES projects(uuid),
    PRIMARY KEY (uuid)
);

CREATE TABLE IF NOT EXISTS projects (
    uuid uuid DEFAULT gen_random_uuid(),
    title text NOT NULL,
    status text NOT NULL,
    possible_points integer NOT NULL,
    course_uuid uuid NOT NULL REFERENCES courses(uuid),
    PRIMARY KEY (uuid)
);

CREATE TABLE IF NOT EXISTS groups (
    uuid uuid DEFAULT gen_random_uuid(),
    id text NOT NULL,
    course_uuid uuid NOT NULL REFERENCES courses(uuid),
    PRIMARY KEY (uuid)
);

CREATE TABLE IF NOT EXISTS assignment_templates (
    uuid uuid DEFAULT gen_random_uuid(),
    title text NOT NULL,
    PRIMARY KEY (uuid)
);

CREATE TABLE IF NOT EXISTS assignments (
    uuid uuid DEFAULT gen_random_uuid(),
    assignment_template_uuid uuid NOT NULL REFERENCES assignment_templates(uuid),
    student_uid text NOT NULL REFERENCES students(uid),
    course_uuid uuid NOT NULL REFERENCES courses(uuid),
    possible_points integer NOT NULL,
    points_earned integer NOT NULL,
    PRIMARY KEY (uuid)
);
