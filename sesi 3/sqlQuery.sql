CREATE TABLE Book(
	ID     int, 
	Name   varchar(250),
	Genre  varchar(250),
	Author varchar(250)
);

insert into Book(id, name, genre, author)
values(6, 'basuki', 'komedi', 'joko');

select * from book;

ALTER TABLE book 
ALTER COLUMN ID SET DATA TYPE integer,
ALTER COLUMN ID SET NOT NULL;


ALTER TABLE book 
ADD COLUMN new_id serial PRIMARY KEY;

UPDATE book
SET new_id = ID;


alter table book
drop column ID;

alter table book
rename column new_id to ID;

ALTER TABLE book 
ALTER COLUMN id SET DATA TYPE serial;

CREATE SEQUENCE book_new_id_seq;

ALTER TABLE book 
ADD COLUMN id INTEGER NOT NULL DEFAULT nextval('book_new_id_seq') PRIMARY KEY;

alter table book
drop column ID;

ALTER TABLE book 
ALTER COLUMN id SET DATA TYPE serial;










