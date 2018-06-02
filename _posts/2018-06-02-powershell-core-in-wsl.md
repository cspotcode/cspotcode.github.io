---
layout: post
title:  "Running pwsh in WSL via ConEmu"
date:   2018-06-02
# excerpt: "Running pwsh in WSL via ConEmu"
# image: "/images/pic02.jpg"
---

I was setting up [ConEmu](https://conemu.github.io/) on Windows to run Powershell Core within my WSL Ubuntu installation.  The latest ConEmu includes [wslbridge](https://github.com/rprichard/wslbridge) and is preconfigured to launch `bash` like this:

```bash
set "PATH=%ConEmuBaseDirShort%\wsl;%PATH%" & %ConEmuBaseDirShort%\conemu-cyg-64.exe --wsl -cur_console:pm:/mnt
```

Behind the scenes this runs wslbridge, which in turn launches the default shell (`bash`) in a TTY.  TTY is only on by default when running the default shell.  If we specify a custom executable, we must also specify `-t` to enable the TTY.  Putting it together, plus a few nice `pwsh` args, we append `-t pwsh -NoLogo -ExecutionPolicy RemoteSigned`.

```bash
set "PATH=%ConEmuBaseDirShort%\wsl;%PATH%" & %ConEmuBaseDirShort%\conemu-cyg-64.exe --wsl -cur_console:pm:/mnt -t pwsh -NoLogo -ExecutionPolicy RemoteSigned
```