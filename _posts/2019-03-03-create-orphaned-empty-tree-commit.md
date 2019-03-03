---
date:   2019-03-03
title:  "Create an orphaned, empty tree branch in Git"
# excerpt: ""
---

This blurb creates a git branch with a single, parent-less commit, with an empty tree.
In other words, it's an orphan with a commit but no files.

You can achieve this by checking out an orphaned branch and then `commit --allow-empty`.  However,
the blurb below doesn't require you to switch your worktree to a new branch.

I've used this trick when I want to commit a subdirectory into a new, empty branch.  I
pass this branch to `git worktree add -b new-branch dir-name empty-orphan`.

```bash
# Create and save an empty tree object
treehash="$( git hash-object -w -t tree /dev/null )"
# Create a parent-less commit with empty tree
commithash="$( git commit-tree "$treehash" -m "Empty tree" )"
# Create an orphan branch at the new commit
git branch empty-orphan "$commithash"
```
