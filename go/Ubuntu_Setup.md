# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run GoLang and connect to Azure SQL.

## Step 1.2 Install GoLang

If you already have Go installed on your machine, skip this step. To install GoLang, follow these commands:

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
   
You may also have to [install git](https://git-scm.com/downloads) on your machine, to make future calls to "go get" work.

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


# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/go/Readme.md) to complete the tutorial.
