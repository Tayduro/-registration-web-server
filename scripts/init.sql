CREATE TABLE IF NOT EXISTS users
(
    user_id BIGINT NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    email character varying(255) unique NOT NULL
);

CREATE TABLE IF NOT EXISTS credentials(
    id_users_data BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY,
    salt VARCHAR,
    hash VARCHAR,
        CONSTRAINT users_fkey FOREIGN KEY (id_users_data) REFERENCES users (user_id)
    )