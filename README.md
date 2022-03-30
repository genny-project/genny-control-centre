# Genny Control Centre

A Command Line Tool used to interact with a local genny system to assist in development.

&nbsp;

## Installing

* **Ensure your you have Golang installed on your device** - https://go.dev/

* **Clone genny-control-centre into genny folder.**
```shell
cd $HOME/projects/genny
git clone https://github.com/genny-project/genny-control-centre.git
```
* **Install Packages and Build from source.**
```shell
cd genny-control-centre
./packages.sh
./build.sh
```
* **Add executable to path**
```shell
./install.sh
```
* **Run the help command to get started**
```shell
gctl help
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
