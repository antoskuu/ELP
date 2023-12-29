package main

import "fmt"

func afficheSudoku(grille [][]int, n int) { //n est la taille du tableau

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%2d ", grille[i][j])
		}
		fmt.Println()
	}
}
func respecteRegles(grille [][]int, test, i, j, n int) bool {
	/*
		test=valeur qu'on regarde si elle respecte;
		i=coordonnée ligne dans le tableau;
		j= coord colonne dans le tableau;
		n= taille du tableau (9 pour un 9x9)
	*/

	for colonne := 0; colonne < n; colonne++ {
		if grille[i][colonne] == test {
			return false
		}
	}
	for ligne := 0; ligne < n; ligne++ {
		if grille[ligne][j] == test {
			return false
		}
	}
	var debutLigne = 3 * (i / 3)
	var debutColonne = 3 * (j / 3)

	for ligne := debutLigne; ligne < debutLigne+3; ligne++ {
		for colonne := debutColonne; colonne < debutColonne+3; colonne++ {
			if test == grille[ligne][colonne] {
				return false
			}
		}
	}
	return true
}

func main() {

	grilleSudoku := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	afficheSudoku(grilleSudoku, 9)
	if respecteRegles(grilleSudoku, 9, 0, 2, 9) == true { //c'est juste un test
		fmt.Printf("ca marche ca")
	} else {
		fmt.Printf("noooon")
	}
}
