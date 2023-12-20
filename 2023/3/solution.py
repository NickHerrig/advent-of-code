import re
from collections import namedtuple
import math


Part = namedtuple("Part", "value line start end")
Symbol = namedtuple("Symbol", "value line start")

def parse_parts(lines):
    parts = []
    pattern = "\d+"
    for i, line in enumerate(lines):
        for match in re.finditer(pattern, line):
            parts.append(Part(match,i ,match.start(), match.end()-1)) 
    return parts

def parse_symbols(lines):
    symbols = []
    pattern = r"[^0-9\.]"
    for i, line in enumerate(lines):
        for match in re.finditer(pattern, line):
            symbols.append(Symbol(match, i, match.start()))
    return symbols

def parse_gears(lines):
    gears = []
    for i, line in enumerate(lines):
        for j, char in enumerate(line):
            if char == "*":
                gears.append((i, j))
    return gears

def part_one(lines):
    parts = parse_parts(lines)
    symbols = parse_symbols(lines)

    part_sum = 0
    for p in parts:
        for s in symbols:
            first_letter = math.dist((p.line, p.start), (s.line, s.start))
            second_letter = math.dist((p.line, p.end), (s.line, s.start))
            if first_letter < 2 or second_letter < 2:
                part_sum += int(p.value.group())

    return part_sum

def part_two(lines):
    parts = parse_parts(lines)
    gears = parse_gears(lines)

    gear_ratio_sum = 0
    for g in gears:
        adjacent_parts = []
        for p in parts:
            first_letter = math.dist((p.line, p.start), (g[0], g[1]))
            second_letter = math.dist((p.line, p.end), (g[0], g[1]))

            if first_letter < 2 or second_letter < 2:
                adjacent_parts.append(p)
        
        if len(adjacent_parts) == 2:
            gear_ratio_sum += int(adjacent_parts[0].value.group()) * int(adjacent_parts[1].value.group())

    return gear_ratio_sum


if __name__ == '__main__':
    with open("input.txt", "r") as f:
        lines = [ line.strip() for line in f.readlines() ]

    print(part_one(lines))
    print(part_two(lines))