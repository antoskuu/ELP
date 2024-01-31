#convertisseur de grille du site top_sudoku pour go

def convert_and_increment(grid):
    rows = grid.split('\n')

    for i in range(len(rows)):
        row = rows[i].split()

    #décommenter la ligne suivante pour unr grille 16x16
        #row = [str(int(num, 16) + 1) if num != '*' else '0' for num in row]

        row_str = ', '.join(row)

        rows[i] = '{' + row_str + '}' + ','

    result = '\n'.join(rows)

    return result


input_grid = """5  0  0   0  0  4   0  7  0 
0  1  0   6  0  0   3  0  0 
0  0  8   0  1  0   0  0  2 
0  0  5   0  8  0   0  0  0 
0  2  0   5  0  0   0  0  0 
9  0  0   0  0  2   0  0  0 
0  0  3   0  0  0   0  0  9 
0  6  0   1  0  0   5  0  0 
8  0  0   0  0  7   0  4  0 """

result = convert_and_increment(input_grid)
print(result)