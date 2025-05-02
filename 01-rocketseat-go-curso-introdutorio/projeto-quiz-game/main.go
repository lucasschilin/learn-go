package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

			question := Question{
				Text:    row[0],
				Options: row[1:5],
				Answer:  toInt(row[5]),
			}

			g.Questions = append(g.Questions, question)
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

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Erro ao converter valor de string para inteiro")
	}

	return i
}

func main() {
	game := &GameState{}
	go game.ProcessCSV()
	game.Init()

	fmt.Println(game.Questions)

}
