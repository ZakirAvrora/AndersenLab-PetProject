CREATE TABLE posts(
    id INT PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    title VARCHAR(150) NOT NULL,
    body TEXT NOT NULL,
    CONSTRAINT fk_post
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);