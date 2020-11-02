# Columnstore

> Now that you have explored the basics, you are ready to see how you can make your app better with Azure SQL. In this module we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.

## Create a new table with 3 million rows using sqlcmd

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -t 60000 -Q "WITH a AS (SELECT * FROM (VALUES(1),(2),(3),(4),(5),(6),(7),(8),(9),(10)) AS a(a))
SELECT TOP(3000000)
ROW_NUMBER() OVER (ORDER BY a.a) AS OrderItemId
,a.a + b.a + c.a + d.a + e.a + f.a + g.a + h.a AS OrderId
,a.a * 10 AS Price
,CONCAT(a.a, N' ', b.a, N' ', c.a, N' ', d.a, N' ', e.a, N' ', f.a, N' ', g.a, N' ', h.a) AS ProductName
INTO Table_with_3M_rows
FROM a, a AS b, a AS c, a AS d, a AS e, a AS f, a AS g, a AS h;"
```

## Create a PHP app that queries this tables and measures the time taken

```terminal
cd ~/
mkdir AzureSqlColumnstoreSample
cd AzureSqlColumnstoreSample
```

Using your favorite text editor, create a new file called columnstore.php in the AzureSqlColumnstoreSample folder, using the example [**columnstore.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/AzureSqlColumnstoreSample/columnstore.php). Update any connection information, save, and close the file.


## Measure how long it takes to run the query

Run your PHP script from the terminal.

```terminal
php columnstore.php
```

## Add a columnstore index to your table

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "CREATE CLUSTERED COLUMNSTORE INDEX Columnstoreindex ON Table_with_3M_rows;"
```

## Measure how long it takes to run the query with a columnstore index

```terminal
php columnstore.php
```
> Congratulations! You just made your PHP app faster using Columnstore Indexes!
