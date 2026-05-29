# Jour 7 : Projet Final - CLI Todo List

Félicitations pour être arrivé au Jour 7 ! Aujourd'hui, vous allez rassembler toutes les connaissances acquises pour construire une application interactive en ligne de commande (CLI) de gestion de tâches (**Todo List**).

Cette application enregistrera les tâches dans un fichier JSON local pour ne pas les perdre lorsque le programme s'arrête.

---

## Nouveaux concepts pour le projet

Pour mener à bien ce projet, nous introduisons deux packages essentiels de la bibliothèque standard de Go :

### 1. La manipulation des fichiers (`os`)
Pour enregistrer et lire des données sur le disque :
*   `os.WriteFile(filename, data, perm)` : Permet d'écrire un tableau d'octets (`[]byte`) dans un fichier.
*   `os.ReadFile(filename)` : Permet de lire le contenu complet d'un fichier sous forme de tableau d'octets.

### 2. Le format JSON (`encoding/json`)
Go dispose d'un excellent support pour la sérialisation (conversion en texte) et la désérialisation (reconstruction en objets Go) du JSON.

```go
type Tache struct {
    Titre string `json:"titre"` // Balise JSON pour spécifier le nom dans le fichier
    Fait  bool   `json:"fait"`
}

// 1. Struct -> JSON (Sérialisation)
liste := []Tache{{Titre: "Apprendre Go", Fait: false}}
donneesJSON, err := json.MarshalIndent(liste, "", "  ") // Format propre indenté

// 2. JSON -> Struct (Désérialisation)
var listeChargee []Tache
err := json.Unmarshal(donneesJSON, &listeChargee)
```

---

##  Fonctionnalités attendues de l'application

Votre application Todo List offrira un menu interactif dans le terminal :
1. **Afficher** la liste des tâches (avec leur numéro, statut `[ ]` ou `[X]`, et titre).
2. **Ajouter** une nouvelle tâche en saisissant son titre.
3. **Marquer une tâche comme terminée** en saisissant son numéro.
4. **Supprimer** une tâche en saisissant son numéro.
5. **Sauvegarder et Quitter**.

---

## Structure du Code de l'Exercice

Le fichier [day07_final_project/main.go](file:///home/ulrich/workspaces/GO_LEARNING/day07_final_project/main.go) contient déjà la structure générale du programme, y compris la boucle interactive principale et les fonctions de lecture de saisie clavier.

Votre travail consiste à implémenter les fonctionnalités clés repérées par les commentaires `// TODO`.

### Pour lancer l'application :
```bash
cd /home/ulrich/workspaces/GO_LEARNING/day07_final_project
go run main.go
```
Puis, laissez-vous guider par le menu interactif !
