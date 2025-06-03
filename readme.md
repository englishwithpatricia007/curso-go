# QuizGo

Este projeto é um quiz de perguntas e respostas em Go, que lê questões de um arquivo CSV e interage com o usuário via terminal.

## Como funciona

- O programa lê as perguntas do arquivo `quizgo.csv`.
- Solicita o nome do usuário e inicia o quiz.
- Para cada pergunta, exibe as opções e solicita a resposta.
- Ao final, mostra a pontuação total.

## Ponteiros

O uso de ponteiros é fundamental neste projeto. A estrutura `GameState` é manipulada por referência (ponteiro) em funções como `Init`, `ProcessCSV` e `Run`. Isso garante que as alterações feitas no estado do jogo (nome do jogador, pontuação, perguntas carregadas) persistam entre as funções, sem a necessidade de retornar valores ou copiar estruturas.

Exemplo:

```go
func main() {
    game := &GameState{Points: 0} // Ponteiro para GameState
    game.Init()
    game.Run()
}
```

## Imports utilizados

O projeto utiliza os seguintes pacotes da biblioteca padrão do Go:

- `bufio`: Para leitura eficiente da entrada do usuário.
- `encoding/csv`: Para leitura do arquivo CSV com as perguntas.
- `errors`: Para tratamento de erros personalizados.
- `fmt`: Para formatação e exibição de mensagens no terminal.
- `os`: Para manipulação de arquivos e entrada/saída.
- `strconv`: Para conversão de strings em inteiros.
- `time`: Para controle de tempo (exemplo: aguardar carregamento das perguntas).

## Execução

Para rodar o projeto:

```sh
go run main.go
```

---

- Veja a implementação principal em `main.go`.
- Questões e respostas estão em `quizgo.csv`.
