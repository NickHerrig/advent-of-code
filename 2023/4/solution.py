
def parse_cards(lines):
    for line in lines:
        winners, my_numbers = line.split("|")
        _, winners = winners.split(":")
        winners = winners.split()
        my_numbers = my_numbers.split()
                                   
        yield (winners, my_numbers)



def part_one(lines):

    cards = list(parse_cards(lines))

    points = 0
    for winners, my_numbers in cards:
        win = 0
        for n in my_numbers:
            if n in winners:
                win += 1

        if win == 0:
            points += 0
        else: 
            points += pow(2, win-1)

    return points

def part_two(lines):
    return "Not Implemented"


if __name__ == '__main__':
    with open("input.txt", "r") as f:
        lines = [ line.strip() for line in f.readlines() ]

    print(part_one(lines))
    print(part_two(lines))