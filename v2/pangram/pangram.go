package pangram

func IsPangram(input string) bool {
  lettersMap := make(map[rune]bool)

  for _, letter := range input {
    if !isLetter(letter) {
      continue
    }

    lettersMap[letterToLowerCase(letter)] = true
  }

  return len(lettersMap) >= 26
}

func isLetter(letter rune) bool {
  return letter >= 97 && letter <= 122 || letter >= 65 && letter <= 90
}

func letterToLowerCase(letter rune) rune {
  if isUpperCase(letter) {
    return letter + 32
  }

  return letter
}

func isUpperCase(letter rune) bool {
  return letter < 96
}
