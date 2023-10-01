package src

import (
	"encoding/json"
	"fmt"
	"os"
)

const structPath string = "./struct.json"
const maxHunger, maxHappines, maxDiscipline, maxAge = 4.0, 4.0, 100, 100

type Gomagotchi struct {
	Name       string
	Age        int8
	Weight     int16
	Hunger     float32
	Happiness  float32
	Poops      int8
	Awake      bool
	Discipline int8
}

// Create a new Gomagotchi
func newGomagotchi() Gomagotchi {
	return Gomagotchi{
		Name:       "Gomagotchi",
		Age:        1,
		Weight:     5,
		Hunger:     3,
		Happiness:  3,
		Poops:      0,
		Awake:      true,
		Discipline: 0,
	}
}

// Either add or decrease the hunger of Gomagotchi
func (gomagotchi *Gomagotchi) adjustHunger(amount float32) {
	if amount > 0 && gomagotchi.Hunger+amount > maxHunger {
		gomagotchi.Hunger = maxHunger
	} else if amount < 0 && gomagotchi.Hunger+amount < 0 {
		gomagotchi.Hunger = 0
	} else {
		gomagotchi.Hunger += amount
	}
	gomagotchi.saveGomagotchi()
}

// Either add or decrease the happiness of Gomagotchi
func (gomagotchi *Gomagotchi) adjustHappiness(amount float32) {
	if amount > 0 && gomagotchi.Happiness+amount > maxHappines {
		gomagotchi.Happiness = maxHappines
	} else if amount < 0 && gomagotchi.Happiness+amount < 0 {
		gomagotchi.Happiness = 0
	} else {
		gomagotchi.Happiness += amount
	}
	gomagotchi.saveGomagotchi()
}

// Pet the Gomagotchi to add 0.25 of happiness to his life :)
func (gomagotchi *Gomagotchi) pet() {
	gomagotchi.adjustHappiness(+0.25)
	gomagotchi.saveGomagotchi()
}

func (gomagotchi *Gomagotchi) sweepPoops() {
	fmt.Println("Sweeping the poops")
	gomagotchi.Poops = 0
	gomagotchi.saveGomagotchi()
}

func (gomagotchi *Gomagotchi) adjustWeight(amount int16) {
	if amount+gomagotchi.Weight < 1 {
		gomagotchi.Weight = 1
	} else {
		gomagotchi.Weight += amount
	}
	gomagotchi.saveGomagotchi()
}

// Adjust the happiness, hunger and weight with food's properties
func (gomagotchi *Gomagotchi) feed(foodie food) {
	if gomagotchi.isFull() {
		fmt.Println("I'm full -.-")
		return
	}
	gomagotchi.adjustHappiness(foodie.happiness)
	gomagotchi.adjustHunger(foodie.hunger)
	gomagotchi.adjustWeight(foodie.weight)
	fmt.Println("Yum :))")
}

func (gomagotchi *Gomagotchi) isFull() bool {
	return gomagotchi.Hunger == maxHunger
}

func (gomagotchi Gomagotchi) String() string {
	return fmt.Sprintf("Name: %q\n\tAge: %d,\n\tWeight: %d,\n\tHunger: %.2f,\n\tHappiness: %.2f,\n\tDiscipline: %d%%\n\tPoops %d,\n\tAwake: %t\n",
		gomagotchi.Name,
		gomagotchi.Age,
		gomagotchi.Weight,
		gomagotchi.Hunger,
		gomagotchi.Happiness,
		gomagotchi.Discipline,
		gomagotchi.Poops,
		gomagotchi.Awake)
}

func (gomagotchi *Gomagotchi) printStats() {
	fmt.Printf("%+v\n", gomagotchi)
}

// Loads Gomagotchi from a JSON file
func loadGomagotchi(gomagotchi *Gomagotchi) bool {
	fileContent, err := os.ReadFile(structPath)
	if err != nil {
		fmt.Printf("Couldn't find %v\n", structPath)
		return false
	}
	unMarshal(fileContent, gomagotchi)
	return true
}

// Saves Gomagotchi to a JSON file
func (gomagotchi *Gomagotchi) saveGomagotchi() {
	serializedGomagotchi := marshal(gomagotchi)
	if err := os.WriteFile(structPath, serializedGomagotchi, 0666); err != nil {
		fmt.Printf("Error saving gomagotchi to file %v", err)
	}
}

func unMarshal(jsonData []byte, gomagotchi *Gomagotchi) {
	err := json.Unmarshal(jsonData, gomagotchi)
	if err != nil {
		panic(err)
	}
}

func marshal(gomagotchi *Gomagotchi) []byte {
	bytes, err := json.Marshal(gomagotchi)
	if err != nil {
		fmt.Printf("Couldn't encode %+v as JSON", gomagotchi)
	}
	return bytes
}
