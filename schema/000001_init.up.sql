CREATE TABLE users
(
    id           serial       not null unique PRIMARY KEY,
    first_name   VARCHAR(255) not null,
    last_name    VARCHAR(255) not null,
    email        VARCHAR(255) not null unique,
    password     VARCHAR(255) not null,
    phone_number VARCHAR(255) not null,
    role         VARCHAR(255) not null
);

CREATE TABLE degrees
(
    id   serial       not null unique PRIMARY KEY,
    name VARCHAR(255) not null
);

CREATE TABLE faculties
(
    id   serial       not null unique PRIMARY KEY,
    name VARCHAR(255) not null
);

CREATE TABLE ep
(
    id         serial       not null unique PRIMARY KEY,
    name       VARCHAR(255) not null,
    faculty_id INT          not null,
    FOREIGN KEY (faculty_id) REFERENCES faculties (id)
);

CREATE TABLE progress
(
    id   serial       not null unique PRIMARY KEY,
    name VARCHAR(255) not null
);

CREATE TABLE types
(
    id   serial       not null unique PRIMARY KEY,
    name VARCHAR(255) not null
);

CREATE TABLE statuses
(
    id   serial       not null unique PRIMARY KEY,
    name VARCHAR(255) not null
);

CREATE TABLE instructors
(
    id      serial not null unique PRIMARY KEY,
    user_id INT    not null,
    about   TEXT   not null,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE students
(
    id         serial       not null unique PRIMARY KEY,
    user_id    INT          not null,
    degree_id  INT          not null,
    ep_id      INT          not null,
    group_name VARCHAR(255) not null,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (degree_id) REFERENCES degrees (id),
    FOREIGN KEY (ep_id) REFERENCES ep (id)
);

CREATE TABLE works
(
    id              serial       not null unique PRIMARY KEY,
    created_at      DATE     not null,
    title           VARCHAR(255) not null unique,
    description     TEXT         not null,
    type_id         INT          not null,
    degree_id       INT          not null,
    instructor_id   INT          not null,
    student_id   INT          not null,
    is_approved     BOOLEAN      not null,
    progress_id     INT,
    FOREIGN KEY (type_id) REFERENCES types (id),
    FOREIGN KEY (degree_id) REFERENCES degrees (id),
    FOREIGN KEY (instructor_id) REFERENCES instructors (id),
    FOREIGN KEY (student_id) REFERENCES students (id),
    FOREIGN KEY (progress_id) REFERENCES progress (id)
);

CREATE TABLE works_eps
(
    work_id INT not null,
    ep_id   INT not null,
    FOREIGN KEY (work_id) REFERENCES works (id),
    FOREIGN KEY (ep_id) REFERENCES ep (id)
);



CREATE TABLE requests
(
    id          serial   not null unique PRIMARY KEY,
    created_at  DATE not null,
    work_id     INT      not null,
    student_id  INT      not null,
    status_id   INT      not null,
    description TEXT     not null,
    FOREIGN KEY (work_id) REFERENCES works (id),
    FOREIGN KEY (student_id) REFERENCES students (id),
    FOREIGN KEY (status_id) REFERENCES statuses (id)
);




