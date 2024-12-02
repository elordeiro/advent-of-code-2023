games = []

with open("../input.txt", "r") as file:
    while (line := file.readline()) and line:
        mins = {
            "red": 0,
            "green": 0,
            "blue": 0,
        }
        
        sets = line[line.index(":")+1:].split(";")
        for set in sets:
            cubes = set.split(",")
            for cube in cubes:
                count, color = cube.split()
                mins[color] = max(mins[color], int(count))

        games.append(mins["red"] * mins["green"] * mins["blue"])
        

print(sum(games))