---
page_type: sample
languages:
- go
products:
- azure-sql-database	
description: "Azure SQL - Getting Started Samples"
---

# Developing applications with Go and Azure SQL 

This repository contains a tutorial that will guide you through the creation of a simple solution using [GO](https://golang.org/) to take advantage of Azure SQL. Azure SQL as many features for developers and learning how to take advantage of it will help you to create secure, scalable and performant modern applications. To learn more about several of the features that Azure SQL provides to developers, read here: [10 reasons to use Azure SQL in your next project](https://devblogs.microsoft.com/azure-sql/10-reasons-to-use-azure-sql-in-your-next-project/).

With the proposed tutorial you will learn how to create a database, use the most common packages to connect to it and, finally, you'll see how performance can be improved *a lot* by using Columnstore Indexes.

No matter which is the platform or the OS you are using, you can happily use Azure SQL. As you can see the tutorial is available to be used with:

- [Windows](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/Windows_Setup.md)
- [Red Hat Enterprise Linux](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/RHEL_Setup.md)
- [Ubuntu](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/Ubuntu_Setup.md)
- [SLES](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/SLES_Setup.md)
- [Mac](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/Mac_Setup.md)

Once you'll be more expert on Azure SQL and Go, you can also take advtange of [several samples](https://docs.microsoft.com/en-us/samples/browse/?expanded=dotnet&products=azure-sql-database&languages=nodejs) that will help you to create Full-Stack solutions or Back-End API, that can be used in project of any size and scale.

# Prerequisites

## Create an Azure SQL DB
All of the above examples require an Azure SQL DB  Please follow these instructions to create one.

Go to this [**site**](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-single-database-get-started?tabs=azure-portal) for instructions on how to set up an Azure Hosted SQL Database.

1.  Perform the Prerequisites steps.

2. Follow steps 1-17 from the section: **Create a Single Database**.

## Set up your machine for Go and Azure SQL

Set up your machine using the instructions for your OS by clicking on the links below, then return here to complete the tutorial.

- [Windows](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/Windows_Setup.md)
- [Red Hat Enterprise Linux](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/RHEL_Setup.md)
- [Ubuntu](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/Ubuntu_Setup.md)
- [SLES](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/SLES_Setup.md)
- [Mac](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/go/Mac_Setup.md)


# Get Started with Go and Azure SQL

## Get Connection Information to use in Connection Strings, and Create a Firewall Rule.

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


## Create a Go app that connects to Azure SQL and executes queries

Create a new project directory and install Go dependencies.

```terminal
    cd ~/

    #Create Project Directory
    mkdir AzureSqlSample
    cd AzureSqlSample

    # Get and install the Azure SQL DB driver for Go
    go get github.com/denisenkom/go-mssqldb
    go install github.com/denisenkom/go-mssqldb
```

Now you will create a simple Go app that connects to Azure SQL DB.

Using [your favorite text editor](https://code.visualstudio.com/), create a file named [**connect.go**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/go/connect.go) in the AzureSqlSample folder.   This sample uses the GoLang Context methods to ensure that there's an active connection to the database server. Don't forget to update the username and password with your own.


Run the application.

```terminal
go run connect.go
```

Using your favorite text editor, create a file called CreateTestData.sql in the AzureSqlSample folder. Copy and paste the following the T-SQL code inside it. This will create a schema, table, and insert a few rows.

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

Connect to the database using sqlcmd and run the SQL script to create the schema, table, and insert some rows.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -i ./CreateTestData.sql
```

Using your favorite text editor, create a new file called [**crud.go**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/go/crud.go) in the AzureSqlSample folder. Update your connection information. This will insert, update, delete, and read a few rows.


Run the crud.go app to see the results

```terminal
go run crud.go
```

## Create a Go app that connects to Azure SQL DB using the popular GORM

Create the app directory and initialize Go dependencies.

```terminal
    cd ~/
    mkdir AzureSqlGormSample
    cd AzureSqlGormSample

    # Get and install the SQL Server driver for Go
    go get github.com/denisenkom/go-mssqldb
    go install github.com/denisenkom/go-mssqldb

   # Get and install GORM
   go get github.com/jinzhu/gorm
   go install github.com/jinzhu/gorm
```

Paste the contents below into a file called [**orm.go**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/go/orm.go). Make sure to replace the connection information.


Run the orm.go app

```terminal
    go run orm.go
```


> Congratulations! You created your first three Go apps with AzureSQL DB! Check out the next section to learn about how you can make your apps faster with Azure SQL DB’s Columnstore feature.


# Improve Performance Using Columnstore

> Now that you have explored the basics, you are ready to see how you can make your app better with Azure SQL. In this module we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.


Note! May features available in SQL Server are also availble in Azure SQL DB.  Please reference **[this link](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-features)** for a full breakdown of which features are available in Azure DB.
Fortunately for us, supported features include innovations that can significantly improve your application’s throughput, latency, and security. Enjoy!

## Create a C# Console Application to explore Columnstore performance enhancements

To showcase the capabilities of Columnstore indexes, let's create a C# application that creates a sample database and a sample table with 3 million rows and then runs a simple query before and after adding a Columnstore index.

## Create a new table with 3 million rows using sqlcmd

Change to your home directory and create a folder for your project.

```terminal
cd ~/
mkdir AzureSqlColumnstoreSample
cd AzureSqlColumnstoreSample
```

Using your favorite text editor, create a new file called CreateSampleTable.sql in the folder AzureSqlColumnstoreSample. Paste the T-SQL code below into your new SQL file. Save and close the file.

```SQL
WITH a AS (SELECT * FROM (VALUES(1),(2),(3),(4),(5),(6),(7),(8),(9),(10)) AS a(a))
SELECT TOP(3000000)
ROW_NUMBER() OVER (ORDER BY a.a) AS OrderItemId
,a.a + b.a + c.a + d.a + e.a + f.a + g.a + h.a AS OrderId
,a.a * 10 AS Price
,CONCAT(a.a, N' ', b.a, N' ', c.a, N' ', d.a, N' ', e.a, N' ', f.a, N' ', g.a, N' ', h.a) AS ProductName
INTO Table_with_3M_rows
FROM a, a AS b, a AS c, a AS d, a AS e, a AS f, a AS g, a AS h;
```

Connect to the database using sqlcmd and run the SQL script to create the table with 3 million rows. This may take a few minutes to run.

```terminal
  sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -i ./CreateSampleTable.sql
```

## Create a Go app that queries this tables and measures the time taken

In your project folder, initialize Go dependencies.

```terminal
    go get github.com/denisenkom/go-mssqldb
    go install github.com/denisenkom/go-mssqldb
```

Using you favorite text editor, create a file called [**columnstore.go**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/go/columnstore.go) in your folder.

## Measure how long it takes to run the query

Run your Go app from the terminal.

```terminal
  go run columnstore.go
```

## Add a columnstore index to your table using SQLCMD.

Run this command to create a Columnstore Index on your table:

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "CREATE CLUSTERED COLUMNSTORE INDEX Columnstoreindex ON Table_with_3M_rows;"
```

## Re-run the columnstore.go script and notice how long the query took to complete this time.

```terminal
  go run columnstore.go
```
> Congratulations! You just made your Go app faster using Columnstore Indexes!
