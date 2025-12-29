# API Alunos (Go + Gin + GORM)

Projeto exemplo em Go que expõe endpoints para gerenciamento de alunos, com persistência em PostgreSQL, execução via Docker Compose e integração contínua configurada em GitHub Actions.

## Estrutura principal

- `main.go`: ponto de entrada da aplicação.
- `database/db.go`: inicializa a conexão com o Postgres (GORM) e faz AutoMigrate.
- `models/alunos.go`: definição do modelo `Aluno` e validações.
- `controllers/controller.go`: handlers HTTP (CRUD e páginas HTML).
- `routes/route.go`: configuração do router Gin e rotas.
- `docker-compose.yml`: define serviços `postgres` e `app` para desenvolvimento.
- `.github/workflows/go.yml`: pipeline de Integração Contínua (lint, testes, build).

## Requisitos

- Docker & Docker Compose (para execução via containers)
- Go 1.22 (para compilar/rodar localmente)

## Variáveis de ambiente (usadas pela aplicação)

As variáveis podem ser definidas no ambiente ou através do `docker-compose`:

- `DB_HOST` (ex.: `postgres` quando via Docker Compose)
- `DB_USER` (ex.: `root`)
- `DB_PASSWORD` (ex.: `root`)
- `DB_NAME` (ex.: `root`)
- `DB_PORT` (ex.: `5432`)

## Executando em desenvolvimento (com Docker Compose)

1. Subir os serviços (Postgres e app):

```bash
docker compose up -d
```

2. A aplicação estará disponível em `http://localhost:8080`.

Observação: o serviço `app` do `docker-compose.yml` usa `go run main.go`, portanto a imagem é `golang:1.22` e monta o diretório do projeto para permitir edição ao vivo.

## Executando localmente sem Docker

Defina as variáveis de ambiente necessárias e rode:

```bash
export DB_HOST=localhost
export DB_USER=root
export DB_PASSWORD=root
export DB_NAME=root
export DB_PORT=5432
go run main.go
```

## Comandos úteis (Makefile)

- `make start` — executa `docker compose up -d`.
- `make lint` — executa `golangci-lint` via container.
- `make test` — executa `go test main_test.go` no container `app`.
- `make ci` — encadeia `start`, `lint` e `test`.

## Endpoints principais

- `GET /alunos` — retorna todos os alunos.
- `GET /alunos/:id` — retorna aluno por ID.
- `POST /alunos` — cria novo aluno (JSON com `nome`, `rg`, `cpf`).
- `PATCH /alunos/:id` — atualiza aluno (JSON) por ID.
- `DELETE /alunos/:id` — deleta aluno por ID.
- `GET /alunos/cpf/:cpf` — busca aluno pelo CPF.
- `GET /index` — renderiza página `index.html` com lista de alunos.
- `GET /:nome` — rota de saudação (ex.: `/rogerio`).

Exemplo de payload para criação:

```json
{
  "nome": "Fulano",
  "rg": "123456789",
  "cpf": "12345678901"
}
```

Observação: os campos `rg` e `cpf` possuem validações de tamanho e apenas dígitos.

## Testes e CI

- O workflow `.github/workflows/go.yml` sobe um container `postgres`, executa `golangci-lint`, roda os testes com variáveis via `secrets` e compila o binário.
- Para rodar os testes localmente via Makefile: `make test`.

## Observações e boas práticas

- A conexão com o banco implementa tentativas de retry para tolerar startup do Postgres.
- Em produção, não use credenciais padrão (`root`/`root`) nem `sslmode=disable` sem avaliar risco. Configure backups e políticas de segurança.

## Próximos passos sugeridos

- Adicionar README mais detalhado com exemplos de requests (curl/Postman).
- Incluir testes de integração para endpoints usando um banco controlado pelos testes.
- Adicionar documentação OpenAPI/Swagger se desejar consumíveis públicos.

## Mais um teste
- Tentando dar um pull request
