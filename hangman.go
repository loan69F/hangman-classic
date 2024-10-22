package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func lireMots(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var mots []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	return mots, scanner.Err()
}

func choisirMot(mots []string) string {
	rand.Seed(time.Now().UnixNano())
	return mots[rand.Intn(len(mots))]
}

func cacherMot(mot string) ([]rune, error) {
	motCache := make([]rune, len(mot))
	for i := range motCache {
		motCache[i] = '_'
	}
	indices := rand.Perm(len(mot))[:2] // Choisir deux indices uniques
	for _, index := range indices {
		motCache[index] = rune(mot[index])
	}
	return motCache, nil
}

func main() {
	mots, err := lireMots("words3.txt")
	if err != nil {
		log.Fatal(err)
	}

	var jouer string
	for {
		fmt.Print("Tapez 'start' pour commencer ou 'stop' pour arrêter: ")
		fmt.Scanln(&jouer)
		if jouer == "stop" {
			fmt.Println("Jeu arrêté.")
			return
		} else if jouer != "start" {
			fmt.Println("Commande invalide. Veuillez réessayer.")
			continue
		}

		motChoisi := choisirMot(mots)
		motCache, _ := cacherMot(motChoisi)
		tentativesRestantes := 10

		for tentativesRestantes > 0 {
			fmt.Printf("Mot: %s | Tentatives restantes: %d\n", string(motCache), tentativesRestantes)
			var guess string
			fmt.Print("Devinez une lettre ou le mot: ")
			fmt.Scanln(&guess)

			if len(guess) == 1 {
				correctGuess := false
				for i, char := range motChoisi {
					if string(char) == guess {
						motCache[i] = char
						correctGuess = true
					}
				}
				if !correctGuess {
					tentativesRestantes--
					fmt.Println("Mauvaise tentative.")
				}
			} else if guess == motChoisi {
				fmt.Printf("Bravo! Vous avez trouvé le mot: %s\n", motChoisi)
				break
			} else {
				fmt.Println("Mauvaise tentative.")
				tentativesRestantes--
			}

			if string(motCache) == motChoisi {
				fmt.Printf("Bravo! Vous avez trouvé le mot: %s\n", motChoisi)
				break
			}
		}

		if tentativesRestantes == 0 {
			fmt.Printf("Dommage! Le mot était: %s\n", motChoisi)
		}
	}
}
