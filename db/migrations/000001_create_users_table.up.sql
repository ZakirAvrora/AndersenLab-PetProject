CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR NOT NULL ,
    username VARCHAR(150) NOT NULL ,
    email VARCHAR(150) NOT NULL ,
    street  VARCHAR,
    suite  VARCHAR,
    city VARCHAR,
    zipcode VARCHAR,
    lat VARCHAR,
    lng VARCHAR,
    phone VARCHAR(50) NOT NULL,
    website VARCHAR,
    company_name VARCHAR,
    catch_phase TEXT,
    BS VARCHAR
);