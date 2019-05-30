---
date:   2019-05-30
title:  "Attach VSCode to container from CLI"
# excerpt: ""
---

VSCode's new "Remote development" features let you attach the editor to a remote system, for example, WSL, an SSH server, or a docker container.

The WSL integration lets you run `code-insiders .` from WSL, and it opens the current directory in a VSCode window automatically connected to WSL.  However, it wasn't obvious how to do the same for attaching to a docker container.  After reading some of the code, I figured it out:

```
auth="attached-container+$(node -p 'Buffer.from(process.argv[1]).toString("hex")' "$containerId")"
code-insiders --remote="$auth" --folder-uri="vscode-remote://$auth/path/inside/container"
```

## How it works:

As far as I can tell, VSCode lets you specify a `--remote` flag of the form `<authority>+<target>`

`<authority>` is one of `dev-container`, `attached-container`, or `wsl`.
`<target>` depends on the type.  For WSL, a target of `default` will attach to the `default` WSL distro.
For attaching to a container, the target should be the hex-encoded container ID.  We can compute this using a node one-liner.

Technically, it looks like the `<target>` can also be a JSON object converted to hex, where the JSON object contains multiple fields.  I think this is used for the `dev-container` authority to specify both a local directory and a docker host.

To attach using a `devcontainer` configuration, the `<authority>` is `dev-container` and the `<target>` should be one of the following hex-encoded:
* path of the project on the host
* JSON of the form:

```
{hostPath: "/home/me/whatever", dockerHost:"I assume a docker host domain name and port goes here"}
```
