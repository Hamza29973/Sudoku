Sudoku en Go
Ce projet implémente un solveur de Sudoku en Go. Il permet de résoudre une grille de Sudoku donnée en utilisant un algorithme de recherche récursive et de backtracking. Le but est de remplir les cases vides d'une grille de Sudoku avec des chiffres de 1 à 9, en respectant les règles du jeu.

Prérequis
Go 1.18 ou supérieur.
Une compréhension basique des règles du Sudoku.
Structure du projet
bash
Copier le code
sudoku/
├── sudoku.go        # Contient les structures et les fonctions pour résoudre le Sudoku
├── main.go          # Exemple d'utilisation du solveur Sudoku
├── README.md        # Documentation du projet
Règles du Sudoku
Une grille de Sudoku est une matrice 9x9, divisée en 9 sous-grilles de 3x3.
Chaque ligne, chaque colonne et chaque sous-grille 3x3 doivent contenir les chiffres de 1 à 9 sans répétition.
Les cases initiales sont remplies avec des chiffres entre 1 et 9, et les autres cases doivent être remplies par le solveur.
Fonctionnalités
Résolution du Sudoku : Le solveur tente de remplir la grille en respectant les règles du Sudoku.
Validation de la grille : Avant de résoudre, le programme peut vérifier si une grille est valide.
Affichage de la grille : La grille peut être affichée dans un format lisible.
