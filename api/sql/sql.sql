CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

GRANT ALL PRIVILEGES ON devbook.* TO 'devbook'@'%';
FLUSH PRIVILEGES;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int primary key auto_increment,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;