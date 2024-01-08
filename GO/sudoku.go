package main

import "fmt"

const TAILLE = 16
const TAILLE_BLOCK = 9

func absentSurLigne(k int, grille [TAILLE][TAILLE]int, ligne int) bool {
	for colonne := 0; colonne < TAILLE; colonne++ {
		if grille[ligne][colonne] == k {
			return false
		}
	}
	return true
}

func absentSurColonne(k int, grille [TAILLE][TAILLE]int, colonne int) bool {
	for ligne := 0; ligne < TAILLE; ligne++ {
		if grille[ligne][colonne] == k {
			return false
		}
	}
	return true
}

func absentSurBlock(k int, grille [TAILLE][TAILLE]int, ligne int, colonne int) bool {
	ligne2 := ligne - ligne%TAILLE_BLOCK
	colonne2 := colonne - colonne%TAILLE_BLOCK

	for x := 0; x < TAILLE_BLOCK; x++ {
		for y := 0; y < TAILLE_BLOCK; y++ {
			if grille[ligne2+x][colonne2+y] == k {
				return false
			}
		}
	}
	return true
}

func solve(grille [TAILLE][TAILLE]int, ligne int, colonne int) ([TAILLE][TAILLE]int, bool) {
	if ligne == TAILLE {
		return grille, true

	} else if colonne == TAILLE {
		return solve(grille, ligne+1, 0)

	} else if grille[ligne][colonne] != 0 {
		return solve(grille, ligne, colonne+1)

	} else {
		for k := 1; k <= TAILLE; k++ {
			if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
				grille[ligne][colonne] = k
				nouvelle_grille, solution := solve(grille, ligne, colonne+1)
				if solution {
					return nouvelle_grille, true
				}
				grille[ligne][colonne] = 0
			}
		}
		return grille, false
	}
}

func main() {
	grille := [TAILLE][TAILLE]int{
		{0, 7, 0, 14, 0, 0, 15, 0, 3, 11, 2, 8, 5, 10, 0, 0},
		{13, 0, 0, 10, 9, 0, 16, 0, 0, 0, 0, 6, 7, 8, 0, 0},
		{0, 6, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0, 1, 11, 3},
		{0, 0, 3, 1, 0, 12, 0, 0, 0, 15, 0, 0, 14, 6, 9, 0},
		{0, 0, 0, 3, 0, 8, 9, 10, 12, 0, 0, 16, 0, 0, 0, 14},
		{0, 0, 0, 0, 0, 0, 0, 13, 7, 8, 0, 15, 0, 12, 1, 6},
		{0, 0, 0, 16, 14, 0, 0, 0, 0, 0, 11, 0, 0, 9, 15, 0},
		{6, 13, 0, 0, 0, 11, 7, 0, 0, 2, 14, 0, 16, 3, 5, 0},
		{11, 0, 0, 5, 4, 14, 0, 0, 13, 0, 9, 0, 0, 0, 0, 0},
		{0, 0, 0, 12, 2, 10, 0, 9, 0, 1, 16, 5, 0, 14, 0, 7},
		{15, 9, 0, 0, 6, 3, 0, 0, 0, 0, 0, 0, 0, 0, 13, 4},
		{10, 0, 0, 0, 7, 0, 0, 0, 0, 0, 4, 0, 0, 11, 0, 0},
		{12, 0, 0, 0, 0, 0, 0, 0, 0, 13, 0, 0, 0, 5, 8, 0},
		{8, 10, 0, 0, 0, 0, 0, 0, 0, 6, 12, 7, 0, 16, 0, 0},
		{7, 16, 15, 0, 11, 0, 8, 0, 0, 0, 0, 10, 0, 13, 0, 0},
		{14, 0, 1, 4, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	fmt.Printf("Avant la modification : \n")
	for _, ligne := range grille {
		fmt.Println(ligne)
	}

	nouvelle_grille, solution := solve(grille, 0, 0)

	if solution {
		fmt.Printf("\nAprès la modification : \n")
		for _, ligne := range nouvelle_grille {
			fmt.Println(ligne)
		}
	} else {
		fmt.Printf("\nErreur pas de solution")
	}

}
