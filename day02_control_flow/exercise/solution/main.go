package main

import "fmt"

func main() {
	secretNumber := 42
	attempts := []int{12, 55, 38, 45, 42, 10} // Liste des tentatives du joueur

	fmt.Println("Le nombre secret est caché. Début de la recherche...")

	// Boucle 'for range' pour parcourir chaque tentative de la liste
	for _, attempt := range attempts {
		if attempt < secretNumber {
			fmt.Printf("Tentative de %d : Trop petit !\n", attempt)
		} else if attempt > secretNumber {
			fmt.Printf("Tentative de %d : Trop grand !\n", attempt)
		} else {
			fmt.Printf("Tentative de %d : Félicitations ! Le nombre secret était bien %d.\n", attempt, secretNumber)
			break // Trouvé ! On arrête le jeu immédiatement
		}
	}
}
