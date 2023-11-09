use std::fs;

fn main() {
    let contents = fs::read_to_string("src/challenge01/input/users.txt").expect("Cannot read users file");

    let users = contents
        .split("\n\n")
        .map(|user| user.replace('\n', " "))
        .filter(|user| user_is_valid(user))
        .collect::<Vec<String>>();

    println!("Total valid users: {}", users.len());
    println!("Last valid user: {}", users[users.len() - 1]);
}

fn user_is_valid(user: &str) -> bool {
    let fields = vec!["usr:", "eme:", "psw:", "age:", "loc:", "fll:"];
    for field in fields {
        if user.matches(field).count() < 1 {
            return false;
        }
    }
    true
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn user_is_valid_if_it_has_all_fields() {
        let user =
            String::from("usr:@midudev\neme:mi@gmail.com\npsw:123456\nage:22\nloc:bcn\nfll:82");
        assert!(user_is_valid(&user));
    }

    #[test]
    fn user_is_not_valid_if_it_misses_a_field() {
        let user =
            String::from("psw:11133\nloc:Canary\nfll:333\nusr:@pheralb\neme:pheralb@gmail.com");
        assert!(!user_is_valid(&user));
    }
}
