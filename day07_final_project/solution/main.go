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

	// Charger les tâches existantes au démarrage
	err := chargerTaches()
	if err != nil {
		fmt.Println("Erreur lors du chargement des tâches :", err)
	}

	fmt.Println("🌟 BIENVENUE DANS VOTRE GESTIONNAIRE DE TÂCHES GO ! 🌟")

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
			fmt.Println("⚠️ Choix invalide. Veuillez saisir un nombre entre 1 et 5.")
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
	for i, t := range listeTaches {
		statut := "[ ]"
		if t.Fait {
			statut = "[X]"
		}
		fmt.Printf("%d. %s %s\n", i+1, statut, t.Titre)
	}
}

// Ajoute une tâche à la liste
func ajouterTache() {
	titre := lireSaisie("Entrez le titre de la tâche : ")
	if titre == "" {
		fmt.Println("⚠️ Le titre ne peut pas être vide.")
		return
	}

	nouvelleTache := Tache{
		Titre: titre,
		Fait:  false,
	}

	listeTaches = append(listeTaches, nouvelleTache)
	fmt.Printf("✅ Tâche \"%s\" ajoutée avec succès !\n", titre)
}

// Marque une tâche comme terminée
func terminerTache() {
	saisie := lireSaisie("Entrez le numéro de la tâche à terminer : ")
	num, err := strconv.Atoi(saisie)
	if err != nil || num <= 0 || num > len(listeTaches) {
		fmt.Println("⚠️ Numéro de tâche invalide.")
		return
	}

	index := num - 1
	listeTaches[index].Fait = true
	fmt.Printf("✅ Tâche \"%s\" marquée comme terminée !\n", listeTaches[index].Titre)
}

// Supprime une tâche de la liste
func supprimerTache() {
	saisie := lireSaisie("Entrez le numéro de la tâche à supprimer : ")
	num, err := strconv.Atoi(saisie)
	if err != nil || num <= 0 || num > len(listeTaches) {
		fmt.Println("⚠️ Numéro de tâche invalide.")
		return
	}

	index := num - 1
	titreSupprime := listeTaches[index].Titre

	// Supprime l'élément à l'index donné
	listeTaches = append(listeTaches[:index], listeTaches[index+1:]...)
	fmt.Printf("❌ Tâche \"%s\" supprimée avec succès !\n", titreSupprime)
}

// Enregistre les tâches dans le fichier JSON
func sauvegarderTaches() error {
	donnees, err := json.MarshalIndent(listeTaches, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fichierTodo, donnees, 0644)
	if err != nil {
		return err
	}

	fmt.Println("💾 Tâches enregistrées avec succès dans", fichierTodo)
	return nil
}

// Charge les tâches depuis le fichier JSON s'il existe
func chargerTaches() error {
	if _, err := os.Stat(fichierTodo); os.IsNotExist(err) {
		// Pas de fichier, on démarre avec une liste vide
		return nil
	}

	donnees, err := os.ReadFile(fichierTodo)
	if err != nil {
		return err
	}

	err = json.Unmarshal(donnees, &listeTaches)
	if err != nil {
		return err
	}

	fmt.Printf("📂 %d tâche(s) chargée(s) depuis le disque.\n", len(listeTaches))
	return nil
}

// Fonction utilitaire pour lire facilement le texte saisi par l'utilisateur
func lireSaisie(invite string) string {
	fmt.Print(invite)
	texte, _ := lecteurClavier.ReadString('\n')
	// Nettoyer les retours à la ligne selon l'OS
	texte = strings.Replace(texte, "\r", "", -1)
	texte = strings.Replace(texte, "\n", "", -1)
	return strings.TrimSpace(texte)
}
