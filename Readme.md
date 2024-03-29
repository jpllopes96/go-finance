Postgres e Docker:
 - Iniciar o docker e instalar o postgres:
  - docker pull postgres:14-alpine
- Rodar a imagem do docker conectando ao banco:
  - docker run --name postgres -p 5433:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
- executar o container:
  - docker exec -it postgres sh
  - psql -U postgres  -- \q to exit

- Configurar no PGADMIN

- criar a db com o comando dentro do postgres - docker exec -it ... não precisa executar o psql
 - createdb --username=postgres --owner=postgres go_finance
  

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
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/go_finance?sslmode=disable" -verbose up
// executar down
//executar migrations
migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/go_finance?sslmode=disable" -verbose down


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
    - Criamos a o util/random para gerar strings aleatórias
    - Testes de categorias
    - Testes Accounts
  - Github workflow - pipeline que roda os testes no github com workflow
   - criar a pasta .github/workflows quando sobe isso ele sabe que tem que trabalhar com a pipeline
   - criar o make file
   - ao dar o comando push ele vai executar e testar o projeto, podemos ver no github em actions

-- API
 - GIN Framework - Instalar - go get -u github.com/gin-gonic/gin
 - Criar pasta API e criar o server
 - Criar as rotas do user e funcs do user
 - criar o main go na raiz
 - adicionar o server no make file

 - Criar API da Category
 - Criar API de Accounts
 - Criar API dos graphs e reports

- Variavel Ambinets .ENV
  - https://github.com/joho/godotenv

- Validações ao criar conta
  - O tipo da categoria ser do mesmo tipo da conta

- Criar o hash de senhas do user password

- Melhorar os filtros
 - fazer uma query para cada filtro no category and account para pegar o array(no arquivo sql)
 - quando é string usa o Like que fica como "opcional" o valor mesmo send AND o %""% vai funcionar
 - quando é INT32 e Data COALESCE(@category_id, a.category_id) vai retornar a cat ID se tiver valor ou true(empty) - Porem no GO isso não vem nulo -- isso muda o tipo da campo para null então hora de chamar temos que fazer assim:
 Date: sql.NullTime{
			Time:  lastAccount.Date,
			Valid: true,
		},

-- Endpoint de login
 - descriptografar senha
 - criar token
 - validar token

-- auto reload ao salvar - live reload
  - https://github.com/cosmtrek/air
 - go install github.com/cosmtrek/air@latest
  - air init


-- CORS para integrar com o frontend
  - Proteção do backend para que outros dominios não acessem a api
  
-- Retornar o agent ID ao logar

-- Alterar o find all para pegar os parametros por query string
 --passa o form na struct e troca de should bind json para shouldBindQuery



