package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Criar uma requisição HTTP GET para "/"
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Criar um ResponseRecorder para gravar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	// Chamar o handler com a requisição e o response recorder
	handler.ServeHTTP(rr, req)

	// Verificar o status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou status code incorreto: recebido %v esperado %v",
			status, http.StatusOK)
	}

	// Verificar o corpo da resposta
	expected := "Bem-vindo à API de Tarefas!"
	if rr.Body.String() != expected {
		t.Errorf("handler retornou corpo inesperado: recebido %v esperado %v",
			rr.Body.String(), expected)
	}
}

func TestHealthHandler(t *testing.T) {
	// Criar uma requisição HTTP GET para "/health"
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Criar um ResponseRecorder para gravar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)

	// Chamar o handler com a requisição e o response recorder
	handler.ServeHTTP(rr, req)

	// Verificar o status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou status code incorreto: recebido %v esperado %v",
			status, http.StatusOK)
	}

	// Verificar se o corpo da resposta contém tarefas
	var health Health
	err = json.Unmarshal(rr.Body.Bytes(), &health)
	if err != nil {
		t.Errorf("Erro ao decodificar resposta JSON: %v", err)
	}

	// Verificar se pelo menos as tarefas iniciais estão presentes
	if health.Status != "OK" {
		t.Errorf("Status de saúde incorreto: recebido %v esperado %v",
			health.Status, "OK")
	}
}

func TestListarTarefas(t *testing.T) {
	// Criar uma requisição HTTP GET para "/api/tarefas"
	req, err := http.NewRequest("GET", "/api/tarefas", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Criar um ResponseRecorder para gravar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tarefasHandler)

	// Chamar o handler com a requisição e o response recorder
	handler.ServeHTTP(rr, req)

	// Verificar o status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou status code incorreto: recebido %v esperado %v",
			status, http.StatusOK)
	}

	// Verificar se o corpo da resposta contém tarefas
	var tarefasRecebidas []Tarefa
	err = json.Unmarshal(rr.Body.Bytes(), &tarefasRecebidas)
	if err != nil {
		t.Errorf("Erro ao decodificar resposta JSON: %v", err)
	}

	// Verificar se pelo menos as tarefas iniciais estão presentes
	if len(tarefasRecebidas) < 2 {
		t.Errorf("Número de tarefas menor que o esperado: recebido %v esperado pelo menos 2",
			len(tarefasRecebidas))
	}
}

func TestCriarTarefa(t *testing.T) {
	// Criar uma nova tarefa para o teste
	novaTarefa := Tarefa{
		Titulo:    "Tarefa de Teste",
		Concluida: false,
	}

	// Converter para JSON
	jsonTarefa, err := json.Marshal(novaTarefa)
	if err != nil {
		t.Fatal(err)
	}

	// Criar uma requisição HTTP POST para "/api/tarefas"
	req, err := http.NewRequest("POST", "/api/tarefas", bytes.NewBuffer(jsonTarefa))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Criar um ResponseRecorder para gravar a resposta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tarefasHandler)

	// Chamar o handler com a requisição e o response recorder
	handler.ServeHTTP(rr, req)

	// Verificar o status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler retornou status code incorreto: recebido %v esperado %v",
			status, http.StatusCreated)
	}

	// Verificar se a tarefa foi criada corretamente
	var tarefaCriada Tarefa
	err = json.Unmarshal(rr.Body.Bytes(), &tarefaCriada)
	if err != nil {
		t.Errorf("Erro ao decodificar resposta JSON: %v", err)
	}

	if tarefaCriada.Titulo != novaTarefa.Titulo {
		t.Errorf("Título da tarefa incorreto: recebido %v esperado %v",
			tarefaCriada.Titulo, novaTarefa.Titulo)
	}

	if tarefaCriada.Concluida != novaTarefa.Concluida {
		t.Errorf("Status de conclusão incorreto: recebido %v esperado %v",
			tarefaCriada.Concluida, novaTarefa.Concluida)
	}
}
