[![Go Report Card](https://goreportcard.com/badge/github.com/Baianoware/trouxa)](https://goreportcard.com/report/github.com/Baianoware/trouxa)
![Github Action Card](https://github.com/Baianoware/trouxa/actions/workflows/go.yml/badge.svg)
# Trouxa

It is an easy and fast tool to install your packages with just one command.

### What means "Trouxa"?

In portuguese, Trouxa means something like a "bundle of clothes", but it is also a pejorative term like mug or muggle. 
Thinking in the first meaning, why not let this bundle carry Packages too? (It does not make sense, I know lol)

### Why use it?
With trouxa, you can install many packages just noting it in a file, `packages.txt`, and selecting your package manager.

Example of `packages.txt`
```text
python
vim
nano
```
and run it selecting the path of `packages.txt`
```sh
trouxa -m pacman -p ~/packages.txt
```
If the file has the same name as `packages.txt` and in the actual execution directory, you do not need to specify it because this is the default value for `-p`.
```sh
trouxa -m pacman
```
You can use also a remote packages' list to install you tools
```sh
trouxa -m pacman -p https://pastebin.com/raw/ysHUVswx
```


## Package managers supported
- pacman
- yay
- apt
- apk

## How to install

### Build
To build you will need the `Go` environment in version `1.18` and `make` installed.
```sh
make build
```

### Install
```sh
sudo make install 
```
The binary compiled will be copied to your `/usr/bin` and available through the command trouxa in your terminal

## Features 

- Portable
- Easy
- Fast
- Simple

## Dump

### apk
You can list all packages from a Alpine container, as exemple, and use this "dump" to install in another:

```sh
apk info > packages.txt
```

### pacman
The same is true to `pacman`.

```sh
pacman -Qe | sed -e 's/\s.*$//' > packages.txt
```

### Extras

- https://github.com/golang-standards/project-layout
