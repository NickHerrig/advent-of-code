import re


def check_possible(subsets):
    # Game Rules
    blue_cubes, red_cubes, green_cubes = (14, 12, 13)

    for subset in subsets:
        number, color = subset.strip().split(" ")
        if color == "blue" and int(number) > blue_cubes:
            return False
        if color == "red" and int(number) > red_cubes:
            return False
        if color == "green" and int(number) > green_cubes:
            return False
    
    return True

def part_one(input):

    solution = 0
    for line in input:
        game_details , subsets = line.split(":")
        gameID = re.findall(r'\d+', game_details)[0]
        subsets = subsets.replace(";", ",").split(",")

        possible = check_possible(subsets)

        if possible:
            solution += int(gameID)
        else:
            continue

    return solution


def max_colors(subsets):
    colors = {}
    for subset in subsets:
        number, color = subset.strip().split(" ")
        if color not in colors:
            colors[color] = []
        colors[color].append(int(number))

    return max(colors["blue"]) * max(colors["red"]) * max(colors["green"])


def part_two(input):

    solution = 0
    for line in input:
        game_details , subsets = line.split(":")
        gameID = re.findall(r'\d+', game_details)[0]
        subsets = subsets.replace(";", ",").split(",")

        power = max_colors(subsets)

        solution += power

    return solution


if __name__ == '__main__':
    with open("input.txt", "r") as f:
        lines = f.readlines()

    print(part_one(lines))
    print(part_two(lines))