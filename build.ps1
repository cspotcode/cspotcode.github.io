#!/usr/bin/env pwsh

function main() {
    Remove-Item -r out/*
    Copy-Item -R src/* out
    Set-Location out
    exec { git add -A }
    exec { git commit -m "Site update" }
    exec { git push }
}

function exec($block) {
    & $block
    if($LASTEXITCODE -ne 0) { throw "Non-zero exit code $LASTEXITCODE" }
}

$ErrorActionPreference = 'Stop'
try {
    pushd $PSScriptRoot
    main
} finally {
    popd
}