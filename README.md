
# API de Gerenciamento de Posts

Esta é uma API desenvolvida em Go (Golang) para realizar as operações CRUD (Create, Read, Update, Delete) em posts. A API utiliza o framework Gin para a criação das rotas, GORM como ORM e Postgres como banco de dados. O Docker é utilizado para containerização da aplicação e do banco de dados.

## Funcionalidades Principais

-   **Criação de posts**
    
-   **Busca de posts (por ID ou lista completa)**
    
-   **Atualização parcial (PATCH)**
    
-   **Atualização completa (PUT)**
    
-   **Deleção de posts**
    

----------

## Tecnologias Utilizadas

-   **Linguagem:** Go (Golang)
    
-   **Framework:** Gin (v1.10.0)
    
-   **ORM:** GORM
    
-   **Driver de Banco de Dados:** Postgres
    
-   **Containerização:** Docker
    

----------

## Como Configurar o Ambiente

### 1. Clone o repositório

```
git clone <URL_DO_REPOSITORIO>
cd <NOME_DO_REPOSITORIO>
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=nome_do_banco
API_PORT=8080
```

### 3. Suba os containers com Docker

Certifique-se de que o Docker está instalado e em execução. Então, execute:

```
docker-compose up -d
```

Isso irá subir os containers da API e do banco de dados Postgres.

### 4. Instale as dependências

Se estiver desenvolvendo localmente:

```
go mod tidy
```

### 5. Execute a API

Para rodar a aplicação localmente, utilize:

```
go run main.go
```

A API estará disponível em `http://localhost:8080` (ou na porta especificada no arquivo `.env`).

----------

## Endpoints da API

### 1. Criar um novo post

**POST**  `/posts`

#### Request Body:

```
{
  "title": "Meu Post",
  "content": "Conteúdo do post"
}
```

#### Response:

-   **201 Created**
    

```
{
  "id": 1,
  "title": "Meu Post",
  "content": "Conteúdo do post",
  "created_at": "2024-12-20T15:04:05Z"
}
```

----------

### 2. Buscar todos os posts

**GET**  `/posts`

#### Response:

-   **200 OK**
    

```
[
  {
    "id": 1,
    "title": "Meu Post",
    "content": "Conteúdo do post",
    "created_at": "2024-12-20T15:04:05Z"
  },
  {
    "id": 2,
    "title": "Outro Post",
    "content": "Mais conteúdo",
    "created_at": "2024-12-20T16:04:05Z"
  }
]
```

----------

### 3. Buscar um post por ID

**GET**  `/posts/:id`

#### Response:

-   **200 OK**
    

```
{
  "id": 1,
  "title": "Meu Post",
  "content": "Conteúdo do post",
  "created_at": "2024-12-20T15:04:05Z"
}
```

-   **404 Not Found**
    

```
{
  "error": "Post not found"
}
```

----------

### 4. Atualizar um post (completo)

**PUT**  `/posts/:id`

#### Request Body:

```
{
  "title": "Novo Título",
  "content": "Novo conteúdo"
}
```

#### Response:

-   **200 OK**
    

```
{
  "id": 1,
  "title": "Novo Título",
  "content": "Novo conteúdo",
  "created_at": "2024-12-20T15:04:05Z"
}
```

----------

### 5. Atualizar um post (parcial)

**PATCH**  `/posts/:id`

#### Request Body:

```
{
  "content": "Atualização parcial"
}
```

#### Response:

-   **200 OK**
    

```
{
  "id": 1,
  "title": "Meu Post",
  "content": "Atualização parcial",
  "created_at": "2024-12-20T15:04:05Z"
}
```

----------

### 6. Deletar um post

**DELETE**  `/posts/:id`

#### Response:

-   **200 OK**
    

```
{
  "message": "Post deleted successfully"
}
```

-   **404 Not Found**
    

```
{
  "error": "Post not found"
}
```

----------

## Estrutura do Projeto

```
.
├── Dockerfile
├── docker-compose.yml
├── main.go
├── models
│   └── post.go
├── routes
│   └── post_routes.go
├── controllers
│   └── post_controller.go
├── repository
│   └── post_repository.go
├── .env
└── go.mod
```

-   `**main.go**`: Ponto de entrada da aplicação.
    
-   `**models/**`: Contém as definições das structs (ex.: Post).
    
-   `**routes/**`: Define as rotas da API.
    
-   `**controllers/**`: Lógica dos endpoints.
    
-   `**repository/**`: Manipulação de dados com o GORM.
    

----------

## Executando Testes

Adicione testes unitários para os controllers, models e repositories.

Para rodar os testes:

```
go test ./...
```

----------

## Contribuições

Contribuições são bem-vindas! Abra um PR com suas sugestões ou melhorias.

----------

## Licença

Este projeto está sob a licença MIT.
