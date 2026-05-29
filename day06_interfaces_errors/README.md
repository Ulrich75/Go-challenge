# 🛡️ Jour 6 : Interfaces & Gestion d'Erreurs

Aujourd'hui, nous abordons deux sujets cruciaux de l'architecture logicielle en Go : le polymorphisme grâce aux **Interfaces** et la philosophie unique de Go concernant le traitement des **Erreurs**.

---

## 1. Les Interfaces en Go

Une interface définit un ensemble de signatures de méthodes. Tout type qui possède (implémente) ces méthodes satisfait automatiquement l'interface.

### La magie de Go : L'implémentation implicite !
Contrairement à Java, TypeScript ou C#, **il n'y a pas de mot-clé `implements` en Go**. Un type implémente une interface simplement en définissant ses méthodes. C'est ce qu'on appelle le "duck typing" : *« Si ça marche comme un canard et que ça cancane comme un canard, alors c'est un canard. »*

### Exemple d'interface :
```go
// Définition de l'interface
type Forme interface {
    Aire() float64
}

// Une structure Carré
type Carre struct {
    Cote float64
}
// Carré implémente Forme de manière implicite car il possède la méthode Aire()
func (c Carre) Aire() float64 {
    return c.Cote * c.Cote
}

// Une structure Cercle
type Cercle struct {
    Rayon float64
}
// Cercle implémente également Forme
func (c Cercle) Aire() float64 {
    return 3.14159 * c.Rayon * c.Rayon
}
```

### Utilisation de l'interface :
Nous pouvons maintenant écrire une fonction générique acceptant n'importe quel type satisfaisant `Forme` :
```go
func AfficherAire(f Forme) {
    fmt.Printf("L'aire de cette forme est : %.2f\n", f.Aire())
}
```

---

## 2. Le type `any` (Interface vide)

L'interface vide `interface{}` ne contient aucune méthode. Par conséquent, **tous les types l'implémentent**. 
Depuis Go 1.18, `any` est un alias officiel pour `interface{}`. Il permet de stocker n'importe quel type de données.

```go
var coffreFort any

coffreFort = "Un texte"
coffreFort = 42
coffreFort = true
```
*Note : Utilisez `any` avec modération car cela désactive la vérification des types par le compilateur.*

---

## 3. La gestion d'erreurs en Go

En Go, il n'y a pas de blocs `try / catch`. Les erreurs sont traitées comme des valeurs de retour normales. La convention veut que les fonctions renvoient l'erreur comme dernière valeur de retour.

Le type `error` est en réalité une interface intégrée très simple :
```go
type error interface {
    Error() string
}
```

### Motif classique de gestion d'erreur :
```go
func verifierAge(age int) error {
    if age < 0 {
        return errors.New("l'âge ne peut pas être négatif")
    }
    return nil // nil signifie qu'aucune erreur ne s'est produite
}

func main() {
    err := verifierAge(-5)
    if err != nil {
        fmt.Println("Erreur détectée :", err)
        return
    }
    fmt.Println("Âge valide !")
}
```

---

## 🏋️ Exercice du Jour

Rendez-vous dans le dossier [day06_interfaces_errors/exercise/main.go](file:///home/ulrich/workspaces/GO_LEARNING/day06_interfaces_errors/exercise/main.go).

### Objectif de l'exercice :
1. Définir une interface `Notificateur` avec une méthode `Envoyer(message string) error`.
2. Créer deux structures : `Email` (contenant un champ `Adresse`) et `SMS` (contenant un champ `Numero`).
3. Implémenter l'interface `Notificateur` pour ces deux structures.
    *   Pour `Email` : si l'adresse ne contient pas le caractère `@`, renvoyer une erreur. Sinon, afficher "Email envoyé à [Adresse] : [message]" et renvoyer `nil`.
    *   Pour `SMS` : si le numéro fait moins de 10 caractères, renvoyer une erreur. Sinon, afficher "SMS envoyé au [Numero] : [message]" et renvoyer `nil`.
4. Dans le `main`, créer une liste (slice) de `Notificateur` contenant des emails et SMS (valides et invalides) et envoyer un message d'alerte à chacun d'entre eux en gérant proprement les erreurs d'envoi.

### Pour lancer l'exercice :
```bash
cd /home/ulrich/workspaces/GO_LEARNING/day06_interfaces_errors/exercise
go run main.go
```
