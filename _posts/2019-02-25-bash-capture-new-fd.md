---
date: 2019-02-25
title: Interactive bash functions: returning text on FD3
---

When writing interactive bash scripts, you often want to write functions that a) interact with the user and b) return data to the caller.

For example, you may write a function that renders an interactive menu, allows the user to choose one or more items, and then returns a list
of selected items to the caller.

Typically bash functions return data by writing to stdout, but interactive functions use stdout to talk to the user,
so they need to return data on another file descriptor.  By convention, I write output to FD3.

This is the redirection shuffle that captures FD3 from a function, leaving stdout connected to the terminal.

```
# Setup for interactive UI at the top of script
exec 4>&1

# Capture FD3 into variable; redirect function's stdout to TTY stdout
results="$( { get-values-from-user 1>&4 ; } 3>&1 )"

# Cleanup (shouldn't be necessary if the process is exiting)
exec 4>&-
```

1. Save terminal stdout as FD4.
2. Invoke function within a variable capture.  This creates a context with a new stdout that is captured into the variable.
3. Redirect function's stdout to the tty via FD4.
4. Pass a valid FD3 to the function, redirecting its output to the variable capture.
