---
date:   2020-01-29
title:  "Send variables to bash session"
# excerpt: ""
---

Sometimes a shell helper really needs to set variables in your shell session.  An external executable can't do this, so these functions need
to live inside your bash shell.

One way around this is to write the helper function in a real language, but put a tiny adapter in your shell.

```
# In your shell profile
helper-function() {
    { eval "$( ./my-helper-function-written-in-a-real-language.ts "$@" 3>&1 1>&4 )" ; } 4>&1
}
```

This adapter will call an external executable but capture instructions from file descriptor 3.  It'll then eval them in the shell.
The external executable can, at its discretion, write valid shell syntax to file descriptor 3, for example to modify the PATH.

```
#!/usr/bin/env ts-script
console.log('Interact via stdout and stdin because those FDs are passed verbatim.');
const environment = prompt('Which environment should be added to PATH?'); // pseudocode, but you get the idea
fs.writeSync(3, `
PATH=${ shellEscape(`${ Path.join(os.homedir(), '.envs', environment) }:${ process.env.PATH }`) }
`);
```

The magic is the switcheroo between file descriptors 3 and 1.  Bash `$()` syntax passes a new stdout file descriptor,
whose output is captured, to the child process.  Instead, we want to capture file descriptor 3, passing through the original
stdin, stdout, and stderr as file descriptors 0-2.
