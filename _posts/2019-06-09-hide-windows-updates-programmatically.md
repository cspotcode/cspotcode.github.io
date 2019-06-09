---
date:   2019-06-09
title:  "Hide Windows Updates Programmatically"
# excerpt: ""
---

My company gives us Office 2010, but I have my personal copy of 2019 installed instead.
For some reason, Windows Update keeps proposing Office 2010 updates which always fail to install.

[`PSWindowsUpdate`](https://www.powershellgallery.com/packages/PSWindowsUpdate/) is a PowerShell module
for interacting with Windows Update interactively or programmatically.  It allows hiding windows updates.

```
install-module -scope currentuser pswindowsupdate
hide-windowsupdate
```

This will query Windows Update for updates then prompt you to hide all of them one by one.
I had to run it a bunch of times, since each batch of update hiding allowed Windows Update
to find others, probably because at this point there are months of Office 2010 updates that
don't apply to me.
