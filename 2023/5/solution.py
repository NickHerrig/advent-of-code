from collections import OrderedDict


def part_one(seeds, converstion_maps):
    locations = []
    for seed in seeds:
        print(f"Original Seed: {seed}")

        # Lop through conversion maps
        for map_name, map_details in converstion_maps.items():

            # Find the conversion for the Map
            for i, row in enumerate(map_details):
                dest_start, source_start, length = row
                source_end = source_start + length -1
                if seed <= source_end and seed >= source_start:
                    #print(f"Found conversion in {map_name} for seed {seed}")
                    seed = (seed - source_start) + dest_start
                    break
                
                elif len(map_details) == i+1:
                    # Seed stays the same
                    break

        locations.append(seed)         

    return min(locations)

def part_two(seeds, converstion_maps):
    return ""

if __name__ == '__main__':
    with open("input.txt", "r") as f:
        content = f.read()
    
    sections = content.strip().split("\n\n")
    seeds = sections.pop(0).split(": ")[-1].split()
    seeds = [int(seed) for seed in seeds]

    converstion_maps = OrderedDict()
    for section in sections:

        lines = section.split("\n")
        map_name = lines.pop(0).split()[0]
        map_details = []
        for line in lines:
            dest_start, source_start, length = line.split()
            map_details.append((int(dest_start), int(source_start), int(length)))

        converstion_maps[map_name] = map_details

    
    print(part_one(seeds, converstion_maps))
    print(part_two(seeds, converstion_maps))