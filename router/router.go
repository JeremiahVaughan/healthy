package router

import (
    "net/http"
    "fmt"
    "context"

    "github.com/JeremiahVaughan/healthy/controllers"
    natsClient "github.com/JeremiahVaughan/healthy/clients/nats"
    "github.com/JeremiahVaughan/healthy/clients"
    "github.com/JeremiahVaughan/healthy/config"
    "github.com/JeremiahVaughan/healthy/ui_util"

    "github.com/nats-io/nats.go"
)

type Router struct {
    mux *http.ServeMux
    nats *natsClient.Client
    subs []Subscription
    pub Topic
}

type Topic struct {
    subject string
    handler func(context.Context)
}

type Subscription struct {
    subject string
    handler nats.MsgHandler
    sub *nats.Subscription
}

func New(
    controllers *controllers.Controllers,
    clients *clients.Clients,
    config config.Config,
) *Router {
    mux := http.NewServeMux()
    mux.HandleFunc("/infra-graph", controllers.InfraGraph.Graph)
    mux.HandleFunc("/dash", controllers.HealthStatus.Dashboard)
    mux.HandleFunc("/healthStatusCheck", controllers.HealthStatus.Check)
    mux.HandleFunc("/clearUnexpectedErrors", controllers.HealthStatus.ClearUnexpectedErrors)
    if config.LocalMode {
        mux.HandleFunc("/hotreload", controllers.TemplateLoader.HandleHotReload)
    }
    mux.HandleFunc("/health", controllers.Health.Check)

    ui_util.InitStaticFiles(mux, config.UiPath + "/static")

    subs := []Subscription{
        {
            subject: "update-health-status",
            handler: controllers.HealthStatus.UpdateHealthStatus,
        },
        {
            subject: "report-unexpected-error",
            handler: controllers.HealthStatus.ReportUnexpectedError,
        },
    }

    pub := Topic{
        subject: controllers.HealthStatus.RefreshAllKey,
        handler: controllers.HealthStatus.RefreshAll,
    }

    return &Router{ 
        mux: mux,
        subs: subs,
        nats: clients.Nats,
        pub: pub,
    }
}

func (r *Router) Run(ctx context.Context) error {

    err := r.Sub()
    if err != nil {
        return fmt.Errorf("error, when Router.Sub() for Router.Run(). Error: %v", err)
    }
    defer r.Unsub()

    go r.pub.handler(ctx)

    err = http.ListenAndServe(":7777", r.mux)
    if err != nil {
        return fmt.Errorf("error, when serving http. Error: %v", err)
    }

    return nil
}

func (r *Router) Sub() error {
    var err error
    for _, s := range r.subs {
        s.sub, err = r.nats.Conn.Subscribe(s.subject, s.handler)
        if err != nil {
            return fmt.Errorf("error, when subscribing with subject: %s. Error: %v", s.subject, err)
        }
    }
    return nil
}

func (r *Router) Unsub() error {
    for _, s := range r.subs {
        err := s.sub.Unsubscribe()
        if err != nil {
            return fmt.Errorf("error, unable to unsubscribe for subject. Error: %s", s.subject)
        }
    }
    return nil
}
