package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/lmittmann/tint"
	"github.com/nichady/golte"
	"github.com/nichady/plateful/build"
)

var (
	recipes    = map[string]Recipe{}
	recipeLock sync.RWMutex
)

var requests atomic.Int32

var log = slog.New(tint.NewHandler(os.Stderr, nil))

func main() {
	go func() {
		requests.Store(300)
		time.Sleep(24 * time.Hour)
	}()

	r := chi.NewRouter()
	r.Use(middleware.RedirectSlashes)

	r.Route("/", pages)
	r.Route("/api", api)

	http.ListenAndServe(":9099", r)
}

func pages(r chi.Router) {
	r.Use(build.Golte)

	r.Use(golte.Layout("layout/layout"))
	r.Get("/", golte.Page("page/index"))
	r.Get("/recipes/{id}", func(w http.ResponseWriter, r *http.Request) {
		recipeLock.RLock()
		defer recipeLock.RUnlock()
		recipe, ok := recipes[chi.URLParam(r, "id")]
		if !ok {
			golte.RenderError(w, r, "Not Found", http.StatusNotFound)
		}

		golte.RenderPage(w, r, "page/recipe", map[string]any{
			"recipe": recipe,
		})
	})
	r.Get("/recipes", func(w http.ResponseWriter, r *http.Request) {
		recipeLock.RLock()
		defer recipeLock.RUnlock()
		golte.RenderPage(w, r, "page/recipes", map[string]any{
			"recipes": recipes,
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		golte.RenderError(w, r, "Not Found", http.StatusNotFound)
	})
}

func api(r chi.Router) {
	r.Post("/generate", func(w http.ResponseWriter, r *http.Request) {
		if requests.Load() <= 0 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		var g Generate
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		requests.Add(-1)

		id := RandomID()
		log := log.With("recipeID", id)
		log.Info("generating recipe", "generate", g)

		recipe, err := generateRecipe(g)
		if err != nil {
			log.Error("could not generate recipe", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Info("successfully generated recipe", "recipe", recipe)

		recipeLock.Lock()
		defer recipeLock.Unlock()
		recipes[id] = recipe
		fmt.Fprintf(w, "/recipes/%s", id)
	})
}

type Generate struct {
	Times              []string
	Flavors            []string
	InludeIngredients  string
	ExcludeIngredients string
	Locations          []string
}

type Recipe struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	Image        string   `json:"image"`
}
