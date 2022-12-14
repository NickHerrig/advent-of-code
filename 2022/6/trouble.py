def parse_data(file):
    return file.pop(0)


def problem_one(stream):

    for start, end in enumerate(range(4, len(stream)+1)):
        compare_chars = stream[start:end]

        if not any(compare_chars.count(c) > 1 for c in compare_chars):
            print(end)
            break


def problem_two(stream):

    for start, end in enumerate(range(14, len(stream)+1)):
        compare_chars = stream[start:end]

        if not any(compare_chars.count(c) > 1 for c in compare_chars):
            print(end)
            break


def main():

    file = []
    with open("input.txt") as f:
        file = f.readlines()

    stream = parse_data(file)
    problem_one(stream)
    problem_two(stream)



if __name__=="__main__":
    main()