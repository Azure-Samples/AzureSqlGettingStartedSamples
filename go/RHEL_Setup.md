# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run GoLang and connect to Azure SQL.

## Install GoLang

If you already have GoLang installed on your machine, skip this step. To install GoLang, follow these commands:

1. Run the following commands:

    ```terminal
    curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
    tar xvf go1.8.linux-amd64.tar.gz
    sudo chown -R root:root ./go
    sudo mv go /usr/local
    ```

1. Using your favorite text editor, add these two lines to the ~/.profile file.

    ```terminal
    export GOPATH=$HOME/work
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
    ```

Then reload your profile, and confirm that go is on the path:

    ```terminal
    source .profile
    which go
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

# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/go/Readme.md) to complete the tutorial.
