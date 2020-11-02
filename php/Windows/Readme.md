Follow all of these steps to start writing PHP applications using Azure SQL.

# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run PHP and connect to Azure SQL.

## Install PHP and Chocolatey

For full details of installing PHP you can reference this [**site.**](https://www.php.net/manual/en/install.windows.php)  

One of the prerequisiites to PHP is having Visual studio installed on your machine.  Installing the latest community edition of Visual Studio from [**here**](https://visualstudio.microsoft.com/downloads/) will be sufficient.

Once you have Visual Studio installed, you will download PHP using the [**Web Platform Installer**](https://www.microsoft.com/web/downloads/platform.aspx). Once you download Web PI, open it up and download the entry which says **'PHP 7.4.1 (x64) for IIS Express'**.

Next, install Chocolatey. Chocolatey is a package manager like apt-get and yum for Windows. We will use Chocolatey later in the tutorial. Use an elevated Command-line session (run as administrator):

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

You can check the version this way:

```terminal
choco -?
```

For Chocolatey to work, you now need to restart the terminal session by closing and opening the terminal.

> You have succesfully installed PHP and Chocolatey on your machine!

## Install the ODBC Driver and SQL Command Line Utility for SQL Server

SQLCMD is a command line tool that enables you to connect to Azure SQL or SQL Server and run queries.

1. Install the [**ODBC Driver**](https://docs.microsoft.com/sql/connect/odbc/download-odbc-driver-for-sql-server).
2. Install the [**SQL Server Command Line Utilities**](https://docs.microsoft.com/sql/tools/sqlcmd-utility).

After installing SQLCMD, you can connect to Azure SQL using the following command from a CMD session, making sure to update your connection information:

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_databsae
1> # You're connected! Type your T-SQL statements here. Use the keyword 'GO' to execute each batch of statements.
```

This how to run a basic inline query. The results will be printed to STDOUT.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "SELECT @@VERSION"
``` 

> You have successfully installed SQL Server Command Line Utilities on your Windows machine, and used them to connect to Azure SQL! 

# Create and Run your first PHP application with Azure SQL

## Get Connection Information to use in Connection Strings, and Create a Firewall Rule.

### Get the connection string info from the Azure Portal

1. Using the Azure Portal, go to your database and look in the panel on the left.  It should say Overview, activity log, …

Under the settings subcategory, find "connection strings"

2. Click Connection Strings, and then take note of the information:  

 ```results
Server=tcp:your_server.database.windows.net,1433;Initial Catalog=your_database;Persist Security Info=False;User ID=your_user ;Password=<THIS IS ACTUALLY NOT RETURNED>;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;
 ```

3. Make a note of the following somewhere for reference in subsequent steps:

 ```results
Server=your_server.database.windows.net

Database=your_database

UserId=your_user

Password=your_password
```

### Create a firewall rule

In order to connect to your Azure SQL database, you will need to create a filrewall rule on the target server.  This allows your application to talk to your Azure SQL Database.

1.  From your database, in the Overview, you can look in the panel on the right, and where it says Server name, click there.
1.  From the server, in the search bar at the top, type "firewall", and select "Firewalls and virtual networks".
1.  In the bar at the top, you sould see three shortcut options: Save, Discard, and + Add Client IP.  Click on + Add Client IP.
1.  Next click Save.  You can close this view now.

## Install the PHP Drivers for Azure SQL DB

Download the Microsoft PHP Drivers for Azure SQL DB from the [download page](https://docs.microsoft.com/sql/connect/php/download-drivers-php-sql-server).

Pick the appropriate dll - for example **php_pdo_sqlsrv_74_nts_x64.dll** for the **PDO Driver** and **php_sqlsrv_74_nts_x64.dll** for the **SQLSRV driver**.

Copy the dll files to the **C:\Program Files\iis express\PHP\v7.4\ext** folder.

Register the dll files in the **php.ini** file.

```terminal
    cd C:\Program^ Files\iis^ express\PHP\v7.4\ext
    echo extension=php_sqlsrv_74_nts_x64.dll >> C:\Program^ Files\iis^ express\PHP\v7.4\php.ini
    echo extension=php_pdo_sqlsrv_74_nts_64.dll >> C:\Program^ Files\iis^ express\PHP\v7.4\php.ini
```

# Create a PHP app that connects to Azure SQL DB and executes queries

```terminal
mkdir AzureSqlSample
cd AzureSqlSample
```

Using your favorite text editor, create a new file called connect.php in the AzureSqlSample folder, using the sample contents from [**connect.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Windows/AzureSqlSample/connect.php).  Update the connection string information, save, and close the file.


Run your PHP script from the terminal.

```terminal
php connect.php
```

Execute the T-SQL scripts below in the terminal with sqlcmd to create a schema, table, and insert a few rows.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "CREATE SCHEMA TestSchema;"
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "CREATE TABLE TestSchema.Employees (Id INT IDENTITY(1,1) NOT NULL PRIMARY KEY, Name NVARCHAR(50), Location NVARCHAR(50));"
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "INSERT INTO TestSchema.Employees (Name, Location) VALUES (N'Jared', N'Australia'), (N'Nikita', N'India'), (N'Tom', N'Germany');"
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "SELECT * FROM TestSchema.Employees;"
```

Using your favorite text editor, create a new file called crud.php in the AzureSqlSample folder, using the contents from [**crud.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Windows/AzureSqlSample/crud.php). Update connection information as needed. This will insert, update, delete, and read a few rows. 



Run your PHP script from the terminal.

```terminal
php crud.php
```

> Congratulations! You have created your first PHP app with Azure SQL DB! Check out the next section to learn about how you can make your PHP faster with Azure SQL DB's Columnstore feature.
 
 # Improve Performance using Columnstore
 
 > In this section we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.

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

## Create a PHP app that queries this table and measures the time taken

```terminal
cd C:/Users/User
mkdir AzureSqlColumnstoreSample
cd AzureSqlColumnstoreSample
```

Using your favorite text editor, create a new file called columnstore.php in the AzureSqlColumnstoreSample folder from the example [**columnstore.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Windows/AzureSqlColumnstoreSample/columnstore.php). Update connection information, save, and close.


 ## Measure how long it takes to run the query

Run your PHP script from the terminal.

```terminal
php columnstore.php
```

```results
Sum: 30000000
QueryTime: 1403.2 ms
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




