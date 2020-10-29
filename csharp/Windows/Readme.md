This page will walk you through everything you need to setup and run the C# samples on your Windows machine.

# Machine Setup

## Install Visual Studio Community Edition and .NET Framework

If you already have Visual Studio installed on your machine, skip the next few steps.

Visual Studio Community edition is a fully-featured, extensible, free IDE for creating modern applications for Android, iOS, Windows, as well as web & database applications and cloud services.

1. Download the installer from **[here**](https://www.visualstudio.com/thank-you-downloading-visual-studio/?sku=Community&rel=15).
1. Run the installer and follow the installation prompts to complete the installation.

## Install The Azure CLI and Login to Azure

1.  Go to this **[site](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)**, and click on the operating system you are currently using.
1.  Install the Azure CLI following the directions.
1.  Open a new instance of the command window and type az login. Follow the instructions that follow to authenticate.

```terminal
az login
```

**Please Note** You may have to login again to your machine after restart or long periods of inactivity.

At this time, you have authenticated yourself and your machine to Azure, so your application can connect. 

# Create and Run your first Azure SQL applications using C# on Windows

## Get Connection Information to use in Connection Strings, and Create a Firewall Rule.


### Get the connection string info from the Azure Portal

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


## Step Create a C# app that connects to Azure SQL and executes queries

**Create a C# console application**

1. Launch Visual Studio Community
1. Click **File -> New -> Project**
1. In the **New project** dialog, click **Windows** located under **Visual C#** in the **Templates** node
1. Click **Console Application Visual C#**
1. Name the project _"AzureSqlSample"_
1. Click **OK** to create the project

Visual Studio creates a new C# Console Application project and opens the file **Program.cs**. Replace the contents of Program.cs with our sample [**Program.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Windows/AzureSqlSample/Program.cs), and update the connection information.

Press **F5** to build and run the project.

## Secure your app by putting Credentials in Azure Key Vault.

### Create an Azure Key Vault and put your Secret into it.

**Create your Key vault and store your credentials**

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

**Update machine settings and program to use KeyVault for authentication** 
These steps are optional if you prefer to just rely on the machine being logged in to azure (using az login).

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
export AZURE_CLIENT_ID <your_client_id>
export AZURE_CLIENT_SECRET <your_client_secret>
export AZURE_TENANT_ID <your_tenantID>
export KEY_VAULT_NAME <your_key_vault_name>
```
### Install Nuget packages to connect your app to Azure Key Vault

1. Open the Package Manager Console in Visual Studio with "Tools-> Nuget Package Manager -> Package Manager Console"
1. Type: "Install-Package Microsoft.Azure.Services.AppAuthentication -Version 1.4.0"
1. Hit Enter

1. Type: "Install-Package Microsoft.Azure.KeyVault -Version 3.0.5"
1. Hit Enter


### Modify your app to use Key Vault

We need to make some changes to our program to access key vault.  This includes making calls to key vault, and adding some asyncrony to the program flow.  

Take the sample code from [**Program_KeyVault.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Windows/AzureSqlSample/Program_KeyVault.cs) and paste it into your Program.cs file.  Now run the program.

## Create a C# app that connects to Azure SQL using the Entity Framework ORM in .NET Framework

### Create a C# console application

1. Launch Visual Studio Community
1. Click **File -> New -> Project**
1. In the **New project** dialog, click **Windows** located under **Visual C#** in the **Templates** node 
1. Click **Console Application Visual C#**
1. Name the project "_AzureSqlEFSample"_
1. Click **OK** to create the project

Visual Studio creates a new C# Console Application project and opens the file **Program.cs**.

### Add Entity Framework and Azure Key Vault dependencies to your project

**Entity Framework**
1. Open the Package Manager Console in Visual Studio with "Tools -> Nuget Package Manager -> Package Manager Console"
1. Type: "Install-Package EntityFramework" 
1. Hit enter


**Authentication**
1. Open the Package Manager Console in Visual Studio with "Tools-> Nuget Package Manager -> Package Manager Console"
1. Type: "Install-Package Microsoft.Azure.Services.AppAuthentication -Version 1.4.0"
1. Hit Enter

**Key Vault**
1. Type: "Install-Package Microsoft.Azure.KeyVault -Version 3.0.5"
1. Hit Enter

Close the Package Manager Console. You have successfully added the required Entity Framework and Azure Key Vault dependencies to your project.

### Create the files for using Entity Framework

For this sample, let's create two tables. The first will hold data about "users" and the other will hold data about “tasks”.

**Create User.cs:**

1. Click **Project -> Add Class**
1. Type "User.cs" in the name field
1. Click **Add** to add the new class to your project

Place the Contents of [**User.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Windows/AzureSqlEFSample/User.cs) in this file.

**Create Task.cs:**

1. Click **Project -> Add Class**
2. Type "Task.cs" in the name field
3. Click **Add** to add the new class to your project

Place the contents of [**Task.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Windows/AzureSqlEFSample/Task.cs) in this file.

**Create EFSampleContext.cs:**

1. Click Project -> Add Class
2. Type "EFSampleContext.cs" in the name field
3. Click Add to add the new class to your project

Place the contents of [**EFSampleContext.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Windows/AzureSqlEFSample/EFSampleContext.cs) in this file.

Replace the code in the **Program.cs** file using [**Program.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Windows/AzureSqlEFSample/Program.cs). Don't forget to update the connection nwith your own. Save and close the file.


Press **F5** to build and run the project.


> Congratulations! You just created a C# app and secured your credentials by placing them in Azure Key Vault! Check out the next section to learn about how you can **make your C# apps faster with SQL Server's Columnstore feature.**


# Improve Performance Using Columnstore

> Now that you have explored the basics, you are ready to see how you can make your app better with Azure SQL. In this module we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.


Note! May features available in SQL Server are also availble in Azure SQL DB.  Please reference **[this link](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-features)** for a full breakdown of which features are available in Azure DB.
Fortunately for us, supported features include innovations that can significantly improve your application’s throughput, latency, and security. Enjoy!

## Create a C# Console Application to explore Columnstore performance enhancements

To showcase the capabilities of Columnstore indexes, let's create a C# application that creates a sample database and a sample table with 5 million rows and then runs a simple query before and after adding a Columnstore index.

### Create a C# console application
1. Launch Visual Studio Community
1. Click **File -> New -> Project**
1. In the **New project** dialog, click **Windows** located under **Visual C#** in the **Templates** node
1. Click **Console Application Visual C#**
1. Name the project "SqlServerColumnstoreSample"
1. Click **OK** to create the project

Visual Studio creates a new C# Console Application project and opens the file **Program.cs**. Replace the contents of **Program.cs** with the contents from [**Program.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Windows/AzureSqlColumnstoreSample/Program.cs), and update the connection properties.

Press **F5** to build and run your project.


> The performance of the query was greatly improved!

Now that you've built a few C# apps with SQL Server and .NET Core, continue checking out other SQL Server features that are available in Azure SQL.
