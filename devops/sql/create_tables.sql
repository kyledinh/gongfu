CREATE SEQUENCE users_seq INCREMENT 1 START 100 CACHE 1;
CREATE TABLE users (
    id int4 NOT NULL DEFAULT nextval('users_seq'::regclass),
    accountid INT,
    name VARCHAR(120),
    email VARCHAR(120),
    password VARCHAR(120),
    active BOOLEAN default 't'
);

CREATE SEQUENCE chirps_seq INCREMENT 1 START 100 CACHE 1;
CREATE TABLE chirps (
    id INT NOT NULL DEFAULT NEXTVAL('chirps_seq'::regclass),
    userid INT,
    type VARCHAR(24),    
    message VARCHAR(140),
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    active BOOLEAN DEFAULT TRUE
);

BEGIN;
INSERT INTO users VALUES ('1', '1', 'Kyle Dinh', 'kyle@redbridge.io', '$2a$10$RHkjXScW9MbtuZN1rcQR0Ouiyte3FICmDjXxtWpxF0KD7..j/vBbC');
INSERT INTO chirps (userid, type, message) VALUES ('1', 'Bot', 'Here we are');
COMMIT;
