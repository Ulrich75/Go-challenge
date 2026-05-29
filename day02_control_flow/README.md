# 🔁 Jour 2 : Les Structures de Contrôle

Dans ce deuxième jour, nous allons voir comment orienter le flux de notre programme. Go propose des structures de contrôle classiques mais avec quelques particularités notables visant à simplifier l'écriture du code.

---

## 🚦 Les Conditions (`if`, `else if`, `else`)

La syntaxe des conditions en Go ressemble à celle des autres langages, à deux différences près :
1. **Pas de parenthèses** autour de la condition.
2. **Les accolades `{}` sont obligatoires**, même s'il n'y a qu'une seule ligne de code.

```go
age := 18

if age >= 18 {
    fmt.Println("Vous êtes majeur.")
} else if age >= 16 {
    fmt.Println("Vous êtes presque majeur.")
} else {
    fmt.Println("Vous êtes mineur.")
}
```

### Initialisation dans le `if` (spécificité Go)
Go permet de déclarer et d'initialiser une variable directement dans la condition du `if`. Sa portée est alors limitée à ce bloc conditionnel.
```go
if score := calculerScore(); score > 50 {
    fmt.Printf("Bravo ! Score de %d suffisant.\n", score)
} // 'score' n'est plus accessible ici
```

---

## 🔄 La Boucle `for` (L'unique boucle de Go)

Contrairement à la plupart des langages, **Go n'a pas de mot-clé `while` ou `do-while`**. Il n'y a qu'une seule boucle : `for`. Elle s'utilise de trois façons différentes :

### 1. La boucle standard (style C)
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 2. Le style "while" (boucle conditionnelle)
Il suffit d'omettre l'initialisation et l'incrémentation.
```go
compteur := 1
for compteur <= 5 {
    fmt.Println(compteur)
    compteur++
}
```

### 3. La boucle infinie (avec `break` et `continue`)
Utile pour les serveurs ou les écoutes d'événements. Vous en sortez à l'aide de `break`.
```go
nombre := 1
for {
    if nombre > 5 {
        break // Quitte la boucle
    }
    fmt.Println(nombre)
	nombre++
}
```

---

## 🎛️ L'instruction `switch`

L'instruction `switch` est une alternative propre à une série de `if else`.
*   **Pas besoin de `break` !** En Go, l'exécution s'arrête automatiquement à la fin du cas correspondant (pas de comportement "fallthrough" par défaut).

```go
jour := "lundi"

switch jour {
case "lundi", "mardi", "mercredi", "jeudi", "vendredi":
    fmt.Println("C'est un jour de semaine.")
case "samedi", "dimanche":
    fmt.Println("C'est le week-end !")
default:
    fmt.Println("Jour invalide.")
}
```

---

## 🏋️ Exercice du Jour

Rendez-vous dans le dossier [day02_control_flow/exercise/main.go](file:///home/ulrich/workspaces/GO_LEARNING/day02_control_flow/exercise/main.go).

### Objectif de l'exercice :
Écrire un programme qui simule un **jeu de devinette de nombre** simple.
1. Le programme définit un "nombre secret" (par exemple `42`).
2. Il utilise une boucle pour simuler des tentatives d'un joueur.
3. À chaque tentative (simulation d'essais successifs fournis dans le code) :
    *   Si le nombre proposé est trop petit, afficher "Trop petit !".
    *   Si le nombre proposé est trop grand, afficher "Trop grand !".
    *   Si le nombre proposé est correct, afficher "Félicitations, vous avez trouvé !" et arrêter la boucle.

### Pour lancer l'exercice :
```bash
cd /home/ulrich/workspaces/GO_LEARNING/day02_control_flow/exercise
go run main.go
```
