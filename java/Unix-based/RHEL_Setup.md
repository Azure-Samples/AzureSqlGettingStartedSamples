# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Java and connect to Azure SQL.

## Install Java

If you already have Java installed on your machine, skip the next two steps. Install the Java Runtime Environment (JRE) using the following command.

```terminal
  sudo rpm --import http://repos.azul.com/azul-repo.key
  sudo curl http://repos.azul.com/azure-only/zulu-azure.repo -o /etc/yum.repos.d/zulu-azure.repo
  sudo yum -q -y update
  sudo yum -q -y install zulu-11-azure-jdk
```

## Install Maven

[Maven](https://maven.apache.org/) can be used to help manage dependencies, build, test and run your Java project.

```terminal
  cd /opt
sudo wget https://www-eu.apache.org/dist/maven/maven-3/3.6.3/binaries/apache-maven-3.6.3-bin.tar.gz
sudo tar xzf apache-maven-3.6.3-bin.tar.gz
sudo ln -s apache-maven-3.6.3 maven
```


Now, create a maven.sh file in /etc/profile.d/maven.sh, and put the following into it:

```terminal
export M2_HOME=/opt/maven
export PATH=${M2_HOME}/bin:${PATH}
```

Load the file, and check the maven version:

```terminal
source /etc/profile.d/maven.sh
mvn -v
```
> You have successfully installed Java and Maven on your Red Hat machine. You now have everything you need to start writing your Java apps with Azure SQL DB!

# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/java/Unix-based#start-writing-apps-with-java-and-azure-sql) to complete the tutorial.
