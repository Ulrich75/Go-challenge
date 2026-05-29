package main

import (
	"errors"
	"fmt"
	"strings"
)

// Étape 1 : Définir l'interface 'Notificateur' contenant la méthode :
// Envoyer(message string) error
// TODO : Déclarer l'interface ici


// Étape 2 : Créer la structure 'Email' (champ : Adresse string) et implémenter la méthode 'Envoyer'.
// Si 'Adresse' ne contient pas de "@" (utilisez strings.Contains), retourner une erreur.
// TODO : Déclarer la structure Email et sa méthode Envoyer ici


// Étape 3 : Créer la structure 'SMS' (champ : Numero string) et implémenter la méthode 'Envoyer'.
// Si 'Numero' a une longueur inférieure à 10 (utilisez len(Numero)), retourner une erreur.
// TODO : Déclarer la structure SMS et sa méthode Envoyer ici


func main() {
	// Étape 4 : Créer un slice de 'Notificateur' contenant :
	// - Un Email valide
	// - Un Email invalide (sans @)
	// - Un SMS valide (au moins 10 chiffres/caractères)
	// - Un SMS invalide (trop court)
	// TODO : Déclarer le slice ici
	

	// Étape 5 : Parcourir le slice avec une boucle 'for range'
	// Appeler la méthode 'Envoyer' avec le message "Alerte système !" pour chaque notificateur.
	// Si une erreur se produit, l'afficher. Sinon, confirmer le bon envoi.
	// TODO : Écrire la boucle ici
}
