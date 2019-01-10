CREATE TABLE IF NOT EXISTS ocha.users (
	`id` INT(11) NOT NULL,
	`name` VARCHAR(20) CHARACTER SET 'utf8' COLLATE 'utf8_spanish2_ci' NOT NULL,
	`password` VARCHAR(255) CHARACTER SET 'utf8' COLLATE 'utf8_spanish2_ci' NOT NULL,
	PRIMARY KEY (`id`, `name`)
);

INSERT INTO users VALUES (101010,'lider1','$2a$10$FS02aspjeb0stqFypP.EIuhzKGpQ5wCnJm5sboCEDfvpFonnmn.xy');
