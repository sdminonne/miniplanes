package www

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"sync"
	"time"

	storagemodels "github.com/amadeusitgroup/miniapp/storage/pkg/gen/models"
	"github.com/amadeusitgroup/miniapp/ui/assets"
	"github.com/gorilla/mux"
)

// Templates
var navigationBarHTML string
var homepageTpl *template.Template
var adminViewTpl *template.Template

func init() {
	navigationBarHTML = assets.MustAssetString("templates/navigation_bar.html")

	homepageHTML := assets.MustAssetString("templates/index.html")
	homepageTpl = template.Must(template.New("homepage_view").Parse(homepageHTML))

	adminViewHTML := assets.MustAssetString("templates/admin_view.html")
	adminViewTpl = template.Must(template.New("admin_view").Parse(adminViewHTML))
}

var username = []byte("admin")
var password = []byte("admin")

func BasicAuth(w http.ResponseWriter, r *http.Request, user, pass []byte) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == string(user) && pair[1] == string(pass)
}

// Config provides basic configuration
type Config struct {
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// HTMLServer represents the web service that serves up HTML
type HTMLServer struct {
	server *http.Server
	wg     sync.WaitGroup
}

// Start launches the HTML Server
func Start(cfg Config) *HTMLServer {
	// Setup Context
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup Handlers
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/admin", AdminHandler)
	router.HandleFunc("/airports", AirportsHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Create the HTML Server
	htmlServer := HTMLServer{
		server: &http.Server{
			Addr:           cfg.Host,
			Handler:        router,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}

	// Add to the WaitGroup for the listener goroutine
	htmlServer.wg.Add(1)

	// Start the listener
	go func() {
		fmt.Printf("\nHTMLServer : Service started : Host=%v\n", cfg.Host)
		htmlServer.server.ListenAndServe()
		htmlServer.wg.Done()
	}()

	return &htmlServer
}

// Stop turns off the HTML Server
func (htmlServer *HTMLServer) Stop() error {
	// Create a context to attempt a graceful 5 second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fmt.Printf("\nHTMLServer : Service stopping\n")

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := htmlServer.server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down: %v", err)
		// Looks like we timed out on the graceful shutdown. Force close.
		//		if err := htmlServer.server.Close(); err != nil {
		//			fmt.Printf("\nHTMLServer : Service stopping : Error=%v\n", err)
		//			return err
		//		}
	}
	// Wait for the listener to report that it is closed.
	htmlServer.wg.Wait()
	fmt.Printf("\nHTMLServer : Stopped\n")
	return nil
}

// Render a template, or server error.
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		fmt.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}

// Push the given resource to the client.
func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}

// HomeHandler handles the home which is the flight research common user
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

// AdminHandler renders the admin view template
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
	}
	// pass from global variables
	if !BasicAuth(w, r, username, password) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Admin protected."`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		return
	}

	render(w, r, adminViewTpl, "admin_view", fullData)
}

// AirportsHandler handles the airports
func AirportsHandler(w http.ResponseWriter, r *http.Request) {
	airports := []storagemodels.Airport{
		storagemodels.Airport{
			AirportID: 507,
			City:      "London",
			Country:   "United Kingdom",
			DST:       "E",
			IATA:      "",
			ICAO:      "",
			Timezone:  "Europe/London",
			TZ:        "",
			Latitude:  51.4706,
			Longitude: -0.461941,
			Name:      "London Heathrow Airport",
		},
	}
	data, err := json.Marshal(airports)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
