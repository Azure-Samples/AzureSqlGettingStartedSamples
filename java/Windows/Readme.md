Follow all of these steps to start writing Java applications using Azure SQL.

# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Java and connect to Azure SQL.

##  Install Java

If you already have Java installed on your machine, skip the next three steps.

Install the **Java Long-terms supoort for Azure** by following the steps below.

1. Download and install Java from [**Microsoft's Java Support Page**](https://docs.microsoft.com/en-us/java/azure/jdk/java-jdk-install?view=azure-java-stable)
2.  Check your Java install worked by opening a command window and typing the following command.

```terminal
  Java --version
```

## Install Maven

[Maven](https://maven.apache.org/) can be used to help manage dependencies, build, test and run your Java project. Follow the instructions below to install Maven.

1. Download the [Maven binary](https://downloads.apache.org/maven/maven-3/3.6.3/binaries/apache-maven-3.6.3-bin.zip)
2. Unzip the installer to a file location on your computer

Add the Maven bin directory to your PATH environment variable and add the JRE to the JAVA_HOME environment variable

1. Press start 
2. Search for "Advanced System Settings" 
3. Click on the "Environment Variables" button 
4. Add the location of the bin folder of the Maven installation to the PATH variable in **System Variables**. The following is a typical value for the PATH variable: C:\WINDOWS\system32;C:\WINDOWS;C:\Maven\bin
5. Create a new System Variable for "JAVA_HOME" and point it to the JDK folder (ex. C:\Program Files\Zulu-11)
6. Check that Maven was installed properly by opening a new command window (to refresh the variables), and then running the following command.

```terminal
  mvn --version
```

## Install The Azure CLI and Login to Azure


1.  Go to this **[site](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest)**, and click on the operating system you are currently using.
1.  Install the Azure CLI following the directions.
1.  Open a new instance of the command window and type az login. Follow the instructions that follow to authenticate.

```terminal
az login
```

**Please Note** You may have to login again to your machine after restart or long periods of inactivity.

At this time, you have authenticated yourself and your machine to Azure, so your application can connect. 


> You have successfully installed Java and Maven on Windows, and authenticated to Azure. You now have everything you need to start writing your Java apps with Azure SQL!

# Start Writing apps with Java and Azure SQL

> In this section you will create two simple Java apps. One of them will perform basic Insert, Update, Delete, and Select, while the second one will make use of [Hibernate](http://hibernate.org/orm/), one of the most popular Java Object-relational mappers, to execute the same operations.

## Get Connection Information to use in Connection Strings, and Create a Firewall Rule.
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
This may cause a number of things to be downloaded.

```terminal
    mvn archetype:generate "-DgroupId=com.sqlsamples" "-DartifactId=AzureSqlSample" "-DarchetypeArtifactId=maven-archetype-quickstart" "-Dversion=1.0.0"
```

Change directories into your newly created project.

```terminal
    cd AzureSqlSample
```

You should already have a file called **pom.xml** in your Maven project located at: _\AzureSqlSample_

Open this file in your favorite text editor and replace the contents using [**pom.xml**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Windows/AzureSqlSample/pom.xml) to add the Microsoft JDBC Driver for SQL Server to your Maven project and specify the version of Java to compile the project against.

Save and close the file.

You should already have a file called **App.java** in your Maven project located at: AzureSqlSample\src\main\java\com\sqlsamples\App.java

Open this file in your favorite text editor and replace the contents using [**App.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Windows/AzureSqlSample/App.java). Don't forget to replace the connection information with your own. Save and close the file.

Build the project and create a jar package using the following command from your project root (AzureSqlSample):

```terminal
    mvn package
```

Now run the application. You can remove the "-q" in the command below to show info messages from Maven.

```terminal
mvn -q exec:java "-Dexec.mainClass=com.sqlsamples.App"
```

>You created your first Java + Azure SQL app with Maven!  Now try securing your credentials in Azure Key Vault.

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

## Update your application to get credentials from Key Vault.

The **pom.xml** from this sample is already updated to have the appropriate libraries for connecting to Azure Key Vault.  However, you need to modify the program by using the contents from [**App_KeyVault.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Windows/AzureSqlSample/App_KeyVault.java), and updating the connection infomation as needed.


Now, build and then run the program:

```terminal
mvn package
```

Run the application. You can remove the "-q" in the command below to show info messages from Maven.

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

Open this file in your favorite text editor and replace the contents using [**pom.xml**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Windows/AzureSqlHibernateSample/pom.xml) to add the Microsoft JDBC Driver for SQL Server and Hibernate to your Maven project and specify the version of Java to compile the project against.

Save and close the file.

For this sample, let's create two tables. The first will hold data about "users". Using [**User.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Windows/AzureSqlHibernateSample/User.java), create your own **User.java** file in your Maven project located at: AzureSqlHibernateSample\src\main\java\com\sqlsamples\User.java.

Similarly, we will create a second table to assign tasks to users.  Using [**Task.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Windows/AzureSqlHibernateSample/Task.java), create your own **Task.java** file in your Maven project located at: AzureSqlHibernateSample\src\main\java\com\sqlsamples\Task.java.

Now, update the code in the **App.java** file in your Maven project located at: AzureSqlHibernateSample\src\main\java\com\sqlsamples\App.java by using the code from [**App.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Windows/AzureSqlHibernateSample/App.java).  Update any connection information, and then make sure all three files are saved.

From the project folder, build the project and create a jar package using the following command:

```terminal
mvn package
```
Now run the application. You can remove the "-q" in the command below to show info messages from Maven.

```terminal
mvn -q exec:java "-Dexec.mainClass=com.sqlsamples.App"
```

> Congratulations! You created your first two Java apps with Azure SQL! Check out the next section to learn about how you can make your Java apps faster with Azure SQL’s Columnstore feature.

# Improve Performance with Columnstore
> Now that you have explored the basics, you are ready to see how you can make your app better with Azure SQL. In this module we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.

## Step 3.1 Create a Java app to demonstrate Columnstore indexes

To showcase the capabilities of Columnstore indexes, let's create a Java application that creates a sample database and a sample table with 5 million rows and then runs a simple query before and after adding a Columnstore index.

Change to your home directory. Create your Maven starter package. This will create the project directory with a basic Maven project and pom.xml file.

```terminal
mvn archetype:generate -DgroupId=com.sqlsamples -DartifactId=AzureSqlColumnstoreSample -DarchetypeArtifactId=maven-archetype-quickstart -Dversion=1.0.0
```

You should already have a file called **pom.xml** in your Maven project located at: _\AzureSqlColumnstoreSample_

Open this file in your favorite text editor and replace the contents using [**pom.xml**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Windows/AzureSqlColumnstoreSample/pom.xml)  to add the Microsoft JDBC Driver for SQL Server to your Maven project and specify the version of Java to compile the project against.

Save and close the file.


Change directories into your newly created project.

```terminal
cd AzureSqlColumnstoreSample
```

You should already have a file called **App.java** in your Maven project located at: \AzureSqlColumnstoreSample\src\main\java\com\sqlsamples\App.java

Open this file in your favorite text editor and replace the contents using [**App.java**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/java/Windows/AzureSqlColumnstoreSample/App.java). Don't forget to update the username and password with your own. Save and close the file.


Build the project and create a jar package using the following command:

```terminal
mvn package
```


Now run the application. You can remove the of "-q" in the command below to show info messages from Maven.

```terminal
mvn -q exec:java -Dexec.mainClass=com.sqlsamples.App
```


> Congratulations! You just made your Java app faster using Columnstore Indexes!




