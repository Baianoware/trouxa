# Trouxa

It is an easy and fast tool to install your packages with just one command.

### What means "Trouxa"?

In portuguese, Trouxa means something like a "bundle of clothes", but it is also a pejorative term like mug, gink, muggle. 
Thinking in the first meaning, why not let this bundle carry Packages too? (It does not make sense, I know lol)

### Why use it?
With trouxa, you can install many packages just noting it in a file, `packages.txt`, and selecting your package manager.

If you are using the a Arch based distro like me, you just need create a file with your needs, and run a command.

```text
python
vim
nano
and many others packages....
```
and run it on the same directory as the `packages.txt`
```sh
trouxa pacman -p packages.txt
```
If the file is with the name `packages.txt`, you cannot specify it because this is the default value of `-p`
```sh
trouxa pacman
```


### Package managers supported
- pacman
- yay

## How to install

### Build
To build you will need the `Go` environment in version `1.16` and `make` installed.
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

### Extras

- https://github.com/golang-standards/project-layout
