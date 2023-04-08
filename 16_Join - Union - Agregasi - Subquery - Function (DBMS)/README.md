1. Contoh sintaksis subquery dalam SQL:
SELECT column_name(s)
FROM table_name
WHERE column_name IN (SELECT column_name FROM table_name WHERE condition);
2. Terdapat beberapa jenis join dalam SQL, antara lain:
a. INNER JOIN: Menghasilkan kumpulan data yang hanya mencakup baris yang memiliki nilai kunci yang cocok di kedua 
tabel yang bergabung.
b. LEFT JOIN: Menghasilkan kumpulan data yang mencakup semua baris dari tabel kiri dan hanya baris yang memiliki 
nilai kunci cocok di tabel kanan.
c. RIGHT JOIN: Menghasilkan kumpulan data yang mencakup semua baris dari tabel kanan dan hanya baris yang memiliki 
nilai kunci cocok di tabel kiri.
d. FULL OUTER JOIN: Menghasilkan kumpulan data yang mencakup semua baris dari kedua tabel, tanpa memandang apakah 
ada nilai kunci yang cocok atau tidak.
3. Contoh sintaksis Union dalam SQL:
SELECT column_name(s) FROM table1
UNION
SELECT column_name(s) FROM table2;
Dalam contoh di atas, Union digunakan untuk menggabungkan hasil query dari dua tabel (table1 dan table2) dengan kolom yang sama, dan menghasilkan kumpulan data tunggal yang tidak berisi duplikat data.