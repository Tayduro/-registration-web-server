CREATE TABLE IF NOT EXISTS users
(
    user_id  SERIAL PRIMARY KEY,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255) unique
    );