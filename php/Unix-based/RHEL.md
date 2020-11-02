# Machine Setup

> In this section, you will install the necessary dependencies to create PHP apps with Azure SQL DB.


Now that your machine is set up, you can proceed with the next steps.

# Start Writing apps with PHP and Azure SQL

## Install PHP and other required packages

To install PHP on Red Hat 7, run the following:

> To install PHP 7.2 or 7.3, replace `remi-php74` with `remi-php72` or `remi-php73` respectively in the following commands.

```terminal
    sudo su
    yum install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
    yum install https://rpms.remirepo.net/enterprise/remi-release-7.rpm
    subscription-manager repos --enable=rhel-7-server-optional-rpms
    yum install yum-utils
    yum-config-manager --enable remi-php74
    yum update
    yum install php php-pdo php-xml php-pear php-devel re2c gcc-c++ gcc
```

To install PHP on Red Hat 8, run the following:

> To install PHP 7.2 or 7.3, replace `remi-7.4` with `remi-7.2` or `remi-7.3` respectively in the following commands.

```terminal
    sudo su
    dnf install https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm
    dnf install https://rpms.remirepo.net/enterprise/remi-release-8.rpm
    dnf install yum-utils
    dnf module reset php
    dnf module install php:remi-7.4
    subscription-manager repos --enable codeready-builder-for-rhel-8-x86_64-rpms
    dnf update
    dnf install php-pdo php-pear php-devel
```

> You have successfully installed PHP on your RHEL machine!

> SELinux is installed by default and runs in Enforcing mode. To allow Apache to connect to a database through SELinux, run the following command: 

```terminal
    sudo setsebool -P httpd_can_network_connect_db 1
```


## Install the ODBC Driver and SQL Command Line Utility for Azure SQL DB

[SQLCMD](https://docs.microsoft.com/sql/linux/sql-server-linux-connect-and-query-sqlcmd) is a command line tool that enables you to connect to Azure SQL DB and run queries.

```terminal
sudo su

curl https://packages.microsoft.com/config/rhel/7/prod.repo > /etc/yum.repos.d/msprod.repo

exit

sudo yum remove mssql-tools unixODBC-utf16-devel
sudo yum install mssql-tools unixODBC-devel

# optional: for bcp and sqlcmd
echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bash_profile
echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bashrc
source ~/.bashrc

# optional: for unixODBC development headers
sudo yum install unixODBC-devel
```

After installing SQLCMD, you can connect to Azure SQL DB using the following command:

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database
1> # You're connected! Type your T-SQL statements here. Use the keyword 'GO' to execute each batch of statements.
```

This how to run a basic inline query. The results will be printed to the STDOUT.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "SELECT @@VERSION"
```


> You have successfully installed SQL Server Command Line Utilities on your Red Hat machine! 


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

## Install the PHP Driver for Azure SQL DB

```terminal
sudo pecl install sqlsrv
sudo pecl install pdo_sqlsrv
sudo su
echo extension=pdo_sqlsrv.so >> `php --ini | grep "Scan for additional .ini files" | sed -e "s|.*:\s*||"`/30-pdo_sqlsrv.ini
echo extension=sqlsrv.so >> `php --ini | grep "Scan for additional .ini files" | sed -e "s|.*:\s*||"`/20-sqlsrv.ini
exit
```


## Create a PHP app that connects to Azure SQL DB and executes queries

```terminal
mkdir AzureSqlSample
cd AzureSqlSample
```

Using your favorite text editor, create a new file called connect.php in the AzureSqlSample folder, using the example from [**connect.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/AzureSqlSample/connect.php). Update pertinent connection information, save, and close.


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

Using your favorite text editor, create a new file called crud.php in the SqlServerSample folder, based on the sample [**crud.php**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/AzureSqlSample/crud.php). Update the connection information, save, and close.

Run your PHP script from the terminal.

```terminal
php crud.php
```


> Congratulations! You have created your first PHP app with Azure SQL DB! Check out the next section to learn about how you can make your PHP faster with Azure SQL DB's Columnstore feature.

# Improve Performance Using Columnstore

This [**tutorial**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/php/Unix-based/Columnstore.md) will walk you through the process of experimenting with columnstore to improve application and query performance.



