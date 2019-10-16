# Command-line client for managing Insights operator

## Description

A simple CLI client for managing the Insights operator. Currently this client supports just basic operations to retrieve and change configuration of operator on selected cluster.

## Supported commands

* `exit`
* `quit`
* `bye`
  * exit from the client

## How to build the CLI client

Use the standard Go command:

```
go build
```

This command should create an executable file named `insights-operator-cli`.

## Start

Just run the executable file created by `go build`:

```
./insights-operator-cli
```

## Configuration

No further configuration is needed at this moment.
