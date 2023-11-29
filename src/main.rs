fn main() {}

#[no_mangle]
pub extern fn add(a: i64, b: i64) -> i64 {
    a + b
}
