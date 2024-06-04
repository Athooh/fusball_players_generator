package players

import (
	"bufio"
	"os"
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
