CREATE TABLE products
(
    identifier  UUID    NOT NULL
        CONSTRAINT products_pk PRIMARY KEY,
    price       NUMERIC NOT NULL,
    title       VARCHAR NOT NULL,
    description TEXT    NOT NULL
);

ALTER TABLE products
    owner TO ba;

INSERT INTO products
VALUES ('434eea72-22a6-4c61-b5ef-945874a5c478', 13.32, 'WSL1 Licence Key', 'An imaginary licence key for wsl'),
       ('6bfd28d2-863a-4cbb-951d-cc9be47bbdc1', 5.00, 'Nemo', 'The fish nemo, safe and sound!'),
       ('06b7af72-451e-482b-9ae1-32422a354146', 31.25, 'Tactical Turtleneck', 'One of the most powerful pieces of clothing in existence.'),
       ('c6a9c488-ecc9-4a8d-a25c-4aabd540ac37', 13.32, 'Diamond Sword', 'Stolen from steve during a nether exploration.'),
       ('6508ba08-78d1-4885-9805-0d4b3f8bc4e3', 2.00, 'Companion Cube', 'A simple cube made out of unknown matter.'),
       ('f6aae282-af6c-4733-9830-fde55bc0a7dc', 999.99, 'Mjolnir Powered Assault Armor', 'The best exoskeleton on the market'),
       ('66503f5d-a486-4c31-b772-3dec26ce0a53', 41.98, 'Gravity Gun', 'Still looking for the the third part of this..'),
       ('1041d793-27a8-4a73-b391-b7652c5a40fe', 666.66, 'BFG-9000', 'Capable of gooping monsters better than anything else!'),
       ('f4294d47-13f5-4cd0-87f7-508b419e2f64', 431.75, 'Blades of Chaos', 'Revenge against the gods has never been easier!'),
       ('0b1e533a-da2f-4535-9a25-eb86a164dd93', 13.32, 'Master Sword',
        'A link into the past sure seems easy with a sword like this.') ON CONFLICT DO NOTHING;
