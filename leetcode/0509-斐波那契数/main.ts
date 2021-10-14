function fib(n: number): number {
    if (n < 2) {
        return n
    }
    let total:number = 0
    let pre: number = 0
    let post: number = 1
    for (let i=1; i < n; i++) {
        total = pre + post
        pre = post
        post = total
    }
    return total
};

console.log(fib(3))