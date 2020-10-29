# Machine Setup for Ubuntu

This page will walk you through everything you need to setup and run the C# samples on your Ubuntu machine.

## Install .NET Core

If you already have .NET Core installed on your machine, skip this step. Otherwise, install it using the following commands.

```terminal
curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg
sudo mv microsoft.gpg /etc/apt/trusted.gpg.d/microsoft.gpg
sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/microsoft-ubuntu-xenial-prod xenial main" > /etc/apt/sources.list.d/dotnetdev.list'
sudo apt-get update
sudo apt-get install dotnet-sdk-3.1
```

##  Install The Azure CLI and Login to Azure

1.  This page follows instructions from this **[site](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-zypper?view=azure-cli-latest)**.
1.  Install the Azure CLI:

```terminal
sudo zypper install -y curl
sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc
sudo zypper addrepo --name 'Azure CLI' --check https://packages.microsoft.com/yumrepos/azure-cli azure-cli
sudo zypper install --from azure-cli azure-cli
```

You may have to reboot your machine for these libraries to take effect.

1.  Open a new instance of the command window and type az login. Follow the instructions that follow to authenticate.

```terminal
az login
```

**Please Note** You may have to login again to your machine after restart or long periods of inactivity.

At this time, you have authenticated yourself and your machine to Azure, so your application can connect. 

> You have successfully installed .NET Core on your Ubuntu machine, and authenticated to Azure. You now have everything you need to start writing your C# apps with Azure SQL!


# Proceed with the tutorial by following the remaining directions [**here**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/Readme.md)
