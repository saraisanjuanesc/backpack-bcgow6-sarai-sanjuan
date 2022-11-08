# Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
SELECT e.nombre, e.puesto, d.localidad
FROM empleado e INNER JOIN departamento d
ON e.depto_nro = d.depto_nro;

#Visualizar los departamentos con más de cinco empleados.
SELECT d.nombre_depto, count(*) AS empleados
FROM empleado e INNER JOIN departamento d
ON e.depto_nro = d.depto_nro
WHERE empleados > 5;

# Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
SELECT concat(e.nombre, " ", e.apellido) as nombre, e.salario, d.nombre_depto
FROM empleado e INNER JOIN departamento d
ON e.depto_nro = d.depto_nro
WHERE nombre = "Mito Barchuk";

# Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
SELECT *
FROM empleado e INNER JOIN departamento d
ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto = "Contabilidad";

# Mostrar el nombre del empleado que tiene el salario más bajo.
SELECT nombre
FROM empleado 
WHERE salario = (SELECT min(salario) FROM empleado );

# Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT *
FROM empleado 
WHERE salario IN (SELECT max(salario) 
		FROM empleado e INNER JOIN departamento d
		ON e.depto_nro = d.depto_nro
        WHERE d.nombre_depto = "Ventas");
