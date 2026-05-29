# Jour 1 : Introduction & Syntaxe de Base

Bienvenue dans le premier jour de votre apprentissage de Go ! Aujourd'hui, nous allons comprendre pourquoi Go existe, comment est structuré un programme, et comment manipuler les variables et types de données fondamentaux.

---

## Pourquoi Go ?

Créé par Google (Robert Griesemer, Rob Pike, et Ken Thompson), Go (ou Golang) a été conçu pour résoudre des problèmes de développement à grande échelle. Ses points forts sont :
1. **La simplicité** : Go possède un nombre minimal de mots-clés (seulement 25), ce qui le rend très rapide à apprendre.
2. **La performance** : C'est un langage compilé directement en code machine (comme C/C++), ce qui le rend extrêmement rapide.
3. **Le typage fort et statique** : Le compilateur détecte la majorité des erreurs avant même que le code ne tourne.
4. **La concurrence native** : Il est conçu pour le multithreading moderne grâce aux "Goroutines" et aux "Channels" (que nous verrons plus tard).

---

## 🔬 Anatomie d'un programme Go

Regardons le fichier d'exemple `hello.go` :
```go
package main

import "fmt"

func main() {
    fmt.Println("Bonjour, le monde !")
}
```

Détaillons chaque ligne :
*   `package main` : Définit le nom du paquet auquel appartient le fichier. Le paquet `main` est spécial : il indique à Go que ce programme produit un exécutable et qu'il contient le point d'entrée.
*   `import "fmt"` : Importe le paquet standard `fmt` (abréviation de *format*), utilisé pour formater le texte et l'afficher dans le terminal.
*   `func main()` : C'est la fonction principale. Tout programme exécutable en Go commence par exécuter le code contenu dans cette fonction.
*   `fmt.Println(...)` : Appelle la fonction `Println` du paquet `fmt` pour afficher du texte suivi d'un retour à la ligne.

---

##  Variables et Constantes

En Go, il y a plusieurs manières de déclarer des variables.

### 1. La déclaration explicite (avec `var`)
```go
var nom string = "Ulrich"
var age int = 25
```

### 2. L'inférence de type
Si vous donnez une valeur initiale, Go peut deviner (inférer) le type automatiquement :
```go
var pays = "France" // Go comprend que c'est un string
```

### 3. La déclaration courte (avec `:=`)
C'est la syntaxe la plus couramment utilisée à l'intérieur des fonctions :
```go
ville := "Paris" // Équivalent à var ville string = "Paris"
note := 18.5     // Équivalent à var note float64 = 18.5
```
> **Attention :** L'opérateur `:=` ne peut être utilisé qu'à l'intérieur d'une fonction (`func`). Pour les variables globales (au niveau du package), vous devez obligatoirement utiliser le mot-clé `var`.

### 4. Les Constantes (avec `const`)
Une constante est une valeur qui ne peut pas être modifiée après sa création.
```go
const Pi = 3.14159
```

---

## Les Types de Données Fondamentaux

*   `string` : Chaînes de caractères. Exemple: `"Bonjour Go"`
*   `int`, `int8`, `int16`, `int32`, `int64` : Nombres entiers (signés).
*   `uint`, `uint8`, `uint16`, `uint32`, `uint64` : Nombres entiers positifs uniquement (non signés).
*   `float32`, `float64` : Nombres à virgule flottante (décimaux).
*   `bool` : Valeurs booléennes (`true` ou `false`).

---

## Formater l'affichage avec `fmt.Printf`

Au lieu de `Println`, vous pouvez utiliser `Printf` pour insérer des variables dans une chaîne grâce à des "verbes" de formatage :
```go
nom := "Alex"
age := 30
fmt.Printf("Mon nom est %s et j'ai %d ans.\n", nom, age)
```
Quelques verbes utiles :
*   `%s` : pour les chaînes de caractères (`string`)
*   `%d` : pour les entiers (`int`)
*   `%f` : pour les décimaux (`float64`) (utilisez `%.2f` pour limiter à 2 chiffres après la virgule)
*   `%T` : affiche le *type* de la variable
*   `%v` : affiche la valeur par défaut (très utile pour débugger)

---

## Exercice du Jour

Rendez-vous dans le dossier [day01_introduction/exercise/main.go](file:///home/ulrich/workspaces/GO_LEARNING/day01_introduction/exercise/main.go) pour réaliser votre premier exercice.

### Objectif de l'exercice :
1. Déclarer plusieurs variables (nom, âge, note en mathématiques, note en français).
2. Calculer la moyenne des notes.
3. Afficher un message formaté présentant l'élève et sa moyenne avec 2 chiffres après la virgule.

### Pour lancer l'exercice :
```bash
cd /home/ulrich/workspaces/GO_LEARNING/day01_introduction/exercise
go run main.go
```
