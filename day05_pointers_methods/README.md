# 📍 Jour 5 : Pointeurs & Méthodes

Aujourd'hui, nous allons aborder une notion fondamentale pour la gestion de la mémoire et la performance en Go : les **Pointeurs**, puis nous verrons comment attacher des comportements (des **Méthodes**) à nos structures.

---

## 1. Qu'est-ce qu'un Pointeur ?

Chaque variable déclarée dans un programme est stockée à une adresse précise dans la mémoire de l'ordinateur. 
*   Un **pointeur** est une variable qui stocke **l'adresse mémoire** d'une autre variable.

### Les deux opérateurs clés :
*   `&` (Adresse de) : Récupère l'adresse mémoire d'une variable.
*   `*` (Valeur pointée) : Accède ou modifie la valeur stockée à l'adresse pointée (c'est la *déréférenciation*).

```go
age := 25
var ptr *int = &age // ptr contient l'adresse mémoire de age

fmt.Println(ptr)  // Affiche l'adresse mémoire (ex: 0xc0000140e8)
fmt.Println(*ptr) // Affiche la valeur pointée : 25

*ptr = 30         // Modifie directement la valeur dans la mémoire de age
fmt.Println(age)  // Affiche maintenant : 30 !
```

---

## 2. Pourquoi utiliser des pointeurs ?

### Éviter les copies coûteuses en mémoire
Par défaut, lorsque vous passez une variable à une fonction en Go, **la valeur est copiée**. Si vous passez une grosse structure, Go va copier toutes ses données. En passant un pointeur (qui fait seulement 8 octets), vous ne copiez que l'adresse.

### Modifier une valeur passée en paramètre
Si vous voulez qu'une fonction modifie une variable externe, vous devez lui passer son pointeur.
```go
func doublerSansPointeur(n int) {
    n = n * 2 // Modifie la copie locale de n
}

func doublerAvecPointeur(n *int) {
    *n = *n * 2 // Modifie la variable d'origine en mémoire
}
```

---

## 3. Les Méthodes sur les Structures

En Go, il n'y a pas de classes comme en programmation orientée objet classique (Java, C++, Python). À la place, on définit des **méthodes** attachées à des structures.

Une méthode est simplement une fonction qui possède un paramètre spécial appelé **receveur** (placé avant le nom de la fonction).

### Déclaration d'une méthode
```go
type Rectangle struct {
    Largeur, Hauteur float64
}

// Le receveur est (r Rectangle)
func (r Rectangle) Aire() float64 {
    return r.Largeur * r.Hauteur
}
```
*Appel de la méthode :*
```go
rect := Rectangle{Largeur: 10, Hauteur: 5}
fmt.Println(rect.Aire()) // Affiche 50
```

---

## 4. Receveurs par Valeur vs Receveurs par Pointeur

Il existe deux types de receveurs pour les méthodes :

### A. Receveur par valeur `(r Rectangle)`
*   La méthode reçoit une copie de la structure.
*   Toute modification des champs à l'intérieur de la méthode ne persiste pas en dehors.

### B. Receveur par pointeur `(r *Rectangle)`
*   La méthode reçoit l'adresse de la structure.
*   **Permet de modifier les champs de la structure d'origine.**
*   Évite de copier la structure en mémoire.

```go
type Joueur struct {
    Nom   string
    Score int
}

// Receveur par pointeur car on veut modifier le score
func (j *Joueur) AugmenterScore(points int) {
    j.Score += points // Modifie le joueur d'origine
}
```

---

## 🏋️ Exercice du Jour

Rendez-vous dans le dossier [day05_pointers_methods/exercise/main.go](file:///home/ulrich/workspaces/GO_LEARNING/day05_pointers_methods/exercise/main.go).

### Objectif de l'exercice :
Modéliser un compte bancaire avec des opérations sécurisées.
1. Définir une structure `CompteBancaire` avec les champs `Titulaire` (string) et `Solde` (float64).
2. Définir une méthode `Afficher()` (receveur par valeur) qui affiche les détails du compte.
3. Définir une méthode `Deposer(montant float64)` (receveur par pointeur) qui ajoute de l'argent au solde.
4. Définir une méthode `Retirer(montant float64) bool` (receveur par pointeur) qui retire de l'argent s'il y a assez de provisions (renvoie `true` si le retrait a réussi, `false` sinon).
5. Tester ces méthodes dans la fonction `main`.

### Pour lancer l'exercice :
```bash
cd /home/ulrich/workspaces/GO_LEARNING/day05_pointers_methods/exercise
go run main.go
```
