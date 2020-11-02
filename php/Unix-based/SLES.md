# Machine Setup

> In this section, you will install the necessary dependencies to create PHP apps with Azure SQL DB.


Now that your machine is set up, you can proceed with the next steps.

# Start Writing apps with PHP and Azure SQL

## Install PHP and other required packages

> In the following instructions, replace `<SuseVersion>` with your version of Suse - if you are using Suse Enterprise Linux 15, it will be SLE_15 or SLE_15_SP1. For Suse 12, use SLE_12_SP4 (or above if applicable). Not all versions of PHP are available for all versions of Suse Linux - please refer to `http://download.opensuse.org/repositories/devel:/languages:/php` to see which versions of Suse have the default version PHP available, or to `http://download.opensuse.org/repositories/devel:/languages:/php:/` to see which other versions of PHP are available for which versions of Suse.

> Packages for PHP 7.4 are not available for Suse 12. To install PHP 7.2, replace the repository URL below with the following URL: `https://download.opensuse.org/repositories/devel:/languages:/php:/php72/<SuseVersion>/devel:languages:php:php72.repo`. To install PHP 7.3, replace the repository URL below with the following URL: `https://download.opensuse.org/repositories/devel:/languages:/php:/php73/<SuseVersion>/devel:languages:php:php73.repo`.

```terminal
sudo zypper ar http://download.opensuse.org/repositories/devel:/languages:/php/openSUSE_Leap_15.1/ php
sudo zypper mr -p 70 php
sudo zypper refresh
sudo zypper in php7 php php7-devel
```

> You have successfully installed PHP on your SLES machine! 

## Install the ODBC Driver and SQL Command Line Utility for Azure SQL DB
[SQLCMD](https://docs.microsoft.com/sql/linux/sql-server-linux-connect-and-query-sqlcmd) is a command line tool that enables you to connect to Azure SQL DB and run queries.

```terminal
sudo su

#Download appropriate package for the OS version
#Choose only ONE of the following, corresponding to your OS version

#SUSE Linux Enterprise Server 11 SP4
#Ensure SUSE Linux Enterprise 11 Security Module has been installed 
zypper ar https://packages.microsoft.com/config/sles/11/prod.repo

#SUSE Linux Enterprise Server 12
zypper ar https://packages.microsoft.com/config/sles/12/prod.repo

#SUSE Linux Enterprise Server 15
zypper ar https://packages.microsoft.com/config/sles/15/prod.repo
#(Only for driver 17.3 and below)
SUSEConnect -p sle-module-legacy/15/x86_64

exit
sudo ACCEPT_EULA=Y zypper install msodbcsql17
# optional: for bcp and sqlcmd
sudo ACCEPT_EULA=Y zypper install mssql-tools
echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bash_profile
echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bashrc
source ~/.bashrc
# optional: for unixODBC development headers
sudo zypper install unixODBC-devel
```

After installing SQLCMD, you can connect to Azure SQL DB using the following command:

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database
1> # You're connected! Type your T-SQL statements here. Use the keyword 'GO' to execute each batch of statements.
```

This how to run a basic inline query. The results will be printed to the STDOUT.

```terminal
sqlcmd -S your_database.database.windows.net -U your_user -P your_password -d your_database -Q "SELECT @@VERSION"
```

> You have successfully installed SQL Server Command Line Utilities on your SLES machine! 

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

## Step 2.2 Install the PHP Driver for SQL Server

> If you get an error message saying `Connection to 'pecl.php.net:443' failed: Unable to find the socket transport "ssl"`, edit the pecl script at /usr/bin/pecl and remove the `-n` switch in the last line. This switch prevents PECL from loading ini files when PHP is called, which prevents the OpenSSL extension from loading.

```terminal
sudo pecl install sqlsrv
sudo pecl install pdo_sqlsrv
sudo su
echo extension=pdo_sqlsrv.so >> `php --ini | grep "Scan for additional .ini files" | sed -e "s|.*:\s*||"`/pdo_sqlsrv.ini
echo extension=sqlsrv.so >> `php --ini | grep "Scan for additional .ini files" | sed -e "s|.*:\s*||"`/sqlsrv.ini
exit
```

## Create a PHP app that connects to Azure SQL DB and executes queries

```terminal
mkdir AzureSqlSample
cd AzureSqlSample
```

Using your favorite text editor, create a new file called connect.php in the AzureSqlSample folder, based on the sample [**connect.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/AzureSqlSample/connect.php), update connection information, save, and close.

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

Using your favorite text editor, create a new file called crud.php in the AzureSqlSample folder, based on the sample [**crud.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/AzureSqlSample/crud.php), update the connection information, save, and close the file.


Run your PHP script from the terminal.

```terminal
php crud.php
```


> Congratulations! You have created your first PHP app with Azure SQL! Check out the next section to learn about how you can make your PHP faster with Azure SQL DB's Columnstore feature.

# Improve Performance Using Columnstore

This [**tutorial**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/Columnstore.md) will walk you through the process of experimenting with columnstore to improve application and query performance.


