CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL
    );

-- default data
INSERT INTO users (name) VALUES ('admin master');

INSERT INTO users (name) VALUES ('normal star');

INSERT INTO users (name) VALUES ('simple fisher');
