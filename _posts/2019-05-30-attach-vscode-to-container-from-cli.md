---
date:   2019-05-30
title:  "Attach VSCode to container from CLI"
# excerpt: ""
---

VSCode's new "Remote development" features let you attach the editor to a remote system, for example, WSL, an SSH server, or a docker container.

The WSL integration lets you run `code-insiders .` from WSL, and it automatically opens a new VSCode window connected to WSL.  However,
I want to do the same for attaching to a docker container.  After reading some of the code, I figured it out:

```
auth="attached-container+$(node -p "Buffer.from(process.argv[1]).toString('hex')" "$containerId")"
code-insiders --remote="$auth" --folder-uri="vscode-remote://$auth/path/inside/container"
```

How it works:

As far as I can tell, VSCode lets you specify a `--remote` flag of the form `<authority>+<target>`

`<authority>` is one of `dev-container`, `attached-container`, or `wsl`.
`<target>` depends on the type.  For WSL, a target of `default` will attach to the `default` WSL distro.
For attaching to a container, the target should be the hex-encoded container ID.  We can compute this using a node one-liner.
