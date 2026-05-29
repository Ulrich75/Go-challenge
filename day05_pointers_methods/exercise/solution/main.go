package main

import "fmt"

// Structure CompteBancaire
type CompteBancaire struct {
	Titulaire string
	Solde     float64
}

// Méthode Afficher avec un receveur par valeur (lecture seule)
func (c CompteBancaire) Afficher() {
	fmt.Printf("Compte de %-8s | Solde : %.2f €\n", c.Titulaire, c.Solde)
}

// Méthode Deposer avec un receveur par pointeur (modification)
func (c *CompteBancaire) Deposer(montant float64) {
	if montant > 0 {
		c.Solde += montant
		fmt.Printf("Dépôt de %.2f € effectué avec succès.\n", montant)
	} else {
		fmt.Println(" Le montant du dépôt doit être positif.")
	}
}

// Méthode Retirer avec un receveur par pointeur (modification + retour de succès)
func (c *CompteBancaire) Retirer(montant float64) bool {
	if montant <= 0 {
		fmt.Println("  Le montant du retrait doit être positif.")
		return false
	}

	if c.Solde >= montant {
		c.Solde -= montant
		fmt.Printf("Retrait de %.2f € effectué avec succès.\n", montant)
		return true
	}

	fmt.Printf(" Retrait de %.2f € refusé : solde insuffisant (%.2f € disponibles).\n", montant, c.Solde)
	return false
}

func main() {
	// Création du compte
	compte := CompteBancaire{
		Titulaire: "Alice",
		Solde:     100.0,
	}

	// Tests des méthodes
	compte.Afficher()

	compte.Deposer(50.0)
	compte.Afficher()

	// Tentative réussie
	reussi := compte.Retirer(30.0)
	fmt.Println("Retrait réussi ?", reussi)
	compte.Afficher()

	// Tentative refusée
	reussi = compte.Retirer(150.0)
	fmt.Println("Retrait réussi ?", reussi)
	compte.Afficher()
}
