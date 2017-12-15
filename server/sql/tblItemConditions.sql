DROP TABLE IF EXISTS tblItemConditions CASCADE;

create table tblItemConditions (
	condition_id SERIAL PRIMARY KEY,
	condition VARCHAR(25)
);
insert into tblItemConditions (condition_id, condition) values (1, 'As New');
insert into tblItemConditions (condition_id, condition) values (2, 'Fine');
insert into tblItemConditions (condition_id, condition) values (3, 'Very Good');
insert into tblItemConditions (condition_id, condition) values (4, 'Good');
insert into tblItemConditions (condition_id, condition) values (5, 'Fair');
insert into tblItemConditions (condition_id, condition) values (6, 'Poor');
insert into tblItemConditions (condition_id, condition) values (7, 'Ex-Library');
insert into tblItemConditions (condition_id, condition) values (8, 'Book Club');
insert into tblItemConditions (condition_id, condition) values (9, 'Binding Copy');
