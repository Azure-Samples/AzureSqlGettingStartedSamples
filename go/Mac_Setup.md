# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run GoLang and connect to Azure SQL.

## Install Homebrew and GoLang

If you already have GoLang installed on your machine, skip this step. Use the following commands to install Homebrew. Make sure to restart your terminal session once you're done.

1. Install Homebrew.

```terminal
  ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

1. Restart the terminal session.

1. Update Homebrew and install GoLang.

    ```terminal
    brew update
    brew install go
    ```

1. Set up the GOPATH, GOROOT and GOBIN environment variables and add these to the PATH with the following commands:

    ```terminal
    echo "export GOPATH=$HOME/golang" >> ~/.bash_profile
    echo "export GOROOT=/usr/local/opt/go/libexec" >> ~/.bash_profile
    echo "export GOBIN=$GOPATH/bin" >> ~/.bash_profile
    echo "export PATH=$PATH:$GOPATH" >> ~/.bash_profile
    echo "export PATH=$PATH:$GOROOT/bin" >> ~/.bash_profile
    ```
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

# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/go/Readme.md) to complete the tutorial.
