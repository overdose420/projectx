-- -- Creation of product table
-- CREATE TABLE IF NOT EXISTS product (
--   product_id INT NOT NULL,
--   name varchar(250) NOT NULL,
--   PRIMARY KEY (product_id)
-- );


CREATE TABLE  users
(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL
);
