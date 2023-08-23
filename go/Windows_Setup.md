# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run GoLang and connect to Azure SQL.

## Install GoLang

If you do not already have Go installed on your machine, run the following command:

```terminal
winget install GoLang.Go 
```
If you want to make sure you have the latest version of Go installed on your machine, run the following command:

```terminal
winget update golang.go
```

## Install the SQL Command Line Utility for SQL Server

SQLCMD is a command line tool that enables you to connect to Azure SQL or SQL Server and run queries. Run the following command to install SQLCMD:

```terminal
winget install sqlcmd 
```

After installing SQLCMD, you can connect to Azure SQL using the following command, after making sure to update your connection information:

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_databsae
1> # You're connected! Type your T-SQL statements here. Use the keyword 'GO' to execute each batch of statements.
```

This how to run a basic inline query. The results will be printed to STDOUT.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "SELECT @@VERSION"
``` 

You have successfully installed SQLCMD on your machine and used it to connect to Azure SQL.

Return to the [**main page**](Readme) to complete the tutorial.
