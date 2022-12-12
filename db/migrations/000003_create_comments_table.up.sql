CREATE TABLE comments(
    id BIGSERIAL PRIMARY KEY  NOT NULL,
    post_id INT NOT NULL ,
    name VARCHAR(150) NOT NULL ,
    email VARCHAR(150) NOT NULL ,
    body TEXT NOT NULL,
    CONSTRAINT fk_post
        FOREIGN KEY (post_id)
            REFERENCES posts(id)
            ON DELETE CASCADE
            ON UPDATE CASCADE
);