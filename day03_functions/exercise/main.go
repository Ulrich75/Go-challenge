package main

import (
	"errors"
	"fmt"
)

// TODO : Déclarer la fonction 'calculer' prenant 'a' (float64), 'b' (float64), et 'operateur' (string)
// Elle doit retourner deux valeurs : un float64 (résultat) et une error.
// Si l'opérateur est inconnu, retourner une erreur : errors.New("opérateur inconnu")
// Si le diviseur est 0 pour une division, retourner : errors.New("division par zéro impossible")
// Votre code de fonction ici :


func main() {
	// Voici une liste de tests pour valider votre fonction.
	// Décommentez le code ci-dessous une fois la fonction 'calculer' écrite.

	/*
	tests := []struct {
		a, b      float64
		operateur string
	}{
		{10, 5, "+"},
		{10, 5, "-"},
		{10, 5, "*"},
		{10, 0, "/"},
		{10, 2, "/"},
		{10, 5, "%"}, // opérateur inconnu
	}

	for _, t := range tests {
		res, err := calculer(t.a, t.b, t.operateur)
		if err != nil {
			fmt.Printf("Erreur pour %.0f %s %.0f : %v\n", t.a, t.operateur, t.b, err)
		} else {
			fmt.Printf("Résultat de %.0f %s %.0f = %.2f\n", t.a, t.operateur, t.b, res)
		}
	}
	*/
}
