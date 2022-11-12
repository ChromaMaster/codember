use std::fs;

fn main() {
    println!("Hello, world!");

    let contents = fs::read_to_string("input/users.txt").expect("Cannot read users file");

    let users = contents
        .split("\n\n")
        .map(|user| user.replace('\n', " "))
        .collect::<Vec<String>>();

    let mut valid_users = 0;
    let mut last_valid_user = String::new();
    for user in users {
        let valid = user_is_valid(&user);
        // println!("{} -> {}\n", user, valid);
        if valid {
            valid_users += 1;
            last_valid_user = user.clone();
        }
    }

    println!("\nTotal valid users: {}\nLast valid user is: {}", valid_users, last_valid_user);
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
    fn user_has_all_fields() {
        let user = String::from("usr:@midudev\neme:mi@gmail.com\npsw:123456\nage:22\nloc:bcn\nfll:82");
        assert!(user_is_valid(&user));
    }
}
