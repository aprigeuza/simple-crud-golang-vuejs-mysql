-- Create crud Database
CREATE DATABASE `crud`;

-- Create Table Contacts
CREATE TABLE `contacts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `phone_no` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


INSERT INTO contacts (name,phone_no) VALUES
	 ('Jhon Wik Wik','+1 812 4421 156'),
	 ('Kim Jong un','+850 512 540213'),
	 ('Teh Yuli','+62 81245788956'),
	 ('Robert Surgey','+9 879 125 313'),
	 ('Putin','+7 9845 12482'),
	 ('Biden','+1 256 12546'),
	 ('Jokowi','+62 111 211 777');
