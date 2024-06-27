CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  email varchar NOT NULL,
  hashedPassword varchar
);
