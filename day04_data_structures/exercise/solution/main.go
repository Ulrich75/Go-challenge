package main

import "fmt"

// Structure Produit
type Produit struct {
	Nom      string
	Prix     float64
	Quantite int
}

func main() {
	// Création du slice de produits
	inventaire := []Produit{
		{Nom: "Ordinateur Portable", Prix: 899.99, Quantite: 5},
		{Nom: "Souris Sans Fil", Prix: 25.50, Quantite: 15},
		{Nom: "Clavier Mécanique", Prix: 79.90, Quantite: 8},
		{Nom: "Écran 27 Pouces", Prix: 199.00, Quantite: 4},
	}

	fmt.Println("=== Inventaire du Magasin ===")
	valeurTotale := 0.0

	// Parcours de l'inventaire avec range
	for _, p := range inventaire {
		valeurProduit := p.Prix * float64(p.Quantite)
		valeurTotale += valeurProduit

		fmt.Printf("- %s | Prix: %.2f € | Quantité: %d | Valeur Stock: %.2f €\n", 
			p.Nom, p.Prix, p.Quantite, valeurProduit)
	}

	fmt.Printf("\nValeur totale de l'inventaire : %.2f €\n", valeurTotale)
}
