#convertisseur de grille du site top_sudoku pour go

def convert_and_increment(grid):
    rows = grid.split('\n')

    for i in range(len(rows)):
        row = rows[i].split()

    #décommenter la ligne suivante pour unr grille 16x16
        row = [str(int(num, 16) + 1) if num != '*' else '0' for num in row]

        row_str = ', '.join(row)

        rows[i] = '{' + row_str + '}' + ','

    result = '\n'.join(rows)

    return result


input_grid = """*  *  *  1   *  *  *  F   *  8  6  *   *  *  C  2 
7  *  4  *   *  *  *  B   0  *  *  5   *  *  9  * 
*  *  *  F   *  *  *  *   9  3  *  4   *  7  E  * 
*  *  5  *   6  *  *  *   7  *  D  1   0  *  4  * 
*  *  3  9   *  *  *  *   *  *  B  2   *  *  0  8 
2  *  *  D   C  4  *  *   *  *  *  8   *  *  *  7 
8  *  *  *   *  B  *  A   E  *  C  *   4  *  *  * 
4  C  *  *   2  3  7  *   *  *  *  *   *  1  *  * 
D  4  8  *   A  *  5  C   *  2  9  *   F  *  7  B 
9  0  *  *   E  1  *  *   *  4  *  *   D  C  A  * 
*  *  2  *   *  D  *  *   *  *  *  *   8  *  *  * 
*  3  *  *   4  *  6  9   *  E  *  *   *  *  *  * 
*  *  *  *   *  6  *  5   F  9  *  *   *  *  *  * 
*  *  1  *   B  7  *  *   *  *  *  *   *  *  D  0 
*  5  *  *   *  *  E  *   B  *  4  D   9  6  *  * 
*  *  *  C   1  *  *  *   *  *  *  E   *  *  5  F """

result = convert_and_increment(input_grid)
print(result)