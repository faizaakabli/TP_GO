package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Question 1
func totalvotes() {
	fichier, err := os.Open("resultats-par-niveau-burvot-t1-france-entiere.txt")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	defer fichier.Close()
	scanner := bufio.NewScanner(fichier)

	if !scanner.Scan() {
		fmt.Println("Erreur : fichier vide")
		return
	}

	totalVotes := 0

	for scanner.Scan() {
		colonnes := strings.Split(scanner.Text(), ";")
		if len(colonnes) > 0 {
			votesStr := colonnes[10]
			votes, err := strconv.Atoi(votesStr)
			if err != nil {
				fmt.Println("Erreur :", err)
				continue
			}

			totalVotes += votes
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("votes total :", totalVotes)
	fmt.Printf("\n")
}

// Question 2
func TotalVotesCandidats() {
	indices := map[string]int{
		"ARTHAUD Nathalie":      25,
		"ROUSSEL Fabien":        32,
		"MACRON Emmanuel":       39,
		"LASSALLE Jean":         46,
		"LE PEN Marine":         53,
		"ZEMMOUR Éric":          60,
		"MÉLENCHON Jean-Luc":    67,
		"HIDALGO Anne":          74,
		"JADOT Yannick":         81,
		"PÉCRESSE Valérie":      88,
		"POUTOU Philippe":       95,
		"DUPONT-AIGNAN Nicolas": 102,
	}

	fichier, err := os.Open("resultats-par-niveau-burvot-t1-france-entiere.txt")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	defer fichier.Close()

	totalVotes := make(map[string]int)

	scanner := bufio.NewScanner(fichier)
	scanner.Scan()

	for scanner.Scan() {
		colonnes := strings.Split(scanner.Text(), ";")

		for personne, indiceVoix := range indices {
			if indiceVoix < len(colonnes) {
				voixStr := colonnes[indiceVoix]
				voix, err := strconv.Atoi(voixStr)
				if err != nil {
					fmt.Printf("Erreur %s: %v\n", personne, err)
					continue
				}

				totalVotes[personne] += voix
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	for personne, total := range totalVotes {
		fmt.Printf("%s votes : %d\n", personne, total)
	}
}

// Question 3
func totalVotesDepartements() {
	fichier, err := os.Open("resultats-par-niveau-burvot-t1-france-entiere.txt")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	defer fichier.Close()

	totalVotesDepartment := make(map[string]int)

	scanner := bufio.NewScanner(fichier)
	scanner.Scan()

	for scanner.Scan() {
		colonnes := strings.Split(scanner.Text(), ";")

		departement := colonnes[1]
		votesStr := colonnes[10]

		votes, err := strconv.Atoi(votesStr)
		if err != nil {
			fmt.Printf("Erreur %s: %v\n", departement, err)
			continue
		}

		totalVotesDepartment[departement] += votes
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var votesPairs []struct {
		departement string
		totalVotes  int
	}
	for departement, totalVotes := range totalVotesDepartment {
		votesPairs = append(votesPairs, struct {
			departement string
			totalVotes  int
		}{departement, totalVotes})
	}

	sort.Slice(votesPairs, func(i, j int) bool {
		return votesPairs[i].totalVotes > votesPairs[j].totalVotes
	})

	for _, pair := range votesPairs {
		fmt.Printf("%s : %d votants\n", pair.departement, pair.totalVotes)
	}
}

func main() {
	// Appeler les fonctions que vous avez définies
	fmt.Println("Total des votes:")
	totalvotes()

	fmt.Println("\nTotal des votes par candidat:")
	TotalVotesCandidats()

	fmt.Println("\nTotal des votes par département:")
	totalVotesDepartements()
}
