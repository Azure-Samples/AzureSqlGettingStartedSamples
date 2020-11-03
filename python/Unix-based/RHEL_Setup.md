# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Python and connect to Azure SQL.

## Install Python and other required packages

```terminal
wget https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sudo rpm -ivh epel-release-latest-7.noarch.rpm
sudo yum install python python-pip python-wheel python-devel
sudo yum group install "Development tools"
```

> You now have Python installed! The next section will walk you through getting the tools to interact with your database.

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

## Install The Azure CLI and Login to Azure

1.  Go to this **[site](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)**, and click on the operating system you are currently using.
1.  Install the Azure CLI following the directions.
1.  Open a new instance of the command window and type az login. Follow the instructions that follow to authenticate.

```terminal
az login
```

**Please Note** You may have to login again to your machine after restart or long periods of inactivity.

At this time, you have authenticated yourself and your machine to Azure, so your application can connect. 

> You have successfully installed the Python Driver on your Ubuntu machine. You now have everything you need to start writing Python apps with Azure SQL DB!

# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/python/Unix-based#start-writing-apps-with-python-and-azure-sql) to complete the tutorial.
