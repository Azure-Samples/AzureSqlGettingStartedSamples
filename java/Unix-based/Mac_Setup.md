# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Java and connect to Azure SQL.

## Install Java

If you already have Java installed on your machine, skip this step. Install the Azul Zulu JDK's for Mac.  Follow instructions [here](https://docs.microsoft.com/en-us/java/azure/jdk/java-jdk-install?view=azure-java-stable).

[Download the 64-bit Azul Zulu JDK 8 as a ZIP file](https://repos.azul.com/azure-only/zulu/packages/zulu-11/11.0.3/zulu-11-azure-jdk_11.31.11-11.0.3-macosx_x64.zip) to a location on your client, such as /Library/Java/JavaVirtualMachines/. (.DMG packages are also provided on Azul's Azure downloads page.)

Launch Finder, navigate to the download directory, and double-click the ZIP file. Alternatively, you can launch a terminal command window, navigate to the directory, and run:

```terminal
unzip <name_of_zulu_package>.zip
```

> You have sucessfully installed Java on your macOS! 

## Install Maven

[Maven](https://maven.apache.org/) can be used to help manage dependencies, build, test and run your Java project.


Download the latest version of Maven from [here](https://maven.apache.org/), selecting the binary tar.gz file.  Extract the archive to your desired location.

```terminal
sudo su
chown -R root:wheel Downloads/apache-maven*
mv Downloads/apache-maven* /opt/apache-maven
exit
```

Now using your favorite text editor to add the following to your /.profile.  This will add the Maven binaries to your path:

```terminal
export PATH=$PATH:/opt/apache-maven/bin
```


Check that you have Maven properly installed by running the following command.

```terminal
mvn -v
```


# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/java/Unix-based#start-writing-apps-with-java-and-azure-sql) to complete the tutorial.
