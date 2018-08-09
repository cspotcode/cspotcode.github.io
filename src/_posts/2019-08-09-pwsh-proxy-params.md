---
date:   2018-08-09
title:  "PowerShell command mimic params of another"
# excerpt: "
---

I wanted to write a PowerShell command that wraps another, accepting (almost) all the same params and passing them to the target, sorta like a command proxy or decorator.  It would be *really* sweet if tab completion magically worked, too.  I got it to work!

The magic here is not passing params to another command.  That's easy: `other-command @PSBoundParameters`  The magic is using `DynamicParams {}` to declare
that our command has all the same parameters of the target, so that tab completion and validation do the right thing.  We `Get-command` the target command,
get its list of parameters via Powershell's runtime metadata / reflection, then declare a set of matching params.

This is still a work-in-progress, but it seems to do the trick.  I haven't verified if it supports parameter sets, and positional params probably need to have their offsets bumped.

## Usage

```powershell
$ ./proxy.ps1 -command set-location - # <-- tab completion here shows list of params for set-location, e.g. -Path
```

## `proxy.ps1`
```powershell
[cmdletbinding()]
param(
    <# Passthrough to this "target" command (for demonstration) #>
    [string] $command
)
dynamicParam {
    # Skip params that are already declared (for example, a bunch are already introduced via CmdletBinding)
    $skipParams = $MyInvocation.MyCommand.Parameters.Keys
    $cmd = get-command $command
    $RuntimeParamDic = [System.Management.Automation.RuntimeDefinedParameterDictionary]::new()
    # For each parameter of target command...
    foreach($pair in $cmd.parameters.GetEnumerator()) {
        $name = $pair.key
        $param = $pair.value
        if(-not ($skipParams -contains $name)) {
            # Declare a runtime parameter matching the target command's parameter
            $runtimeParam = [System.Management.Automation.RuntimeDefinedParameter]::new($name, $param.parametertype, $param.attributes)
            $RuntimeParamDic.add($name, $runtimeParam)
        }
    }
    $RuntimeParamDic
}

process {
    # Filter params in some way.  For example, we cannot pass '-command'
    # to our target.
    $noPassParams = @('command')
    $passParams = @{ }
    foreach($pair in $PSBoundParameters.GetEnumerator()) {
        if(-not ($noPassParams -contains $pair.key)) {
            $passParams.$( $pair.key ) = $pair.Value
        }
    }

    # Invoke the target command
    & $command @passParams
}
```