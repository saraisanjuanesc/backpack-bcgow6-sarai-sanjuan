# Listar los datos de los autores.
SELECT *
FROM autor;

# Listar nombre y edad de los estudiantes
SELECT nombre, edad
FROM estudiante;

# ¿Qué estudiantes pertenecen a la carrera informática?
SELECT *
FROM estudiante
WHERE carrera = "Informática";

# ¿Qué autores son de nacionalidad francesa o italiana?
SELECT nombre
FROM autor
WHERE nacionalidad = "Francesa" or nacionalidad = "Italiana";

# ¿Qué libros no son del área de internet?
SELECT titulo
FROM libro
WHERE area = "Internet";

# Listar los libros de la editorial Salamandra.
SELECT titulo
FROM libro
WHERE editorial = "Salamandra";

# Listar los datos de los estudiantes cuya edad es mayor al promedio.
SELECT *
FROM estudiante
WHERE edad >= (SELECT avg(edad) from estudiante);

# Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
SELECT *
FROM estudiante 
WHERE apellido LIKE "G%";

# Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
SELECT a.nombre
FROM autor a INNER JOIN libroautor la
ON a.idAutor = la.idAutor
WHERE la.idLibro IN (SELECT idLibro FROM libro WHERE titulo = "El Universo: Guía de viaje");

# ¿Qué libros se prestaron al lector “Filippo Galli”?
SELECT l.titulo
FROM libro l INNER JOIN prestamo p INNER JOIN estudiante e
ON l.idLibro = p.idLibro AND e.idLector = p.idLector
WHERE "Filippo Galli" IN (SELECT concat(nombre, " ", apellido) FROM estudiante);

# Listar el nombre del estudiante de menor edad.
SELECT nombre
FROM estudiante 
WHERE edad IN (SELECT min(edad) FROM estudiante);

# Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
SELECT e.nombre
FROM estudiante e INNER JOIN prestamo p
ON  e.idLector = p.idLector
WHERE p.idLibro IN (SELECT idLibro FROM libro WHERE titulo = "Base de Datos");

# Listar los libros que pertenecen a la autora J.K. Rowling.
SELECT l.titulo 
FROM libro l INNER JOIN libroautor la
ON l.idLibro = la.idLibro
WHERE la.idAutor = (SELECT idAutor FROM autor WHERE nombre = "J. K. Rowling");

# Listar títulos de los libros que debían devolverse el 16/07/2021.
SELECT l.titulo 
FROM libro l INNER JOIN prestamo p
ON l.idLibro = p.idLibro
WHERE p.fechaDevolucion = '2021-07-16 00:00:00';

# ---------------------------- CREACIÓN DE LA BASE DE DATOS -----------------------
CREATE DATABASE biblioteca;
USE biblioteca;

CREATE TABLE `autor`(
	 `idAutor` int(10) unsigned NOT NULL AUTO_INCREMENT,
     `nombre` varchar (100) COLLATE utf8_unicode_ci NOT NULL,
     `nacionalidad` varchar (100) COLLATE utf8_unicode_ci DEFAULT NULL,
     PRIMARY KEY (`idAutor`)
);

CREATE TABLE `libro`(
	`idLibro` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`titulo` varchar (100) COLLATE utf8_unicode_ci NOT NULL,
	`editorial` varchar (100) COLLATE utf8_unicode_ci NOT NULL,
    `area` varchar (100) COLLATE utf8_unicode_ci DEFAULT NULL,
    PRIMARY KEY (`idLibro`)
);

CREATE TABLE `libroautor`(
	`idAutor` int(10) unsigned NOT NULL,
    `idLibro` int(10) unsigned NOT NULL,
    PRIMARY KEY (`idAutor`,`idLibro`),
    KEY `libroautor_idAutor_foreign`(`idAutor`),
	KEY `libroautor_idLibro_foreign`(`idLibro`),
	CONSTRAINT `libroautor_idAutor_foreign` FOREIGN KEY (`idAutor`) REFERENCES `autor` (`idAutor`),
    CONSTRAINT `libroautor_idLibro_foreign` FOREIGN KEY (`idAutor`) REFERENCES `libro` (`idLibro`)
);

CREATE TABLE `estudiante`(
	`idLector` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `nombre` varchar (100) COLLATE utf8_unicode_ci NOT NULL,
    `apellido` varchar (100) COLLATE utf8_unicode_ci NOT NULL,
    `direccion` varchar (100) COLLATE utf8_unicode_ci NOT NULL,
    `carrera` varchar (100) COLLATE utf8_unicode_ci NOT NULL,
    `edad` int(10) unsigned DEFAULT NULL,
    PRIMARY KEY (`idLector`)
);

CREATE TABLE `prestamo`(
	`idLector` int(10) unsigned NOT NULL,
    `idLibro` int(10) unsigned NOT NULL,
    `fechaPrestamo` timestamp NOT NULL,
    `fechaDevolucion`timestamp NOT NULL,
    `devuelto` BIT NOT NULL,
    PRIMARY KEY (`idLector`,`idLibro`),
    KEY `prestamo_idLector_foreign` (`idLector`),
    KEY `prestamo_idLibro_foreign` (`idLibro`),
    CONSTRAINT `prestamo_idLector_foreign` FOREIGN KEY (`idLector`) REFERENCES `estudiante` (`idLector`),
    CONSTRAINT `prestamo_idLibro_foreign` FOREIGN KEY (`idLibro`) REFERENCES `libro` (`idLibro`)
);

INSERT INTO `autor` VALUES 
(1, 'Gabriel García Márquez', 'Colombiana'),
(2, 'Oscar Wilde', 'Británica irlandesa'),
(3, 'J. K. Rowling', 'Británica'),
(4, 'J. D. Salinger', 'Estadounidense'),
(5, 'Victor Marie Hugo', 'Francesa'),
(6, 'Oliver Berry', 'Inglesa'),
(7, 'Valerie Stimac', 'Italiana');

INSERT INTO `libro` VALUES 
(1, 'Cien años de soledad', 'Diana', 'Internet'),
(2, 'El Universo: Guía de viaje', 'Lonely Planet', 'Internet'),
(3, 'Harry Potter and the Cursed Child', 'Salamandra', null),
(4, 'El Ickabog', 'Salamandra', null),
(5, 'El Retrato De Dorian Gray', 'Porrúa', 'Internet'),
(6, 'Los miserables', 'NoBook', null),
(7, 'El Guardian Entre El Centeno', 'Alianza ED', 'Internet');

INSERT INTO `libroautor` VALUES 
(1, 1), (2,5),(3,3),(3,4),(4,7),(5,6),(6,2),(7,2);

INSERT INTO `estudiante` VALUES 
(1, 'Louis', 'Homes', 'Saldiva 543', 'Informática',23),
(2, 'Gaston', 'Torres', 'Av. Independencia 467', 'Industrial',20),
(3, 'Filippo', 'Galli', 'Los olmos 284', 'Medicina',29),
(4, 'Mina', 'Gala', 'Malaga 483', 'Informática',25);

INSERT INTO `prestamo` VALUES 
(1,3,'2021-07-04 00:00:00', '2021-07-16 00:00:00',1),
(2,7,'2022-10-01 00:00:00', '2022-10-15 00:00:00',0),
(3,4,'2022-11-01 00:00:00', '2022-11-15 00:00:00',1),
(3,2,'2022-11-01 00:00:00', '2022-11-15 00:00:00',0);
