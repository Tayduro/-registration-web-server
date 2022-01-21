CREATE TABLE IF NOT EXISTS users
(

        first_name character varying(255) NOT NULL,
        last_name character varying(255) NOT NULL,
        email character varying(255) unique NOT NULL,
        password character varying(255) NOT NULL,
        PRIMARY KEY (first_name, last_name, email, password)

)