package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Player represents a football player.
type Player struct {
	Name string
}

// Pair represents a pairing of a striker and a defender.
type Pair struct {
	Striker  *Player
	Defender *Player
}

// ReadPlayersFromFile reads a list of players from a file and returns them as a slice.
func ReadPlayersFromFile(filename string) ([]*Player, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var players []*Player
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		playerName := scanner.Text()
		player := &Player{Name: playerName}
		players = append(players, player)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return players, nil
}

// PairPlayers pairs strikers with defenders.
func PairPlayers(strikers, defenders []*Player) ([]*Pair, error) {
	if len(strikers) != len(defenders) {
		return nil, fmt.Errorf("number of strikers and defenders must be equal")
	}

	pairs := make([]*Pair, 0)
	usedStrikers := make(map[int]bool)
	usedDefenders := make(map[int]bool)

	for len(pairs) < len(strikers) {
		rand.Seed(time.Now().UnixNano())
		sIndex := rand.Intn(len(strikers))
		dIndex := rand.Intn(len(defenders))

		if !usedStrikers[sIndex] && !usedDefenders[dIndex] {
			pairs = append(pairs, &Pair{
				Striker:  strikers[sIndex],
				Defender: defenders[dIndex],
			})
			usedStrikers[sIndex] = true
			usedDefenders[dIndex] = true
		}
	}

	return pairs, nil
}

// WritePairsToFile writes the pairs to a .txt file.
func WritePairsToFile(pairs []*Pair, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, pair := range pairs {
		_, err := fmt.Fprintf(file, "%s - %s\n", pair.Striker.Name, pair.Defender.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	strikers, err := ReadPlayersFromFile("strikers.txt")
	if err != nil {
		panic(err)
	}

	defenders, err := ReadPlayersFromFile("defenders.txt")
	if err != nil {
		panic(err)
	}

	pairs, err := PairPlayers(strikers, defenders)

	err = WritePairsToFile(pairs, "pairs.txt")
	if err != nil {
		panic(err)
	}
}
