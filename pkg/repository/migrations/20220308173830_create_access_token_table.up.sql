CREATE TABLE IF NOT EXISTS access_token(
                                           user_id int,
                                           token VARCHAR,
                                           FOREIGN KEY (user_id)  REFERENCES users(user_id)
    )