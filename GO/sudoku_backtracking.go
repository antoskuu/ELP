package main

import "fmt"

const TAILLE = 9
const TAILLE_BLOCK = 3

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

func solve(grille *[TAILLE][TAILLE]int, ligne int, colonne int) bool {
	if ligne == TAILLE {
		return true

	} else if colonne == TAILLE {
		return solve(grille, ligne+1, 0)

	} else if (*grille)[ligne][colonne] != 0 {
		return solve(grille, ligne, colonne+1)

	} else {
		for k := 1; k <= TAILLE; k++ {
			if absentSurBlock(k, *grille, ligne, colonne) && absentSurColonne(k, *grille, colonne) && absentSurLigne(k, *grille, ligne) {
				(*grille)[ligne][colonne] = k
				if solve(grille, ligne, colonne+1) {
					return true
				}
				(*grille)[ligne][colonne] = 0
			}
		}
		return false
	}
}

func main() {
	grille := [TAILLE][TAILLE]int{
		{0, 0, 7, 0, 3, 0, 5, 0, 0},
		{3, 6, 0, 0, 0, 0, 0, 4, 1},
		{0, 0, 0, 4, 1, 6, 0, 0, 0},
		{0, 0, 0, 2, 8, 7, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 1, 0},
		{0, 8, 0, 0, 0, 0, 0, 5, 0},
		{0, 0, 4, 0, 0, 0, 3, 0, 0},
		{0, 0, 2, 0, 0, 0, 8, 0, 0},
		{0, 0, 0, 3, 9, 8, 0, 0, 0},
	}

	fmt.Printf("Avant la modification : \n")
	for _, ligne := range grille {
		fmt.Println(ligne)
	}

	if solve(&grille, 0, 0) {
		fmt.Printf("\nAprès la modification : \n")
		for _, ligne := range grille {
			fmt.Println(ligne)
		}
	} else {
		fmt.Printf("\nErreur pas de solution")
	}

}
