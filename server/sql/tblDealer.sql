DROP TABLE IF EXISTS tblDealer CASCADE;

create table tblDealer (
	dealer_id SERIAL PRIMARY KEY,
	first_name VARCHAR(25),
	last_name VARCHAR(25),
	email VARCHAR(40),
	phone_number VARCHAR(20),
	dealer_address VARCHAR(40),
	dealer_city VARCHAR(30),
	dealer_state VARCHAR(30),
	dealer_country VARCHAR(30)
);
insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Bruce', 'King', 'bking@latimes.com', '+1 513 380 5466', '134 Center Terrace', 'Cincinnati', 'Ohio', 'United States');
insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Brandon', 'Johns', 'brandonJ@mayoclinic.com', '+44 889 132 8214', '4 Scoville Pass', 'Hatton', 'England', 'United Kingdom');
insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Jake', 'Holden', 'JakeHolden@gov.uk', '+1 229 182 0449', '59 Lotheville Junction', 'Albany', 'Georgia', 'United States');
-- insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Lola', 'Glasscoe', 'lglasscoe3@archive.org', '+1 917 152 0825', '327 Towne Crossing', 'New York City', 'New York', 'United States');
-- insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Martino', 'Pucker', 'mpucker4@msn.com', '+1 757 462 9467', '29010 Sage Road', 'Virginia Beach', 'Virginia', 'United States');
-- insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Elaine', 'Hodge', 'ehodge5@omniture.com', '+1 915 330 0761', '556 Steensland Place', 'El Paso', 'Texas', 'United States');
-- insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Robena', 'Symon', 'rsymon6@bbc.co.uk', '+1 469 698 8225', '35 Westport Plaza', 'Dallas', 'Texas', 'United States');
-- insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Sileas', 'Vanezis', 'svanezis7@bluehost.com', '+1 551 670 3130', '87 Boyd Center', 'Bridgewater', 'Nova Scotia', 'Canada');
-- insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Benedetto', 'Pankethman', 'bpankethman8@canalblog.com', '+1 901 473 8290', '3179 Melody Hill', 'Wingham', 'Ontario', 'Canada');
-- insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Sauveur', 'Hinzer', 'shinzer9@rakuten.co.jp', '+31 691 459 0731', '04 Onsgard Street', 'Haarlem', 'Provincie Noord-Holland', 'Netherlands');
-- insert into tblDealer (first_name, last_name, email, phone_number, dealer_address, dealer_city, dealer_state, dealer_country) values ('Jamima', 'Goodenough', 'jgoodenougha@yandex.ru', '+31 887 197 0751', '73 Buell Drive', 'Hilversum', 'Provincie Noord-Holland', 'Netherlands');