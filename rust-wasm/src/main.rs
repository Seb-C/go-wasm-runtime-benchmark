fn main() {}

#[no_mangle]
pub extern fn add(a: i64, b: i64) -> i64 {
    a + b
}

// From https://github.com/eliovir/rust-examples/blob/master/fibonacci.rs#L65C1-L83C2
#[no_mangle]
pub extern fn fibonacci(n: i64) -> i64 {
    if n < 0 {
        panic!("{} is negative!", n);
    } else if n == 0 {
        panic!("zero is not a right argument to fibonacci()!");
    } else if n == 1 {
        return 1;
    }

    let mut sum = 0;
    let mut last = 0;
    let mut curr = 1;
    for _i in 1..n {
        sum = last + curr;
        last = curr;
        curr = sum;
    }
    sum
}
