CREATE DATABASE empresa_internet;
USE empresa_internet;

CREATE TABLE `planinternet`(
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `velocidad`int(10) unsigned NOT NULL,
    `precio` decimal(8,2) NOT NULL,
    `descuento` decimal(5,2) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `cliente`(
	`dni` varchar(9) NOT NULL,
    `nombre` varchar (100) COLLATE utf8_unicode_ci NOT NULL,
    `apellido` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
    `fecha_nacimiento` timestamp NOT NULL,
    `provincia` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
    `ciudad` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
    `id_planinternet` int(10) unsigned DEFAULT NULL,
    PRIMARY KEY (`dni`),
    KEY `cliente_planinternet_id_foreign` (`id_planinternet`),
    CONSTRAINT `cliente_planinternet_id_foreign` FOREIGN KEY (`id_planinternet`) REFERENCES `planinternet` (`id`)
);

INSERT INTO `planinternet` VALUES 
(1,20, 2750.00, 5.00),
(2,50, 3300.00, 5.00),
(3,100, 4772.00, 15.00),
(4,150, 5911.00, 15.00),
(5,300, 6272.00, 20.00);

INSERT INTO `cliente` VALUES 
('00523821F', 'Mario', 'Gonzalez','1996-11-04 00:00:00', 'Misiones', 'Posadas',1),
('20123821D', 'Marcela', 'Tapia','1992-07-14 00:00:00', 'Cordoba', 'Cordoba',2),
('35578823A', 'Nicolas', 'Zilli','1998-02-07 00:00:00', 'Santa Fe', 'Rosario',1),
('72390654G', 'Valeria', 'Estrada','1995-10-05 00:00:00', 'Salta', 'Salta',3),
('37514267Q', 'Montserrat', 'Gutierrez','1997-12-10 00:00:00', 'Entre Rios', 'Parana',5),
('93678201E', 'Rosario', 'Marcial','1999-09-09 00:00:00', 'San Luis', 'San Luis',2),
('02839719A', 'Luis', 'Sarmiento','1997-06-14 00:00:00', 'Buenos Aires', 'La Plata',3),
('23910847D', 'Angel', 'Gomez','1993-01-04 00:00:00', 'Mendoza', 'Godoy Cruz',4),
('93098738L', 'Nicole', 'Sosa','1994-09-22 00:00:00', 'Buenos Aires', 'San Miguel',4),
('73109832F', 'Elena', 'Torres','1991-07-23 00:00:00', 'Formosa', 'Formosa',5);



#################### CONSULTAS ######################
# 1. Mostrar el total de clientes que tienen un plan de velocidad con id = 5
SELECT   count(*) as total
FROM cliente
WHERE id_planinternet = 5;

# 2. Ordenar los clientes de acuerdo a su fecha de nacimiento en orden descendente
SELECT *
from cliente
ORDER BY fecha_nacimiento DESC;

# 3. Mostrar los primeros 3 clientes ordenados de acuerdo a su apellido
SELECT *
FROM cliente
ORDER BY apellido LIMIT 3;

# 4. Mostrar las provincias de donde son los clientes sin mostrar duplicados.
SELECT distinct provincia
FROM cliente;

# 5. Sumar los descuentos de los planes de internet
SELECT sum(descuento) as total
FROM planinternet;

# 6. Mostrar los clientes que su nombre inicie con la letra M
SELECT *
FROM cliente
WHERE nombre LIKE 'M%';

# 7. Mostrar los clientes que nacieron entre los años 1994 y 1997
SELECT * 
FROM cliente
WHERE year(fecha_nacimiento) BETWEEN 1994 AND 1997;

# 8. Mostrar el nombre y apellido de los clientes que tengan en su DNI el numero 38
SELECT nombre, apellido
FROM cliente
WHERE dni LIKE '%38%';

# 9. Mostrar los planes de interent que tienen un descuento de entre 5.00 y 17.00
SELECT *
FROM planinternet
WHERE descuento BETWEEN 5.00 AND 17.00;

# 10. Mostrar la velocidad de los planes de internet que su precio sea menor a 5000.00
SELECT velocidad
FROM planinternet
WHERE precio < 5000.00;