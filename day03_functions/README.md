# 🛠️ Jour 3 : Les Fonctions

Aujourd'hui, nous allons étudier l'un des piliers de Go : les fonctions. Go apporte des fonctionnalités modernes et pratiques comme les retours multiples et la planification d'exécution avec `defer`.

---

## Déclarer une fonction de base

En Go, une fonction est déclarée avec le mot-clé `func`. La syntaxe générale est la suivante :

```go
func nomFonction(parametre1 type1, parametre2 type2) typeRetour {
    // Corps de la fonction
    return valeur
}
```

### Exemple simple :
```go
func additionner(a int, b int) int {
    return a + b
}
```

### Raccourci de syntaxe :
Si plusieurs paramètres consécutifs ont le même type, vous pouvez ne spécifier le type qu'à la fin :
```go
func additionner(a, b int) int { // a et b sont tous deux des int
    return a + b
}
```

---

## Retours multiples (La force de Go !)

Contrairement à beaucoup de langages qui nécessitent de créer un objet ou un tableau pour renvoyer plusieurs valeurs, **Go prend en charge les retours multiples nativement**. C'est particulièrement utilisé pour renvoyer un résultat ET une erreur.

```go
func diviser(a, b float64) (float64, bool) {
    if b == 0 {
        return 0, false // Retourne 0 et false (division invalide)
    }
    return a / b, true  // Retourne le résultat et true
}
```

### Utilisation et ignorer une valeur (`_`) :
Lorsqu'une fonction retourne plusieurs valeurs, vous devez toutes les réceptionner. Si l'une d'elles ne vous intéresse pas, vous pouvez l'ignorer en utilisant le caractère blanc `_`.
```go
resultat, ok := diviser(10, 2)

// Si l'on ne veut pas de la variable 'ok' :
resultatSeul, _ := diviser(20, 4)
```

---

## Retours nommés

Go permet de nommer les variables de retour directement dans la signature de la fonction. Elles se comportent alors comme des variables locales initialisées à leur valeur par défaut.
```go
func rectangleInfos(largeur, hauteur float64) (perimetre float64, aire float64) {
    perimetre = 2 * (largeur + hauteur)
    aire = largeur * hauteur
    return // "Naked return" : retourne automatiquement perimetre et aire
}
```
*Conseil : N'utilisez les retours nommés que pour les fonctions courtes pour ne pas nuire à la lisibilité.*

---

## L'instruction `defer`

Le mot-clé `defer` permet de reporter (différer) l'exécution d'une fonction jusqu'à la fin de la fonction appelante. Les appels différés sont exécutés dans l'ordre inverse de leur déclaration (LIFO - Last In, First Out).

Cela est couramment utilisé pour nettoyer les ressources (fermer un fichier, déverrouiller une base de données, etc.).

```go
func exemple() {
    defer fmt.Println("Étape 3 : Exécuté à la fin !")
    fmt.Println("Étape 1")
    fmt.Println("Étape 2")
}
// Affichage dans l'ordre : Étape 1, Étape 2, puis Étape 3.
```

---

## Exercice du Jour

Rendez-vous dans le dossier [day03_functions/exercise/main.go](file:///home/ulrich/workspaces/GO_LEARNING/day03_functions/exercise/main.go).

### Objectif de l'exercice :
Créer une mini-calculatrice modulaire :
1. Écrire une fonction `calculer(a, b float64, operateur string) (float64, error)`.
2. Gérer les cas d'addition (`+`), soustraction (`-`), multiplication (`*`), et division (`/`).
3. Pour la division, retourner une erreur si le diviseur est égal à `0` (utilisez `errors.New` du package standard `errors` ou `fmt.Errorf`).
4. Dans la fonction `main`, tester différents cas et afficher le résultat ou l'erreur retournée.

### Pour lancer l'exercice :
```bash
cd /home/ulrich/workspaces/GO_LEARNING/day03_functions/exercise
go run main.go
```
