Follow all of these steps to start writing Python applications using Azure SQL.

# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Python and connect to Azure SQL.

## Step 1.2 Install Python

Download and run the installer [**here**](https://www.python.org/downloads/)

Next, add Python to your path

1. Press start 
2. Search for "Advanced System Settings" 
3. Click on the "Environment Variables" button 
4. Add the location of the Python27 folder to the PATH variable in System Variables. The following is a typical value for the PATH variable: C:\Python27, so is: C:\Users\User\AppData\Local\Programs\Python\Python38-32\

> You have succesfully installed Python on your machine!

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

## Install Pyodbc

 These instructions are based off of the ones found [**here:**](https://docs.microsoft.com/en-us/sql/connect/python/pyodbc/step-1-configure-development-environment-for-pyodbc-python-development?view=sql-server-ver15)

```terminal
> cd C:\<The path you used when you added python to your PATH variable>\Scripts  
> pip install pyodbc
```

# Start Writing apps with Python and Azure SQL

> In this section you will create a simple Python app. The Python app will perform basic Insert, Update, Delete, and Select.

## Get Connection Information to use in Connection Strings, and Create a Firewall Rule.

1. Using the Azure Portal, go to your database and look in the panel on the left.  It should say Overview, activity log, …

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

### Create a firewall rule

In order to connect to your Azure SQL database, you will need to create a filrewall rule on the target server.  This allows your application to talk to your Azure SQL Database.

1.  From your database, in the Overview, you can look in the panel on the right, and where it says Server name, click there.
1.  From the server, in the search bar at the top, type "firewall", and select "Firewalls and virtual networks".
1.  In the bar at the top, you sould see three shortcut options: Save, Discard, and + Add Client IP.  Click on + Add Client IP.
1.  Next click Save.  You can close this view now.

## Create a Python app that connects to Azure SQL and executes queries

Create a new folder for the sample

```terminal
mkdir AzureSqlSample
cd AzureSqlSample
```

Execute the T-SQL scripts below in the terminal with sqlcmd to a table and insert some row.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "CREATE TABLE Employees (Id INT IDENTITY(1,1) NOT NULL PRIMARY KEY, Name NVARCHAR(50), Location NVARCHAR(50));"
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -Q "INSERT INTO Employees (Name, Location) VALUES (N'Jared', N'Australia'), (N'Nikita', N'India'), (N'Tom', N'Germany');"
```

Using your favorite text editor, create a new file called crud.py in the AzureSqlSample folder based on the file [**crud.py**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/python/Windows/AzureSqlSample/crud.py). Update the connection information. This will insert, update, delete, and read a few rows.

Run your Python script from the terminal.

```terminal
python crud.py
```


# Secure your credentials using Azure Key Vault


## Create your Key vault and store your credentials

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

## Update machine settings and program to use KeyVault for authentication

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

2. Set environment variables.  You can do this from the command line in the following way:

```terminal
setx AZURE_CLIENT_ID <your_client_id>
setx AZURE_CLIENT_SECRET <your_client_secret>
setx AZURE_TENANT_ID <your_tenantID>
setx KEY_VAULT_NAME <your_key_vault_name>
```

## Install the python libraries

The reference documentation about the python sdk for Azure can be found [**here.**](https://docs.microsoft.com/en-us/python/api/overview/azure/key-vault?view=azure-python)

1. Execute the following from an Administrator command window:

```terminal 
pip install azure-keyvault-secrets 
pip install azure-keyvault-keys
pip install azure-keyvault-certificates
pip install azure-identity
```

## Update your crud.py to use the Key Vault for Authentication

The sample [**crud_KeyVault.py**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/python/Windows/AzureSqlSample/crud_KeyVault.py) shows how to modify the previous program to use the Key Vault to store and retreive passwords. Be sure to check that all of the other connection information is correct, save, and close the file.

Then run again:

```terminal
python crud.py
```

> Congratulations! You created your first Python app with Azure SQL, and secured your credentials in Azure Key Vault! Check out the next section to learn about how you can make your Python app faster with Azure SQL's Columnstore feature.

# Improve Performance with Columnstore
> Now that you have explored the basics, you are ready to see how you can make your app better with Azure SQL. In this module we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.

## Create a new table with 3 million rows using sqlcmd

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_database -t 60000 -Q "WITH a AS (SELECT * FROM (VALUES(1),(2),(3),(4),(5),(6),(7),(8),(9),(10)) AS a(a))
SELECT TOP(3000000)
ROW_NUMBER() OVER (ORDER BY a.a) AS OrderItemId
,a.a + b.a + c.a + d.a + e.a + f.a + g.a + h.a AS OrderId
,a.a * 10 AS Price
,CONCAT(a.a, N' ', b.a, N' ', c.a, N' ', d.a, N' ', e.a, N' ', f.a, N' ', g.a, N' ', h.a) AS ProductName
INTO Table_with_3M_rows
FROM a, a AS b, a AS c, a AS d, a AS e, a AS f, a AS g, a AS h;"
```

##  Create a Python app that queries this tables and measures the time taken

```terminal
mkdir AzureSqlColumnstoreSample
cd AzureSqlColumnstoreSample
```

Using your favorite text editor, create a new file called columnstore.py, based on the sample [**columnstore.py**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/python/Windows/AzureSqlColumnstoreSample/columnstore.py) in the SqlServerColumnstoreSample folder. 


##  Measure how long it takes to run the query

Run your Python script from the terminal.

```terminal
python columnstore.py
```

## Add a columnstore index to your table.

```terminal
sqlcmd -S your_server.database.windows.net -U your_user -P your_password -d your_db -Q "CREATE CLUSTERED COLUMNSTORE INDEX Columnstoreindex ON Table_with_3M_rows;"
```

## Measure how long it takes to run the query with a columnstore index

```terminal
python columnstore.py
```


> Congratulations! You just made your Python app faster using Columnstore Indexes!

 

