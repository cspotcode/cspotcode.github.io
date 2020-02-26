---
date:   2020-02-26
title:  "Configuring a basic TypeScript composite project"
# excerpt: ""
---

Sometimes I forget the cleanest way to configure a tsconfig composite project with separate tsconfigs for src and test.  Here's the
latest template that I use.  `tsc --showConfig` is helpful to validate that it's doing the right thing.

### tsconfig.json

```
{
    // If in a monorepo, extend from the base config
    "extends": "../../tsconfig.json",
    "references": [
        {
            "path": "src"
        },
        {
            "path": "test"
        }
    ],
    // Exclude references and emitted output
    "exclude": ["src", "test", "dist"],
    "compilerOptions": {
        "composite": true,
        "rootDir": ".",
        
        // We can't pass "noEmit" in incremental mode, but this is close enough.
        "outDir": ".tmp/tsc/root",
        "emitDeclarationOnly": true
    }
}
```

### src/tsconfig.json

```
{
    "extends": "../tsconfig.json",
    "exclude": [],
    "compilerOptions": {
        "rootDir": ".",
        "outDir": "../dist",
        "emitDeclarationOnly": false
    }
}
```

### test/tsconfig.json

```
{
    "extends": "../tsconfig.json",
    "references": [
        {
            "path": "../src"
        }
    ],
    "exclude": [],
    "compilerOptions": {
        "rootDir": ".",
        "outDir": "../.tmp/tsc/test"
    }
}
```

