# Machine Setup 

In this part of the tutorial, you will install the necessary dependencies to run Ruby and connect to Azure SQL.

## Install prerequisites for Ruby

Use the commands below to install prerequisites.

```terminal
sudo apt-get update
sudo apt-get install autoconf bison build-essential libssl-dev libyaml-dev libreadline6-dev zlib1g-dev libncurses5-dev libffi-dev libgdbm-dev
sudo apt-get install git
```

##  Install rbenv and ruby-build

If you already have rbenv and ruby-build installed on your machine, skip this step. Use the following commands to install prerequisites for Ruby.

```terminal
cd
git clone https://github.com/rbenv/rbenv.git ~/.rbenv
echo 'export PATH="$HOME/.rbenv/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(rbenv init -)"' >> ~/.bashrc
exec $SHELL
git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
echo 'export PATH="$HOME/.rbenv/plugins/ruby-build/bin:$PATH"' >> ~/.bashrc
exec $SHELL
```

##  Install Ruby

Use the commands below to install Ruby using rbenv and check the version.

```terminal
rbenv install 2.7.0
rbenv global 2.7.0
ruby -v
```

## Install FreeTDS

FreeTDS is a driver that enables you to connect to Azure SQL DB. It is a prerequisite for the connector you'll get later in the tutorial to connect to SQL Server. Run the following commands to install FreeTDS:

```terminal
wget ftp://ftp.freetds.org/pub/freetds/stable/freetds-1.00.27.tar.gz
tar -xzf freetds-1.00.27.tar.gz
cd freetds-1.00.27
./configure --prefix=/usr/local --with-tdsver=7.3
sudo make
sudo make install
```

> You have successfully installed Ruby on your Ubuntu machine. You now have everything you need to start writing your Ruby apps with Azure SQL!

# Return to the [**main page**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/tree/master/ruby#get-started-with-ruby-and-azure-sql) to complete the tutorial.

