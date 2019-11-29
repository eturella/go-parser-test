-- # ------------------- #
select * from tabella

-- # ------------------- #
select * from mydb.tabella

-- # ------------------- #
select col1, col2, col3 from tabella

-- # ------------------- #
select col1, col2, col3 from tabella ORDER BY col1 DESC LIMIT 0, 25

-- # ------------------- #
select col1, lower(col2), col3 * col5 from tabella

-- # ------------------- #
select col1 al1, col2 + col3 al2 from tabella

-- # ------------------- #
select col1, col2 from tabella WHERE col3 = 5 AND (char_length(col6) > 3 OR left(col7, 1) = 'A') AND col7 BETWEEN 3 AND 45

-- # ------------------- #
select col1, max(col2) FROM tabella GROUP BY col1

-- # ------------------- #
select col1, max(col2) FROM tabella GROUP BY col1 HAVING max(col2) < 5

-- # ------------------- #
select a.col1 al1, a.col2 + col3 al2 from tabella a

-- # ------------------- #
select a.col1, b.col2, max(b.col3) tot from sk1.tabella a JOIN sk2.secondaTabella b ON a.id = b.id WHERE b.col9 <= 3

-- # ------------------- #
select a.col1, b.col2, max(b.col3) tot from sk1.tabella a NATURAL JOIN sk2.secondaTabella b WHERE b.col9 <= 3

-- # ------------------- #
select col1, (select max(col2) FROM secondaTabella b where a.id = b.id) from tabella a

-- # ------------------- #
select a.col1, b.col2, max(b.col3) tot from tabella a JOIN (select distinct id, col2, col3 FROM secondaTabella) b ON a.id = b.id WHERE b.col3 <= 3

