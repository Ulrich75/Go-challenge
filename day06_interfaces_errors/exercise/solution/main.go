package main

import (
	"errors"
	"fmt"
	"strings"
)

// Interface Notificateur
type Notificateur interface {
	Envoyer(message string) error
}

// Structure Email
type Email struct {
	Adresse string
}

// Implémentation de Notificateur pour Email
func (e Email) Envoyer(message string) error {
	if !strings.Contains(e.Adresse, "@") {
		return fmt.Errorf("adresse email invalide : %s (symbole @ manquant)", e.Adresse)
	}
	fmt.Printf("[Email] Envoyé à <%s> : %s\n", e.Adresse, message)
	return nil
}

// Structure SMS
type SMS struct {
	Numero string
}

// Implémentation de Notificateur pour SMS
func (s SMS) Envoyer(message string) error {
	if len(s.Numero) < 10 {
		return fmt.Errorf("numéro SMS invalide : %s (doit faire au moins 10 caractères)", s.Numero)
	}
	fmt.Printf("[SMS] Envoyé au %s : %s\n", s.Numero, message)
	return nil
}

func main() {
	// Création du slice de notificateurs (contient des emails et des sms)
	notificateurs := []Notificateur{
		Email{Adresse: "thomas@exemple.com"},
		Email{Adresse: "adresse_invalide.com"},
		SMS{Numero: "+33612345678"},
		SMS{Numero: "0612"}, // numéro trop court
	}

	messageAlerte := "Alerte système : Fuite d'eau dans la salle des serveurs !"

	fmt.Println("=== Envoi des Notifications ===")
	for i, n := range notificateurs {
		fmt.Printf("\nDestinataire #%d :\n", i+1)
		err := n.Envoyer(messageAlerte)
		if err != nil {
			fmt.Printf("❌ Échec de l'envoi : %v\n", err)
		} else {
			fmt.Println("✅ Message envoyé avec succès !")
		}
	}
}
