import re
def replace_words_with_digits(s):
    word_to_digit = {
        'one': '1', 'two': '2', 'three': '3', 'four': '4',
        'five': '5', 'six': '6', 'seven': '7', 'eight': '8', 'nine': '9'
    }
    
    # Find all occurrences of the words and their start positions
    occurrences = [(word, m.start()) for word, digit in word_to_digit.items() for m in re.finditer(word, s)]
    
    # Sort by start position
    occurrences.sort(key=lambda x: x[1])
    
    # Replace each occurrence, adjusting for changes in string length
    shift = 0
    for word, pos in occurrences:
        s = s[:pos + shift] + word_to_digit[word] + s[pos + shift + len(word):]
        shift += 1 - len(word)

    return s

# Test case
print(replace_words_with_digits("eightwothree"))  # Expected output: 8wo3
