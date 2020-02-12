---
date:   2020-02-12
title:  "Emulate JavaScript's proposed pipeline operator with a helper function"
# excerpt: ""
---

JavaScript's proposed pipeline operator has been stuck at stage 1 for a long time, and I'm skeptical it'll ever be added to the language.

It's easy to emulate with a helper function.  You have to type a few extra characters, but once the code's written, it's very readable and gets the job done.

[Playground link](https://www.typescriptlang.org/play/?ssl=1&ssc=1&pln=25&pc=12#code/JYOwLgpgTgZghgYwgAgArAA7QDwBUB8yA3gLABQyly2ASvgBQIBGAXMvQG5u4CUyAvIRo826LFFr4A3OSrIOcADYBXCNxlkAvuRjKQCMMAD2IZBkwR6ItBYl6AJhBigI96Tr0Hjp81jwMFFTVkXlFbfw1dfUMTMwtOJVUAfjY4EABPazT04lkqKK9YhAALOFBGVmRsvlIKOSooCDBlKB945gSgnh4NOW06yhKykAA6QNUBeUSIXoamltMh0CqAZyqMjX6PaO8q+3t6ODYQZQBbJmgAGmRKk-PoGrzKRubWquQAahvN8kcERTgjWQBRiphWyiYYCgiDAh2OZwuUGutwRD1yA2QLwW7wAtN9yFsyAgTCswJjJr5LDwnuwAPoCQhwfb0ABM1wAzN0aZwGchwZDoQZONcWVyMWNplIgA)

```
interface Pipeline<T> {
    <R>(cb: (v: T) => R): Pipeline<R>;
    value: T;
}
function pipe(): Pipeline<undefined>;
function pipe<T>(value: T): Pipeline<T>;
function pipe(value?: any): any {
    function chain(cb: any) {
        return pipe(cb(value));
    }
    chain.value = value;
    return chain as any;
}

function add(a: number, b: number) {
    return a + b;
}
declare function subtract(a: number, b: number) {
    return a - b;
}

const r = pipe()
    (_ => add(2, 3))
    (v => subtract(v, 2))
    .value;
```
