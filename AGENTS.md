# AGENTS.md — rest-api-with-jwt

Diretrizes para agentes de IA atuando no repositório `rest-api-with-jwt`.

## Contexto do Projeto

API REST em Go (Gin) demonstrando autenticação e autorização via JSON Web Tokens (JWT) com Access e Refresh tokens e invalidação de sessão com Redis.

- Para documentação dos endpoints (`/login`, `/logout`, `/token/refresh`, `/todo`) e instruções de execução local, consulte o [README.md](README.md).
- Especificação OpenAPI disponível em `openapi.yaml`.

## Estrutura do Repositório

```
rest-api-with-jwt/
├── main.go            # Ponto de entrada e configuração do servidor HTTP
├── main_test.go       # Testes de integração/rotas
├── router/            # Definição de rotas e handlers HTTP (Gin)
├── middleware/        # Middlewares de autenticação e validação de tokens JWT
├── model/             # Modelos de domínio (User, Todo, TokenDetails)
├── redis/             # Cliente Redis para persistência de tokens e metadados
├── openapi.yaml       # Especificação OpenAPI 3.0 dos contratos de API
├── .env.example       # Modelo de variáveis de ambiente (ACCESS_SECRET, REFRESH_SECRET)
└── README.md          # Documentação e guia de uso da API
```

## Regras e Comandos

1. **Testes e Build:**
   - Executar testes: `go test -v ./...`
   - Checagem de build: `go build -o /dev/null .`
   - Formatação: `gofmt -s -w .`
2. **Ambiente Local:**
   - Requer Redis rodando na porta 6379: `docker run -d -p 6379:6379 redis:6`
   - Variáveis em `.env` (baseadas em `.env.example`).
3. **Padrões de Desenvolvimento:**
   - Respeitar a especificação em `openapi.yaml` ao alterar ou adicionar endpoints.
   - Commits seguindo Conventional Commits (`feat:`, `fix:`, `refactor:`, `test:`).
   - Comunicação com o usuário em Português do Brasil; código e documentação técnica em Inglês.
