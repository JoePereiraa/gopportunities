# Gopportunities

API REST em Go (Gin) para cadastro e gestão de **vagas de emprego** (job openings). Os dados são persistidos em **SQLite** (`./db/main.db`).

## Requisitos

- [Go](https://go.dev/dl/) (versão do projeto: ver `go.mod`)

## Como executar

```bash
make run
# ou: go run main.go
```

O servidor sobe em **http://localhost:8080**.

## Documentação interativa (Swagger)

Com a aplicação em execução:

- **http://localhost:8080/swagger/index.html**

Para regenerar os artefatos Swagger após alterar comentários `swag`:

```bash
make docs
```

## Base da API

| Item        | Valor              |
| ----------- | ------------------ |
| Prefixo     | `/api/v1`          |
| Formato     | JSON               |
| Banco       | SQLite em `./db/` |

## Endpoints

| Método | Caminho            | Descrição                          |
| ------ | ------------------ | ---------------------------------- |
| `GET`  | `/api/v1/openings` | Lista todas as vagas               |
| `GET`  | `/api/v1/opening`  | Busca uma vaga por `id` (query)    |
| `POST` | `/api/v1/opening`  | Cria uma vaga                      |
| `PUT`  | `/api/v1/opening`  | Atualiza uma vaga por `id` (query) |
| `DELETE` | `/api/v1/opening` | Remove uma vaga por `id` (query)   |

### Query `id`

Usado em **GET**, **PUT** e **DELETE** em `/opening`:

- Exemplo: `/api/v1/opening?id=1`

### Corpo JSON — criar vaga (`POST /opening`)

Campos obrigatórios:

| Campo      | Tipo    | Descrição        |
| ---------- | ------- | ---------------- |
| `role`     | string  | Cargo / função   |
| `company`  | string  | Empresa          |
| `location` | string  | Localização      |
| `remote`   | boolean | Trabalho remoto  |
| `link`     | string  | URL da vaga      |
| `salary`   | número  | Salário (`> 0`) |

### Corpo JSON — atualizar vaga (`PUT /opening?id=...`)

É necessário enviar **pelo menos um** campo válido para atualizar. Campos não enviados (ou vazios) permanecem como estão.

Atualização parcial suportada pelo código para: `role`, `company`, `location`, `remote`, `salary`.

### Respostas de sucesso (HTTP 200)

Formato geral:

```json
{
  "message": "operation from handler: <operação> successfull",
  "data": { }
}
```

- **Listagem** (`GET /openings`): `data` é um **array** de vagas.
- **Demais operações**: `data` é um **objeto** vaga.

Em erro, a API responde com JSON contendo `message` e `errorCode` (código HTTP usado como número).

### Códigos HTTP usuais

| Código | Situação típica                          |
| ------ | ---------------------------------------- |
| 200    | Sucesso                                  |
| 400    | JSON inválido, validação ou `id` faltando |
| 404    | Vaga não encontrada                      |
| 500    | Erro interno (banco, etc.)               |

## Exemplos rápidos (curl)

**Listar vagas**

```bash
curl -s http://localhost:8080/api/v1/openings
```

**Criar vaga**

```bash
curl -s -X POST http://localhost:8080/api/v1/opening \
  -H "Content-Type: application/json" \
  -d '{"role":"Dev Go","company":"ACME","location":"SP","remote":true,"link":"https://example.com/job","salary":12000}'
```

**Buscar por id**

```bash
curl -s "http://localhost:8080/api/v1/opening?id=1"
```

**Atualizar**

```bash
curl -s -X PUT "http://localhost:8080/api/v1/opening?id=1" \
  -H "Content-Type: application/json" \
  -d '{"salary":15000}'
```

**Remover**

```bash
curl -s -X DELETE "http://localhost:8080/api/v1/opening?id=1"
```

## Outros comandos (Makefile)

| Comando    | Ação                    |
| ---------- | ----------------------- |
| `make build` | Compila binário `gopportunities` |
| `make test`  | Executa testes          |
| `make clean` | Remove o binário e a pasta `./docs` (conforme o `makefile`) |

---

Projeto: **gopportunities** — API de oportunidades de emprego.
