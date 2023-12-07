def parse_data(file):

    grid = []
    for line in file:
        for char in line:
            grid.append(char)

    return grid


def problem_one(data):

    return 0

def problem_two(data):

    return 0 



def main():
    file = []
    with open("input.txt") as f:
        file = f.readlines()

    data = parse_data(file)


    rows = 99
    cols = 100

    for i in range(rows):
        for j in range(cols):
            print(data[i*cols + j])


    print(problem_one(data))
    print(problem_two(data))


if __name__=="__main__":
    main()