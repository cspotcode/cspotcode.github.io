---
date:   2018-06-03
title:  "Powershell Script Boilerplate"
# excerpt: "Pow
---

I'm a big fan of PowerShell, but scripts can require some annoying boilerplate.  Here's what I'm using lately for new scripts:

```powershell
#!/usr/bin/env pwsh
param(
    # cli args declared here
)

function main() {
    # Script goes here

    # To call external binaries and fail on non-zero exit codes, wrap in `exec {}`:
    exec { git merge foo/bar }
}

function exec($block) {
    & $block
    if($LASTEXITCODE -ne 0) { throw "Non-zero exit code $LASTEXITCODE" }
}

$ErrorActionPreference = 'Stop'
try {
    $_pwd = $pwd
    cd $PSScriptRoot
    main
} finally {
    cd $_pwd
}
```

## Breaking it down

The shebang is for Linux; does nothing on Windows.

Setting `$ErrorActionPreference` is like bash's `set -e` but for PowerShell commands.  When a command fails, it throws an error, effectively ending the script unless caught.
It can be overridden on a command-by-command basis with the `-ErrorAction` flag.

`try {} finally` allows us to `cd` to the script's directory and guarantee we'll go back upon termination.  PowerShell runs scripts a bit differently
from bash.  In bash, the default is to spawn a new process unless you explicitly dot-source something.  In PowerShell the default is to run in-process unless you explicitly spawn a new shell.
`cd`ing in a script affects the shell's `$pwd` unless we explicitly `cd` back to where we started.

The `exec` function exists to run external binaries and convert non-zero exit codes into thrown errors.  Again, this is a big difference between bash and PowerShell.  In bash almost every command is an external
process, and non-zero exit codes are interpreted as an error via `set -e`.  In PowerShell most commands are in-process and have proper failure handling -- they throw errors --
but exit codes of external processes are ignored.  PowerShell inherits this quirk from Windows, where exit codes do not carry the significance they do on Unix.
To get bash-like behavior for external processes, we use this wrapper function.  Technically,
it will execute any script block, but by convention we pass a script block that runs a single external command.

## Gotchas

`$PSScriptRoot` is only defined in files with a `.ps1` extension, *even with a shebang, even on Linux.*  I dunno why; the file needs that extension or else that variable will not be set.
