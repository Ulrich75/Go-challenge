# 📊 Jour 4 : Les Structures de Données

Aujourd'hui, nous allons étudier comment regrouper et structurer nos données en Go. Nous aborderons les Tableaux, les Slices (tranches), les Maps (dictionnaires), et les Structs (structures).

---

## 1. Les Tableaux (Arrays)

En Go, un tableau a une **taille fixe** déterminée à sa création, et tous ses éléments doivent avoir le même type.
```go
var tab [5]int // Un tableau de 5 entiers, initialisés à 0
tab[0] = 10
tab[1] = 20

// Initialisation directe
nombres := [3]string{"Go", "Python", "Rust"}
```
> En pratique, les tableaux fixes sont assez peu utilisés directement. On leur préfère les **Slices**.

---

## 2. Les Slices (Tranches)

Un Slice est une vue dynamique sur un tableau sous-jacent. Sa taille peut changer à tout moment.

### Déclaration d'un Slice
```go
// Directement avec des valeurs
notes := []float64{12.5, 15.0, 18.0}

// Avec la fonction make()
// make([]type, longueur, capacite_optionnelle)
panier := make([]string, 3) 
```

### Ajouter des éléments avec `append`
La fonction `append` ajoute des éléments à la fin d'un slice et renvoie le nouveau slice.
```go
var fruits []string // Slice vide (nil)
fruits = append(fruits, "Pomme")
fruits = append(fruits, "Banane", "Orange")
```

### Le parcours de collections avec `range`
Pour boucler sur un slice (ou un tableau), on utilise `range` qui fournit à chaque itération l'index et la valeur de l'élément :
```go
for index, valeur := range fruits {
    fmt.Printf("Index %d : %s\n", index, valeur)
}
// Pour ignorer l'index :
for _, valeur := range fruits { ... }
```

---

## 3. Les Maps (Dictionnaires / Clé-Valeur)

Une `map` associe des clés uniques à des valeurs.

### Déclaration et initialisation
```go
// Déclaration vide via make
ages := make(map[string]int)
ages["Thomas"] = 28
ages["Sophie"] = 31

// Déclaration littérale
capitales := map[string]string{
    "France":  "Paris",
    "Espagne": "Madrid",
}
```

### Accéder et tester l'existence d'une clé
Si une clé n'existe pas, la map renvoie la valeur par défaut du type (ex: 0 pour un int, "" pour un string). Pour vérifier l'existence réelle d'une clé :
```go
capitale, existe := capitales["Italie"]
if existe {
    fmt.Println("La capitale est :", capitale)
} else {
    fmt.Println("Clé non trouvée.")
}
```

### Supprimer une clé
```go
delete(capitales, "Espagne")
```

---

## 4. Les Structures (Structs)

Une structure (`struct`) est une collection de champs nommés. Elle permet de modéliser des objets du monde réel.

### Définition
```go
type Utilisateur struct {
    Nom    string
    Age    int
    Actif  bool
}
```

### Instanciation
```go
u1 := Utilisateur{
    Nom:   "Thomas",
    Age:   28,
    Actif: true,
}

// Accéder aux champs
fmt.Println(u1.Nom)
u1.Age = 29
```

---

## 🏋️ Exercice du Jour

Rendez-vous dans le dossier [day04_data_structures/exercise/main.go](file:///home/ulrich/workspaces/GO_LEARNING/day04_data_structures/exercise/main.go).

### Objectif de l'exercice :
Créer un gestionnaire d'inventaire de produits.
1. Définir une structure `Produit` ayant un `Nom` (string), un `Prix` (float64) et une `Quantite` (int).
2. Dans le `main`, créer un inventaire sous forme de slice contenant plusieurs produits.
3. Calculer et afficher la valeur totale estimée de tout le stock en magasin (somme des `Prix * Quantite` de chaque produit).
4. Afficher la liste complète des produits sous forme de tableau propre.

### Pour lancer l'exercice :
```bash
cd /home/ulrich/workspaces/GO_LEARNING/day04_data_structures/exercise
go run main.go
```
