# How to use these tutorials

Follow these step by step instructions to develop applications in Java to interact with Azure SQL.

# Machine Setup

First, set up your machine based on your OS by following the instructions in the links below.  Then, return to this page and follow the remainder of the tutorial

- [**Ubuntu**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/Ubuntu_Setup.md)
- [**RHEL**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/RHEL_Setup.md)
- [**SLES**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/SLES_Setup.md)
- [**MacOs**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/Mac_Setup.md)

Now that your machine is set up, you can proceed with the next steps.

# Start Writing apps with Java and Azure SQL

> In this section you will create two simple Java apps. One of them will perform basic Insert, Update, Delete, and Select, while the second one will make use of [Hibernate](http://hibernate.org/orm/), one of the most popular Java Object-relational mappers, to execute the same operations.

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

## Create a Java app that connects to Azure SQL and executes queries


In your home directory, create your Maven starter package. This will create the project directory with a basic Maven project and pom.xml file. This step can also be performed in an IDE such as NetBeans or Eclipse.

```terminal
    mvn archetype:generate "-DgroupId=com.sqlsamples" "-DartifactId=AzureSqlSample" "-DarchetypeArtifactId=maven-archetype-quickstart" "-Dversion=1.0.0"
```

Change directories into your newly created project.

```terminal
    cd AzureSqlSample
```

You should already have a file called **pom.xml** in your Maven project located at: _~/AzureSqlSample_

Open this file in your favorite text editor and replace the contents using [**pom.xml**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/AzureSqlSample/pom.xml) to add the Microsoft JDBC Driver for SQL Server to your Maven project and specify the version of Java to compile the project against.

Save and close the file.


You should already have a file called **App.java** in your Maven project located at: AzureSqlSample/src/main/java/com/sqlsamples/App.java

Open this file in your favorite text editor and replace the contents using [**App.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/AzureSqlSample/App.java). Don't forget to replace the username and password with your own. Save and close the file.


Build the project and create a jar package using the following command:

```terminal
    mvn package
```


Now run the application. You can remove the "-q" in the command below to show info messages from Maven.

```terminal
mvn -q exec:java "-Dexec.mainClass=com.sqlsamples.App"
```

>You created your first Java + Azure SQL app with Maven! Now try securing your credentials in Azure Key Vault.

# Secure your app by putting Credentials in Azure Key Vault

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

The **Pom.xml** file in this sample already includes references required to authenticate to Azure Key Vault. 


## Update your App.java to use the Key Vault for Authentication

Now we add calls to the Key Vault libraries.  You can update your version of App.java using [**App_KeyVault.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/AzureSqlSample/App_KeyVault.java). Update your database connection info and keyvault name afterwards.

Now, build and run the program:

```terminal
mvn package
```


```terminal
mvn -q exec:java "-Dexec.mainClass=com.sqlsamples.App"
```

> Now, you have secured your credentials in azure key vault, and fetched them for use in your application!  Check out the next section to create a Java App using an ORM!

# Create a Java app that connects to SQL Server using the popular framework Hibernate

In your home directory, create your Maven starter package. This will create the project directory with a basic Maven project and pom.xml file. This step can also be performed in an IDE such as NetBeans or Eclipse.

```terminal
mvn archetype:generate "-DgroupId=com.sqlsamples" "-DartifactId=AzureSqlHibernateSample" "-DarchetypeArtifactId=maven-archetype-quickstart" "-Dversion=1.0.0"
```


Change directories into your newly created project.

```terminal
cd AzureSqlHibernateSample
```

You should already have a file called **pom.xml** in your Maven project located at: _\AzureSqlHibernateSample_

Open this file in your favorite text editor and replace the contents using [**pom.xml**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/AzureSqlHibernateSample/pom.xml) to add the Microsoft JDBC Driver for SQL Server and Hibernate to your Maven project and specify the version of Java to compile the project against.

Save and close the file.

For this sample, let's create two tables. The first will hold data about "users". Using [**User.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/AzureSqlHibernateSample/User.java), create a **User.java** file in your Maven project located at: AzureSqlHibernateSample\src\main\java\com\sqlsamples\User.java.


Let's create a second table to assign tasks to users. Using [**Task.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/AzureSqlHibernateSample/Task.java), create a **Task.java** file in your Maven project located at: AzureSqlHibernateSample\src\main\java\com\sqlsamples\Task.java.


Finally, replace the code in the **App.java** file in your Maven project located at: AzureSqlHibernateSample\src\main\java\com\sqlsamples\App.java with the template [**App.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/AzureSqlHibernateSample/App.java), and update any relevant connection information.


From the project folder, build the project and create a jar package using the following command:

```terminal
mvn package
```

Now run the application. You can remove the "-q" in the command below to show info messages from Maven.

```terminal
mvn -q exec:java "-Dexec.mainClass=com.sqlsamples.App"
```


> Congratulations! You created your first two Java apps with SQL Server! Check out the next section to learn about how you can make your Java apps faster with SQL Server’s Columnstore feature.


# Improve Performance using Columnstore

> Now that you have explored the basics, you are ready to see how you can make your app better with Azure SQL. In this module we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.


## Create a Java app to demonstrate Columnstore indexes

To showcase the capabilities of Columnstore indexes, let's create a Java application that creates a sample database and a sample table with 5 million rows and then runs a simple query before and after adding a Columnstore index.

Change to your home directory. Create your Maven starter package. This will create the project directory with a basic Maven project and pom.xml file.

```terminal
cd ~/
mvn archetype:generate -DgroupId=com.sqlsamples -DartifactId=AzureSqlColumnstoreSample -DarchetypeArtifactId=maven-archetype-quickstart -Dversion=1.0.0
```

You should already have a file called **pom.xml** in your Maven project located at: _\AzureSqlColumnstoreSample_

Open this file in your favorite text editor and replace the contents using [**pom.xml**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/AzureSqlColumnstoreSample/pom.xml) to add the Microsoft JDBC Driver for SQL Server to your Maven project and specify the version of Java to compile the project against.

Save and close the file.
Change directories into your newly created project.

```terminal
cd AzureSqlColumnstoreSample
```

You should already have a file called **App.java** in your Maven project located at: \AzureSqlColumnstoreSample\src\main\java\com\sqlsamples\App.java

Open this file in your favorite text editor and replace the contents using [**App.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Unix-based/AzureSqlColumnstoreSample/App.java). Don't forget to update connection information. Save and close the file.

Build the project and create a jar package using the following command:

```terminal
cd ~/AzureSqlColumnstoreSample
mvn package
```

Now run the application. You can remove the of "-q" in the command below to show info messages from Maven.

```terminal
mvn -q exec:java -Dexec.mainClass=com.sqlsamples.App
```

> Congratulations! You just made your Java app faster using Columnstore Indexes!




Save and close all three files.

