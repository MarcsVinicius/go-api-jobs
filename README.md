
# Go Jobs API

Este projeto é uma API moderna de oportunidades de trabalho construída usando Golang.

---

## Instalação

Para usar o projeto siga esses passos

1. Clone o repositório: `git clone https://github.com/username/repo-name.git`
2. Instale as dependências: `go mod download`
3. Build a aplicação: `go build`
4. Rode a aplicação:  `go run main`

## Makefile Comandos

O projeto inclui um Makefile para ajudá-lo a gerenciar tarefas comuns com mais facilidade. Aqui está uma lista dos comandos disponíveis e uma breve descrição do que eles fazem:

- `make run`: Executa a aplicação sem gerar a documentação da API.
- `make run-with-docs`: Gere a documentação da API usando o Swag e execute o aplicativo.
- `make build`: Compile o aplicativo e crie um arquivo executável chamado `gopportunities`.
- `make test`: Executa testes para todos os pacotes do projeto.
- `make docs`: Gera a documentação da API usando Swag.
- `make clean`: Remova o executável `gopportunities` e exclua o diretório `./docs`.

## Tecnlogias Utilizadas


Este projeto utiliza as seguintes ferramentas:

- [Golang](https://golang.org/) para desenvolvimento de back-end
- [Go-Gin](https://github.com/gin-gonic/gin) para gerenciamento de rota
- [GoORM](https://gorm.io/) para comunicação de banco de dados
- [SQLite](https://www.sqlite.org/index.html) como banco de dados
- [Swagger](https://swagger.io/) para documentação e teste da API


## Uso

Após a execução da API, você pode usar a interface do usuário do Swagger para interagir com os endpoints para pesquisar, criar, editar e excluir oportunidades de trabalho. A API pode ser acessada em `http://localhost:$PORT/swagger/index.html`.

Padrão $PORT se não for fornecido=8080.
