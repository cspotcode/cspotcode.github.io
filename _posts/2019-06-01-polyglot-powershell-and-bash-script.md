---
date:   2019-06-01
title:  "Polyglot PowerShell and bash script"
# excerpt: ""
---

I wanted to put some bash and some PowerShell into the same script file, so if I ran it in bash it would do one thing,
and if I ran in PowerShell it would do something else.  Here's what I came up with:

```
#!/bin/bash
echo `# <#`

# Bash goes here
set -euo pipefail

ls -al doing bash stuff

exit
#> > $null

# PowerShell goes here

$val = (
    'powershell',
    'stuff',
    'here'
)

foreach($v in $val) {
    # yadda yadda
}
```

The magic is all here:

```
echo `# <#`
...
#> > $null
```

On bash, it executes what's between the backticks, which is just a comment.  So it succeeds silently and passes through to the bash script.

On PowerShell the backtick is an escape character, so `\`#` is the string `"#"`.  It's followed by a multiline comment wrapped in `<# #>`,
so PowerShell skips over all of our bash code.  We redirect to `$null` to suppress the echoed `#`.
