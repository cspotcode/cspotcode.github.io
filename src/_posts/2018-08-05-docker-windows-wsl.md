---
date:   2018-08-05
title:  "Connect WSL Docker CLI to Docker for Windows"
# excerpt: "
---

Use case: You want to run docker commands from WSL bash.  You're running Windows and have Docker for Windows installed.

For example, most of my team at work runs Mac and Linux, so all our scripts are written against Unix / bash.  I work within WSL
but have Docker for Windows installed.  As far as I know, you can't run a docker engine within WSL, but you *can* install and run docker's CLI.

*NOTE: my username in both WSL and Windows is "abradley"; substitute your own.*

## Installation 
1. Install docker-ce, go, and socat in WSL Ubuntu
1. Install npiperelay and docker-relay script

### Install docker-ce, go, and socat in WSL Ubuntu

Docker installation is documented on [Docker's website](https://docs.docker.com/install/linux/docker-ce/ubuntu/).

Go installation is documented on [Go's website](https://github.com/golang/go/wiki/Ubuntu).

socat is installed with apt-get: `sudo apt-get install socat`

### Install npiperelay and docker-relay script

[npiperelay](https://github.com/jstarks/npiperelay)

```bash
# In WSL bash
go get -d github.com/jstarks/npiperelay
GOOS=windows go build -o /mnt/c/Users/abradley/bin/npiperelay.exe github.com/jstarks/npiperelay
```

`go get` checks out the npiperelay source code into /home/abradley/go/src/github.com/jstarks/npiperelay  
`go build` builds the Windows version of npiperelay and saves it at $HOME/go/bin/npiperelay.exe *in Windows*. (not WSL)

```bash
# In WSL bash
sudo adduser abradley docker
mkdir -p $HOME/bin
cp /home/abradley/go/src/github.com/jstarks/npiperelay/docker-relay $HOME/bin/
chmod +x $HOME/bin/docker-relay
```

... except that I had issues running `docker-relay` because socat running as root was unable to find "npiperelay.exe"  So I modified it to hard-code the path:

```bash
#!/bin/sh

exec socat UNIX-LISTEN:/var/run/docker.sock,fork,group=docker,umask=007 EXEC:"/mnt/c/Users/abradley/bin/npiperelay.exe -ep -s //./pipe/docker_engine",nofork
```

...and I created another script, `start-docker-relay`, to spawn the first in the background:

```bash
#!/bin/bash
sudo bash -c "$HOME/bin/docker-relay 1>/dev/null 2>/dev/null & disown"
```

*NOTE: I have `$HOME/bin` in my `$PATH`.*

## Usage

Start docker relay.

```bash
start-docker-relay
```

At this point `docker` commands should work.  They'll attempt to connect to unix socket `/var/run/docker.sock`, which is hosted by `socat`.  Socat will start and pipe the connection to `npiperelay.exe`, which connects to Windows named pipe `//./pipe/docker_engine`.

*NOTE: If you've previously configured DOCKER_HOST or other docker environment variables, you'll probably have to undo that.  This setup relies on docker's default behavior, connecting to Unix socket /var/run/docker.sock*