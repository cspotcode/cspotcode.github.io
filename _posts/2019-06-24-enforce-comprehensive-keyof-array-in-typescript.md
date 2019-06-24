---
date:   2019-06-24
title:  "Enforce comprehensive keyof array in Typescript"
# excerpt: ""
---

I had a situation where I wanted to ensure an array of `keyof T` included *all* keys of that interface.
This could be useful in situations where you need to keep some runtime type information in sync with a compile-time interface.

Here's the exact code in question, but this trick could be applied to other situations.  In this situation, I want to normalize a `notifier` object.
It may or may not contain any number of callbacks; I need to replace all the missing ones with `noop` functions.  Otherwise I'll get
runtime errors when I try to fire those callbacks elsewhere.

```
function createNotifier<T, Key extends keyof T>(n: T | undefined): never;
function createNotifier<T, Key extends keyof T>(n: T | undefined, ...keys: Array<Key>):  keyof T extends Key ? Required<NonNullable<T>> : never;
function createNotifier(n: any, ...keys: any) {
    const n2 = n ? n : {} as any;
    for(const key of keys) {
        if(!n2[key]) n2[key] = noop as any;
    }
    return n2 as any;
}
```

The function accepts an array of `keyof T`.
To *prove* that this array is comprehensive (doesn't omit any properties) I use a conditional type in the return.  Basically I flip the extends backwards: `keyof T extends Key` instead of `Key extends keyof T`.  If that check fails, the function returns `never` which is sure to show up as an error elsewhere in my code.  I suppose I could also return `"error": "HELPFUL ERROR MESSAGE"}`

Multiple function signatures are necessary to account for the case where zero `keys` are passed to the function.  In that case,
inference sets `Key` equal to `keyof T`, which allows our validation in the second function signature to erroneously pass.  We
avoid that by adding the first function signature.  If zero `keys` are passed, the function returns `never`.
