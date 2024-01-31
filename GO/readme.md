# Projet Go

Découvrez notre solveur de sudoku en backtracking en parallèle à l'aide de goroutine

## Comment le lancer 

1) Cloner le github sur votre machine :

``` bash
git clone https://github.com/antoskuu/ELP.git
```

2) Deplacer vous dans le dossier :

``` bash
cd ELP/GO/code
```

3) Voici les 3 programmes :

- sudoko_backtracking.go : solveur de sudoku en séquentiel
- sudoku_backtracking_parallele.go : solveur de sudoku en parallèle avec des goroutines
- sudoku_worker.go : solveur de sudoku en parallèle avec des workers

Des grilles sont aussi présentes dans le dossier : grille 
Si vous souhaitez prendre d'autres grilles, un convertisseur des grilles se trouvant sur https://www.top-sudoku.com/sudoku/fr/choisir-une-grille-sudoku.php est à votre disposition.
Une fois sur le site, choisissez une grille, puis enregistrez la grille. Copiez ensuite la grille dans le convertisseur. 
