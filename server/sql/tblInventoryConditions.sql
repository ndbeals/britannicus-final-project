DROP TABLE IF EXISTS tblInventoryConditions CASCADE;

create table tblInventoryConditions (
	condition_id SERIAL PRIMARY KEY,
	condition VARCHAR(25)
);
insert into tblInventoryConditions (condition_id, condition) values (1, 'As New');
insert into tblInventoryConditions (condition_id, condition) values (2, 'Fine');
insert into tblInventoryConditions (condition_id, condition) values (3, 'Very Good');
insert into tblInventoryConditions (condition_id, condition) values (4, 'Good');
insert into tblInventoryConditions (condition_id, condition) values (5, 'Fair');
insert into tblInventoryConditions (condition_id, condition) values (6, 'Poor');
insert into tblInventoryConditions (condition_id, condition) values (7, 'Ex-Library');
insert into tblInventoryConditions (condition_id, condition) values (8, 'Book Club');
insert into tblInventoryConditions (condition_id, condition) values (9, 'Binding Copy');
