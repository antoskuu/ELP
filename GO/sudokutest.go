package main

import (
	"fmt"
	"sync"
)

const TAILLE = 9
const TAILLE_BLOCK = 3

var wg sync.WaitGroup

func Copie(grille [TAILLE][TAILLE]int) [TAILLE][TAILLE]int {
	var copie [TAILLE][TAILLE]int
	for i := 0; i < TAILLE; i++ {
		for j := 0; j < TAILLE; j++ {
			copie[i][j] = grille[i][j]
		}
	}
	return copie
}

func absentSurLigne(k int, grille [TAILLE][TAILLE]int, ligne int, ch chan bool) {
	for colonne := 0; colonne < TAILLE; colonne++ {
		if grille[ligne][colonne] == k {
			ch <- false
		}
	}
	ch <- true
}

func absentSurColonne(k int, grille [TAILLE][TAILLE]int, colonne int, ch chan bool) {
	for ligne := 0; ligne < TAILLE; ligne++ {
		if grille[ligne][colonne] == k {
			ch <- false
		}
	}
	ch <- true
}

func absentSurBlock(k int, grille [TAILLE][TAILLE]int, ligne int, colonne int, ch chan bool) {
	ligne2 := ligne - ligne%TAILLE_BLOCK
	colonne2 := colonne - colonne%TAILLE_BLOCK

	for x := 0; x < TAILLE_BLOCK; x++ {
		for y := 0; y < TAILLE_BLOCK; y++ {
			if grille[ligne2+x][colonne2+y] == k {
				ch <- false
			}
		}
	}
	ch <- true
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
			ch := make(chan bool, 3)
			wg.Add(3)
			go absentSurBlock(k, grille, ligne, colonne, ch)

			go absentSurColonne(k, grille, colonne, ch)

			go absentSurLigne(k, grille, ligne, ch)

			go func(ch chan bool) {
				wg.Wait()
				close(ch)
			}(ch)
			toutVrai := true
			for i := 0; i < 3; i++ {
				result := <-ch
				if !result {
					toutVrai = false
				}
			}
			if toutVrai {
				grille[ligne][colonne] = k
				nouvelle_grille, solution := solve(grille, ligne, colonne+1)
				if solution {
					return nouvelle_grille, true
				}
				grille[ligne][colonne] = 0

			}
			wg.Done()
		}
		return grille, false
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
