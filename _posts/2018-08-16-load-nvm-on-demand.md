---
date:   2018-08-16
title:  "Load nvm on-demand"
# excerpt: ""
---

nvm startup is slow on WSL.  This is WSL's fault, not nvm's.  All the process forking in nvm's startup is slow.  My work machine's corporate virus scanner might be exacerbating the problem.

Regardless, it's easy to load nvm on-demand via bash profile.  I rarely invoke the nvm command; I just need node and npm in my $PATH.

Default nvm loader:

```bash
# nvm puts this in your .bashrc
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
```

Change it to this:

```bash
export NVM_DIR="$HOME/.nvm"
function nvm {
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
    [ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
    nvm "$@"
}
PATH="/home/abradley/.nvm/versions/node/$( cat /home/abradley/.nvm/alias/default )/bin:$PATH"
```

There might be some corner-cases not handled when I set the PATH, but it works for me.
