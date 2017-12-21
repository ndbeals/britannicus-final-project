DROP TABLE IF EXISTS tblInventoryConditions CASCADE;

create table tblInventoryConditions (
	condition_id SERIAL PRIMARY KEY,
	condition VARCHAR(25)
);
insert into tblInventoryConditions (condition) values ('As New');
insert into tblInventoryConditions (condition) values ('Fine');
insert into tblInventoryConditions (condition) values ('Very Good');
insert into tblInventoryConditions (condition) values ('Good');
insert into tblInventoryConditions (condition) values ('Fair');
insert into tblInventoryConditions (condition) values ('Poor');
insert into tblInventoryConditions (condition) values ('Ex-Library');
insert into tblInventoryConditions (condition) values ('Book Club');
insert into tblInventoryConditions (condition) values ('Binding Copy');
