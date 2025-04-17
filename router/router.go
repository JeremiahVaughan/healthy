package router

import (
    "net/http"
    "fmt"

    "github.com/JeremiahVaughan/healthy/controllers"
)

type Router struct {
    mux *http.ServeMux
}

func New(controllers *controllers.Controllers) *Router {
    mux := http.NewServeMux()
    mux.HandleFunc("/infra-graph", controllers.InfraGraph.Graph)
    mux.HandleFunc("/hotreload", controllers.TemplateLoader.HandleHotReload)
    mux.HandleFunc("/health", controllers.Health.Check)
    return &Router{ 
        mux: mux,
    }
}

func (r *Router) Run() error {
    err := http.ListenAndServe(":7777", r.mux)
    if err != nil {
        return fmt.Errorf("error, when serving http. Error: %v", err)
    }
    return nil
}
