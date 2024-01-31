package main

import (
	"fmt"
	"sync"
	"time"
)

const TAILLE = 9
const TAILLE_BLOCK = 3

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup

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

func solve(grille [TAILLE][TAILLE]int, ligne int, colonne int, is_solve *bool) ([TAILLE][TAILLE]int, bool) {
	if !(*is_solve) {
		if ligne == TAILLE {
			fmt.Printf("\nAprès la modification : \n")
			for _, ligne := range grille {
				fmt.Println(ligne)
			}
			*is_solve = true
			return grille, true

		} else if colonne == TAILLE {
			return solve(grille, ligne+1, 0, is_solve)

		} else if grille[ligne][colonne] != 0 {
			return solve(grille, ligne, colonne+1, is_solve)

		} else {
			for k := 1; k <= TAILLE; k++ {
				if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
					grille[ligne][colonne] = k
					nouvelle_grille, solution := solve(grille, ligne, colonne+1, is_solve)
					if solution {
						return nouvelle_grille, true
					}
					grille[ligne][colonne] = 0
				}
			}
			return grille, false
		}
	} else {
		return grille, false
	}
}

func solvesudokupartial(grille [TAILLE][TAILLE]int, ligne int, colonne int, value int, is_solve *bool) {
	grille[ligne][colonne] = value
	solve(grille, 0, 0, is_solve)
}

func solve_para_case2(grille [TAILLE][TAILLE]int, ligne, colonne int, value int, is_solve *bool) {
	defer wg1.Done()
	for grille[ligne][colonne] != 0 {
		colonne = colonne + 1
		if colonne == TAILLE {
			colonne = 0
			ligne = ligne + 1
		}
	}
	for k := 1; k <= TAILLE; k++ {
		if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) && !(*is_solve) {
			fmt.Printf("\nSTART goroutine pour : %d - %d\n", value, k)
			wg2.Add(1)
			value_case2 := k
			grille[ligne][colonne] = value_case2
			go func(grille_case2 [TAILLE][TAILLE]int) {
				solve(grille_case2, 0, colonne+1, is_solve)
				defer wg2.Done()
			}(grille)
		}
	}
	wg2.Wait()
}

func solve_para_case1(grille [TAILLE][TAILLE]int, ligne int, colonne int) {
	is_solve := false
	for grille[ligne][colonne] != 0 {
		colonne = colonne + 1
		if colonne == TAILLE {
			colonne = 0
			ligne = ligne + 1
		}
	}
	for k := 1; k <= TAILLE; k++ {
		if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
			grille[ligne][colonne] = k
			wg1.Add(1)
			fmt.Printf("\nSTART gouroutine pour : %d", k)
			go solve_para_case2(grille, ligne, colonne+1, k, &is_solve)
		}
	}
	wg1.Wait()
}

func main() {
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
	solve_para_case1(grille, 0, 0)
	fin := time.Now()
	fmt.Printf("\nTemps d'execution : %.6fs \n", fin.Sub(debut).Seconds())
}
