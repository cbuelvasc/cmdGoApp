DROP DATABASE IF EXISTS `go_db`;

CREATE DATABASE `go_db`;

USE `go_db`;

DROP TABLE IF EXISTS `clients`;

CREATE TABLE `clients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `created_date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO `clients` (`name`,`email`,`phone`,`created_date`) VALUES
	 ('Carmelo Buelvas Comas','cbc@gmail.com','3114389075',now()),
	 ('Raquel Buelvas','rbq@gmail.com','3001234567',now()),
	 ('Juan Perez Castillo','jpc@mail.com','+573001234567',now());