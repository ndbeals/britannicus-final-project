DROP TABLE IF EXISTS tblUser CASCADE;

create table tblUser (
	user_id SERIAL PRIMARY KEY,
	user_name VARCHAR(25),
	user_email VARCHAR(40),
	user_password VARCHAR(80)
);
insert into tblUser (user_name, user_email, user_password ) values ('Admin', 'admin@britannicus.com', 'password');
insert into tblUser (user_name, user_email, user_password ) values ('User1', 'user1@britannicus.com', '$2a$10$hajTTtG1HVzs.ckfX8wZVOx8xynbcjT83HxVaxxjioay6D/xpBQOe');
-- insert into tblUser (customer_id, user_name, user_email, user_password ) values (3, 'Admin', 'admin@britannicus.com', '$2a$10$iRYt5JApjFKvs0nhRaaY0ueh0t5sigzpA81B2oQ7T5Sro4X5KZbLq');
