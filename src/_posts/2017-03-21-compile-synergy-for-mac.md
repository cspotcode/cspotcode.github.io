---
layout: post
date:   2017-03-21
title:  "Compile Synergy for Mac"
excerpt: "How to compile Synergy KVM software from source on Mac"
# image: "/images/pic02.jpg"
---

```bash
brew install cmake
# install Qt from web installer

# add Qt to PATH (e.g. /Users/abradley/Qt/5.8/clang_64/bin)
# Test it
which qmake
qmake --version

git clone https://github.com/cspotcode/synergy
cd synergy
git checkout no-license-manager

# Edit ./ext/toolchain/commands1.py; change "frameworkRootDir" to be the path to your Qt lib directory (e.g. /Users/abradley/Qt/5.8/clang_64/lib)

./hm.sh conf -g 1 --mac-sdk 10.12 --mac-identity Sierra
./hm.sh build

# Mac app is at ./bin/Synergy.app

# If you make a mistake or need to start over, clean up first:
./hm.sh clean
rm -r ./build
rm -r ./bin
```
