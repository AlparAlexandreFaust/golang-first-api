# API de Tarefas em Go

Esta é uma API REST simples para gerenciamento de tarefas, desenvolvida em Go.

## Requisitos

- Go 1.21 ou superior

## Como executar

1. Clone este repositório
2. Navegue até a pasta do projeto
3. Execute o comando:

```bash
go run main.go
```

O servidor será iniciado na porta 8080.

## Endpoints disponíveis

### Página inicial
- **GET /** - Retorna uma mensagem de boas-vindas

### Tarefas
- **GET /health** - Retorna o Health Check
- **GET /api/tarefas** - Lista todas as tarefas
- **POST /api/tarefas** - Cria uma nova tarefa
- **GET /api/tarefas/{id}** - Retorna uma tarefa específica
- **PUT /api/tarefas/{id}** - Atualiza uma tarefa existente
- **DELETE /api/tarefas/{id}** - Remove uma tarefa

## Exemplos de uso

### Obter o Health Check
```bash
curl http://localhost:8080/health
```

### Criar uma nova tarefa
```bash
curl -X POST http://localhost:8080/api/tarefas \
  -H "Content-Type: application/json" \
  -d '{"titulo": "Minha nova tarefa", "concluida": false}'
```

### Listar todas as tarefas
```bash
curl http://localhost:8080/api/tarefas
```

### Obter uma tarefa específica
```bash
curl http://localhost:8080/api/tarefas/1
```

### Atualizar uma tarefa
```bash
curl -X PUT http://localhost:8080/api/tarefas/1 \
  -H "Content-Type: application/json" \
  -d '{"titulo": "Tarefa atualizada", "concluida": true}'
```

### Remover uma tarefa
```bash
curl -X DELETE http://localhost:8080/api/tarefas/1
``` 