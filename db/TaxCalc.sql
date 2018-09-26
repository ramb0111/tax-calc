/*
User table is not used for now as we are not actually inserting users
*/
DROP TABLE IF EXISTS users cascade;
CREATE TABLE users
(
	id SERIAL PRIMARY KEY,
	name CHARACTER VARYING(100) NOT NULL DEFAULT '',
	phone_no CHARACTER VARYING(20) NOT NULL,
	email CHARACTER VARYING(100) NOT NULL DEFAULT '',
	password CHARACTER VARYING(100) NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);


/* 
In case the `code` is changed for any `type`,
to make entries consistent in other tables,
I am using `id`.
*/
DROP TABLE IF EXISTS tax cascade;
CREATE TABLE tax
(
	id SERIAL PRIMARY KEY,
	code int NOT NULL,
	type CHARACTER VARYING(100) NOT NULL 
);

INSERT INTO tax VALUES (1,1,'food');
INSERT INTO tax VALUES (2,2,'tobacco');
INSERT INTO tax VALUES (3,3,'entertainment');


DROP TABLE IF EXISTS item cascade;
CREATE TABLE item
(
	id SERIAL PRIMARY KEY,
	name CHARACTER VARYING(100) NOT NULL,
	amount NUMERIC(18,2) NOT NULL,
	tax_amount NUMERIC(18,2) NOT NULL,
	total_amount NUMERIC(18,2) NOT NULL,
	tax_id bigint NOT NULL,
	FOREIGN KEY (tax_id) REFERENCES tax (id)
);


/*
User id foriegn key is commented as we are assuming user is already inserted
*/
DROP TABLE IF EXISTS user_item_mapping;
CREATE TABLE user_item_mapping(
	id SERIAL PRIMARY KEY,
	user_id bigint NOT NULL,
	item_id bigint NOT NULL,
	-- FOREIGN KEY (user_id) REFERENCES users (id),
	FOREIGN KEY (item_id) REFERENCES item (id)
);

/* Commands
psql postgres
create user postgres with password 'postgres' createdb superuser; 
create database tax_calc;
\c tax_calc 
\i /path/to/file.sql
*/