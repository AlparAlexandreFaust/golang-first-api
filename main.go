package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Tarefa representa uma tarefa na nossa API
type Tarefa struct {
	ID        string `json:"id"`
	Titulo    string `json:"titulo"`
	Concluida bool   `json:"concluida"`
}

type Health struct {
	Status string `json:"status"`
}

// Banco de dados em memória para armazenar tarefas
var tarefas = []Tarefa{
	{ID: "1", Titulo: "Aprender Go", Concluida: false},
	{ID: "2", Titulo: "Criar uma API REST", Concluida: false},
}

func main() {
	// Definir rotas
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/tarefas", tarefasHandler)
	http.HandleFunc("/api/tarefas/", tarefaHandler)

	// Iniciar o servidor
	porta := ":8080"
	fmt.Printf("Servidor iniciado na porta %s\n", porta)
	log.Fatal(http.ListenAndServe(porta, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo à API de Tarefas!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	var health = Health{Status: "OK"}
	health.Status = "OK"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func tarefasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Retornar todas as tarefas
		json.NewEncoder(w).Encode(tarefas)
	case "POST":
		// Adicionar uma nova tarefa
		var novaTarefa Tarefa
		err := json.NewDecoder(r.Body).Decode(&novaTarefa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Gerar ID simples (em produção, use algo mais robusto)
		novaTarefa.ID = fmt.Sprintf("%d", len(tarefas)+1)
		tarefas = append(tarefas, novaTarefa)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(novaTarefa)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func tarefaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extrair ID da URL (formato: /api/tarefas/{id})
	id := r.URL.Path[len("/api/tarefas/"):]

	// Encontrar a tarefa pelo ID
	var tarefaEncontrada *Tarefa
	var indice int

	for i, t := range tarefas {
		if t.ID == id {
			tarefaEncontrada = &tarefas[i]
			indice = i
			break
		}
	}

	if tarefaEncontrada == nil {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		// Retornar a tarefa específica
		json.NewEncoder(w).Encode(tarefaEncontrada)
	case "PUT":
		// Atualizar a tarefa
		var tarefaAtualizada Tarefa
		err := json.NewDecoder(r.Body).Decode(&tarefaAtualizada)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Manter o ID original
		tarefaAtualizada.ID = id
		tarefas[indice] = tarefaAtualizada

		json.NewEncoder(w).Encode(tarefaAtualizada)
	case "DELETE":
		// Remover a tarefa
		tarefas = append(tarefas[:indice], tarefas[indice+1:]...)
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
