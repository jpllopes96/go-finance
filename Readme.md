Instalar o migrate cli: https://github.com/golang-migrate/migrate
    -Instalar o scope: https://scoop.sh/ rodando os comandos no powershell
        - Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
        - irm get.scoop.sh | iex
    - Depois de ter o scoop instlar o migrate
        - scoop install migrate

Comandos:
//criar migrations
migrate create -ext sql -dir db/migration -seq user_table
//executar migrations
migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/go_finance?sslmode=disable" -verbose up
// executar down
//executar migrations
migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/go_finance?sslmode=disable" -verbose up


Instalar o psql https://docs.sqlc.dev/en/stable/overview/install.html
 - go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
 - criar o arquivo sqlc.yaml
 - criar as folders query e sqlc em db 
 - executar para gerar o codigo com o comando no powersehell
    - docker run --rm -v C:\Users\jplopes\Documents\gofinance:/src -w /src sqlc/sqlc generate

- isso vai gerar os codigos para as ações com os bancos
- Vamos criar as store para poder acessar os metodos, fica dentro do sqlc - é uma interface que tem o Querier( que tem todas funções do banco de dados)

- Testes unitarios
     - instalar postegres de teste lib pq https://github.com/lib/pq
        - go get github.com/lib/pq
    - instalar a testfy https://github.com/stretchr/testify
        - go get github.com/stretchr/testify
    - Criar o main_test.go dentro de sqlc para testar connec com banco
    - Criar o user_test e fazer os testes