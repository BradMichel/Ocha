CREATE TABLE IF NOT EXISTS ocha.thematics (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `nombre` VARCHAR(27) CHARACTER SET 'utf8' COLLATE 'utf8_spanish2_ci',
    `idtype` INT(11),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`idtype`) REFERENCES typesOfThematic(`id`)
);

INSERT INTO thematics VALUES (NULL,'Conflictos en el territorio',1);
INSERT INTO thematics VALUES (NULL,'Derechos de las víctimas',1);
INSERT INTO thematics VALUES (NULL,'Movilización social',1);
INSERT INTO thematics VALUES (NULL,'Capacidad institucional',1);
INSERT INTO thematics VALUES (NULL,'Seguridad ciudadana',1);
INSERT INTO thematics VALUES (NULL,'Iniciativas posacuerdo',1);
INSERT INTO thematics VALUES (NULL,'Pobreza',2);
INSERT INTO thematics VALUES (NULL,'Mercado laboral',2);
INSERT INTO thematics VALUES (NULL,'Educación',2);
INSERT INTO thematics VALUES (NULL,'Salud',2);
INSERT INTO thematics VALUES (NULL,'Vivienda',2);
INSERT INTO thematics VALUES (NULL,'Sostenibilidad',2);
INSERT INTO thematics VALUES (NULL,'Perspectiva de género',2);
INSERT INTO thematics VALUES (NULL,'Concentración de la tierra',2);
INSERT INTO thematics VALUES (NULL,'Desplazamiento forzado',3);
INSERT INTO thematics VALUES (NULL,'Minas antipersonal',3);
INSERT INTO thematics VALUES (NULL,'Niñez y conflicto',3);
INSERT INTO thematics VALUES (NULL,'Comunidades étnicas',3);
INSERT INTO thematics VALUES (NULL,'Confinamiento y acceso hum',3);
INSERT INTO thematics VALUES (NULL,'Desastres naturales',3);
INSERT INTO thematics VALUES (NULL,'Cultura política',4);
INSERT INTO thematics VALUES (NULL,'Frontera',4);
INSERT INTO thematics VALUES (NULL,'Corrupción',4);
INSERT INTO thematics VALUES (NULL,'Empoderamiento ciudadano',4);
INSERT INTO thematics VALUES (NULL,'Mercado interno',4);
INSERT INTO thematics VALUES (NULL,'Gestión del conocimiento',4);
INSERT INTO thematics VALUES (NULL,'Etnodesarrollo',4);
