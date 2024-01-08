package main

import (
	"fmt"
	"sync"
)

const TAILLE = 9
const TAILLE_BLOCK = 3

var wg sync.WaitGroup
var mu sync.Mutex
var solutionFound bool

func Copie(grille [TAILLE][TAILLE]int) [TAILLE][TAILLE]int {
	var copie [TAILLE][TAILLE]int
	for i := 0; i < TAILLE; i++ {
		for j := 0; j < TAILLE; j++ {
			copie[i][j] = grille[i][j]
		}
	}
	return copie
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

func solve(grille [TAILLE][TAILLE]int, ligne int, colonne int) {
	defer wg.Done()

	if solutionFound {
		return
	}

	if ligne == TAILLE {
		mu.Lock()
		solutionFound = true
		mu.Unlock()
		fmt.Println("Solution trouvée par une goroutine!")
		for _, ligne := range grille {
			fmt.Println(ligne)
		}
		return
	}

	if colonne == TAILLE {
		solve(grille, ligne+1, 0)
		return
	}

	if grille[ligne][colonne] != 0 {
		solve(grille, ligne, colonne+1)
		return
	}

	for k := 1; k <= TAILLE; k++ {
		if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
			grille[ligne][colonne] = k
			wg.Add(1)
			go solve(Copie(grille), ligne, colonne)
		}
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

	wg.Add(1)
	go solve(grille, 0, 0)

	wg.Wait()
}
