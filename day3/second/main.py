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

def isgear(i, j, mat):
    adjs = {} 

    if j > 0 and mat[i][j-1] > 0:
        adjs[mat[i][j-1]] = True
    if j < len(mat[0])-1 and mat[i][j+1] > 0:
        adjs[mat[i][j+1]] = True

    if i > 0:
        if mat[i-1][j] > 0:
            adjs[mat[i-1][j]] = True
        if j > 0 and mat[i-1][j-1] > 0:
            adjs[mat[i-1][j-1]] = True
        if j < len(mat[0])-1 and mat[i-1][j+1] > 0:
            adjs[mat[i-1][j+1]] = True

    if i < len(mat)-1:
        if mat[i+1][j] > 0:
            adjs[mat[i+1][j]] = True
        if j > 0 and mat[i+1][j-1] > 0:
            adjs[mat[i+1][j-1]] = True
        if j < len(mat[0])-1 and mat[i+1][j+1] > 0:
            adjs[mat[i+1][j+1]] = True

    if len(adjs) != 2: 
        return None 
    return list(adjs.keys()) 

def main():
    matrix = []

    with open("../input.txt") as file:
        while (line := file.readline()):
            matrix.append(line.strip())

    n, m = len(matrix), len(matrix[0])
    parts = [[0] * m for _ in range(n)]
    idToVal = {}

    partsId = 1
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
                for j in range(l, r):
                    parts[i][j] = partsId
                idToVal[partsId] = int(matrix[i][l:r])
                partsId += 1
            r += 1

    sum = 0
    for i in range(n):
        j = 0
        while j < m:
            if matrix[i][j] != '*':
                j += 1
                continue
            ids = isgear(i, j, parts)
            if ids:
                sum += idToVal[ids[0]] * idToVal[ids[1]] 
            j += 1

    print(sum)

main()
