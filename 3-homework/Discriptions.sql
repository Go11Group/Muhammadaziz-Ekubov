--UPDATE
    --Ma'lumotlarni yangilash uchun ishlatiladi.

--Sintaksis:

UPDATE jadval_nomi
SET ustun1 = qiymat1, ustun2 = qiymat2, ...
WHERE shart;
--Misol:

UPDATE talaba
SET yosh = 22
WHERE ism = 'Ali';

--DELETE

--Ma'lumotlarni o'chirish uchun ishlatiladi.

--Sintaksis:

DELETE FROM jadval_nomi
WHERE shart;

--Misol:


DELETE FROM talaba
WHERE ism = 'Ali';
GROUP BY

--Ma'lumotlarni guruhlash uchun ishlatiladi.

  --  Sintaksis:

SELECT ustun, COUNT(*)
FROM jadval_nomi
GROUP BY ustun;

--Misol:

SELECT fakul, COUNT(*)
FROM talaba
GROUP BY fakul;
ORDER BY

--Ma'lumotlarni tartiblash uchun ishlatiladi.

--Sintaksis:

SELECT ustun1, ustun2, ...
FROM jadval_nomi
ORDER BY ustun1 [ASC|DESC];

--Misol:


SELECT ism, yosh
FROM talaba
ORDER BY yosh DESC;

--JOIN

--Ikkita yoki undan ortiq jadvalni birlashtirish uchun ishlatiladi.

--Sintaksis:


SELECT jadval1.ustun1, jadval2.ustun2, ...
FROM jadval1
JOIN jadval2
ON jadval1.ustun = jadval2.ustun;

--Misol:


SELECT talaba.ism, kurs.nomi
FROM talaba
JOIN kurs
ON talaba.kurs_id = kurs.id;