def isvalid(c):
    if (not c.isnumeric()) and c != '.':
        return True
    return False

def ispart(i, l, r, mat):
    if l > 0 and mat[i][l-1] != '.':
        return True
    
    if r < len(mat) - 1 and mat[i][r+1] != '.':
        return True
    
    for j in range(l-1, r+2):
        if j < 0 or j >= len(mat[0]):
            continue
        if i > 0 and isvalid(mat[i-1][j]):
            return True
        if i < len(mat) - 1 and isvalid(mat[i+1][j]):
            return True
    
    return False

def main():
    matrix = []

    with open("../input.txt") as file:
        while (line := file.readline()):
            matrix.append(line.strip())

    n, m = len(matrix), len(matrix[0])
    sum = 0

    for i in range(n):
        r = 0
        while r < m:
            if not matrix[i][r].isnumeric():
                r += 1
                continue
            l = r
            while r < m and matrix[i][r].isnumeric():
                r += 1
            if ispart(i, l, r-1, matrix):
                sum += int(matrix[i][l:r])
            r += 1

    print(sum)

main()
