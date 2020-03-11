use sample_db;

CREATE TABLE products (
  id int(10) unsigned not null auto_increment,
  title varchar(255) not null,
  price double not null,
  description text not null,
  imageUrl varchar(255) not null,
  primary key (id),
  unique index id_unique (id asc) visible
);

-- sample data
insert into products (
  title, price, description, imageUrl
  ) values (
    'Book', 
    '11', 
    'this is a good book!', 
    'https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1267&q=80'
  )