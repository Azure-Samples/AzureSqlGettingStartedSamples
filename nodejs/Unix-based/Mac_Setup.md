# Machine Setup for MacOs

This page will walk you through everything you need to setup and run the Node.js samples on your Mac machine.

## Install [Homebrew](https://brew.sh/) by following the instructions on the site

# Restart the terminal session.

## Install Node.js

    ```terminal
      brew install node
    ```

> You now have Node.js installed! The next section will walk you through getting the tools to interact with your database.

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

# Proceed with the tutorial by following the remaining directions [**here**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/nodejs/Unix-based/Readme.md)
