# Sudoku Solver

Un résolveur de Sudoku simple basé sur un algorithme de **backtracking**. Ce projet permet de résoudre des grilles de Sudoku en utilisant la méthode de recherche exhaustive. Le programme peut résoudre des grilles de Sudoku valides et renvoie la solution si elle existe.

## Fonctionnalités

- Résolution automatique de grilles de Sudoku de 9x9.
- Validation des règles du Sudoku : chaque ligne, chaque colonne et chaque sous-grille 3x3 doivent contenir les chiffres de 1 à 9 sans répétition.
- Interface simple en ligne de commande pour une utilisation facile.
- Prise en charge de grilles de Sudoku partiellement remplies.

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

Le Sudoku est un puzzle où l'objectif est de remplir une grille de 9x9 avec des chiffres de 1 à 9. Chaque ligne, chaque colonne, et chaque sous-grille 3x3 doivent contenir tous les chiffres de 1 à 9 sans répétition. Le résolveur de ce projet utilise un algorithme de **backtracking** qui explore toutes les combinaisons possibles pour trouver la solution valide.

## Prérequis

Avant d'exécuter le programme, vous devez vous assurer que vous avez les éléments suivants installés sur votre machine :

- **Python 3.x** ou une version supérieure.
- Bibliothèque `numpy` pour la gestion des matrices (facultatif, selon votre implémentation).

Vous pouvez installer les dépendances nécessaires avec `pip` :

```bash
pip install numpy
