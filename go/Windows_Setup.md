# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run GoLang and connect to Azure SQL.

## Install GoLang

If you do not already have Go installed on your machine, run the following command:

```terminal
winget install GoLang.Go 
```
If you want to make sure you have the latest version of Go installed on your machine, run the following command:

```terminal
winget update GoLang.Go
```

## Install the Az Cli

The Azure Command-Line Interface (CLI) is a cross-platform command-line tool that can be installed locally on Windows computers. You can use the Azure CLI for Windows to connect to Azure and execute administrative commands on Azure resources. The Azure CLI for Windows can also be used from a browser through the Azure Cloud Shell or run from inside a Docker container.

```terminal
winget install -e --id Microsoft.AzureCLI
```

## Install the SQL Command Line Utility for SQL Server

SQLCMD is a command line tool that enables you to connect to Azure SQL or SQL Server and run queries. Run the following command to install SQLCMD:

```terminal
winget install sqlcmd 
```

After installing SQLCMD, close and re-open your terminal window. You can now connect to Azure SQL using the following command (after making sure to update your connection information). To simplify the process, we will use az login to authenticate to Azure AD and SQLCMD can re-use that connection. It only needs to be run once per terminal window but is included in each step below in case someone is starting from a fresh terminal window. Feel free to skip running it for later tasks.

```terminal
az login
sqlcmd -S your_server.database.windows.net -d your_databsae
1> # You're connected! Type your T-SQL statements here. Use the keyword 'GO' to execute each batch of statements.
```

This how to run a basic inline query. The results will be printed to STDOUT.

```terminal
az login
sqlcmd -S your_server.database.windows.net -d your_database -Q "SELECT @@VERSION"
``` 

You have successfully installed SQLCMD on your machine and used it to connect to Azure SQL.


Return to the [**main page**](Readme.md) to complete the tutorial.
