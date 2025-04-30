package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type Player struct {
	Name  string
	Email string
}

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Player    *Player
	Questions []Question
	Score     float32
}

func (g *GameState) Init() {
	fmt.Println("Olá, seja bem vindo ao Quiz da Copa Do Mundo!")

	player := Player{}

	fmt.Println("Qual é o seu nome:")
	name := ReadInteraction()
	player.Name = name

	fmt.Println("Qual é o seu endereço de e-mail:")
	email := ReadInteraction()
	player.Email = email

	g.Player = &player

	g.ProcessCSV()

	fmt.Printf("Vamos ao jogo, %s", player.Name)
}

func (g *GameState) ProcessCSV() {
	file, err := os.Open("questions.csv")
	if err != nil {
		panic("Erro ao abrir CSV")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Erro ao ler conteúdo do CSV")
	}

	for index, row := range records {
		if index != 0 {
			fmt.Println(row)
			// TODO: Create Question{} and add into GameState{}
		}
	}
}

func ReadInteraction() string {
	fmt.Print(">> ")
	reader := bufio.NewReader(os.Stdin)
	interaction, err := reader.ReadString('\n')
	if err != nil {
		panic("Erro ao ler a string")
	}

	return interaction
}

func main() {
	game := &GameState{}
	game.Init()
}
