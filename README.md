# Sudoku Solver en Go

Ce projet implémente un résolveur de Sudoku en utilisant le langage Go (Golang). Il prend une grille de Sudoku valide en entrée, vérifie sa validité, puis utilise un algorithme de **backtracking** pour résoudre la grille. Si la grille est valide et résoluble, la solution est affichée ; sinon, un message d'erreur est renvoyé.

## Fonctionnalités

- Vérification de la validité de la grille de Sudoku.
- Résolution de la grille avec un algorithme de **backtracking**.
- Affichage de la solution si la grille est résolue.
- Gestion des erreurs pour des entrées invalides.

## Table des matières

1. [Aperçu](#aperçu)
2. [Prérequis](#prérequis)
3. [Installation](#installation)
4. [Utilisation](#utilisation)
5. [Exemples](#exemples)
6. [Structure du projet](#structure-du-projet)
7. [Contribuer](#contribuer)
8. [Licence](#licence)

## Aperçu

Le Sudoku est un puzzle logique où l'on doit remplir une grille de 9x9 cases avec les chiffres de 1 à 9. Chaque ligne, chaque colonne et chaque sous-grille 3x3 doivent contenir tous les chiffres de 1 à 9 sans répétition.

Le résolveur de ce projet prend une grille de Sudoku sous forme de neuf chaînes de neuf caractères. Le programme valide l'entrée, résout le Sudoku si la grille est valide et affiche la solution, ou affiche un message d'erreur si quelque chose ne va pas.

### Exemple de grille d'entrée

Une grille de Sudoku doit être entrée sous forme de neuf chaînes de neuf caractères, comme suit :

Le programme valide la grille, vérifie si elle respecte les règles du Sudoku, puis résout le puzzle.

## Prérequis

Pour exécuter ce programme, vous devez avoir installé Go (version 1.18 ou supérieure) sur votre machine.

- **Go (Golang)** : Vous pouvez télécharger Go depuis le site officiel [https://golang.org/dl/](https://golang.org/dl/).

## Installation

1. Clonez ce repository sur votre machine :
    ```bash
    git clone https://github.com/votre-utilisateur/sudoku-solver.git
    ```

2. Accédez au dossier du projet :
    ```bash
    cd sudoku-solver
    ```

3. Compilez et exécutez le programme :
    ```bash
    go run main.go "534678912" "672195348" "198342567" "859761423" "426853791" "713924856" "961537284" "287419635" "345286179"
    ```

   Ou compilez-le pour une exécution plus rapide :
    ```bash
    go build -o sudoku-solver main.go
    ./sudoku-solver "534678912" "672195348" "198342567" "859761423" "426853791" "713924856" "961537284" "287419635" "345286179"
    ```

## Utilisation

### Entrée du programme

Le programme attend en entrée une grille de Sudoku sous la forme de 9 chaînes de caractères. Chaque chaîne doit contenir exactement 9 chiffres représentant une ligne du Sudoku. Par exemple, pour la grille :

Vous pouvez l'exécuter en ligne de commande comme ceci :

```bash
go run main.go "534678912" "672195348" "198342567" "859761423" "426853791" "713924856" "961537284" "287419635" "345286179"


