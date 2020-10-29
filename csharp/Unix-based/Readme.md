# How to use these tutorials

## Prerequisite: Machine Setup

First, set up your machine based on your OS by following the instructions in the links below.  Then, return to this page and follow the remainder of the tutorial

- [**Ubuntu**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/Ubuntu_Setup.md)
- [**RHEL**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/RHEL_Setup.md)
- [**SLES**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/SLES_Setup.md)
- [**MacOs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/Mac_Setup.md)

Now that your machine is set up, you can proceed with the next steps.

# Create your first C# Projects

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

## Create a C# app that connects to Azure SQL and executes queries

Change to your home directory. Create a new .NET Core project. This will create the project directory with a basic .NET Core Program.cs and csproj file.

```terminal
cd ~/
dotnet new console -o AzureSqlSample
```

```results
Getting ready...
The template "Console Application" was created successfully.

Processing post-creating actions...
Running 'dotnet restore' on AzureSqlSample/AzureSqlSample.csproj...
  Restore completed in 126.89ms for /home/usr1/AzureSqlSample.csproj.

Restore succeeded.
```

You should already have a file called **AzureSqlSample.csproj** in your .NET Core project located at: ~/AzureSqlSample

Open this file in your favorite text editor and replace the contents with the code below to add System.Data.SqlClient to your project. Save and close the file.

```xml
<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>netcoreapp3.1</TargetFramework>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="System.Data.SqlClient" Version="4.4.0" />
  </ItemGroup>

</Project>
```

You should already have a file called **Program.cs** in your .NET Core project located at: ~/AzureSqlSample.  Replace this file (or its contents) using [**Program.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/AzureSqlSample/Program.cs).  Update the connection string information.

Change directories into the project folder and restore the dependencies in the csproj by running the following commands.

```terminal
cd ~/AzureSqlSample
dotnet restore
```
Now build and run.

```terminal
dotnet run
```

> You created your first C# + Azure SQL app with .NET Core on Ubuntu! Check out the next section to secure this app by putting credentials in Azure Key Vault.

## Secure your credentials using Azure Key vault

### Create an Azure Key Vault and put your Secret into it.

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

### Set up your environment to Authenticate to Azure Key Vault

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

### Add required Nuget packages to support Azure Key Vault

```xml
<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>netcoreapp3.1</TargetFramework>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="System.Data.SqlClient" Version="4.4.0" />
    <PackageReference Include="Microsoft.Azure.Services.AppAuthentication" Version="1.4.0" />
    <PackageReference Include="Microsoft.Azure.KeyVault" Version="3.0.5" />
  </ItemGroup>

</Project>
```

Change directories into the project folder and restore the dependencies in the csproj by running the following commands.

```terminal
cd ~/AzureSqlSample
dotnet restore
```

### Update your Code to access Key Vault

Replace your Program.cs with the file (or contents) of [**Program_KeyVault.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/AzureSqlSample/Program_KeyVault.cs), update your connection info, and then run the program.  

## Create a C# Application that connects to Azure SQL using the Entity Framework ORM in .Net Framework

Change to your home directory. Create a new .NET Core project. This will create the project directory with a basic .NET Core Program.cs and csproj file.

```terminal
cd ~/
dotnet new console -o AzureSqlEFSample
```


You should now have a file called **AzureSqlEFSample.csproj** in your .NET Core project located at: ~/AzureSqlEFSample
Open this file in your favorite text editor and replace the contents with the code below to add Entity Framework Core to your project. Save and close the file.

### Add Entity Framework and Azure Key Vault dependencies to your project
```xml
<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>netcoreapp3.1</TargetFramework>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="EntityFramework" Version="6.4.0" />
    <PackageReference Include="Microsoft.Azure.KeyVault" Version="3.0.5" />
    <PackageReference Include="Microsoft.Azure.Services.AppAuthentication" Version="1.4.0" />
    <PackageReference Include="Microsoft.EntityFrameworkCore" Version="3.1.3" />
    <PackageReference Include="Microsoft.EntityFrameworkCore.SqlServer" Version="3.1.3" />
  </ItemGroup>

</Project>
```
For this sample, let’s create two tables. The first will hold data about “users”. Create a [**User.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/AzureSqlEFSample/User.cs) file in your .NET Core project located at: ~/AzureSqlEFSample/User.cs

Let’s create a second table to assign tasks to users. Create a [**Task.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/AzureSqlEFSample/Task.cs) file in your .NET Core project located at: ~/AzureSqlEFSample/Task.cs

**Create EFSampleContext.cs:**

Let's also create a class for the Entity Framework Database context. Use your favorite text editor to create the file [**EFSampleContext.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/AzureSqlEFSample/EFSampleContext.cs) file in your .NET Core project located at: ~/AzureSqlEFSample/EFSampleContext.cs

Replace the code in the **Program.cs** file in your .NET Core project located at: ~/AzureSqlEFSample/Program.cs with [**Program.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/AzureSqlEFSample/Program.cs). Don't forget to update the connection information and key vault name with your own. Save and close the file.


Change directories into the project folder and restore the dependencies in the csproj by running the following commands.

```terminal
cd ~/SqlServerEFSample
dotnet restore
```

Now build and run.

```terminal
dotnet run
```

> Congratulations! You just created two C# apps! Check out the next section to learn about how you can **make your C# apps faster with Azure SQl and SQL Server's Columnstore feature**.

# Improve Performance Using Columnstore

> Now that you have explored the basics, you are ready to see how you can make your app better with Azure SQL. In this module we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.


Note! May features available in SQL Server are also availble in Azure SQL DB.  Please reference **[this link](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-features)** for a full breakdown of which features are available in Azure DB.
Fortunately for us, supported features include innovations that can significantly improve your application’s throughput, latency, and security. Enjoy!

## Create a C# Console Application to explore Columnstore performance enhancements

o showcase the capabilities of Columnstore indexes, let's create a C# application that creates a sample database and a sample table with 5 million rows and then runs a simple query before and after adding a Columnstore index.

Change to your home directory. Create a new .NET Core project. This will create the project directory with a basic .NET Core Program.cs and csproj file.

```terminal
cd ~/
dotnet new console -o AzureSqlColumnstoreSample
```

You should already have a file called **AzureSqlColumnstoreSample.csproj** in your .NET Core project located at: _~/AzureSqlColumnstoreSample_

Open this file in your favorite text editor and replace the contents with the code below to add System.Data.SqlClient to your project. Save and close the file.

```xml
<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>netcoreapp3.1</TargetFramework>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="System.Data.SqlClient" Version="4.4.0" />
  </ItemGroup>

</Project>
```

You should already have a file called **Program.cs** in your .NET Core project located at: _~/AzureSqlColumnstoreSample_

Open this file in your favorite text editor and replace the contents using [**Program.cs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/AzureSqlColumnstoreSample/Program.cs). Don't forget to replace the connection info with your own. Save and close the file.


Change directories into the project folder and restore the dependencies in the csproj by running the following commands.

```terminal
cd ~/AzureSqlColumnstoreSample
dotnet restore
```

Now build and run.
```terminal
dotnet run
```


> The performance of the query was greatly improved! 
Now that you've built a few C# apps with Azure SQL and .NET Core, continue checking out other Azure SQL features.





