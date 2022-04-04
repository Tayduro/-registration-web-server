CREATE TABLE IF NOT EXISTS credentials(
                                          user_id int,
                                          salt VARCHAR,
                                          hash VARCHAR,
                                          FOREIGN KEY (user_id)  REFERENCES users(user_id)
    );