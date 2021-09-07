package main

import "fmt"

type Gamer struct{
	Name string
	Games []string
}

func (gamer *Gamer) addGame(game string){
	gamer.Games = append(gamer.Games, game)
}

func main(){
	games := Gamer{Name: "Zelda"}

	games.addGame("FIFA")
	games.addGame("PES")

	for _, game := range games.Games{
		fmt.Println(game)
	}
}