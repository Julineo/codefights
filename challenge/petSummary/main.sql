/*
https://app.codesignal.com/challenge/g3vs3meJBTykq36q8

A couple of years ago, you decided to open your own nursery for homeless and neglected pets. Now your nursery is growing very fast, so you decided to make a summary of the pets living there. For now you have pets of only four species - cats, dogs, weasels, and chinchillas (poor things!) - info about them is stored in the table pets.

Your goal is to return a new table with columns cats, dogs, weasels and chinchillas with names of each pet placed in the corresponding column. Names in each column should be ordered alphabetically. Print NULL when there are no more names corresponding to a species.

The pets table contains the following columns:

name - the unique name of the pet;
species - the species of the pet ("cat", "dog", "weasel" or "chinchilla").
Return the summary table as described above.

Example

For given table pets

name	species
Barker	weasel
Chevy	chinchilla
Fay	chinchilla
Grit	cat
Hex	dog
Lunar	dog
Mojo	chinchilla
Peter	weasel
Shogun	cat
Sierra	cat
the output should be

cats	dogs	weasels	chinchillas
Grit	Hex	Barker	Chevy
Shogun	Lunar	Peter	Fay
Sierra	NULL	NULL	Mojo
The first column is an alphabetically ordered list of cats, second - of dogs, third - of weasels and fourth - of chinchillas. The empty cell data for columns with less than the maximum number of names per species (in this case, dogs and weasels) are filled with NULL values.
*/
CREATE PROCEDURE petSummary()
BEGIN

    SELECT c.name cats, d.name dogs, w.name weasels, h.name chinchillas FROM
    (select @l := @l + 1 AS i from pets, (SELECT @l := 0) r)a
    LEFT JOIN 
    (select @i := @i + 1 AS i, name from pets, (SELECT @i := 0) r WHERE species = 'cat')c
    on a.i = c.i
    LEFT JOIN
    (select @j := @j + 1 AS i, name from pets, (SELECT @j := 0) r WHERE species = 'dog')d
    ON a.i = d.i
    LEFT JOIN
    (select @m := @m + 1 AS i, name from pets, (SELECT @m := 0) r WHERE species = 'weasel')w
    on a.i = w.i
    LEFT JOIN
    (select @n := @n + 1 AS i, name from pets, (SELECT @n := 0) r WHERE species = 'chinchilla')h
    ON a.i = h.i
    WHERE NOT (c.name IS NULL
    AND d.name IS NULL
    AND w.name IS NULL
    AND h.name IS NULL);

END
