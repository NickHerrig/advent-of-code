from collections import deque

def parse_data(data):

    for index, line in enumerate(data):
        if line.isspace():
            stacks, columns, moves = data[:index-1], data[index-1], data[index+1:]


    num_stacks = max([ int(number) for number in columns.split() ])

    parsed_stacks = [[] for i in range(num_stacks)]
    for row in stacks:
        for stack, index in enumerate(range(1, 4*num_stacks, 4),1):
            if row[index].isalpha():
                parsed_stacks[stack-1].append(row[index])


    parsed_moves = []
    for move in moves:
        move_tuple = []
        for val in move.split():
            if val.isnumeric():
                move_tuple.append(int(val))
        parsed_moves.append(move_tuple)

    return parsed_stacks, parsed_moves

def problem_one(stacks, moves):

    for move in moves:
        crates_to_move, source, destination = move
        for crate in range(crates_to_move):
            stacks[destination-1].insert(0,stacks[source-1].pop(0))

    return "".join([stack.pop(0) for stack in stacks])


def problem_two(stacks, moves):
    for move in moves:
        num_crates_to_move, source, destination = move
        crates_to_move = stacks[source-1][:num_crates_to_move]
        stacks[source-1] = stacks[source-1][num_crates_to_move:]
        stacks[destination-1] = crates_to_move + stacks[destination-1]

    return "".join([stack.pop(0) for stack in stacks])


def main():

    file = []
    with open("input.txt") as f:
        file = f.readlines()

    stacks, moves = parse_data(file)
    print(problem_one(stacks, moves))

    stacks, moves = parse_data(file)
    print(problem_two(stacks, moves))




if __name__=="__main__":
    main()