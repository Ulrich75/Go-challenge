#!/bin/bash
# Script de vérification automatique de la compilation des exercices et solutions Go

# Arrêter le script à la première erreur
set -e

# Chemin absolu de Go configuré sur le système
GO_BIN="/usr/local/go/bin/go"

echo "=== DÉBUT DE LA VÉRIFICATION DES FICHIERS GO ==="

# Liste des fichiers Go à vérifier
fichiers=(
  "day01_introduction/hello.go"
  "day01_introduction/exercise/main.go"
  "day01_introduction/exercise/solution/main.go"
  "day02_control_flow/exercise/main.go"
  "day02_control_flow/exercise/solution/main.go"
  "day03_functions/exercise/main.go"
  "day03_functions/exercise/solution/main.go"
  "day04_data_structures/exercise/main.go"
  "day04_data_structures/exercise/solution/main.go"
  "day05_pointers_methods/exercise/main.go"
  "day05_pointers_methods/exercise/solution/main.go"
  "day06_interfaces_errors/exercise/main.go"
  "day06_interfaces_errors/exercise/solution/main.go"
  "day07_final_project/main.go"
  "day07_final_project/solution/main.go"
)

for f in "${fichiers[@]}"; do
  echo -n "Validation de la compilation : $f... "
  # Tenter de compiler sans générer de binaire (sortie vers /dev/null)
  if $GO_BIN build -o /dev/null "$f"; then
    echo "✅ Réussite"
  else
    echo "❌ Échec de la compilation"
    exit 1
  fi
done

echo "================================================"
echo "🎉 Tous les fichiers Go compilent avec succès !"
echo "================================================"
