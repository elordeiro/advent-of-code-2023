from collections import namedtuple

Sdr = namedtuple('SDR', ['src', 'dst', 'rng'])

seeds = []
maps = [[] for _ in range(7)]

with open("../input.txt") as file:
    seeds = list(map(int, file.readline().lstrip("seeds:").strip().split()))

    idx = -1
    while (line := file.readline()):
        line = line.strip()
        if line == "":
            continue
        if not line[0].isnumeric():
            idx += 1
            continue 
        
        dsr = list(map(int, line.split()))
        maps[idx].append(Sdr(dsr[1], dsr[0], dsr[2]))

closest = float('inf')

# Brute force. Takes too long
# Go solution takes about 90s. Same logic
for i in range(0, len(seeds), 2):
    for seed in range(seeds[i], seeds[i] + seeds[i+1]):
        old_seed = seed
        for map in maps:
            for sdr in map:
                if sdr.src <= seed <= sdr.src+sdr.rng:
                    seed = sdr.dst + (seed - sdr.src)
                    break
        closest = min(closest, seed)
        seed = old_seed

print(closest)