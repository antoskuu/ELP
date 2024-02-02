//création de la pioche

const fs = require('fs');

fs.writeFile('historique.txt', '', err => {
    if (err) {
        console.error(err);
        return;
    }
    // Le fichier est maintenant vide
});

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

function ecrireDansFichier(texte) {
    fs.appendFile('historique.txt', texte + '\n', (err) => {
        if (err) throw err;
    });
}

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





// function peutFormerMot(lettres, mot, grille) {
//     if (mot.length > 2 && mot.length <= 9) {
//         console.log("mot.length "+mot.length)
//         let copieLettres = [...lettres]; // Créer une copie de la liste de lettres pour ne pas la modifier
//     for (let lettre of mot) {
//         let index = copieLettres.indexOf(lettre);
//         if (index === -1) {
//             // Si la lettre n'est pas dans la liste, retourner false
//             console.log("premier")
//             return false;
//         } else {
//             // Sinon, supprimer la lettre de la liste
//             copieLettres.splice(index, 1);
//         }
//     }
//     return true;
//         } else {
        
//         console.log("La longueur du mot ne satisfait pas les conditions.");
//         return false; // Le mot n'a pas été ajouté
        
//     }
    
// }
    
function peutFormerMot(lettres, mot) {
    if (mot.length > 2 || mot.length <= 9) {
        for (let lettre of mot) {
            if (!lettres.includes(lettre)) {
                // Si une lettre du mot n'est pas dans la liste de lettres, retourner false
                return false;
            }
        }
        // Si toutes les lettres du mot sont dans la liste de lettres, retourner true
        return true;
    } else {
        // Si le mot a 2 caractères ou moins, retourner false
        return false;
    }
}


function ajoutMotAGrille(grille, mot, ligne ,main) {
    // Vérifier si la longueur du mot est supérieure à 3 et inférieure à la taille maximale de la grille
    if (mot.length > 2 && mot.length <= grille[0].length) {
        // Vérifier si la ligne spécifiée est vide
        
            // Ajouter le mot à la ligne de la grille
            for (let i = 0; i < mot.length; i++) {
                grille[ligne][i] = mot[i];
                retirerLettreDeMain(main, mot[i]);
            }
            return true; // Mot ajouté avec succès
        
    } 
}

// function ajoutMotAGrilleJarnac(grille, mot, ligne, mot_de_base, main_joueur) {
//     // Vérifiez si mot_de_base est une sous-chaîne de mot
//     if (mot.includes(mot_de_base)) {
//         // Obtenez les lettres supplémentaires dans mot
//         let lettresSupplementaires = mot.replace(new RegExp(`[${mot_de_base}]`, 'g'), '');
        
//         // Vérifiez si toutes les lettres supplémentaires sont dans main_joueur
//         for (let lettre of lettresSupplementaires) {
//             if (!main_joueur.includes(lettre)) {
//                 return false; // Si une lettre n'est pas dans main_joueur, ne faites rien et quittez la fonction
//             }
//         }

//         // Si toutes les lettres supplémentaires sont dans main_joueur, ajoutez mot à grille
//         grille[ligne] = mot;
//         return true
//     }
// }
// function ajoutMotAGrilleJarnac(grille, mot, ligne, mot_de_base, main_joueur) {
//     // Convertir mot_de_base en une chaîne pour la recherche de sous-chaîne
//     let mot_de_base_str = mot_de_base.join('');

//     // Vérifiez si mot_de_base est une sous-chaîne de mot
//     if (mot.includes(mot_de_base_str)) {
//         // Obtenez les lettres supplémentaires dans mot
//         let lettresSupplementaires = mot.split('').filter(lettre => !mot_de_base.includes(lettre));
        
//         // Vérifiez si toutes les lettres supplémentaires sont dans main_joueur
//         for (let lettre of lettresSupplementaires) {
//             if (!main_joueur.includes(lettre)) {
//                 return false; // Si une lettre n'est pas dans main_joueur, ne faites rien et quittez la fonction
//             }
//         }

//         // Si toutes les lettres supplémentaires sont dans main_joueur, ajoutez mot à grille
//         grille[ligne] = mot;
//         return true;
//     }

//     // Si mot ne contient pas mot_de_base, retourner false
//     return false;
// }
function ajoutMotAGrilleJarnac(grille, mot, ligne, mot_de_base, main_joueur) {
    // Convertir mot_de_base en une chaîne pour la recherche de sous-chaîne
    let mot_de_base_str = mot_de_base.join('');

    // Vérifiez si mot_de_base est une sous-chaîne de mot
    if (mot.includes(mot_de_base_str)) {
        // Obtenez les lettres supplémentaires dans mot
        let lettresSupplementaires = mot.split('').filter(lettre => !mot_de_base.includes(lettre));
        
        // Vérifiez si toutes les lettres supplémentaires sont dans main_joueur
        for (let lettre of lettresSupplementaires) {
            if (!main_joueur.includes(lettre)) {
                return false; // Si une lettre n'est pas dans main_joueur, ne faites rien et quittez la fonction
            }
        }

        // Si toutes les lettres supplémentaires sont dans main_joueur, ajoutez mot à grille
        grille[ligne] = mot;
        return true;
    }

    // Si mot ne contient pas mot_de_base, retourner false
    return false;
}

function retirerLettreDeMain(main, lettre) {
    const index = main.indexOf(lettre);
    if (index > -1) {
        main.splice(index, 1);
    }
}

// function jeuEstTermine(grille) {
//     // Parcourir chaque ligne de la grille
//     for (let ligne of grille) {
//         // Si la ligne contient une cellule vide, retourner false
//         if (ligne.includes("") || ligne.includes(undefined)) {
//             return false;
//         }
//     }

//     // Si aucune ligne ne contient de cellule vide, retourner true
//     return true;
// }

function jeuEstTermine(listeDeListes) {
    // Obtenez la dernière liste de la liste de listes
    let derniereListe = listeDeListes[listeDeListes.length - 1];

    // Vérifiez si le premier élément de la dernière liste est différent de ' '
    if (derniereListe[0] !== '') {
        return true;''
    }

    return false;
}






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


function peutFormerMotAvecTransformation(main, nouveauMot, motExistant) {
    // Créer une copie de la main du joueur
    let lettresDisponibles = main.slice();

    // Créer un ensemble des lettres du nouveau mot
    const ensembleLettresNouveauMot = new Set(nouveauMot);

    let new_word = nouveauMot.slice()

    let copi_motExistant = [...motExistant]

    let ancien_mot = copi_motExistant.join('')
    console.log("Voici le mot existant")
    console.log(new_word)


    // Vérifier si toutes les lettres du mot existant sont présentes dans le nouveau mot
    for (const lettre of motExistant) {
        if (!ensembleLettresNouveauMot.has(lettre) && lettre !== '') {
            return false;
        }
    }

    // Ajouter les lettres non vides du mot existant à la copie de la main
    lettresDisponibles = lettresDisponibles.concat(motExistant.filter(lettre => lettre !== ''));

    // Vérifier si le nouveau mot peut être formé en utilisant les lettres disponibles
    for (let i = 0; i < nouveauMot.length; i++) {
        const lettre = nouveauMot[i];
        const index = lettresDisponibles.indexOf(lettre);

        // Si la lettre n'est pas dans les lettres disponibles, le mot ne peut pas être formé
        if (index === -1) {
            return false;
        }

        // Retirer la lettre utilisée de la copie de la main
    }

    // Retirer les lettres de ancien_mot de new_word
    for (let j = 0; j < ancien_mot.length; j++) {
        new_word = new_word.replace(ancien_mot[j], '');
    }

    // Retirer les lettres restantes de new_word de la main
    for (let j = 0; j < new_word.length; j++) {
        retirerLettreDeMain(main, new_word[j]);
    }

    return true;
}



function peutFormerMotAvecTransformation(main, nouveauMot, motExistant) {
    // Créer une copie de la main du joueur
    let lettresDisponibles = main.slice();

    // Créer un ensemble des lettres du nouveau mot
    const ensembleLettresNouveauMot = new Set(nouveauMot);

    // Vérifier si toutes les lettres du mot existant sont présentes dans le nouveau mot
    for (const lettre of motExistant) {
        if (!ensembleLettresNouveauMot.has(lettre) && lettre !== '') {
            return false;
        }
    }

    // Ajouter les lettres non vides du mot existant à la copie de la main
    lettresDisponibles = lettresDisponibles.concat(motExistant.filter(lettre => lettre !== ''));

    // Vérifier si le nouveau mot peut être formé en utilisant les lettres disponibles
    for (let i = 0; i < nouveauMot.length; i++) {
        const lettre = nouveauMot[i];
        const index = lettresDisponibles.indexOf(lettre);

        // Si la lettre n'est pas dans les lettres disponibles, le mot ne peut pas être formé
        if (index === -1) {
            return false;
        }

        // Retirer la lettre utilisée de la copie de la main
        lettresDisponibles.splice(index, 1);
    }

    return true;
}


function jarnacSupprimerLigne(grille, reponse) {
    grille.splice(reponse, 1);
    grille.push(Array(9).fill(""));
}

function jarnacSupprimerLettres(main_joueur, mot) {
    // Convertir le mot en tableau de lettres
    let lettres = mot.split('');

    // Supprimer les lettres du mot de la main du joueur
    for (let lettre of lettres) {
        let index = main_joueur.indexOf(lettre);
        if (index !== -1) {
            main_joueur.splice(index, 1);
        }
    }
}


let mains = [];
let plateaux = [];
let nombreDeJoueurs = 2; // Remplacez par le nombre de joueurs que vous voulez

// Initialiser les mains et les plateaux pour chaque joueur
for (let i = 0; i < nombreDeJoueurs; i++) {
    mains[i] = [];
    tirageMain(mains[i]);
    plateaux[i] = grilleInit(""); // Supposons que vous ayez une fonction pour créer un nouveau plateau
}

ligne=[1, 1]


function calculPoints(plateau1) {
    let points1 = 0;
    
    for (let i = 1; i < plateau1.length; i++) {
        compteur=0
        for (let j = 1; j < plateau1[i].length; j++) {
            if (plateau1[i][j] !== ``) {
                compteur++;
            }
        
        }
        points1 = points1 + plateau1[0][compteur];
        
    }
    return points1;   
}


function poserQuestion(numero_ligne, joueur) {

    // console.log(plateaux[joueur])
    if (jeuEstTermine(plateaux[joueur])) {
        console.log("Le jeu est terminé. Le joueur " + (joueur + 1) + " a rempli toute sa grille!");
        console.log("Le joueur 1 a marqué " + calculPoints(plateaux[0]) + " points.");
        console.log("Le joueur 2 a marqué " + calculPoints(plateaux[1]) + " points.");
        rl.close();
        return;
    }

    console.log(`C'est au tour du joueur ${joueur + 1}.`);

    console.log(`Voici votre main : ${mains[joueur]}`);

    rl.question('Tapez:\n 1 pour mettre un mot, \n 2 pour en modifier un, \n 3 pour ne rien faire, \n 4 pour JARNAC! \n ', (reponse1) => {


        if (reponse1 == 1) {
        console.log('Vous avez choisi de mettre un mot');
        console.log('Pour revenir en arrière, tapez R');
        rl.question('Entrez un mot : ', (reponse) => {
            if (peutFormerMot(mains[joueur], reponse, plateaux[joueur])) {
                console.log(`Vous pouvez former le mot ${reponse} avec les lettres de votre main.`);
                ecrireDansFichier(`Joueur ${joueur + 1} a formé le mot ${reponse} avec les lettres de sa main.`)
                reponse=reponse.toString();
                ajoutMotAGrille(plateaux[joueur], reponse, numero_ligne[joueur], mains[joueur]);
                letter=tirerLettreAleatoire(pioche);
                console.log('Vous avez pioché la lettre ' + letter);
                mains[joueur].push(letter);
                ecrireDansFichier(`Joueur ${joueur + 1} a pioché la lettre ${letter}.`)
                affichagePlateau(plateaux[joueur]);
                numero_ligne[joueur] = numero_ligne[joueur] + 1;
                poserQuestion(numero_ligne, joueur); // Appeler poserQuestion à nouveau pour le prochain tour
            } else if (reponse == "R") {
                poserQuestion(numero_ligne, joueur)
            } else {
                console.log(`Vous ne pouvez pas former le mot ${reponse} avec les lettres de votre main.`);
                poserQuestion(numero_ligne, joueur); // Si le mot n'est pas valide, le même joueur essaie à nouveau
            }
        });
    } else if (reponse1 == 2) {
        console.log('Vous avez choisi de modifier un mot');
        console.log('Pour revenir en arrière, tapez R');
// ...

        rl.question('Entrez le numéro de la ligne que vous souhaitez modifier : ', (reponse) => {
            // Vérifiez si la ligne existe
            reponse = parseInt(reponse);
            if (isNaN(reponse) || reponse < 0 || reponse >= numero_ligne[joueur]) {
                console.log("Numéro de ligne invalide. Veuillez réessayer.");
                poserQuestion(numero_ligne, joueur);
                return;
            }

            rl.question('Entrez le nouveau mot : ', (nouveauMot) => {
                if (peutFormerMotAvecTransformation(mains[joueur], nouveauMot, plateaux[joueur][reponse])) {
                    console.log(`Vous pouvez former le mot ${nouveauMot} avec les lettres de votre main et du mot existant.`);
                    let ancienMot = plateaux[joueur][reponse];
                    let ancienMotString = ancienMot.join('');
                    ecrireDansFichier(`Joueur ${joueur + 1} a remplacé le mot ${ancienMotString} de la ligne ${reponse} en ${nouveauMot}.`)
                    // Modifier le mot sur la ligne spécifiée
                    plateaux[joueur][reponse] = nouveauMot;
                    let lettresAjoutees = nouveauMot.split('');
                    for (let lettre of ancienMot) {
                        let index = lettresAjoutees.indexOf(lettre);
                        if (index !== -1) {
                            lettresAjoutees.splice(index, 1);
                        }
                    }
                
                    // Supprimer les lettres ajoutées de la main du joueur
                    for (let lettre of lettresAjoutees) {
                        retirerLettreDeMain(mains[joueur], lettre);
                    }
                    letter=tirerLettreAleatoire(pioche);
                    console.log('Vous avez pioché la lettre ' + letter);
                    mains[joueur].push(letter);
                    ecrireDansFichier(`Joueur ${joueur + 1} a pioché la lettre ${letter}.`)
                    affichagePlateau(plateaux[joueur]);
                    poserQuestion(numero_ligne, joueur); // Appeler poserQuestion à nouveau pour le prochain tour
                } else {
                    console.log(`Vous ne pouvez pas former le mot ${nouveauMot} avec les lettres de votre main et du mot existant.`);
                    poserQuestion(numero_ligne, joueur); // Si le mot n'est pas valide, le même joueur essaie à nouveau
                }
            });
        });
    
    } else if (reponse1 == 3) {
        ecrireDansFichier(`Joueur ${joueur + 1} a passé son tour.`);
        console.log('Vous avez choisi de passer votre tour');
        letter=tirerLettreAleatoire(pioche);
        console.log('Vous avez pioché la lettre ' + letter);
        mains[joueur].push(letter);
        ecrireDansFichier(`Joueur ${joueur + 1} a pioché la lettre ${letter}.`)
        let prochainJoueur = (joueur + 1) % nombreDeJoueurs;
        affichagePlateau(plateaux[(joueur+1)%2]);
        poserQuestion(numero_ligne, prochainJoueur);


    } else if (reponse1 == 4)   {
        console.log('Vous avez choisi de faire un JARNAC');
        console.log('Pour revenir en arrière, tapez R');
        rl.question('Quelle ligne voulez vous voler?', (reponse) => {
            if (reponse=="R") {
                poserQuestion(numero_ligne, joueur)
            }
            else if (parseInt(reponse)<= 0 || parseInt(reponse)>=numero_ligne[(joueur + 1) % nombreDeJoueurs]) {
                console.log('Vous devez entrer une réponse valide');
                poserQuestion(numero_ligne, joueur)
            }

            else if (parseInt(reponse)> 0 && parseInt(reponse)<numero_ligne[(joueur + 1) % nombreDeJoueurs]){
            main_temporaire=[]
            mot_de_base=[]
            main_totale=[]
            longueur=0
            for (let element of plateaux[(joueur + 1) % nombreDeJoueurs][reponse]) {
                if (element !== "") {
                    longueur++;
                    mot_de_base.push(element);
                }
            }
            for (let element of mains[(joueur + 1) % nombreDeJoueurs]) {
                main_temporaire.push(element);
            }
            main_temporaire.pop()
            for (let element of mot_de_base) {
                main_totale.push(element);
            }
            for (let element of main_temporaire) {
                main_totale.push(element);
            }
            
            console.log(`Vous devez former un mot de plus de ${longueur} lettres avec:`);
            console.log(mot_de_base + " obligatoirement");
            console.log("et au moins une lettre de " +main_temporaire);
            rl.question('Quel mot formez vous? ', (reponse2) => {
                console.log("reponse2.length "+reponse2.length)
                console.log("longueur " +longueur)
                console.log(main_totale)
                console.log("peutformermot "+peutFormerMot(main_totale, reponse2))


              if (reponse2.length>longueur && peutFormerMot(main_totale, reponse2))
               { 
                console.log("Le nouveau mot est plus long que l'ancien");
                console.log(ajoutMotAGrilleJarnac(plateaux[joueur], reponse2, numero_ligne[joueur], mot_de_base, main_temporaire))
                if (ajoutMotAGrilleJarnac(plateaux[joueur], reponse2, numero_ligne[joueur], mot_de_base, main_temporaire)){

                jarnacSupprimerLigne(plateaux[(joueur + 1) % nombreDeJoueurs], reponse);

                jarnacSupprimerLettres(mains[(joueur + 1) % nombreDeJoueurs], reponse2, reponse2);
                
                console.log(`Vous avez volé la ligne ${reponse} et formé le mot ${reponse2}`);
                ecrireDansFichier(`Joueur ${joueur + 1} a volé la ligne ${reponse} et a formé le mot ${reponse2}.`)
                console.log(`Voici votre nouveau plateau:`);
                affichagePlateau(plateaux[joueur]);
                console.log(`Voici le plateau de l'adversaire:`);
                affichagePlateau(plateaux[(joueur + 1) % nombreDeJoueurs]);
                numero_ligne[joueur] = numero_ligne[joueur] + 1;
                numero_ligne[(joueur + 1) % nombreDeJoueurs] = numero_ligne[(joueur + 1) % nombreDeJoueurs] - 1;
                poserQuestion(numero_ligne, joueur)
               }
              }
              else { 
                console.log("Le nouveau mot n'est pas plus long que l'ancien");
                poserQuestion(numero_ligne, joueur)}
            });
            }

            else {
                console.log('Vous devez entrer une réponse valide');
                poserQuestion(numero_ligne, joueur)
                
            } 



        
        // else if (reponse == "R") {
        //     poserQuestion(numero_ligne, joueur)
        // }
    });

    }
    
    
    
    else {
        console.log('Vous devez entrer une réponse valide');
        poserQuestion(numero_ligne, joueur)
    } 
}
    )
}

// Commencer le jeu avec le tour 1 et le joueur 0
poserQuestion(ligne, 0);



