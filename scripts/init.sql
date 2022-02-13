CREATE TABLE IF NOT EXISTS users
(
    user_id  SERIAL PRIMARY KEY,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    email character varying(255) unique NOT NULL
);

CREATE TABLE IF NOT EXISTS credentials(
    user_id int,
    salt VARCHAR,
    hash VARCHAR,
    FOREIGN KEY (user_id)  REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS access_token(
     user_id int,
     token VARCHAR,
     FOREIGN KEY (user_id)  REFERENCES users(user_id)
)