package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type projectsResponse struct {
	Projects []project `json:"projects"`
}

var mockProjects = projectsResponse{
	Projects: []project{
		{
			ID:          "nutriscan-mvp",
			Name:        "NutriScan MVP",
			Description: "Сканирование штрихкода и понятная оценка состава продукта.",
			Status:      "in_progress",
			CreatedAt:   time.Date(2026, time.January, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          "product-catalog",
			Name:        "Каталог продуктов",
			Description: "Каталог продуктов, составов, аллергенов и пищевых ограничений.",
			Status:      "planned",
			CreatedAt:   time.Date(2026, time.February, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          "store-prices",
			Name:        "Агрегация цен",
			Description: "Сравнение предложений ближайших магазинов.",
			Status:      "planned",
			CreatedAt:   time.Date(2026, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
	},
}

// New registers the HTTP routes served by the API.
func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /projects", getProjects)
	return securityHeaders(mux)
}

func getProjects(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(mockProjects); err != nil {
		http.Error(w, "encode response", http.StatusInternalServerError)
	}
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
