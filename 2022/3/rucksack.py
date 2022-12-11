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

    groups = []
    for i in range(0, len(rucksacks), 3):
        group = (rucksacks[i], rucksacks[i+1], rucksacks[i+2])
        groups.append(group)
    
    badge_total = 0
    for group in groups:
        common_badge = set(group[0]) & set(group[1]) & set(group[2])
        badge_total += priority_dict[common_badge.pop()]

    print(total)
    print(badge_total)
    


if __name__=="__main__":
    main()