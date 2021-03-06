---
date:   2018-08-05
title:  "Connect WSL Docker CLI to Docker for Windows"
# excerpt: "
---

Use case: You want to run docker commands from WSL bash.  You're running Windows and have Docker for Windows installed.

For example, most of my team at work runs Mac and Linux, so all our scripts are written against Unix / bash.  I work within WSL on a Windows host,
so I have Docker for Windows installed.

As far as I know, you can't run a docker engine within WSL, but you *can* install and run docker's CLI.
So you simply need to configure WSL's docker CLI to talk to Windows' Docker engine.  The Windows Docker engine accepts connections
on a Windows named pipe, which WSL processes cannot access.  Fortunately, there's a program --
[`npiperelay.exe`](https://github.com/jstarks/npiperelay) -- which can create a bridge between a Windows
named pipe and a Unix domain socket.  Connections to the Unix socket are proxied to the Windows named pipe.

*NOTE: my username in both WSL and Windows is "abradley"; substitute your own.*

## Installation

1. Install docker-ce, go, and socat in WSL Ubuntu
1. Install npiperelay and docker-relay script

### Install docker-ce, go, and socat in WSL Ubuntu

*I use Ubuntu; I'm sure other distros will work, too*

Docker installation is documented on [Docker's website](https://docs.docker.com/install/linux/docker-ce/ubuntu/).

Go installation is documented on [Go's website](https://github.com/golang/go/wiki/Ubuntu).

socat is installed with apt-get: `sudo apt-get install socat`

### Install npiperelay and docker-relay script

[npiperelay](https://github.com/jstarks/npiperelay)

Use `go get` to check out npiperelay source code into "/home/abradley/go/src/github.com/jstarks/npiperelay"
Use `go build` to build the Windows version of npiperelay and saves it at "$HOME/bin/npiperelay.exe" *in Windows*. (not WSL)

```bash
# In WSL bash
go get -d github.com/jstarks/npiperelay
GOOS=windows go build -o /mnt/c/Users/abradley/bin/npiperelay.exe github.com/jstarks/npiperelay
```

npiperelay includes a script to setup a docker pipe.  Copy it to `$HOME/bin` *in WSL* and set
execute permissions.

```bash
# In WSL bash
sudo adduser abradley docker
mkdir -p $HOME/bin
cp /home/abradley/go/src/github.com/jstarks/npiperelay/docker-relay $HOME/bin/
chmod +x $HOME/bin/docker-relay
```

... except I had issues running `docker-relay`, because socat running as root is unable to find "npiperelay.exe".  So I modified it to use an absolute path:

```bash
#!/bin/sh

# Notice the full path to npiperelay.exe
exec socat UNIX-LISTEN:/var/run/docker.sock,fork,group=docker,umask=007 EXEC:"/mnt/c/Users/abradley/bin/npiperelay.exe -ep -s //./pipe/docker_engine",nofork
```

...and I created another script, `start-docker-relay`, to spawn the first in the background:

```bash
#!/bin/bash
sudo bash -c "$HOME/bin/docker-relay 1>/dev/null 2>/dev/null & disown"
```

*NOTE: I have `$HOME/bin` in my `$PATH`.*

## Usage

Start the pipe relay.

```bash
start-docker-relay
```

At this point `docker` commands should work.  They'll attempt to connect to unix socket `/var/run/docker.sock`, which is hosted by `socat`.  Socat will start and pipe the connection to `npiperelay.exe`, which connects to Windows named pipe `//./pipe/docker_engine`.

*NOTE: If you've previously configured DOCKER_HOST or other docker environment variables, you'll probably have to undo that.  This setup relies on docker's default behavior, connecting to Unix socket /var/run/docker.sock*

## Misc

Docker for Windows has an option to accept TCP connections without any encryption.  This is less secure than the named pipe.

There are tutorials for setting up Windows Docker to use TLS.  These are for Windows server and do not work for a
developer machine.
