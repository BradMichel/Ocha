CREATE TABLE IF NOT EXISTS ocha.typesOfThematics (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `nombre` VARCHAR(11) CHARACTER SET 'utf8' COLLATE 'utf8_spanish2_ci',
    PRIMARY KEY (`id`)
);

INSERT INTO typesOfThematic VALUES (NULL, 'Paz');
INSERT INTO typesOfThematic VALUES (NULL, 'Desarrollo');
INSERT INTO typesOfThematic VALUES (NULL, 'Humanitario');
INSERT INTO typesOfThematic VALUES (NULL, 'Coyuntura');
