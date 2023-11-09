use std::fs;

fn main() {
    let contents = fs::read_to_string("src/challenge02/input/encrypted.txt").expect("Cannot read users file");

    let chars = split_into_ascii_values(&contents)
        .iter()
        .map(|&int_val| int_val as char)
        .collect::<Vec<char>>();

    println!("Message: {}", chars.iter().collect::<String>());
}

fn split_into_ascii_values(text: &str) -> Vec<u8> {
    let ascii_values = text
        .chars()
        .fold(String::new(), |acc, chr| {
            if chr.is_whitespace() {
                return format!("{}{} ", acc, "32");
            }

            // Split elements by ascii values
            let elements: Vec<String> = acc.split(' ').map(|elem| elem.to_owned()).collect();

            // Append the last processes char and check the integer value of it
            let tmp = format!("{}{}", elements[elements.len() - 1], chr);
            let int_value = tmp.parse::<u8>().unwrap();

            // If the two digits ascii has meaning, split it.
            if int_value < 97 {
                format!("{}{}", acc, chr)
            } else {
                format!("{}{} ", acc, chr)
            }
        });

    ascii_values
        .trim()
        .split(' ')
        .map(|ascii| ascii.to_owned().parse::<u8>().unwrap())
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_ascii_string_is_correctly_split_into_individual_characters() {
        let ascii_text = String::from("109105100117");

        let got = split_into_ascii_values(&ascii_text);
        let want = vec![109, 105, 100, 117];

        assert_eq!(got, want)
    }

    #[test]
    fn test_space_character_is_correctly_inserted_if_found() {
        let ascii_text = String::from("9911110010110998101114 109105100117");

        let got = split_into_ascii_values(&ascii_text);
        let want = vec![99, 111, 100, 101, 109, 98, 101, 114, 32, 109, 105, 100, 117];

        assert_eq!(got, want)
    }
}
