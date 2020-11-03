# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Python and connect to Azure SQL.

## Install Homebrew and Python

1. Install Homebrew.

```terminal
  ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

2. Restart the terminal session.

3. Install Python

    ```terminal
    brew install python
    ```

    ```results
    ==> Downloading https://homebrew.bintray.com/bottles/python-2.7.12.el_capitan.bottle.tar.gz

    ...

    ==> Caveats
    Pip and setuptools have been installed. To update them
      pip install --upgrade pip setuptools

    You can install Python packages with
      pip install

    ==> Summary
    🍺  /usr/local/Cellar/python/2.7.12: 3,476 files, 46.7M
    ```

> You now have Python installed! The next section will walk you through getting the tools to interact with your database.

## Install the ODBC Driver and SQL Command Line Utility for Azure SQL DB

[SQLCMD](https://docs.microsoft.com/sql/linux/sql-server-linux-setup-tools){:target="_blank"} is a command line utility that enables you to connect to SQL Server and run queries.

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

## Step 1.4 Install The Azure CLI and Login to Azure


1.  Go to this **[site](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)**, and click on the operating system you are currently using.
1.  Install the Azure CLI following the directions.
1.  Open a new instance of the command window and type az login. Follow the instructions that follow to authenticate.

```terminal
az login
```

**Please Note** You may have to login again to your machine after restart or long periods of inactivity.

At this time, you have authenticated yourself and your machine to Azure, so your application can connect. 

> You have successfully installed the Python Driver on your Mac. You now have everything you need to start writing your Python apps with Azure SQL DB!

# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/python/Unix-based#start-writing-apps-with-python-and-azure-sql) to complete the tutorial.
