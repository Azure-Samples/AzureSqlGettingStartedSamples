# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Python and connect to Azure SQL.

##  Install Python and other required packages

```terminal
sudo apt-get install python python-pip gcc g++ build-essential
```

> You now have Python installed! The next section will walk you through getting the tools to interact with your database.

## Install the ODBC Driver and SQL Command Line Utility for SQL Server

[**SQLCMD**](https://docs.microsoft.com/en-us/sql/linux/sql-server-linux-setup-tools?view=sql-server-ver15#ubuntu) is a command line tool that enables you to connect to Azure SQL and run queries.

```terminal

curl https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add -

Download appropriate package for the OS version
#Choose only ONE of the following, corresponding to your OS version

#Ubuntu 16.04
curl https://packages.microsoft.com/config/ubuntu/16.04/prod.list | sudo tee /etc/apt/sources.list.d/msprod.list

#Ubuntu 18.04
curl https://packages.microsoft.com/config/ubuntu/18.04/prod.list | sudo tee /etc/apt/sources.list.d/msprod.list

#Ubuntu 19.10
curl https://packages.microsoft.com/config/ubuntu/19.10/prod.list | sudo tee /etc/apt/sources.list.d/msprod.list


sudo apt-get update 
sudo apt-get install mssql-tools unixodbc-dev


echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bash_profile

echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bashrc
source ~/.bashrc

sudo su
curl https://packages.microsoft.com/keys/microsoft.asc | apt-key add -
```

After installing SQLCMD, you can connect to Azure SQL using the following command:

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database
1> # You're connected! Type your T-SQL statements here. Use the keyword 'GO' to execute each batch of statements.
```

This how to run a basic inline query. The results will be printed to the STDOUT.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "SELECT @@VERSION"
```

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
