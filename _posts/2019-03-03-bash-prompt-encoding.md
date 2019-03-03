---
date:   2019-03-03
title:  "Bash prompt zero-width spans"
# excerpt: ""
---

I had a lot of trouble setting up a bash prompt that renders within an external function, rather than inline within the PS1 string.
Zero-width sequences from this external function were instead rendering `\[` and `\]` on the console.

```
PROMPT_COMMAND=prompt
function prompt {
    status_info="< generate a bunch of status info here>"
}
PS1='$status_info'
> '
```

Zero-width byte sequences in PS1 need to be wrapped with `\[` and `\]` so that bash can compute the correct visual width of the prompt.

It wasn't immediately clear that bash special-cases `\[` and `\]` character sequences within PS1, converting them into bytes 1 and 2.
External variables and functions being interpolated into the PS1 string need to contain 1 and 2 byte values, not the 2-character `\[` `\]`
sequences.

https://unix.stackexchange.com/questions/504150/emit-zero-width-bash-prompt-sequence-from-external-binary
