# Machine Setup

> In this section, you will install the necessary dependencies to create PHP apps with Azure SQL DB.


Now that your machine is set up, you can proceed with the next steps.

# Start Writing apps with PHP and Azure SQL

## Install [Homebrew](https://brew.sh/), PHP, and other required packages

```terminal
  ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

2. Restart the terminal session.

3. Install PHP.

> To install PHP 7.2 or 7.3, replace `php@7.4` with `php@7.2` or `php@7.3` respectively in the following commands.

```terminal
    brew tap 
    brew tap homebrew/core
    brew install php@7.4
```

PHP should now be in your path -- run `php -v` to verify that you are running the correct version of PHP. If PHP is not in your path or it is not the correct version, run the following:

```terminal
    brew link --force --overwrite php@7.4
```

4. Install other required packages.

```terminal
    brew install autoconf automake libtool
```

> You have successfully installed PHP on your macOS!

## Install the ODBC Driver and SQL Command Line Utility for Azure SQL DB

[SQLCMD](https://docs.microsoft.com/sql/linux/sql-server-linux-setup-tools) is a command line utility that enables you to connect to SQL Server and run queries.

```terminal
brew tap microsoft/mssql-release https://github.com/Microsoft/homebrew-mssql-release
brew update
HOMEBREW_NO_ENV_FILTERING=1 ACCEPT_EULA=Y brew install msodbcsql17 mssql-tools
```

After installing SQLCMD, you can connect to Azure SQL DB using the following command:

```terminal
sqlcmd -S your_database.database.windows.net -U your_user -P your_password -d your_database
1> # You're connected! Type your T-SQL statements here. Use the keyword 'GO' to execute each batch of statements.
```

This how to run a basic inline query. The results will be printed to the STDOUT.

```terminal
sqlcmd -S your_database.database.windows.net -U your_user -P your_password -d your_database -Q "SELECT @@VERSION"
```

> You have successfully installed SQL Server Command Line Utilities on your macOS machine!

# Create and Run your first PHP apps

Now we will starting writing PHP applications that can connect to Azure SQL.

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

## Install the PHP Drivers for Azure SQL DB

```terminal
    sudo pecl install pdo_sqlsrv
    sudo pecl install sqlsrv
```

## Create a PHP app that connects to Azure SQL DB and executes queries

```terminal
mkdir AzureSqlSample
cd AzureSqlSample
```

Using your favorite text editor, create a new file called connect.php in the AzureSqlSample folder, based on the sample [**connect.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/AzureSqlSample/connect.php). Update your connection information, then save and close the file.

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

Using your favorite text editor, create a new file called crud.php in the AzireSqlSample folder, based on the template from [**crud.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/AzureSqlSample/crud.php), taking care to correct the connection information.  
This will insert, update, delete, and read a few rows.

Run your PHP script from the terminal.

```terminal
php crud.php
```

> Congratulations! You have created your first PHP app with Azure SQL DB! Check out the next section to learn about how you can make your PHP faster with Azure SQL DB's Columnstore feature.

# Improve Performance Using Columnstore

This [**tutorial**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/Columnstore.md) will walk you through the process of experimenting with columnstore to improve application and query performance.
