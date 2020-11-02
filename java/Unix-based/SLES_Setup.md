# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Java and connect to Azure SQL.

## Install Java

If you already have Java installed on your machine, skip the next two steps. Install the Java Runtime Environment (JRE) and the Java Development Kit (JDK) using the following commands.

```terminal
    sudo zypper update
    sudo zypper install java-10-openjdk
    sudo zypper install java-10-openjdk-devel
```

## Step Install Maven

[Maven](https://maven.apache.org/) can be used to help manage dependencies, build, test and run your Java project.

```terminal
  wget -c https://downloads.apache.org/maven/maven-3/3.6.3/binaries/apache-maven-3.6.3-bin.zip
unzip apache-maven-3.6.3-bin.zip
sudo mv apache-maven-3.6.3 /opt/maven
```

In your favorite text editor, add the following to a new file, maven.sh, in /etc/profile.d/.

```terminal
export M2_HOME=/opt/maven/
export M2=$M2_HOME/bin
export PATH=$M2:$PATH
```

Next, refresh your terminal session or create a new one, and check your versions:

```terminal
mvn -v
```

# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/java/Unix-based#start-writing-apps-with-java-and-azure-sql) to complete the tutorial.
