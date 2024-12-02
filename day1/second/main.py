atoi = {
    "one":   '1',
    "two":   '2',
    "three": '3',
    "four":  '4',
    "five":  '5',
    "six":   '6',
    "seven": '7',
    "eight": '8',
    "nine":  '9',
}

def main():
    values = []
    with open("../input.txt", "r") as file:
        while (line := file.readline()) and line:
            l, r = '', ''
            
            def assign(c):
                nonlocal l, r
                if l == '':
                    l = c
                r = c
            
            for i, c in enumerate(line):
                if c.isnumeric():
                    assign(c)
                else:
                    for k, v in atoi.items():
                        if line[i:].startswith(k):
                            assign(v)
            
            values.append(int(l + r))

    print(sum(values))

main()