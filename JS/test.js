//création de la pioche

const entree = "4A 4B 7C 5D 19E 2F 4G 2H 11I 1J 1K 6L 5M 9N 8O 4P 1Q 10R 7S 9T 8U 2V 1W 1X 1Y 2Z";

function creerPioche(entrée) {
    const pioche = [];
    const parties = entrée.split(' ');
    parties.forEach(partie => {
        const quantite = parseInt(partie.slice(0, -1), 10);
        const lettre = partie.slice(-1);
        pioche.push({ lettre, quantite });
    });

    return pioche;
}

const pioche = creerPioche(entree);



function tirerLettreAleatoire(pioche) {
    let lettresDisponibles = pioche.filter(item => item.quantite > 0);

    if (lettresDisponibles.length > 0) {
        let lettreAleatoireIndex = Math.floor(Math.random() * lettresDisponibles.length);
        let lettreAleatoire = lettresDisponibles[lettreAleatoireIndex].lettre;

        // Mettre à jour la quantité de la lettre tirée
        let lettreIndex = pioche.findIndex(item => item.lettre === lettreAleatoire);
        pioche[lettreIndex].quantite--;

        return lettreAleatoire;
    } else {
        console.log("Il n'y a plus de lettres disponibles dans la pioche.");
        return null;
    }
}

//pile ou face
function choixjoueur(arr) {
    var randomIndex = Math.floor(Math.random() * arr.length);
    return arr[randomIndex];
  }
  
var firstplayerplay = choixjoueur([1, 2]);
console.log(firstplayerplay);


//création de la matrice des
function grilleInit(valeurInitiale) {
    const grille = [];
    for (let i = 0; i < 8; i++) {
        const ligne = Array(9).fill(valeurInitiale);
        grille.push(ligne);
    }
    
    ajoutScore(grille, 0, [0, 0, 9, 16, 25, 36, 49, 64, 81]);
    return grille;
}

//affichage du plateau
function affichagePlateau(grille) {
    const lignes = grille.length;
    const colonnes = grille[0].length;

    // Afficher les indices des colonnes


    // Afficher la grille avec les indices de ligne et les tirets
    for (let i = 0; i < lignes; i++) {
        let ligne = `${i}\t|`;
        for (let j = 0; j < colonnes; j++) {
            ligne += ` ${grille[i][j] || ' '} \t|`;
        }
        console.log(ligne);
    }
}

    // Si toutes les lettres du mot sont dans la liste, retourner true



//ajout d'un score
function ajoutScore(grille, ligne, suite) {
    // Vérifier si la suite de nombres est un tableau
    if (Array.isArray(suite)) {
        // Vérifier si la ligne spécifiée est vide
        if (grille[ligne].every(cellule => cellule === "" || cellule === undefined)) {
            // Ajouter les numéros à la ligne de la grille
            for (let i = 0; i < suite.length; i++) {
                grille[ligne][i] = suite[i];
            }
            return true; // Numéros ajoutés avec succès
        } else {
            console.log("La ligne n'est pas vide. Choisissez une ligne vide.");
        }
    } else {
        console.log("La suite de nombres n'est pas un tableau.");
    }

    return false; // Les numéros n'ont pas été ajoutés
}
//ajout d'un mot à la grille


// Exemple d'utilisation pour créer une grille 3x4 avec la valeur initiale ""
const maGrille = grilleInit("");





function peutFormerMot(lettres, mot, grille) {
    if (mot.length > 2 && mot.length <= grille[0].length) {
        let copieLettres = [...lettres]; // Créer une copie de la liste de lettres pour ne pas la modifier
    for (let lettre of mot) {
        let index = copieLettres.indexOf(lettre);
        if (index === -1) {
            // Si la lettre n'est pas dans la liste, retourner false
            return false;
        } else {
            // Sinon, supprimer la lettre de la liste
            copieLettres.splice(index, 1);
        }
    }
    return true;
        } else {
        console.log("La longueur du mot ne satisfait pas les conditions.");
        return false; // Le mot n'a pas été ajouté
    }
    
}
    

function ajoutMotAGrille(grille, mot, ligne) {
    // Vérifier si la longueur du mot est supérieure à 3 et inférieure à la taille maximale de la grille
    if (mot.length > 2 && mot.length <= grille[0].length) {
        // Vérifier si la ligne spécifiée est vide
        if (grille[ligne].every(cellule => cellule === "" || cellule === undefined)) {
            // Ajouter le mot à la ligne de la grille
            for (let i = 0; i < mot.length; i++) {
                grille[ligne][i] = mot[i];
            }
            return true; // Mot ajouté avec succès
        } else {
            console.log("La ligne n'est pas vide. Choisissez une ligne vide.");
        }
    } 
}
function jeuEstTermine(numero_tour){
    return numero_tour > 7; // Le jeu se termine après le tour 6
} 

// function poserQuestion(numero_tour, numero_plateau) {
//     if (jeuEstTermine(numero_tour)) {
//         console.log("Le jeu est terminé.");
//         rl.close();
//         return;
//     }

//     rl.question('Entrez un mot : ', (reponse) => {
//         if (peutFormerMot(main1, reponse, numero_plateau)) {
//             console.log(`Vous pouvez former le mot ${reponse} avec les lettres de votre main.`);
//             reponse=reponse.toString();
//             ajoutMotAGrille(numero_plateau, reponse, numero_tour);
//             console.log("lettre au hasard: " + tirerLettreAleatoire(pioche));
//             affichagePlateau(numero_plateau);
//             poserQuestion(numero_tour + 1, numero_plateau); // Appeler poserQuestion à nouveau pour le prochain tour
//         } else {
//             console.log(`Vous ne pouvez pas former le mot ${reponse} avec les lettres de votre main.`);
//             poserQuestion(numero_tour, numero_plateau);
//         }
//     });
// }







function tirageMain(main) {
    for (let i = 0; i < 7; i++) {
        main.push(tirerLettreAleatoire(pioche));
    }
}

// tirageMain(main1);

// console.log(main1);
// const readline = require('readline');


// affichagePlateau(plateau1);

// poserQuestion(1, plateau1);



plateau1=grilleInit("")
plateau2=grilleInit("")
plateau3=grilleInit("")




const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

// let mains = [];
// let nombreDeJoueurs = 2; // Remplacez par le nombre de joueurs que vous voulez




// // Initialiser les mains pour chaque joueur
// for (let i = 0; i < nombreDeJoueurs; i++) {
//     mains[i] = [];
//     tirageMain(mains[i]);
// }









let mains = [];
let plateaux = [];
let nombreDeJoueurs = 2; // Remplacez par le nombre de joueurs que vous voulez

// Initialiser les mains et les plateaux pour chaque joueur
for (let i = 0; i < nombreDeJoueurs; i++) {
    mains[i] = [];
    tirageMain(mains[i]);
    plateaux[i] = grilleInit(""); // Supposons que vous ayez une fonction pour créer un nouveau plateau
}

function poserQuestion(numero_tour, joueur) {
    if (jeuEstTermine(Math.floor(numero_tour))) {
        console.log("Le jeu est terminé.");
        rl.close();
        return;
    }

    console.log(`C'est au tour du joueur ${joueur + 1}.`);
    console.log(`Voici votre main : ${mains[joueur]}`);
    rl.question('Entrez un mot : ', (reponse) => {
        if (peutFormerMot(mains[joueur], reponse, plateaux[joueur])) {
            console.log(`Vous pouvez former le mot ${reponse} avec les lettres de votre main.`);
            reponse=reponse.toString();
            ajoutMotAGrille(plateaux[joueur], reponse, Math.floor(numero_tour));
            console.log("lettre au hasard: " + tirerLettreAleatoire(pioche));
            affichagePlateau(plateaux[joueur]);
            let prochainJoueur = (joueur + 1) % nombreDeJoueurs; // Alternez le joueur
            poserQuestion(numero_tour + 0.5, prochainJoueur); // Appeler poserQuestion à nouveau pour le prochain tour
        } else {
            console.log(`Vous ne pouvez pas former le mot ${reponse} avec les lettres de votre main.`);
            poserQuestion(numero_tour, joueur); // Si le mot n'est pas valide, le même joueur essaie à nouveau
        }
    });
}

// Commencer le jeu avec le tour 1 et le joueur 0
poserQuestion(1, 0);