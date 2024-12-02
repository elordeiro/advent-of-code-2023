values = []

with open("../input.txt", "r") as file:
    while (line := file.readline()) and line:
        l, r = '', ''
        for c in line:
            if c.isnumeric():
                if l == '':
                    l, r = c, c
                else:
                    r = c
        values.append(int(l+r))

print(sum(values))