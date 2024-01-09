package main

import (
	"fmt"
	"sync"
)

const TAILLE = 9
const TAILLE_BLOCK = 3

var wg sync.WaitGroup

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
		fmt.Printf("\nAprès la modification : \n")
		for _, ligne := range grille {
			fmt.Println(ligne)
		}
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
	defer wg.Done()
	grille[ligne][colonne] = value
	_, solution := solve(grille, 0, 0, wg)
	if solution {
		fmt.Printf("\nSolution pour : %d\n", value)
	} else {
		fmt.Printf("\nPas de solution pour : %d\n", value)
	}
}

func solve_para(grille [TAILLE][TAILLE]int, ligne int, colonne int) {
	for k := 1; k <= TAILLE; k++ {
		if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
			fmt.Printf("\n%d\n", k)
			grille[ligne][colonne] = k
			for l := 1; l <= TAILLE; l++ {
				if absentSurBlock(l, grille, ligne, colonne+1) && absentSurColonne(l, grille, colonne+1) && absentSurLigne(l, grille, ligne) {
					grille[ligne][colonne+1] = l
					for m := 1; m <= TAILLE; m++ {
						if absentSurBlock(m, grille, ligne, colonne+3) && absentSurColonne(m, grille, colonne+3) && absentSurLigne(m, grille, ligne) {
							grille[ligne][colonne+3] = m
							for n := 1; n <= TAILLE; n++ {
								if absentSurBlock(n, grille, ligne, colonne+4) && absentSurColonne(n, grille, colonne+4) && absentSurLigne(n, grille, ligne) {
									wg.Add(1)
									fmt.Printf("\nSTART %d %d %d %d", k, l, m, n)
									go solvesudokupartial(grille, ligne, colonne+5, n, &wg)
									grille[ligne][colonne+4] = 0
								}
							}
						}
						grille[ligne][colonne+3] = 0
					}
				}
				grille[ligne][colonne+1] = 0
			}

		}
		grille[ligne][colonne] = 0
	}
	wg.Wait()
}

func main() {
	grille := [TAILLE][TAILLE]int{
		{0, 0, 4, 0, 0, 9, 0, 0, 6},
		{0, 7, 0, 1, 0, 0, 8, 0, 0},
		{3, 0, 0, 0, 7, 0, 0, 5, 0},
		{7, 0, 0, 0, 3, 0, 0, 9, 0},
		{0, 0, 6, 0, 0, 1, 0, 0, 4},
		{0, 0, 0, 5, 0, 0, 2, 0, 0},
		{9, 0, 0, 0, 8, 0, 0, 0, 0},
		{0, 1, 0, 9, 0, 0, 0, 0, 0},
		{0, 0, 8, 0, 0, 2, 3, 0, 0},
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
