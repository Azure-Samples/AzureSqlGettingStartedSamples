# Machine Setup for Red Hat

This page will walk you through everything you need to setup and run the Node.js samples on your RHEL machine.

## Install Node.js

Add [Extra Packages for Enterprise Linux (EPEL)](https://fedoraproject.org/wiki/EPEL) to your list of repos and update. If you already have Node.js installed on your machine, skip this step.

```terminal
wget https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sudo rpm -ivh epel-release-latest-7.noarch.rpm
sudo yum update
```

Install Node.js by first adding the necessary node repositories, and then installing the nodejs package.

```terminal
sudo yum update
sudo yum -y install nodejs
```

You can verify your Node.js installation using this command:

```terminal
node -v
```


> You now have Node.js installed! The next section will walk you through getting the tools to interact with your database.

## Install the ODBC Driver and SQL Command Line Utility for Azure SQL DB

[SQLCMD](https://docs.microsoft.com/sql/linux/sql-server-linux-connect-and-query-sqlcmd){:target="_blank"} is a command line tool that enables you to connect to Azure SQL DB and run queries.

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
