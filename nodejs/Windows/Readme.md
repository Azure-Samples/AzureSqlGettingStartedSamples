This page will walk you through everything you need to setup and run the above Node.js samples on your Windows machine.

# Machine Setup

If you already have Node.js and choco installed on your machine, skip this step. Install Chocolatey using this command in an elevated Command prompt (run as administrator).

```terminal
@powershell -NoProfile -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
```

For choco to work, you now need to restart the terminal session by closing and opening the command prompt. Open an elevated Command prompt (run as administrator) and run the following commands to install node:

```terminal
choco install -y nodejs
```

 > You have succesfuly installed Node.js and Chocolatey on your machine!

## Install the ODBC Driver and SQL Command Line Utility for SQL Server

SQLCMD is a command line tool that enables you to connect to Azure SQL or SQL Server and run queries.

1. Install the [**ODBC Driver**](https://docs.microsoft.com/sql/connect/odbc/download-odbc-driver-for-sql-server).
2. Install the [**SQL Server Command Line Utilities**](https://docs.microsoft.com/sql/tools/sqlcmd-utility).

After installing SQLCMD, you can connect to Azure SQL using the following command from a CMD session, making sure to update your connection information:

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_databsae
1> # You're connected! Type your T-SQL statements here. Use the keyword 'GO' to execute each batch of statements.
```

This how to run a basic inline query. The results will be printed to STDOUT.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "SELECT @@VERSION"
``` 

> You have successfully installed SQL Server Command Line Utilities on your Windows machine, and used them to connect to Azure SQL! 


## Install The Azure CLI and Login to Azure


1.  Go to this **[site](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)**, and click on the operating system you are currently using.
1.  Install the Azure CLI following the directions.
1.  Open a new instance of the command window and type az login. Follow the instructions that follow to authenticate.

```terminal
az login
```

**Please Note** You may have to login again to your machine after restart or long periods of inactivity.

At this time, you have authenticated yourself and your machine to Azure, so your application can connect.

> You have now authenticated your machine to Azure.  

# Create and Run your first Node.js application with Azure SQL

## Get Connection Information to use in Connection Strings, and Create a Firewall Rule.

1. Get the connection string info from the Azure Portal

Go to your database and look in the panel on the left.  It should say Overview, activity log, …

Under the settings subcategory, find "connection strings"

2. Click Connection Strings, and then take note of the information:  

 ```results
Server=tcp:your_server.database.windows.net,1433;Initial Catalog=your_database;Persist Security Info=False;User ID=your_user ;Password=<THIS IS ACTUALLY NOT RETURNED>;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;
 ```

3. Make a note of the following somewhere for reference in subsequent steps:

 ```results
Server=your_server.database.windows.net

Database=your_database

UserId=your_user

Password=your_password
```

**Create a firewall rule**

In order to connect to your Azure SQL database, you will need to create a filrewall rule on the target server.  This allows your application to talk to your Azure SQL Database.

1.  From your database, in the Overview, you can look in the panel on the right, and where it says Server name, click there.
1.  From the server, in the search bar at the top, type "firewall", and select "Firewalls and virtual networks".
1.  In the bar at the top, you sould see three shortcut options: Save, Discard, and + Add Client IP.  Click on + Add Client IP.
1.  Next click Save.  You can close this view now.

## Create a Node.js app that connects to Azure SQL and executes queries

Create a new project directory and initialize Node dependencies.

```terminal
    cd ~/
    mkdir AzureSqlSample
    cd AzureSqlSample
    npm init -y
    #Install tedious and async module in your project folder
    npm install tedious
    npm install async
```


Now you will create a simple Node.js app that connects to Azure SQL.

Using your favorite editor, create a file named connect.js in the AzureSqlSample folder. Copy the contents of [**connect.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/main/nodejs/Windows/AzureSqlSample/connect.js) into the file.  Be sure to update the connection string information and save.

Run the application.

```terminal
  node connect.js
```

You have succesfully run your node.js application and connected to Azure SQL!

## Create and manipulate some data

Using your favorite text editor, create a file called CreateTestData.sql in the AzureSqlSample folder. Copy and paste the following the T-SQL code inside it. This will create a schema, table, and insert a few rows.

```sql
CREATE SCHEMA TestSchema;
GO

CREATE TABLE TestSchema.Employees (
  Id INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name NVARCHAR(50),
  Location NVARCHAR(50)
);
GO

INSERT INTO TestSchema.Employees (Name, Location) VALUES
(N'Jared', N'Australia'),
(N'Nikita', N'India'),
(N'Tom', N'Germany');
GO

SELECT * FROM TestSchema.Employees;
GO
```

Connect to the database using sqlcmd and run the SQL script to create the schema, table, and insert some rows.

```terminal
  sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -i ./CreateTestData.sql
```

Using your favorite text editor, create a new file called [**crud.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/main/nodejs/Windows/AzureSqlSample/crud.js)  in the AzureSqlSample folder. Copy and paste the following code inside it. Using your favorite text editor, create a new file called crud.js in the AzureSqlSample folder. Copy and paste the following code inside it. This will insert, update, delete, and read a few rows. This will insert, update, delete, and read a few rows.

Run the crud.js app to see the results

```terminal
  node crud.js
```

## Secure your credentials using Azure Key Vault

This section takes you through the steps described [**on this site**](https://docs.microsoft.com/en-us/azure/key-vault/secrets/quick-create-node) to set up your machine for authentication to the key vault.  You need to do this to use the **DefaultAzureCredentialBuilder()**.

1. Open a command window and execute **az login** if you have not already.
1. Create a service prinicpal (make sure you take note of the output, as you will use it in the next two steps.):

```terminal
az ad sp create-for-rbac -n "http://mySP" --sdk-auth
```

1. Give the serpvice prinicpal access to your key vault.

```terminal
az keyvault set-policy -n <your-unique-keyvault-name> --spn <clientId-of-your-service-principal> --secret-permissions delete get list set --key-permissions create decrypt delete encrypt get list unwrapKey wrapKey
```

2. Set environment variables.  You can do this from the command line in the following way:

```terminal
setx AZURE_CLIENT_ID <your_client_id>
setx AZURE_CLIENT_SECRET <your_client_secret>
setx AZURE_TENANT_ID <your_tenantID>
setx KEY_VAULT_NAME <your_key_vault_name>
```

**Update your crud.js to use the Key Vault for Authentication** 

Now, you need to update your code to talk to KeyVault.  Because the Key Vault calls are asyncronous, wrap the main logic of the previous program in a main method, and then call it.  The example code for this is in [**crud_KeyVault.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/main/nodejs/Windows/AzureSqlSample/crud_KeyVault.js) 

Update your connection information, and run the program:

```terminal
  node crud_KeyVault.js
```

## Modify your application to use the popular Sequelize ORM

Create the app directory and initialize Node dependencies.

```terminal
    cd ~/
    mkdir AzureSqlSequelizeSample
    cd AzureSqlSequelizeSample
    npm init -y
    #Install tedious and Sequelize module in your project folder
    npm install tedious
    npm install sequelize
    npm install @azure/keyvault-secrets
    npm install @azure/identity
```

a. Open your favourite text editor and create the file orm.js in the directory AzureSqlSequelizeSample. 
b. Paste the contents below into [**orm.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/main/nodejs/Windows/AzureSqlSequelizeSample/orm.js)
c. Update the connection information to point to your server and database. 

Run the orm.js app

```terminal
    node orm.js
```

# Improve performance using ColumnStore Indexes

In this section we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.

## Create a new table with 3 million rows using sqlcmd

Change to your home directory and create a folder for your project.

```terminal
cd ~/
mkdir AzureSqlColumnstoreSample
cd AzureSqlColumnstoreSample
```

Using your favorite text editor, create a new file called CreateSampleTable.sql in the folder AzureSqlColumnstoreSample. Paste the T-SQL code below into your new SQL file. Save and close the file.

```SQL
WITH a AS (SELECT * FROM (VALUES(1),(2),(3),(4),(5),(6),(7),(8),(9),(10)) AS a(a))
SELECT TOP(3000000)
ROW_NUMBER() OVER (ORDER BY a.a) AS OrderItemId 
,a.a + b.a + c.a + d.a + e.a + f.a + g.a + h.a AS OrderId
,a.a * 10 AS Price
,CONCAT(a.a, N' ', b.a, N' ', c.a, N' ', d.a, N' ', e.a, N' ', f.a, N' ', g.a, N' ', h.a) AS ProductName
INTO Table_with_3M_rows
FROM a, a AS b, a AS c, a AS d, a AS e, a AS f, a AS g, a AS h;
```

Connect to the database using sqlcmd and run the SQL script to create the table with 3 million rows. This may take a few minutes to run.

```terminal
  sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -i ./CreateSampleTable.sql
```

## Create a Node.js that queries this tables and measures the time taken

In your project folder, initialize Node dependencies.

```terminal
npm init -y
npm install tedious
npm install node-uuid
npm install async
npm install @azure/keyvault-secrets
npm install @azure/identity
```

Using you favorite text editor, create a file called [**columnstore.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/main/nodejs/Windows/AzureSqlColumnstoreSample/columnstore.js) in the AzuresSqlColumnstoreSample folder.

##  Measure how long it takes to run the query

Run your Node.js app from the terminal, and record the result.

```terminal
  node columnstore.js
```


## Add a columnstore index to your table.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "CREATE CLUSTERED COLUMNSTORE INDEX Columnstoreindex ON Table_with_3M_rows;"
```

## Step 3.5 Re-run the columnstore.js script and notice how long the query took to complete this time.


```terminal
  node columnstore.js
```

Compare the times.

> Congratulations! You just made your Node.js app faster using Columnstore Indexes!



