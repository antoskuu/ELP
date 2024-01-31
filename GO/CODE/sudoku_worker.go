package main

import (
	"fmt"
	"sync"
	"time"
)

// constante à changer selon la taille de la grille
const TAILLE = 9
const TAILLE_BLOCK = 3

// Structure de données pour notre channel de pbs
type problem struct {
	Grille  [TAILLE][TAILLE]int // Grille à réssoudre
	ligne   int                 // Position en (i,j) du traitement
	colonne int
}

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

func printProb(prob problem) {
	for _, ligne := range prob.Grille {
		fmt.Println(ligne)
	}
	fmt.Printf("Colonne : %d\nLigne : %d\n", prob.colonne, prob.ligne)
}
func printProbChan(probChan chan problem) {
	fmt.Printf("Début du printProbChan\n")
	for prob := range probChan {
		for _, ligne := range prob.Grille {
			fmt.Println(ligne)
		}
		fmt.Printf("Colonne : %d\nLigne : %d\n", prob.colonne, prob.ligne)
	}
	fmt.Printf("Fin du printProbChan\n")

}
func printCase(p problem, id int) {
	fmt.Printf("id: %d, Grille[%d] = %d\n", id, p.ligne, p.Grille[p.ligne])
}
func solve(prob problem, problemChan chan problem, quit chan struct{}) {

	if prob.ligne == TAILLE {
		printProb(prob)
		close(quit)
		return

	} else if prob.colonne == TAILLE {

		go func() {
			problemChan <- problem{
				Grille:  prob.Grille,
				ligne:   prob.ligne + 1,
				colonne: 0,
			}
		}()
		return

	} else if prob.Grille[prob.ligne][prob.colonne] != 0 {
		go func() {
			problemChan <- problem{
				Grille:  prob.Grille,
				ligne:   prob.ligne,
				colonne: prob.colonne + 1,
			}
		}()
		return

	} else {
		for k := 1; k <= TAILLE; k++ {

			if absentSurBlock(k, prob.Grille, prob.ligne, prob.colonne) && absentSurColonne(k, prob.Grille, prob.colonne) && absentSurLigne(k, prob.Grille, prob.ligne) {

				go func() {
					prob.Grille[prob.ligne][prob.colonne] = k
					problemChan <- problem{
						Grille:  prob.Grille,
						ligne:   prob.ligne,
						colonne: (prob.colonne + 1),
					}
				}()
			}
		}
		return
	}
}

func worker(id int, ch chan problem, quit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case x := <-ch:
			// Received value from the channel
			solve(x, ch, quit)
			time.Sleep(time.Second * 2)

		case <-quit:
			// Received signal to quit
			fmt.Printf("Worker %d quitting\n", id)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan problem)
	quit := make(chan struct{})

	numWorkers := 8

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
	prob_base := problem{
		Grille:  grille,
		ligne:   0,
		colonne: 0,
	}
	go func() {
		ch <- prob_base
	}()
	// Start workers
	debut := time.Now()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, quit, &wg)
		time.Sleep(time.Second)
	}

	// Wait for all workers to finish
	wg.Wait()
	fin := time.Now()
	fmt.Printf("\nTemps d'execution : %.6fs \n", fin.Sub(debut).Seconds())
	fmt.Println("Main function finished")
}
