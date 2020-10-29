# Machine Setup for Ubuntu

This page will walk you through everything you need to setup and run the Node.js samples on your Ubuntu machine.

## Install Node.js

Search for the version of node you want to use.

```terminal
sudo zypper refresh
sudo zypper search nodejs
```

Then pick the one you want from the list and install it:

```terminal   
    sudo zypper install nodejs8
```

You can verify your Node.js installation using this command:

```terminal
  node -v
```

> You now have Node.js installed! The next section will walk you through getting the tools to interact with your database.

## Install the ODBC Driver and SQL Command Line Utility for Azure SQL DB
[SQLCMD](https://docs.microsoft.com/sql/linux/sql-server-linux-connect-and-query-sqlcmd) is a command line tool that enables you to connect to Azure SQL DB and run queries.

```terminal
sudo su

#Download appropriate package for the OS version
#Choose only ONE of the following, corresponding to your OS version

#SUSE Linux Enterprise Server 11 SP4
#Ensure SUSE Linux Enterprise 11 Security Module has been installed 
zypper ar https://packages.microsoft.com/config/sles/11/prod.repo

#SUSE Linux Enterprise Server 12
zypper ar https://packages.microsoft.com/config/sles/12/prod.repo

#SUSE Linux Enterprise Server 15
zypper ar https://packages.microsoft.com/config/sles/15/prod.repo
#(Only for driver 17.3 and below)
SUSEConnect -p sle-module-legacy/15/x86_64

exit
sudo ACCEPT_EULA=Y zypper install msodbcsql17
# optional: for bcp and sqlcmd
sudo ACCEPT_EULA=Y zypper install mssql-tools
echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bash_profile
echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bashrc
source ~/.bashrc
# optional: for unixODBC development headers
sudo zypper install unixODBC-devel
```

After installing SQLCMD, you can connect to Azure SQL DB using the following command:

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database
1> # You're connected! Type your T-SQL statements here. Use the keyword 'GO' to execute each batch of statements.
```

This how to run a basic inline query. The results will be printed to the STDOUT.

```terminal
sqlcmd -S your_database.database.windows.net -U your_user -P your_password -d your_database -Q "SELECT @@VERSION"
```

> You have successfully installed SQL Server Command Line Utilities on your SLES machine! 

