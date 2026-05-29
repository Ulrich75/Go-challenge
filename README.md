# Apprendre le langage Go en 7 jours

Bienvenue dans votre espace d'apprentissage du langage Go ! Ce cours est conçu pour vous guider de manière progressive, depuis les concepts fondamentaux jusqu'à la création d'une application complète.

---

## Configuration de l'environnement

 Pour pouvoir exécuter la commande `go` simplement dans votre terminal sans saisir le chemin complet, vous pouvez ajouter Go à votre variable d'environnement `PATH`.

### Option A : Utilisation temporaire (dans la session de terminal actuelle)
Exécutez cette commande dans votre terminal :
```bash
export PATH=$PATH:/usr/local/go/bin
```

### Option B : Utilisation permanente (recommandée)
1. Ouvrez votre fichier de configuration de terminal (généralement `~/.bashrc` ou `~/.zshrc`) :
   ```bash
   nano ~/.bashrc
   ```
2. Ajoutez cette ligne à la toute fin du fichier :
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```
3. Sauvegardez le fichier (`Ctrl+O`, puis `Entrée`, et quittez avec `Ctrl+X`).
4. Rechargez le fichier de configuration :
   ```bash
   source ~/.bashrc
   ```

Pour vérifier que la configuration fonctionne, lancez la commande :
```bash
go version
```
Vous devriez voir s'afficher : `go version go1.26.1 linux/amd64`.

---

## Programme d'Apprentissage

Cliquez sur les liens ci-dessous pour accéder au cours de chaque journée :

*   **[Jour 1 : Introduction & Syntaxe de Base](file:///home/ulrich/workspaces/GO_LEARNING/day01_introduction/README.md)**
    *   *Au programme :* Pourquoi Go ? Structure d'un projet, Hello World, variables, constantes et types de données de base.
    *   *Exercice :* Déclaration de variables et calcul.

*   **[Jour 2 : Structures de Contrôle](file:///home/ulrich/workspaces/GO_LEARNING/day02_control_flow/README.md)**
    *   *Au programme :* Conditions (`if/else`), boucles (`for`), et sélection (`switch`).
    *   *Exercice :* Jeu de devinette de nombre.

*   **[Jour 3 : Les Fonctions](file:///home/ulrich/workspaces/GO_LEARNING/day03_functions/README.md)**
    *   *Au programme :* Déclarer des fonctions, paramètres, retours multiples et nommés, et le mot-clé `defer`.
    *   *Exercice :* Calculatrice modulaire robuste.

*   **[Jour 4 : Structures de Données](file:///home/ulrich/workspaces/GO_LEARNING/day04_data_structures/README.md)**
    *   *Au programme :* Tableaux (Arrays), tranches (Slices), dictionnaires (Maps) et Structures (Structs).
    *   *Exercice :* Gestionnaire d'inventaire de produits.

*   **[Jour 5 : Pointeurs & Méthodes](file:///home/ulrich/workspaces/GO_LEARNING/day05_pointers_methods/README.md)**
    *   *Au programme :* Comprendre les adresses mémoire, les pointeurs (`&` et `*`), et ajouter des méthodes à nos structures.
    *   *Exercice :* Simulation de transactions de compte bancaire.

*   **[Jour 6 : Interfaces & Gestion d'Erreurs](file:///home/ulrich/workspaces/GO_LEARNING/day06_interfaces_errors/README.md)**
    *   *Au programme :* Le polymorphisme avec les interfaces et la gestion idiomatique des erreurs avec le type `error`.
    *   *Exercice :* Calculateur de géométrie avec détection d'erreurs.

*   **[Jour 7 : Projet Final - CLI Todo List](file:///home/ulrich/workspaces/GO_LEARNING/day07_final_project/README.md)**
    *   *Au programme :* Manipulation de fichiers avec le package `os`, sérialisation JSON avec `encoding/json`, et synthèse de tous les jours précédents.
    *   *Exercice :* Compléter le code de l'application To-Do List interactive.

---

## Comment exécuter vos fichiers Go ?

1. Se déplacer dans le répertoire contenant le fichier :
   ```bash
   cd day01_introduction
   ```
2. Compiler et exécuter en une seule commande :
   ```bash
   go run hello.go
   ```
3. Ou compiler un exécutable binaire :
   ```bash
   go build hello.go
   ./hello
   ```
