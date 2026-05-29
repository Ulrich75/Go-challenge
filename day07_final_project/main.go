package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Définition de la structure d'une Tâche
type Tache struct {
	Titre string `json:"titre"`
	Fait  bool   `json:"fait"`
}

const fichierTodo = "todos.json"

// Variables globales pour stocker nos tâches et le lecteur clavier
var listeTaches []Tache
var lecteurClavier *bufio.Reader

func main() {
	lecteurClavier = bufio.NewReader(os.Stdin)

	// Étape 1 : Charger les tâches existantes
	err := chargerTaches()
	if err != nil {
		fmt.Println("Erreur lors du chargement des tâches :", err)
	}

	fmt.Println("BIENVENUE DANS VOTRE GESTIONNAIRE DE TÂCHES GO ! ")

	for {
		afficherMenu()
		choix := lireSaisie("Votre choix : ")

		switch choix {
		case "1":
			listerTaches()
		case "2":
			ajouterTache()
		case "3":
			terminerTache()
		case "4":
			supprimerTache()
		case "5":
			fmt.Println("Sauvegarde des tâches...")
			err := sauvegarderTaches()
			if err != nil {
				fmt.Println("Erreur lors de la sauvegarde :", err)
			}
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println(" Choix invalide. Veuillez saisir un nombre entre 1 et 5.")
		}
		fmt.Println("\n-------------------------------------------")
	}
}

// Affiche le menu dans la console
func afficherMenu() {
	fmt.Println("\nMenu :")
	fmt.Println("1. Lister les tâches")
	fmt.Println("2. Ajouter une tâche")
	fmt.Println("3. Marquer une tâche comme terminée")
	fmt.Println("4. Supprimer une tâche")
	fmt.Println("5. Sauvegarder et Quitter")
}

// Affiche toutes les tâches de la liste
func listerTaches() {
	if len(listeTaches) == 0 {
		fmt.Println("Aucune tâche dans la liste. C'est le moment de se détendre !")
		return
	}

	fmt.Println("\n=== Vos Tâches ===")
	// TODO : Parcourir le slice 'listeTaches'
	// Afficher chaque tâche sous le format : "Index. [ ] Titre" ou "Index. [X] Titre"
	// Note : Les index affichés pour l'utilisateur doivent commencer à 1 (index + 1)
	// Votre code ici :

}

// Ajoute une tâche à la liste
func ajouterTache() {
	titre := lireSaisie("Entrez le titre de la tâche : ")
	if titre == "" {
		fmt.Println("Le titre ne peut pas être vide.")
		return
	}

	// TODO : Créer une nouvelle instance de 'Tache'
	// L'ajouter à la fin du slice 'listeTaches' en utilisant 'append'
	// Afficher un message de confirmation
	// Votre code ici :

}

// Marque une tâche comme terminée
func terminerTache() {
	saisie := lireSaisie("Entrez le numéro de la tâche à terminer : ")
	num, err := strconv.Atoi(saisie)
	if err != nil || num <= 0 || num > len(listeTaches) {
		fmt.Println(" Numéro de tâche invalide.")
		return
	}

	// TODO : Modifier le champ 'Fait' de la tâche sélectionnée pour le passer à true
	// Rappel : L'utilisateur saisit un nombre commençant à 1, l'index du slice commence à 0.
	// Afficher un message de confirmation
	// Votre code ici :

}

// Supprime une tâche de la liste
func supprimerTache() {
	saisie := lireSaisie("Entrez le numéro de la tâche à supprimer : ")
	num, err := strconv.Atoi(saisie)
	if err != nil || num <= 0 || num > len(listeTaches) {
		fmt.Println("Numéro de tâche invalide.")
		return
	}

	index := num - 1

	// TODO : Supprimer la tâche du slice 'listeTaches'
	// Astuce pour supprimer un élément d'un slice en Go à l'index 'i' :
	// listeTaches = append(listeTaches[:index], listeTaches[index+1:]...)
	// Afficher un message de confirmation
	// Votre code ici :

}

// Enregistre les tâches dans le fichier JSON
func sauvegarderTaches() error {
	// TODO : Sérialiser le slice 'listeTaches' en JSON.
	// Utilisez json.MarshalIndent(listeTaches, "", "  ") pour un formatage lisible.
	// Écrire les données obtenues dans le fichier spécifié par 'fichierTodo'
	// en utilisant os.WriteFile avec les permissions 0644.
	// Votre code ici :

	return nil // Remplacez ou ajustez selon votre implémentation
}

// Charge les tâches depuis le fichier JSON s'il existe
func chargerTaches() error {
	// Vérifier si le fichier existe
	if _, err := os.Stat(fichierTodo); os.IsNotExist(err) {
		// Pas d'erreur, le fichier n'existe juste pas encore
		return nil
	}

	// TODO : Lire le contenu complet du fichier 'fichierTodo' en utilisant os.ReadFile.
	// Désérialiser les données lues (JSON) dans notre variable globale 'listeTaches' en utilisant json.Unmarshal.
	// Votre code ici :

	return nil // Remplacez ou ajustez selon votre implémentation
}

// Fonction utilitaire pour lire facilement le texte saisi par l'utilisateur
func lireSaisie(invite string) string {
	fmt.Print(invite)
	texte, _ := lecteurClavier.ReadString('\n')
	// Nettoyer les retours à la ligne selon l'OS (\r\n pour Windows, \n pour Linux)
	texte = strings.Replace(texte, "\r", "", -1)
	texte = strings.Replace(texte, "\n", "", -1)
	return strings.TrimSpace(texte)
}
