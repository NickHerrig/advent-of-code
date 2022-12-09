import string

priority_dict = {}
for i, letter in enumerate(string.ascii_letters, start=1):
    priority_dict[letter] = i


def main():

    rucksacks = []
    with open('input.txt', 'r') as f:
        rucksacks = [ line.strip() for line in f.readlines() ]
    
    total = 0
    for sack in rucksacks:
        mid_index = len(sack) // 2
        first_half = sack[:mid_index]
        second_half = sack[mid_index:]
        common_chars = set(first_half) & set(second_half)
        total += priority_dict[common_chars.pop()]

    print(total)


if __name__=="__main__":
    main()