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

const entree = "4A 4B 7C 5D 19E 2F 4G 2H 11I 1J 1K 6L 5M 9N 8O 4P 1Q 10R 7S 9T 8U 2V 1W 1X 1Y 2Z";
const pioche = creerPioche(entree);
console.log(pioche);


// Exemple de comment tirer une lettre de la pioche avec mise à jour de la quantité
function tirerLettre(pioche, lettre) {
    const lettreIndex = pioche.findIndex(item => item.lettre === lettre);

    if (lettreIndex !== -1 && pioche[lettreIndex].quantite > 0) {
        pioche[lettreIndex].quantite--;
        return lettre;
    } else {
        console.log(`La lettre ${lettre} n'est plus disponible dans la pioche.`);
        return null;
    }
}

function choixjoueur(arr) {
    var randomIndex = Math.floor(Math.random() * arr.length);
    return arr[randomIndex];
  }
  
var firstplayerplay = choixjoueur([1, 2]);
console.log(firstplayerplay);

function grilleInit() {

}


function grilleInit(valeurInitiale) {
    const grille = [];

    // Boucler pour créer chaque ligne de la grille
    for (let i = 0; i < 8; i++) {
        // Initialiser chaque ligne avec la valeur initiale spécifiée
        const ligne = Array(9).fill(valeurInitiale);
        
        // Ajouter la ligne à la grille
        grille.push(ligne);
    }

    return grille;
}

// function affichagePlateau(grille) {
//     for (let i = 0; i < grille.length; i++) {
//         let ligne = "";
//         for (let j = 0; j < grille[i].length; j++) {
//             if (grille[i][j]!=0) {
//                 ligne += grille[i][j]; // Ajouter une tabulation entre chaque valeur
//             }
//         }
//         console.log(ligne);
//     }
// }


function affichagePlateau(grille) {
    const lignes = grille.length;
    const colonnes = grille[0].length;

    // Afficher les indices des colonnes
    let entete = "\t";
    for (let j = 0; j < colonnes; j++) {
        entete += `${j}\t`;
    }
    console.log(entete);

    // Afficher la grille avec les indices de ligne et les tirets
    for (let i = 0; i < lignes; i++) {
        let ligne = `${i}\t|`;
        for (let j = 0; j < colonnes; j++) {
            ligne += ` ${grille[i][j] || '-'} \t|`;
        }
        console.log(ligne);
        console.log("\t" + "-".repeat((colonnes + 1) * 6 - 1)); // Ligne de tirets
    }
}




function ajoutMotAGrille(grille, mot, ligne) {
    // Vérifier si la longueur du mot est supérieure à 3 et inférieure à la taille maximale de la grille
    if (mot.length > 3 && mot.length <= grille[0].length) {
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
    } else {
        console.log("La longueur du mot ne satisfait pas les conditions.");
    }

    return false; // Le mot n'a pas été ajouté
}

// Exemple d'utilisation pour créer une grille 3x4 avec la valeur initiale 'X'
const maGrille = grilleInit("");

// Afficher la grille résultante
affichagePlateau(maGrille)



plateau1=grilleInit()
plateau2=grilleInit()


const motAInserer = "TIGRE";
const ligneAInserer = 1;

if (ajoutMotAGrille(maGrille, motAInserer, ligneAInserer)) {
    console.log("Mot ajouté avec succès !");
} else {
    console.log("Le mot n'a pas été ajouté.");
}

affichagePlateau(maGrille);


let lettreTiree = tirerLettre(pioche, 'A');
console.log(`Lettre tirée : ${lettreTiree}`);
