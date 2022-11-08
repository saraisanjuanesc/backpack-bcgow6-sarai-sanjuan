# 1. Mostrar el título y el nombre del género de todas las series
SELECT se.title, ge.name as genero
FROM series as se INNER JOIN genres as ge
ON se.genre_id = ge.id;

# 2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT ep.title, ac.first_name, ac.last_name
FROM episodes ep INNER JOIN actor_episode ae INNER JOIN actors ac
ON ep.id = ae.episode_id AND ac.id = ae.actor_id; 

# 3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT se.title, count(*) 
FROM series se INNER JOIN seasons t
ON se.id = t.serie_id
GROUP BY (se.title);

# 4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT g.name, count(*) as cantidad
FROM movies m INNER JOIN genres g ON m.genre_id = g.id
GROUP BY (g.name)
HAVING cantidad >= 3;

# 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT distinct ac.first_name, ac.last_name
FROM actors ac INNER JOIN actor_movie acm INNER JOIN movies m
ON ac.id = acm.actor_id AND m.id = acm.movie_id
WHERE m.title LIKE "La Guerra de las galaxias%"
