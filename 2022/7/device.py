def parse_data(data):


    file_system = {"parent": None, "size": 0}
    for line in data:
        match line.split():

            case ["$", "cd", "/"]:
                working_dir = file_system
                            
            case ["$", "ls"]: 
                pass

            case ["dir", dir_name]:
                pass

            case [size, file_name ] if size.isnumeric():
                working_dir[file_name] = size
                size = int(size)
                p = working_dir
                while p:
                    p["size"] += size
                    p = p["parent"]
                
            case ["$", "cd", ".."]:
                working_dir = working_dir["parent"]

            case ["$", "cd", dir_name]:
                if dir_name not in working_dir:
                    working_dir[dir_name] = {"parent": working_dir, "size": 0}
                working_dir = working_dir[dir_name]

    return file_system

def problem_one(file_system):
    
    def sum_dirs(fs):

        total_size = 0
        if fs["size"] <= 100000:
            total_size += fs["size"]

        for key, value in fs.items():
            if isinstance(value, dict) and key != "parent":
                total_size += sum_dirs(value)

        return total_size
    
    return sum_dirs(file_system)

def problem_two(file_system):
    TOTAL_DISK_SPACE = 70000000
    NEEDED = 30000000
    OPEN_SPACE = TOTAL_DISK_SPACE - file_system["size"]

    print(OPEN_SPACE)
    space_to_delete = None
    def sum_dirs(fs):

        nonlocal space_to_delete

        current_size = fs["size"]
        if not space_to_delete:
            space_to_delete = current_size

        if  OPEN_SPACE + current_size >= NEEDED and current_size < space_to_delete:
            space_to_delete = current_size

        for key, value in fs.items():
            if type(value) == dict and key != "parent":
                sum_dirs(value)
        
    sum_dirs(file_system)

    return space_to_delete


def main():

    data = []
    with open("input.txt") as f:
        data = f.readlines()

    file_system = parse_data(data)

    print(problem_one(file_system))
    print(problem_two(file_system))

if __name__=="__main__":
    main()