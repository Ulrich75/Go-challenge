package main

import "fmt"

func main() {
	// Étape 1 : Déclarer une variable 'nomEleve' (string) avec pour valeur "Thomas".
	nomEleve := "Thomas"

	// Étape 2 : Déclarer deux variables 'noteMaths' (float64) et 'noteFrancais' (float64)
	// avec des notes de votre choix (ex: 14.5 et 16.0).
	noteMaths := 14.5
	noteFrancais := 16.0

	// Étape 3 : Calculer la moyenne de ces deux notes dans une variable 'moyenne' (float64).
	moyenne := (noteMaths + noteFrancais) / 2.0

	// Étape 4 : Afficher le résultat formaté dans la console.
	// Le message attendu est : "L'élève Thomas a obtenu une moyenne de 15.25/20."
	fmt.Printf("L'élève %s a obtenu une moyenne de %.2f/20.\n", nomEleve, moyenne)
}
