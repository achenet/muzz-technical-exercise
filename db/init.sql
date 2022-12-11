CREATE DATABASE muzz;
USE muzz;
DROP TABLE IF EXISTS profiles;
CREATE TABLE profiles
    (id int AUTO_INCREMENT,
      email text,
      password text,
      name text,
      gender text,
      age int,
    PRIMARY KEY (id));
