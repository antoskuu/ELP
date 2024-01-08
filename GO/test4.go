package main

import (
	"fmt"
	"sync"
)

const TAILLE = 16
const TAILLE_BLOCK = 4

type Result struct {
	Grille   [TAILLE][TAILLE]int
	Solution bool
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

func solve(grille [TAILLE][TAILLE]int, ligne int, colonne int, solutionCh chan Result) {
	if ligne == TAILLE {
		solutionCh <- Result{Grille: grille, Solution: true}
		return
	} else if colonne == TAILLE {
		solve(grille, ligne+1, 0, solutionCh)
		return
	} else if grille[ligne][colonne] != 0 {
		solve(grille, ligne, colonne+1, solutionCh)
		return
	} else {
		for k := 1; k <= TAILLE; k++ {
			if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
				grille[ligne][colonne] = k
				solve(grille, ligne, colonne+1, solutionCh)
				grille[ligne][colonne] = 0
			}
		}
		return
	}
}

func solvesudokupartial(grille [TAILLE][TAILLE]int, ligne int, colonne int, value int, solutionCh chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	grille[ligne][colonne] = value
	solution := Result{}
	select {
	case solution = <-solutionCh:
	default:
	}
	if solution.Solution {
		fmt.Printf("Solution pour : %d\n", value)
		afficherGrille(solution.Grille)
	}
}

func solve_para(grille [TAILLE][TAILLE]int, ligne int, colonne int, solutionCh chan Result) {
	var wg sync.WaitGroup
	for k := 1; k <= TAILLE; k++ {
		if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
			wg.Add(1)
			go solvesudokupartial(grille, ligne, colonne, k, solutionCh, &wg)
		}
	}
	wg.Wait()
}

func afficherGrille(grille [TAILLE][TAILLE]int) {
	fmt.Printf("\nAprès la modification : \n")
	for _, ligne := range grille {
		fmt.Println(ligne)
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

	solutionCh := make(chan Result)
	go solve(grille, 0, 0, solutionCh)
	solution := <-solutionCh

	if solution.Solution {
		fmt.Printf("\nAprès la modification : \n")
		afficherGrille(solution.Grille)
	} else {
		fmt.Printf("\nErreur : pas de solution\n")
	}
}
