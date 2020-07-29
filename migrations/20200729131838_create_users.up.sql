CREATE TABLE users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    username VARCHAR NOT NULL UNIQUE,
    encrypted_password VARCHAR NOT NULL,
    role_id INT NOT NULL,
    CONSTRAINT fk_role
        FOREIGN KEY(role_id) 
            REFERENCES roles(role_id)
);