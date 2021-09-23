<div align="center">
  <img alt="gt-tweeter Logo" src="https://github.com/marcusolsson/gophers/raw/master/orbiter-gopher.png?v=3&s=300" height="350" />
  <h3 align="center">gt-tweeter</h3>
  <p align="center">A helper CLI that facilitates to manage Twitter account from terminal</p>

</div>

---

[![asciicast](https://asciinema.org/a/JASrDapMHcE40P5OvXFEQZMyt.svg)](https://asciinema.org/a/JASrDapMHcE40P5OvXFEQZMyt)

---

## Installation

### Go

If you have Go 1.16+, you can directly install by running:

```shell
$ go install github.com/GoTurkiye/training/104-cli/gt-tweeter@latest
```

and the resulting binary will be placed at **_$HOME/go/bin/gt-tweeter_**.

## Quick Start

```shell
$ gt-tweeter --help
gt-tweeter is a CLI library for Go that facilitates to use Twitter from your terminal.
It basically shows account information, list and get followers, or event you can send tweet, and
display the status of the tweet.

Usage:
  gt-tweeter [command]

Available Commands:
  account     Helps you to view account information
  completion  Generate completion script
  follower    Helps you to list and get followers
  help        Help about any command
  tweet       Helps you to send tweet and show the status of the tweet

Flags:
  -c  --config string   config file (default is $HOME/.gt-tweeter.yaml)
  -h, --help            help for gt-tweeter

Use "gt-tweeter [command] --help" for more information about a command.
```