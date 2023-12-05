import re

def part_one(lines):

    sum = 0 
    for line in lines:
        digits = re.findall(r'\d', line)

        if len(digits) == 0:
            continue
        
        first = digits[0]
        last = digits[-1]
        num = f'{first}{last}'
        sum += int(num)

    return sum

def part_two(lines):
    number_map = {
        'one': '1', 'two': '2', 'three': '3', 'four': '4', 
        'five': '5', 'six': '6', 'seven': '7', 'eight': '8', 'nine': '9'
    }

    # Look ahead regex for word digits and digits.
    pattern = r'(?=(one|two|three|four|five|six|seven|eight|nine|\d))'
    sum = 0 
    for line in lines:

        matches = list(re.finditer(pattern, line))
        firstMatch = matches[0].group(1)
        lastMatch = matches[-1].group(1)

        # If word in map, replace with digit
        if firstMatch in number_map:
            firstMatch = number_map[firstMatch]
        if lastMatch in number_map:
            lastMatch = number_map[lastMatch]

        # concat digit then sum
        num = f'{firstMatch}{lastMatch}'
        sum += int(num)

    return sum

if __name__ == "__main__":

    with open("input.txt", "r") as f:
        lines = f.readlines()
    
    print(part_one(lines))
    print(part_two(lines))