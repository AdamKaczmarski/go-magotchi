package src

import (
	"fmt"
	"gomagotchi/src/game"
	"time"
)

var gomagotchi Gomagotchi

const gameScoreHappinessMultiplier = 0.2
const PET, STATS, GAME, SWEEP, START, STOP = "pet", "stats", "game", "sweep", "start", "stop"

// I don't like it perhaps the map solution might be better but idk
// how to pass it to the ActOnCOmmand without impacting anyone else
// perhaps I'd have to do a feed command that cmd.feed would call
const FEED_BURGER, FEED_CAKE = "feed_burger", "feed_cake"

func ActOnCommand(cmd string) {
	switch cmd {
	case START:
		initGomagotchi()
		gomagotchi.printStats()
	case STATS:
		gomagotchi.printStats()
	case PET:
		gomagotchi.pet()
	case GAME:
		score := game.Game()
		gomagotchi.adjustHappiness(float32(score) * gameScoreHappinessMultiplier)
	case SWEEP:
		gomagotchi.sweepPoops()
	case FEED_BURGER:
		gomagotchi.feed(feedingFoods.burger)
	case FEED_CAKE:
		gomagotchi.feed(feedingFoods.cake)
	case STOP:
		fmt.Println("Saving and stopping the Gomagotchi")
		gomagotchi.saveGomagotchi()
	default:
		fmt.Println("Unknown command")
	}
}

func run() {
	time.Sleep(10000)

	fmt.Println("shutting down")
}

func initGomagotchi() {
	//Load struct from fiile
	fmt.Println("Loading gomagotchi from file")
	loaded := loadGomagotchi(&gomagotchi)
	if !loaded {
		//If file doesnt exist, create new
		fmt.Println("Creating a new gomagotchi")
		gomagotchi = newGomagotchi()
		gomagotchi.saveGomagotchi()
	}
	//Run in the background
	fmt.Println("Running it in background")
	//daemonize the program
	// go run()
}
