---
page_type: sample
languages:
- go
products:
- azure-sql-database	
description: "Azure SQL - Getting Started Samples"
---

# Developing applications with Go and Azure SQL 

In this tutorial you will learn how to create a database, use [Go](https://golang.org/) to connect to it, and improve performance with columnstore indexes.

Regardless of platform or operating system, you can use Azure SQL. As you can see the tutorial is available to be used with:

- [Windows](Windows_Setup.md)
- [Red Hat Enterprise Linux](RHEL_Setup.md)
- [Ubuntu](Ubuntu_Setup.md)
- [SLES](SLES_Setup.md)
- [Mac](Mac_Setup.md)

Once you've gotten familiar with Azure SQL and Go, you can take advantage of [several samples](https://learn.microsoft.com/samples/browse/?expanded=dotnet&products=azure-sql-database&languages=go). The samples will help you to create Full-Stack solutions or Back-End APIs that can be used in projects of any size and scale.

## Prerequisites

### Create an Azure SQL database

All of these examples require an Azure SQL database. You can follow these instructions to create one.

For instructions on how to set up an Azure SQL database, see [Azure SQL Database](https://learn.microsoft.com/azure/sql-database/sql-database-single-database-get-started?tabs=azure-portal).

### Set up your machine for Go and Azure SQL

Set up your machine using the instructions for your OS by clicking on the links below, then return here to complete the tutorial.

- [Windows](Windows_Setup.md)
- [Red Hat Enterprise Linux](RHEL_Setup.md)
- [Ubuntu](Ubuntu_Setup.md)
- [SLES](SLES_Setup.md)
- [Mac](Mac_Setup.md)

## Get Started with Go and Azure SQL

### Get connection information to use in connection strings, and create a firewall rule

1. Using the Azure portal, go to your database and look in the panel on the left. Under the settings subcategory, find "connection strings".

1. Click Connection Strings, and then take note of the information:  

   ```results
   Server=tcp:your_server.database.windows.net,1433;Initial Catalog=your_database;Persist Security Info=False;User ID=your_user ;Password=<THIS IS ACTUALLY NOT RETURNED>;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;
   ```

1. Make a note of the following somewhere for reference in subsequent steps:

   ```results
   Server=your_server.database.windows.net
   Database=your_database
   UserId=your_user
   Password=your_password
   ```

#### Create a firewall rule

In order to connect to your Azure SQL database, you will need to create a filrewall rule on the target server.  This allows your application to talk to your Azure SQL Database.

1.  From your database, in the **Overview**, navigate to **Server name**.
1.  From the server, in the search bar at the top, type **firewall**, and select **Firewalls and virtual networks**.
1.  Select **+ Add Client IP**.
1.  Select **Save**. 

## Create a Go app that connects to Azure SQL and executes queries

Create a new folder for your project called **AzureSqlSample** and switch to that folder from the command line.

Next we will create a Go app that connects to Azure SQL database.

Create a file named [**connect.go**](connect.go) in the AzureSqlSample folder. This sample uses the GoLang Context methods to ensure that there's an active connection to the database server. Don't forget to update the username and password with your own.

Prepare and run your Go app from the **AzureSqlSample** folder in the terminal.

```terminal
go mod init azure-demo
go mod tidy

go run connect.go
```

Create a file called CreateTestData.sql in the **AzureSqlSample** folder. Copy and paste the following the T-SQL code inside it. This will create a schema, table, and insert some sample rows.

```sql
CREATE SCHEMA TestSchema;
GO

CREATE TABLE TestSchema.Employees (
  Id INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name NVARCHAR(50),
  Location NVARCHAR(50)
);
GO

INSERT INTO TestSchema.Employees (Name, Location) VALUES
(N'Jared', N'Australia'),
(N'Nikita', N'India'),
(N'Tom', N'Germany');
GO

SELECT * FROM TestSchema.Employees;
GO
```

From the **AzureSqlSample** folder, connect to the database using **sqlcmd** and run the SQL script to create the schema, table, and insert some rows.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -i ./CreateTestData.sql
```

Create a new file called [**crud.go**](crud.go) in the AzureSqlSample folder. Update your connection information. This will insert, update, delete, and read a few rows.

Prepare and run the 'crud.go' app from the **AzureSqlSample folder in the terminal.

```terminal
go run crud.go
```

## Create a Go app that connects to Azure SQL DB using the popular GORM

Create a new folder for your project called **AzureSqlGormSample** and switch to that folder from the command line.

Paste the contents below into a file called [**orm.go**](orm.go). Make sure to replace the connection information.

Prepare and run the 'orm.go' app from the **AzureSQLGormSample** in the terminal.

```terminal
go mod init AzureSqlGormSample
go mod tidy

go run orm.go
```

## Improve performance using columnstore

In this section we will show you an example of [Columnstore Indexes](https://learn.microsoft.com/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore indexes can dramatically improve performance on analytical workloads, and offer better data compression than traditional rowstore indexes.

This example creates a table with 3 million rows and then uses a Go application to query that data, capturing the run time before and after adding a columnstore index.

### Create a new table with 3 million rows using sqlcmd

Create a new folder for your project called **AzureSqlColumnstoreSample** and switch to that folder from the command line.

Create a new file called **CreateSampleTable.sql** in the folder **AzureSqlColumnstoreSample**. Paste the T-SQL code below into your new SQL file. Save and close the file.

```SQL
WITH
    a
    AS
    (
        SELECT *
        FROM (VALUES(1), (2), (3), (4), (5), (6), (7), (8), (9), (10)) AS a(a)
    )
SELECT TOP(3000000)
    ROW_NUMBER() OVER (ORDER BY a.a) AS OrderItemId,
    a.a + b.a + c.a + d.a + e.a + f.a + g.a + h.a AS OrderId,
    a.a * 10 AS Price,
    CONCAT(a.a, N' ', b.a, N' ', c.a, N' ', d.a, N' ', e.a, N' ', f.a, N' ', g.a, N' ', h.a) AS ProductName
INTO Table_with_3M_rows
FROM a, a AS b, a AS c, a AS d, a AS e, a AS f, a AS g, a AS h;
```

From the **AzureSqlColumnstoreSample' folder, connect to the database using sqlcmd and run the SQL script to create the table with 3 million rows. This may take a few minutes to run.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -i ./CreateSampleTable.sql
```

### Create a Go app that queries this tables and measures the time taken

In the **AzureSqlColumnstoreSample** folder, create a file called [**columnstore.go**](columnstore.go).

Prepare and run your Go app from the **AzureSqlColumnstoreSample** folder in the terminal.

```terminal
go mod init AzureSqlColumnstoreSample
go mod tidy

go run columnstore.go
```

### Add a columnstore index to your table using SQLCMD

In the **AzureSqlColumnstoreSample** folder, run this command to create a columnstore index on your table:

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "CREATE CLUSTERED COLUMNSTORE INDEX Columnstoreindex ON Table_with_3M_rows;"
```

## Re-run the columnstore.go script and notice how long the query took to complete this time

Run your Go app from the **AzureSqlColumnstoreSample** folder in the terminal.

```terminal
go run columnstore.go
```
Once you are finished with your code, remember to clean up your folders.
