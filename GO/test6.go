package main

import (
	"fmt"
	"sync"
)

const TAILLE = 9
const TAILLE_BLOCK = 3

var wg sync.WaitGroup

func afficheSudoku(grille [9][9]int, n int) { //n est la taille du tableau

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%2d ", grille[i][j])
		}
		fmt.Println()
	}
}

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

func solve(grille [TAILLE][TAILLE]int, ligne int, colonne int, wg *sync.WaitGroup) ([TAILLE][TAILLE]int, bool) {

	if ligne == TAILLE {

		return grille, true

	} else if colonne == TAILLE {
		return solve(grille, ligne+1, 0, wg)

	} else if grille[ligne][colonne] != 0 {
		return solve(grille, ligne, colonne+1, wg)

	} else {
		for k := 1; k <= TAILLE; k++ {
			if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
				grille[ligne][colonne] = k
				nouvelle_grille, solution := solve(grille, ligne, colonne+1, wg)
				if solution {
					return nouvelle_grille, true

				}
				grille[ligne][colonne] = 0
			}
		}
		return grille, false
	}
}

func solvesudokupartial(grille [TAILLE][TAILLE]int, ligne int, colonne int, value int, wg *sync.WaitGroup) {

	var grille2 [TAILLE][TAILLE]int
	defer wg.Done()
	grille[ligne][colonne] = value
	grille2, solution := solve(grille, 0, 0, wg)
	if solution {
		fmt.Printf("Solution pour : %d\n", value)
		afficheSudoku(grille2, TAILLE)
	} else {
		fmt.Printf("Pas de solution pour : %d\n", value)
	}

}

func solve_para(grille [TAILLE][TAILLE]int, ligne int, colonne int) {
	for k := 1; k <= TAILLE; k++ {
		if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
			wg.Add(1)
			go solvesudokupartial(grille, ligne, colonne, k, &wg)

		}
	}
	wg.Wait()
}

func main() {
	grille := [TAILLE][TAILLE]int{
		{0, 0, 8, 0, 0, 4, 0, 9, 0},
		{0, 7, 0, 1, 0, 0, 5, 0, 0},
		{5, 0, 0, 0, 6, 0, 0, 0, 3},
		{1, 0, 0, 0, 4, 0, 0, 0, 8},
		{0, 8, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 1, 0},
		{0, 6, 0, 2, 0, 0, 0, 0, 0},
		{3, 0, 0, 0, 8, 0, 0, 0, 5},
		{0, 0, 7, 0, 0, 9, 0, 4, 0},
	}

	fmt.Printf("Avant la modification : \n")
	for _, ligne := range grille {
		fmt.Println(ligne)
	}

	solve_para(grille, 0, 0)

	//nouvelle_grille, solution := solve(grille, 0, 0)

	//if solution {
	//	fmt.Printf("\nAprès la modification : \n")
	//	for _, ligne := range nouvelle_grille {
	//		fmt.Println(ligne)
	//	}
	//} else {
	//	fmt.Printf("\nErreur pas de solution")
	//}

}
