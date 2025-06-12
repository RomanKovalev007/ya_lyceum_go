package main

import (
	"fmt"
	"slices"
)


type Player struct {
    Name      string
    Goals     int
    Misses    int
    Assists   int
    Rating    float64
}

func (p *Player) calculateRating(){
	rating := (float64(p.Goals) + float64(p.Assists) / 2)
	if p.Misses != 0{
		rating /= float64(p.Misses)
	}
	p.Rating = rating
	
}

func NewPlayer(name string, goals, misses, assists int) Player{
	player := Player{Name: name, Goals: goals, Misses: misses, Assists: assists}
	player.calculateRating()
	return player
}

func goalsSort(players []Player) []Player{
	slices.SortFunc(players, func(a, b Player) int {
		if a.Goals != b.Goals{
			return b.Goals - a.Goals
		}
		switch a.Name < b.Name{
			case true:
				return -1
			case false:
				return 1
			default:
				return 0
			}
	})
	return players
}

func ratingSort(players []Player) []Player{
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating < b.Rating{
			return 1
		} else if a.Rating > b.Rating{
			return -1
		}
		switch a.Name < b.Name{
		case true:
			return -1
		case false:
			return 1
		default:
			return 0
		}
	})
	return players
}

func gmSort(players []Player) []Player{
	slices.SortFunc(players, func(a, b Player) int {
        ratioA := float64(a.Goals)
        if a.Misses != 0 {
            ratioA /= float64(a.Misses)
        }

        ratioB := float64(b.Goals)
        if b.Misses != 0 {
            ratioB /= float64(b.Misses)
        }

        if ratioA < ratioB {
            return 1
        } else if ratioA > ratioB {
            return -1
        }
        switch a.Name < b.Name{
		case true:
			return -1
		case false:
			return 1
		default:
			return 0
		}
	})
	return players
}

func main(){
	players := []Player{
		NewPlayer("Player1", 10, 5, 3),
		NewPlayer("Player2", 15, 7, 2),
		NewPlayer("Player3", 8, 2, 5),
		NewPlayer("Player4", 15, 7, 2),
	}
	fmt.Println(goalsSort(players))
	fmt.Println(ratingSort(players))
	fmt.Println(gmSort(players))
}