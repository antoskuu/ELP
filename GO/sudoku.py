def absentsurligne(k,grille,ligne):
    for colonne in range (0,9):
        if grille[ligne][colonne] == k:
            return False
    return True


def absentsurcolonne(k,grille,colonne):
    for ligne in range (0,9):
        if grille[ligne][colonne] == k:
            return False
    return True

def absentsurbloc(k, grille, ligne, colonne):
    ligne2 = ligne - ligne % 3
    colonne2 = colonne - colonne % 3
    for x in range(3):
        for y in range(3):
            if grille[ligne2 + x][colonne2 + y] == k:
                return False
    return True


def solve(grille,ligne=0,colonne=0):
    if ligne==9:
        return True
    elif colonne == 9:
        return solve(grille,ligne+1,0)
    elif grille[ligne][colonne]!=0:
        return solve(grille,ligne,colonne+1)
    else:
        for k in range (1,10):
            if absentsurbloc(k,grille,ligne,colonne) and absentsurcolonne(k,grille,colonne) and absentsurligne(k,grille,ligne):
                grille[ligne][colonne] = k
                if solve(grille,ligne,colonne+1):
                    return True
                grille[ligne][colonne] = 0
        return False
    
    
def main ():
    grille = [[0,0,7,0,3,0,5,0,0],[3,6,0,0,0,0,0,4,1],[0,0,0,4,1,6,0,0,0],[0,0,0,2,8,7,0,0,0],[0,9,0,0,0,0,0,1,0],[0,8,0,0,0,0,0,5,0],[0,0,4,0,0,0,3,0,0],[0,0,2,0,0,0,8,0,0],[0,0,0,3,9,8,0,0,0]]
    
    print("Avant :")
    for row in grille:
        print(row)


    solve(grille,ligne=0,colonne=0)
    
    if solve(grille,ligne=0,colonne=0):
        print("\nAprès :")
        for row in grille:
            print(row)
    else:
        print("\nErreur pas de solution !")
    
main()