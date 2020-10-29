# How to use these tutorials

## Prerequisite: Machine Setup

First, set up your machine based on your OS by following the instructions in the links below.  Then, return to this page and follow the remainder of the tutorial

- [**Ubuntu**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/nodejs/Unix-based/Ubuntu_Setup.md)
- [**RHEL**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/nodejs/Unix-based/REHL_Setup.md)
- [**SLES**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/nodejs/Unix-based/SLES_Setup.md)
- [**MacOs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/nodejs/Unix-based/Mac_Setup.md)

Now that your machine is set up, you can proceed with the next steps.

# Create and Run your first Node.js apps

> In this section you will create three simple Node.js apps. One of them will just connect to the database, the second will perform basic Insert, Update, Delete, and Select, and the third one will make use of Sequelize, one of the most popular Node.js Object-relational mappers, to execute the same operations.

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

Using your favorite editor, create a file named [**connect.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/main/nodejs/Unix-based/AzureSqlSample/connect.js) in the AzureSqlSample folder. Be sure to update your connection string information and save the file.


Run the application.

```terminal
  node connect.js
```

## Create and Manipulate Data

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
Using your favorite text editor, create a new file called [**crud.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/nodejs/Unix-based/AzureSqlSample/crud.js) in the AzureSqlSample folder.  Update the connection info, and run it.


```terminal
  node crud.js
```

## Secure your credentials using Azure Key vault

**Create an Azure Key Vault and put your Secret into it.**

First, you need to create an Azure Key Vault.  
It is recommended that you create a Key Vault in the same resource group as your database and server.

1. From the Azure Portal, select "+ Create a Resource".  Search for "Key Vault" and select that.
1. In the "Create key vault" page, fill out the resource group and key vault name.
1. Select "Review and Create", then "Create".

For future reference, there are more details [**here**](https://docs.microsoft.com/en-us/azure-stack/user/azure-stack-key-vault-manage-portal?view=azs-2002).

Now that you have created the Key Vault, you need to add a secret called **AppSecret** to your vault.

1. From the Azure Portal dashboard, select All resources, select the key vault that you created earlier, and then select the Keys tile.
1. In the Keys pane, select Generate/Import.
1. Name your key AppSecret, then make the secret your password.
1. You can leave the values for Content Type, activation date, expiration date, and Enabled (Yes) as the defaults.
1. Select Create to start the deployment.

**Add required dependencies to allow the program to connect.**

```terminal
  npm install @azure/keyvault-secrets
  npm install @azure/identity
```

**Set up your environment to Authenticate to Azure Key Vault**

This section takes you through the steps described [**on this site**](https://docs.microsoft.com/en-us/azure/key-vault/secrets/quick-create-node) to set up your machine for authentication to the key vault.  You need to do this to use the **DefaultAzureCredentialBuilder()**.

1. Open a command window and execute **az login** if you have not already.
1. Create a service prinicpal (make sure you take note of the output, as you will use it in the next two steps.):

```terminal
az ad sp create-for-rbac -n "http://mySP" --sdk-auth
```

1. Give the service prinicpal access to your key vault.

```terminal
az keyvault set-policy -n <your-unique-keyvault-name> --spn <clientId-of-your-service-principal> --secret-permissions delete get list set --key-permissions create decrypt delete encrypt get list unwrapKey wrapKey
```

1. Set environment variables.  You can do this from the command line in the following way:

```terminal
export AZURE_CLIENT_ID=your_client_id

export AZURE_CLIENT_SECRET=your_client_secret

export AZURE_TENANT_ID=your_tenant_id

export KEY_VAULT_NAME=your_keyvault_name
```

**Update your program to use the Key Vault for Authentication**

Copy-paste the following into your [**crud_KeyVault.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/nodejs/Unix-based/AzureSqlSample/crud_KeyVault.js), and update the keyvault name, and other connection information.

Now run the program.

```terminal
  node crud_KeyVault.js
```

## Create a Node.js app that connects to SQL Server using the popular Sequelize ORM

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

1. Open your favorite text editor and create the file orm.js in the directory AzureSqlSequelizeSample. 
1. Paste the contents of [**orm.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/nodejs/Unix-based/AzureSqlSequelizeSample/orm.js) 
1. Update the variable for password to use your own password specified in the first module. 
1. Save and close orm.js

Run the orm.js app

```terminal
    node orm.js
```
> Congratulations! You created your first two Node.js apps with Azure SQL! Check out the next section to learn about how you can make your Node.js apps faster with SQL Server’s Columnstore feature

# Improve Performance using Columnstore

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

Using you favorite text editor, create a file called [**columnstore.js**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/nodejs/Unix-based/AzureSqlColumnstoreSample/columnstore.js) in your folder.  Update your connection information, save, and close.

## Measure how long it takes to run the query

Run your Node.js app from the terminal, and take note of the time.

```terminal
  node columnstore.js
```

## Add a columnstore index to your table using SQLCMD.

Run this command to create a Columnstore Index on your table:

```terminal
sqlcmd -S localhost -U sa -P your_password -d SampleDB -Q "CREATE CLUSTERED COLUMNSTORE INDEX Columnstoreindex ON Table_with_3M_rows;"
```

## Re-run the columnstore.js script and notice how long the query took to complete this time.


```terminal
  node columnstore.js
```

> Congratulations! You just made your Node.js app faster using Columnstore Indexes!


