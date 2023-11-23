CREATE TYPE e_sex AS ENUM (
    'female',
    'male'
);

CREATE TABLE IF NOT EXISTS clients
(
  id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  full_name varchar(255) NOT NULL,
  age int NOT NULL,
  sex e_sex NOT NULL
);

CREATE TABLE IF NOT EXISTS workers
(
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    full_name varchar(255) NOT NULL,
    age int NOT NULL,
    sex e_sex  NOT NULL,
    position_id int
);

CREATE TABLE IF NOT EXISTS positions
(
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    title varchar(255) NOT NULL,
    salary int NOT NULL,
    description varchar(255)
);

ALTER TABLE IF EXISTS workers
ADD CONSTRAINT workers_fk FOREIGN KEY(position_id) REFERENCES positions(id)
ON DELETE SET NULL;
