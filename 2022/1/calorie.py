
def main():

    lines = []
    with open('input.txt', 'r') as f:
        lines = [line.rstrip() for line in f]


    calorie_sum = 0
    calorie_list = []
    for line in lines:
        try:
            calorie_sum += int(line)
        except ValueError:
            calorie_list.append(calorie_sum)
            calorie_sum = 0
            continue

    calorie_list.sort(reverse=True)

    print(calorie_list[0])
    print(sum(calorie_list[:3]))



if __name__=="__main__":
    main()
