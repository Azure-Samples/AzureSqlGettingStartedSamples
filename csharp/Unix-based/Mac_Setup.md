# Machine Setup for MacOs

This page will walk you through everything you need to setup and run the C# samples on your MacOs machine.

## Install .NET Core

Follow the instructions from [**this site**](https://docs.microsoft.com/en-us/dotnet/core/install/macos) to install .Net Core on your Mac Machine.


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

> You have successfully installed .NET Core on your Mac machine, and authenticated to Azure. You now have everything you need to start writing your C# apps with Azure SQL!


# Proceed with the tutorial by following the remaining directions [**here**](https://github.com/Azure-Samples/AzureSqlGettingStartedSamples/blob/master/csharp/Unix-based/Readme.md)
