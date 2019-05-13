package main

import (
	"flag"
	"fmt"
)

func main() {

	teamPtr := flag.Bool("team", false, "set this to true to select a list of teams")
	//playerPtr := flag.String("player", "", "name a player to list their stats")

	flag.Parse()

	if *teamPtr {
		teams, err := getTeams()
		if err != nil {
			fmt.Errorf("%v\n",err)
		}

		for _,team := range teams {
			fmt.Println(team.name)
		}
	}

}