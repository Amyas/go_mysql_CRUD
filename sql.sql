create table user_info(
  id int(11) primary key auto_increment,
  username varchar(20),
  departname varchar(20),
  create_time timestamp default current_timestamp
);