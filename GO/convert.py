def convert_and_increment(grid):
    # Split the input grid into rows
    rows = grid.split('\n')

    # Process each row
    for i in range(len(rows)):
        row = rows[i].split()

        # Replace '*' with '0' and convert hex numbers to decimal, then increment each number by 1
        #row = [str(int(num, 16) + 1) if num != '*' else '0' for num in row]

        # Join the numbers with commas
        row_str = ', '.join(row)

        # Surround the row with curly braces
        rows[i] = '{' + row_str + '}' + ','

    # Join the rows with line breaks
    result = '\n'.join(rows)

    return result

# Example usage with your provided grid
input_grid = """0  0  4   0  0  9   0  0  6 
0  7  0   1  0  0   8  0  0 
3  0  0   0  7  0   0  5  0 
7  0  0   0  3  0   0  9  0 
0  0  6   0  0  1   0  0  4 
0  0  0   5  0  0   2  0  0 
9  0  0   0  8  0   0  0  0 
0  1  0   9  0  0   0  0  0 
0  0  8   0  0  2   3  0  0 """

result = convert_and_increment(input_grid)
print(result)
