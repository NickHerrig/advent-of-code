def calculate_distances(time, record):
    
    count = 0
    for hold in range(int(time)+1):
        dist = hold * (int(time)-hold)
        if dist >int(record):
            count += 1
            
    return count
            

def part_one(times, record_dist):

    solution = 1
    for time, record in zip(times, record_dist):
        solution *= calculate_distances(time, record)

    return solution

def part_two(times, record_dist):
    solution = calculate_distances(times, record_dist)
    return solution

if __name__=="__main__":

    with open('input.txt', 'r') as file:
        lines = file.readlines()

    times = lines[0].split()[1:]
    record_dist = lines[1].split()[1:]
    print(part_one(times, record_dist))

    print(part_two("".join(times), "".join(record_dist)))