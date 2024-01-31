package main

import (
	"fmt"
	"time"
)

// constante à changer selon la taille de la grille
const TAILLE = 9
const TAILLE_BLOCK = 3

// fonction qui verifie si un chiffre est absent sur une ligne
func absentSurLigne(k int, grille [TAILLE][TAILLE]int, ligne int) bool {
	for colonne := 0; colonne < TAILLE; colonne++ {
		if grille[ligne][colonne] == k {
			return false
		}
	}
	return true
}

// fonction qui vérifie si un nombre est absent sur une colonnne
func absentSurColonne(k int, grille [TAILLE][TAILLE]int, colonne int) bool {
	for ligne := 0; ligne < TAILLE; ligne++ {
		if grille[ligne][colonne] == k {
			return false
		}
	}
	return true
}

// fonction qui vérifie si un nombre est absent dans un block
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

// fonction de résolution récursive
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
	// grille initiale
	grille := [TAILLE][TAILLE]int{
		{0, 0, 0, 0, 7, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 6, 0, 7, 0},
		{0, 7, 0, 0, 9, 0, 8, 0, 0},
		{4, 0, 0, 0, 0, 0, 0, 5, 0},
		{0, 0, 2, 3, 0, 0, 0, 9, 6},
		{0, 0, 3, 2, 0, 0, 0, 0, 7},
		{0, 5, 0, 0, 8, 0, 9, 0, 0},
		{1, 0, 0, 0, 0, 9, 0, 6, 0},
	}

	fmt.Printf("Avant la modification : \n")
	for _, ligne := range grille {
		fmt.Println(ligne)
	}
	debut := time.Now()
	nouvelle_grille, sol := solve(grille, 0, 0)

	//résolution de la grille
	if sol {
		fmt.Printf("\nAprès la modification : \n")
		for _, ligne := range nouvelle_grille {
			fmt.Println(ligne)
		}
	} else {
		fmt.Printf("\nErreur pas de solution")
	}
	fin := time.Now()
	fmt.Printf("Temps d'execution : %.6fs", fin.Sub(debut).Seconds())
}
