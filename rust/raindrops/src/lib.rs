pub fn raindrops(y: i32) -> String {
    let mut answer = String::new();
    if y % 3 == 0 { answer.push_str("Pling"); }
    if y % 5 == 0 { answer.push_str("Plang"); }
    if y % 7 == 0 { answer.push_str("Plong"); }
    if answer.is_empty() { answer = y.to_string(); }
    return answer
}
