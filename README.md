# Genny Control Centre

A Command Line Tool used to interact with a local genny system to assist in development.

Written in Golang, this tool aims to be a simple, user friendly way of interacting with a local Genny system, and possibly in future with remote Genny systems.

&nbsp;

## Installing

&nbsp;

#### Download the zipped executable: 

```gctl-bin.zip```

https://github.com/genny-project/genny-control-centre/releases

&nbsp;

#### Unzip and add executable to path

```shell
unzip gctl-bin.zip
cd gctl-bin
./install.sh
```

&nbsp;

## Setup

&nbsp;

#### Add env locations to .bashrc (or .zshrc)

```shell
# GENNY Env Vars
export GENNY_HOME="${HOME}/projects/genny"
export GENNY_MAIN="${GENNY_HOME}/genny-main"
export ENV_FILE="${HOME}/projects/genny/genny-main/genny.env"
export GENNY_ENV_FILE="${HOME}/.genny/.env"

source $GENNY_ENV_FILE
```

&nbsp;

#### Create your central .env file

```shell
touch $HOME/.genny/.env
```

Remember to add all important Genny Envs to this file.

&nbsp;

#### Run the help command to get started

```shell
gctl help
```

&nbsp;

## Build from source

#### Ensure your you have Golang installed on your device - https://go.dev/

&nbsp;

#### Clone genny-control-centre into genny folder.

```shell
cd $HOME/projects/genny
git clone https://github.com/genny-project/genny-control-centre.git
```

&nbsp;

#### Install Packages and Build from source.

```shell
cd genny-control-centre
./packages.sh
./build.sh
```

&nbsp;

#### Add executable to path

```shell
./install.sh
```


&nbsp;

## Examples

### Token

```shell
gctl get token
```

### Cache

```shell
gctl read cache PER_USER
```

```shell
gctl watch db PER_USER
```

### Genny

```shell
gctl status
```

```shell
gctl clone
```

```shell
gctl pull
```

```shell
gctl build
```

```shell
gctl start
```

```shell
gctl stop
```

```shell
gctl restart
```
