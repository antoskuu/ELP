# Projet JS

Bienvenue dans le monde captivant de Jarnac, un jeu de mots passionnant implémenté en JavaScript ! Ce jeu offre une expérience textuelle stimulante où les joueurs rivalisent pour créer des mots astucieux et stratégiques sur un plateau.

## Comment jouer

1) Cloner le github sur votre machine :

``` bash
git clone https://github.com/antoskuu/ELP.git
```

2) Déplacez vous dans le dossier :

``` bash
cd ELP/JS
```

3) Lancer le jeu :

``` bash
node jarnac.js
```

Si vous souhaitez voir l'historique des coups, ils sont stockés dans le fichier historique.txt à chaque nouvelle partie !
Le fichier historique.txt est strcturé de la manière suivante : 

- une ligne pour chaque action : le joueur a mis un mot, remplacer un mot, fait un jarnac ou passer son tour
- une ligne dès que le joueur pioche une lettre

Sur chaque ligne, le numéro du joueur qui a réalisé l'action ou pioché la lettre est affiché en spécifiant l'action réalisé et les éventuels changements du jeu.

