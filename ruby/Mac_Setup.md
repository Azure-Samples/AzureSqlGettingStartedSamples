# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Ruby and connect to Azure SQL.

## Install Homebrew 

Ruby is already installed on your Mac. If you already have Homebrew on your machine, skip this step. Install Homebrew using the following commands. Once you have installed Homebrew, make sure to restart the terminal session.

```terminal
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

## Install rbenv and ruby-build

If you already have rbenv and ruby-build installed on your machine, skip this step. Use the following commands to install prerequisites for Ruby.

```terminal
brew update
brew install rbenv ruby-build
echo 'if which rbenv > /dev/null; then eval "$(rbenv init -)"; fi' >> ~/.bash_profile
source ~/.bash_profile
```

## Install Ruby

Use the commands below to install Ruby using rbenv and check the version.

```terminal
rbenv install 2.4.0
rbenv global 2.4.0
ruby -v
```

## Install FreeTDS

FreeTDS is a driver that enables you to connect to SQL Server. It is a prerequisite for the connector you'll get later in the tutorial to connect to SQL Server. Run the following commands to install FreeTDS:

```terminal
brew install FreeTDS
```

> You have successfully installed Ruby on your Mac. You now have everything you need to start writing your Ruby apps with Azure SQL DB!


# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/ruby#get-started-with-ruby-and-azure-sql) to complete the tutorial.
