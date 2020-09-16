create table users (
id BIGINT(20)  not null auto_increment,
first_name VARCHAR(100) NULL,
last_name VARCHAR(100) NULL,
email VARCHAR(100) not null ,
date_created VARCHAR(45) null,
primary key ( id),
unique index email_unique (email asc));


