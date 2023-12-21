from collections import defaultdict


def parse_cards(lines):
    for line in lines:
        winners, my_numbers = line.split("|")
        _, winners = winners.split(":")
        winners = winners.split()
        my_numbers = my_numbers.split()
                                   
        yield (winners, my_numbers)


def calulate_wins(winners, my_numbers):
    win = 0
    for n in my_numbers:
        if n in winners:
            win += 1
    return win

def part_one(lines):

    cards = list(parse_cards(lines))

    points = 0
    for winners, my_numbers in cards:
        win = calulate_wins(winners, my_numbers)

        if win == 0:
            points += 0
        else: 
            points += pow(2, win-1)

    return points

def part_two(lines):
    # card instances {1: 1} aka card one has one instance. 
    card_copies = defaultdict(int)
    cards = list(parse_cards(lines))
    for i, (winners, my_numbers) in enumerate(cards):
        win = calulate_wins(winners, my_numbers)
        card_copies[str(i)] += 1
        copies = card_copies[str(i)]
        for w in range(1, win+1):
            card_copies[str(i+w)] += copies
    return sum(card_copies.values())


if __name__ == '__main__':
    with open("input.txt", "r") as f:
        lines = [ line.strip() for line in f.readlines() ]

    print(part_one(lines))
    print(part_two(lines))