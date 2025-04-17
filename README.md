# SEBAS

Sebas is a general purpose tool which aims to help (me at least) automate tasks per application project.
I mainly work on projects which are based of framework with dotenv and cli commands. The main goal is to allow me to switch configuration as I please without having the comment/uncomment/modify variables or write commands by hand

> Why Sebas ? Because I need my own personal butler :D


## Disclaimer

This tool is not intended for **PRODUCTION** usage.

## Features

### Mod Context

> TBC: is it something we need ?

Context would be a top level abstraction which could be used by other mod \
in order to target the correct project and the correct env

### Mod Environment Variable Manager

1. Create, Read, Update, Delete environment variables from a `.env` file
    - [ ] Read `.env` file
    - [ ] Read a variable
    - [ ] Create a variable
    - [ ] Update a variable
    - [ ] Delete a variable

2. Save preset per project and environment (prod, staging, dev, etc...)
    - [ ] Create new preset
      - [ ] Create preset from existing `.env` file
    - [ ] Read preset
    - [ ] Update preset
    - [ ] Delete preset

### Mod Command Launcher

- [ ] Run a command
- [ ] Save a command for reuse
- [ ] Save a collection of command per project and environment

### Mod Container Registry

- [ ] Connect to a registry
- [ ] List packages
- [ ] Push a package
- [ ] Delete a package

## Development
### Install

```sh
make
```

or

```sh
make install
```

### Run tests

```sh
make test
```

### Launch Desktop application

```sh
make desktop
```
