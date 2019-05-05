package main

import (
	"bytes"
	"strings"
)

//Gere les erreurs d'espace
func handleStringErrors(s string) string {
	var buffer bytes.Buffer

	for x := 0; x != len(s)-1; x++ {
		if s[x] == '\n' && s[x+1] != ' ' {
			buffer.WriteByte(s[x])
			buffer.WriteByte(' ')
		} else {
			buffer.WriteByte(s[x])
		}
	}
	return buffer.String()
}

//Justifie et saute la ligne d'une chaine
func justifytxt(s string, n int) string {
	var comptlen, comptmotbyline, comptwordbyline int
	//Vérifie si il y a bien un espace entre chaque mot puis split la chaine
	splitstr := strings.Split(handleStringErrors(s), " ")

	//Justifie et saute la ligne de l'array
	for x := 0; x != len(splitstr); x++ {
		if comptlen+len([]rune(splitstr[x])) >= n {
			space := n - comptlen
			for comptwordbyline = comptmotbyline - 1; space >= 0; comptwordbyline-- {
				if comptwordbyline == 0 {
					comptwordbyline = comptmotbyline - 1
				}
				splitstr[(x-1)-comptwordbyline] = splitstr[(x-1)-comptwordbyline] + " "
				space--
			}
			splitstr[x-1] = splitstr[x-1] + "\n"
			comptlen, comptmotbyline = 0, 0
		}

		//compte la longueur de la chaine + 1 pour l'espace
		comptlen = comptlen + len([]rune(splitstr[x])) + 1
		comptmotbyline++

		//Si il trouve un retour à la ligne, remet les compteurs à 0
		if strings.Contains(splitstr[x], "\n") == true {
			comptlen, comptmotbyline = 0, 0
		}
	}
	//Remet tout l'array dans la chaine en laissant les espaces
	joinedstr := strings.Join(splitstr, " ")
	newjoinedstr := strings.ReplaceAll(joinedstr, "\n ", "\n")

	return newjoinedstr
}
