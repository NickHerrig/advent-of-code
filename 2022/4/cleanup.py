
def main():

    with open('input.txt', 'r') as f:

        overlap = 0
        qty_overlap = 0

        for line in f:
            elf_one, elf_two = line.split(",")
            elf_one = elf_one.split("-")
            elf_two = elf_two.split("-")
            elf_one = list(range(int(elf_one[0]),int(elf_one[1]) + 1))
            elf_two = list(range(int(elf_two[0]),int(elf_two[1]) + 1))

            elf_one_included = all(item in elf_one for item in elf_two)
            elf_two_included = all(item in elf_two for item in elf_one)

            if elf_one_included or elf_two_included:
                overlap += 1
            
            common_elements = [i for i in elf_one if i in elf_two]
            if len(common_elements) > 0:
                qty_overlap += 1

        print(overlap)
        print(qty_overlap)


if __name__=="__main__":
    main()
