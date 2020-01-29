---
date:   2020-01-29
title:  "Implement bash profile functions as external executables"
# excerpt: ""
---

Sometimes a shell helper needs to set variables in your shell session, for example to modify your PATH.  An external executable can't do this, so these functions need to live inside your bash shell.  Typically this means writing the functions in bash (gross) or dot-sourcing a large, external chunk of bash, which slows shell startup.

One way around this is to write the helper function in a real language and put a tiny adapter in your shell profile.

```shell
# In your shell profile
add-env-to-path() {
    { eval "$( ~/scripts/add-env-to-path.ts "$@" 3>&1 1>&4 )" ; } 4>&1
}
```

This adapter will call an external executable but capture instructions from file descriptor 3.  It'll then eval them in the shell.
The external executable can, at its discretion, write valid shell syntax to file descriptor 3, for example to modify the PATH.

```typescript
#!/usr/bin/env ts-script
// ~/scripts/add-env-to-path.ts
console.log('Interact via stdout and stdin because those FDs are passed verbatim.');
const environment = prompt('Which environment should be added to PATH?'); // pseudocode, but you get the idea
fs.writeSync(3, `
PATH=${ shellEscape(`${ Path.join(os.homedir(), '.envs', environment) }:${ process.env.PATH }`) }
`);
```

The magic is the switcheroo between file descriptors 3 and 1.  Bash `$()` syntax passes a new stdout file descriptor,
whose output is captured, to the child process.  Instead, we want to capture the child process's output to file descriptor 3, passing the original stdin, stdout, and stderr as file descriptors 0-2.
