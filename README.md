# Automation test with Golang

This repository for automation test with Golang. 

## Installation

Use the package manager [Makefile](https://en.wikipedia.org/wiki/Make_(software)) to install Selenoid.

```bash
make docker-clean
make pull-chrome
make pull-firefox
make selenoid-hub
make selenoid-ui
```

## Example Run Test

```bash
$ go test -test.run=TestPageCura
```
## Codebase Structure (Coming Soon)
