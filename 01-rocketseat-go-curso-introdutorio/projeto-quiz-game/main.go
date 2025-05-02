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

type Game struct {
	Player    Player
	Questions []Question
	Score     float32
}

func (g *Game) ProcessCSV() {
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
			correctAnswer, err := toInt(row[5])
			if err != nil {
				panic("Resposta inválida carregada do CSV")
			}
			question := Question{
				Text:    row[0],
				Options: row[1:5],
				Answer:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *Game) Init() {
	fmt.Println("Olá, seja bem vindo ao Quiz da Copa Do Mundo!")

	player := Player{}

	fmt.Println("Qual é o seu nome:")
	name := ReadInteraction()
	player.Name = name

	fmt.Println("Qual é o seu endereço de e-mail:")
	email := ReadInteraction()
	player.Email = email

	g.Player = player

	fmt.Printf("Vamos ao jogo, %s", player.Name)
}

func (g *Game) Run() {
	for indexQ, question := range g.Questions {
		fmt.Println(toStr(indexQ+1)+". -", question.Text)
		for indexO, option := range question.Options {
			fmt.Println("["+toStr(indexO+1)+"]", option)
		}
		fmt.Println("Escolha um alterantiva:")
		var playerAnswer int
		for {
			interaction := ReadInteraction()
			interaction = interaction[:len(interaction)-1]
			answer, err := toInt(interaction)
			if err != nil || answer > len(question.Options) || answer < 1 {
				fmt.Printf("A resposta deve ser um inteiro entre %d e %d", 1, len(question.Options))
			} else {
				playerAnswer = answer
				break
			}
		}
		fmt.Printf("\n\n\n")

		if playerAnswer == question.Answer {
			g.Score += 10
		}
	}

	fmt.Printf("Score = %f\n", g.Score/float32(len(g.Questions)))
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

func toInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func toStr(i int) string {
	return strconv.Itoa(i)
}

func main() {
	game := &Game{}
	go game.ProcessCSV()
	game.Init()
	game.Run()

}
