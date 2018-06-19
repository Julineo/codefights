"""In the popular Minesweeper game you have a board with some mines and those cells that don't contain a mine have a number in it that indicates the total number of mines in the neighboring cells. Starting off with some arrangement of mines we want to create a Minesweeper game setup.

Example

For

matrix = [[true, false, false],
          [false, true, false],
          [false, false, false]]
the output should be

minesweeper(matrix) = [[1, 2, 1],
                       [2, 1, 1],
                       [1, 1, 1]]
"""

def minesweeper(matrix):
    m = len(matrix)
    n = len(matrix[0])
    tmp = []
    
    #creating structure larger then matrix to avoid index boundaries
    for i in range(m+2):
        row = []
        for j in range(n+2):
            row.append(bool(0))
        tmp.append(row)
    
    #feel with the data from original matrix
    for i in range(m):
        for j in range(n):
            tmp[i+1][j+1] = matrix[i][j]
                 
    #form output
    res = []
    for i in range(m):
        row = []
        for j in range(n):
            row.append(minescount(i+1, j+1, tmp))
        res.append(row)
                       
    return res

#helper function to count surrounding mines
def minescount(i, j, tmp):
    cnt = int(tmp[i-1][j-1]) + int(tmp[i-1][j]) + int(tmp[i-1][j+1]) + int(tmp[i][j-1]) + int(tmp[i][j+1]) + int(tmp[i+1][j-1]) + int(tmp[i+1][j]) + int(tmp[i+1][j+1])
    return cnt
