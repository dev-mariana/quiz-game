package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text string
	Options []string
	Answer int
}

type GameState struct {
	Name string
	Points int
	Questions []Question
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	
	if err != nil {
		return 0, errors.New("Não é permitido caractere diferente de número. Por favor digite um número")
	}

	return i, nil
}

func (g *GameState) Init() {
	fmt.Println("Seja bem-vindo(a) ao quiz")
	fmt.Println("Escreva seu nome:")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = name

	fmt.Printf("Vamos ao jogo %s", g.Name)
}

func (g *GameState) ProccessCSV() {
	f, err := os.Open("quiz-go.csv")

	if err != nil {
		panic("Erro ao ler o arquivo")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Erro ao ler o arquivo csv")
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Question{
				Text: record[0],
				Options: record[1:5],
				Answer: correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) Run() {
	for index, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Println("Digite uma alternativa")

		var answer int
		var err error 

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')
		
			answer, err = toInt(read[:len(read)-1])

			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			break

		}

		if answer == question.Answer {
			fmt.Println("Parabéns, você acertou!!")
			g.Points += 10 
		} else {
			fmt.Println("Ops! Errou!")
			fmt.Println("------------------------------------")
		}
	}

}

func main() {
	game := &GameState{Points: 0}

	go game.ProccessCSV()

	game.Init()
	game.Run()

	fmt.Printf("Fim de jogo, sua pontuação foi %d pontos\n", game.Points)
}