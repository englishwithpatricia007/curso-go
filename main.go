package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

type Question struct {
	Text    string
	Options []string
	Answer  int
}

func (g *GameState) Init() {
	fmt.Println("Seja bem vindo ao quiz")
	fmt.Println("Escreva o seu nome:")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = name

	fmt.Printf("Vamos ao jogo %s", g.Name)
}

func (g *GameState) ProcessCSV() {
	f, err := os.Open("quizgo.csv")

	if err != nil {
		panic("Erro ao ler a string")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Erro ao ler a string")
	}

	for index, record := range records {
		//fmt.Println(record)
		correctAnswer, _ := toInt(record[5])
		if index > 0 {
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}
			g.Questions = append(g.Questions, question)
		}
	}

}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("não é permitido caractere diferente de número")
	}
	return i, nil
}

func (g *GameState) Run() {

	for index, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Println("Digite a alternativa:")

		var answer int
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			fmt.Println(read[:len(read)-1])
			answer, err = toInt(read[:len(read)-2])

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		if answer == question.Answer {
			fmt.Println("Parabéns, você acertou!")
			g.Points += 10
		} else {
			fmt.Println("Ops! Errou! ")
			fmt.Println("-----------------------")
		}
	}
}

func main() {
	game := &GameState{Points: 0}
	go game.ProcessCSV()
	time.Sleep(1 * time.Millisecond)

	game.Init()
	game.Run()

	fmt.Printf("Fim de jogo, você fez %d pontos\n", game.Points)

}
