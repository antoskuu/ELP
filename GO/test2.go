package main

import (
	"fmt"
	"sync"
)

const TAILLE = 16
const TAILLE_BLOCK = 4

var solutions chan [TAILLE][TAILLE]int

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

func solve(grille [TAILLE][TAILLE]int, ligne int, colonne int, wg *sync.WaitGroup) {
	defer wg.Done()

	if ligne == TAILLE {
		// Envoyer la solution sur le canal
		solutions <- grille
		return
	}

	if colonne == TAILLE {
		var newWg sync.WaitGroup
		newWg.Add(1)
		go solve(grille, ligne+1, 0, &newWg)
		newWg.Wait()
		return
	}

	if grille[ligne][colonne] != 0 {
		var newWg sync.WaitGroup
		newWg.Add(1)
		go solve(grille, ligne, colonne+1, &newWg)
		newWg.Wait()
		return
	}

	for k := 1; k <= TAILLE; k++ {
		if absentSurBlock(k, grille, ligne, colonne) && absentSurColonne(k, grille, colonne) && absentSurLigne(k, grille, ligne) {
			// Créer une copie de la grille pour éviter de modifier l'original
			newGrille := makeCopy(grille)
			newGrille[ligne][colonne] = k
			var newWg sync.WaitGroup
			newWg.Add(1)
			go solve(newGrille, ligne, colonne+1, &newWg)
			newWg.Wait()
		}
	}
}

func makeCopy(original [TAILLE][TAILLE]int) [TAILLE][TAILLE]int {
	var copy [TAILLE][TAILLE]int
	for i := 0; i < TAILLE; i++ {
		for j := 0; j < TAILLE; j++ {
			copy[i][j] = original[i][j]
		}
	}
	return copy
}

func main() {
	grille := [TAILLE][TAILLE]int{
		{0, 0, 0, 1, 6, 0, 0, 0, 0, 7, 0, 3, 0, 0, 0, 0},
		{0, 6, 0, 0, 0, 5, 0, 4, 14, 0, 1, 0, 2, 11, 0, 13},
		{0, 0, 0, 7, 0, 3, 0, 13, 8, 0, 16, 4, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 6, 10, 0, 11, 3, 0, 8, 14},
		{0, 2, 0, 0, 0, 1, 11, 0, 3, 0, 10, 0, 0, 8, 0, 12},
		{0, 0, 12, 0, 0, 0, 6, 9, 0, 14, 0, 0, 10, 0, 5, 0},
		{0, 10, 0, 13, 0, 0, 3, 15, 0, 9, 0, 2, 0, 6, 0, 16},
		{1, 4, 7, 0, 13, 0, 0, 5, 0, 6, 0, 0, 9, 0, 0, 11},
		{0, 7, 0, 5, 9, 6, 1, 0, 2, 8, 3, 10, 0, 14, 0, 4},
		{0, 0, 8, 0, 0, 0, 0, 3, 5, 0, 0, 15, 0, 13, 0, 10},
		{6, 3, 0, 4, 0, 15, 0, 8, 7, 0, 0, 1, 0, 12, 0, 2},
		{0, 1, 0, 0, 4, 11, 0, 2, 0, 16, 0, 0, 8, 3, 6, 7},
		{0, 0, 0, 3, 0, 0, 2, 10, 0, 13, 0, 6, 0, 5, 0, 0},
		{5, 0, 0, 2, 0, 8, 0, 6, 10, 1, 0, 7, 0, 0, 12, 9},
		{7, 9, 1, 6, 0, 14, 0, 11, 0, 3, 0, 5, 0, 0, 10, 8},
		{0, 0, 0, 0, 0, 0, 9, 1, 4, 0, 0, 8, 0, 7, 2, 3},
	}

	fmt.Printf("Avant la modification : \n")
	for _, ligne := range grille {
		fmt.Println(ligne)
	}

	solutions = make(chan [TAILLE][TAILLE]int, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go solve(grille, 0, 0, &wg)
	wg.Wait()
	close(solutions)

	select {
	case solution := <-solutions:
		fmt.Printf("\nAprès la modification : \n")
		for _, ligne := range solution {
			fmt.Println(ligne)
		}
	default:
		fmt.Printf("\nErreur pas de solution")
	}
}
