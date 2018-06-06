---
date:   2018-06-06
title:  "Accurately typing Object.assign in TypeScript"
# excerpt: "Pow
---

`Object.assign({}, a, b)` -- or any other code doing essentially the same thing -- can be thought to merge `a` and `b`, so naively the return type is `A & B`.  But that's not accurate when `b` is overwriting properties of `a` with values of different types.

For example:

```
const a = {foo: 'hello', bar: 'world'};
const b = {bar: false};
const c = Object.assign({}, a, b);
// c's type should be {foo: string; bar: boolean}
```

Obviously, whether or not you really want to be doing this sort of thing depends on the situation.

Regardless, here's how I accurately type the return value:
```
type Overlay<A, B> = Pick<A, Exclude<keyof A, keyof B>> & Pick<B, keyof B>;

function overlay<A, B>(a, b) {
    return Object.assign({}, a, b) as Overlay<A, B>;
}
```

This should also be useful for modeling certain React HOCs, since some appear to *replace* props of the wrapped component with new semantics and types.
