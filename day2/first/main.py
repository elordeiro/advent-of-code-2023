games = []

maxes = {
    "red": 12,
    "green": 13,
    "blue": 14,
}

with open("../input.txt", "r") as file:
    game_num = 1
    while (line := file.readline()) and line:
        possible = True
        sets = line[line.index(":")+1:].split(";")
        for set in sets:
            cubes = set.split(",")
            for cube in cubes:
                count, color = cube.split()
                if int(count) > maxes[color]:
                    possible = False
        if possible:
            games.append(game_num)
        game_num += 1

print(sum(games))