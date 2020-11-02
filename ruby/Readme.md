---
page_type: sample
languages:
- ruby
products:
- azure-sql-database	
description: "Azure SQL - Getting Started Samples"
---
# Developing applications with Ruby and Azure SQL 

This repository contains a tutorial that will guide you through the creation of a simple solution using [Ruby](https://www.ruby-lang.org/en/) to take advantage of Azure SQL. Azure SQL as many features for developers and learning how to take advantage of it will help you to create secure, scalable and performant modern applications. To learn more about several of the features that Azure SQL provides to developers, read here: [10 reasons to use Azure SQL in your next project](https://devblogs.microsoft.com/azure-sql/10-reasons-to-use-azure-sql-in-your-next-project/).

With the proposed tutorial you will learn how to create a database, use the most common packages to connect to it and, finally, you'll see how performance can be improved *a lot* by using Columnstore Indexes.

No matter which is the platform or the OS you are using, you can happily use Azure SQL. As you can see the tutorial is available to be used with:

- [Red Hat Enterprise Linux](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/ruby/RHEL_Setup.md)
- [Ubuntu](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/ruby/Ubuntu_Setup.md)
- [Mac](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/ruby/Mac_Setup.md)

Once you'll be more expert on Azure SQL and Ruby, you can also take advtange of [several samples](https://docs.microsoft.com/en-us/samples/browse/?expanded=dotnet&products=azure-sql-database&languages=ruby) that will help you to create Full-Stack solutions or Back-End API, that can be used in project of any size and scale.

# Prerequisites

## Create an Auzre SQL DB
All of the above examples require an Azure SQL DB  Please follow these instructions to create one.

Go to this [**site**](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-single-database-get-started?tabs=azure-portal) for instructions on how to set up an Azure Hosted SQL Database.

1.  Perform the Prerequisites steps.

2. Follow steps 1-17 from the section: **Create a Single Database**.

## Set up your machine for Ruby and Azure SQL

Set up your machine using the instructions for your OS by clicking on the links below, then return here to complete the tutorial.

- [Red Hat Enterprise Linux](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/ruby/RHEL_Setup.md)
- [Ubuntu](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/ruby/Ubuntu_Setup.md)
- [Mac](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/ruby/Mac_Setup.md)

# Get Started with Ruby and Azure SQL

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

## Create a Ruby app that connects to Azure SQL and executes queries


## Create a Ruby app that connects to Azure SQL and executes queries

Create a new project directory and install [TinyTDS](https://github.com/rails-sqlserver/tiny_tds). TinyTDS is used to connect Ruby applications to Azure SQL DB.

```terminal
cd ~/
mkdir AzureSqlSample
cd AzureSqlSample
gem install tiny_tds
```

Using your favorite editor, create a file named **connect.rb** in the AzureSqlSample folder, based on the sample [**connect.rb**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/ruby/AzureSqlSample/connect.rb). Don't forget to update the username and password with your own. Save and close the file.


Run the Ruby script from the terminal.

```terminal
ruby connect.rb
```


Using your favorite text editor, create a new file called **crud.rb** in the AzureSqlSample folder, based on the sample [**crud.rb**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/ruby/AzureSqlSample/crud.rb). This will insert, update, delete, and read a few rows. Don't forget to update the username and password with your own. Save and close the file.


Run the Ruby script from the terminal.

```terminal
ruby crud.rb
````

> You created your first Ruby + Azure SQL DB app! Check out the next section to create an app using Ruby on Rails!

## Step 2.3 Create a Ruby app that connects to Azure SQL DB using the Active Record ORM 

Create a new project directory and install the gem dependencies required to connect Ruby to Azure SQL DB using the [Active Record ORM](http://guides.rubyonrails.org/active_record_basics.html#active-record-as-an-orm-framework). You'll also need [TinyTDS](https://github.com/rails-sqlserver/tiny_tds) and the [activerecord-sqlserver-adapter](https://github.com/rails-sqlserver/activerecord-sqlserver-adapter). 

```terminal
cd ~/
mkdir AzureSqlActiveRecordSample
cd AzureSqlActiveRecordSample
gem install active_record tiny_tds activerecord-sqlserver-adapter
```

Using your favorite editor, create a file named **activerecordcrud.rb** in the AzureSqlSample folder based on the sample [**activerecordcrud.rb**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/ruby/AzureSqlSample/activerecordcrud.rb). Don't forget to update the username and password with your own. Save and close the file.


Run the Ruby script from the terminal.

```terminal
ruby activerecordcrud.rb
```

> Congratulations! You created your first two Ruby apps with Azure SQL DB! Check out the next section to learn about how you can make your Ruby apps faster with Azure SQL DB's Columnstore feature.

# Improve Performance with Columnstore 

> Now that you have explored the basics, you are ready to see how you can make your app better with Azure SQL. In this module we will show you a simple example of [Columnstore Indexes](https://docs.microsoft.com/en-us/sql/relational-databases/indexes/columnstore-indexes-overview) and how they can improve data processing speeds. Columnstore Indexes can achieve up to 100x better performance on analytical workloads and up to 10x better data compression than traditional rowstore indexes.

## Create a Ruby app to demonstrate Columnstore indexes

Create a new directory for your project and install [TinyTDS](https://github.com/rails-sqlserver/tiny_tds).

```terminal
cd ~/
mkdir AzrueSqlColumnstoreSample
cd AzureSqlColumnstoreSample
gem install tiny_tds
```

Using your favorite text editor, create a file called **columnstore.rb** in the AzureSqlColumnstoreSample folder, based on the sample [**columnstore.rb**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/ruby/AzureSqlColumnstoreSample/columnstore.rb). Save and close the file.


Now run the program.

```terminal
ruby columstore.rb
```

> Congratulations! You just made your Ruby app faster using Columnstore Indexes!







