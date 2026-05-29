package main

import (
	"errors"
	"fmt"
)

// Fonction calculer qui prend a, b et un operateur et retourne le résultat et l'erreur associée.
func calculer(a, b float64, operateur string) (float64, error) {
	switch operateur {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro impossible")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("opérateur inconnu: %s", operateur)
	}
}

func main() {
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
}
