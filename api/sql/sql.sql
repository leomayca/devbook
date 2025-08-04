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

DROP TABLE IF EXISTS seguidores;

CREATE TABLE seguidores(
    usuario_id int not null,
    seguidor_id int not null,
    primary key (usuario_id, seguidor_id),
    foreign key (usuario_id) references usuarios(id) on delete cascade,
    foreign key (seguidor_id) references usuarios(id) on delete cascade
) ENGINE=INNODB;

DROP TABLE IF EXISTS publicacoes;

CREATE TABLE publicacoes(
    id int primary key auto_increment,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_id int not null,
    curtidas int default 0,
    criadoEm timestamp default current_timestamp(),
    foreign key (autor_id) references usuarios(id) on delete cascade
) ENGINE=INNODB;