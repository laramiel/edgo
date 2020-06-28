package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime"
	"strings"

	"./assets"
	"./edgo/journal"
	"./edgo/watch"

	"github.com/zserge/lorca"
)

var (
	ErrClosed = errors.New("app: UI closed")
)

type app struct {
	events   chan interface{}
	ready    chan net.Addr
	exec     chan string
	shutdown watch.Shutdown
	navroute string
}

// registered under /api
func (app *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving: /api%s  %s", r.URL, app.navroute)

	w.Header().Set("Cache-Control", "no-cache")
	fmt.Fprintf(w, app.navroute)
}

// Lorca uses an in-process http server to serve the application
// ui resources
func (app *app) runHttpServer() {
	mux := http.NewServeMux()

	// Serve static files from the FS.
	fs := http.FileServer(assets.FS)
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.Handle("/api/", http.StripPrefix("/api", app))

	// Start HTML server
	log.Println("starting server...")
	srv := &http.Server{Handler: mux}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		app.shutdown.Kill(err)
		return
	}
	select {
	case app.ready <- ln.Addr():
		/* noop */
	case <-app.shutdown.Dying():
		ln.Close()
		return
	}

	// Spawn a shutdown function.
	go func() {
		<-app.shutdown.Dying()
		srv.Shutdown(context.TODO())
		ln.Close()
	}()

	// always returns error. ErrServerClosed on graceful close
	if err := srv.Serve(ln); err != http.ErrServerClosed {
		// unexpected error. port in use?
		app.shutdown.Kill(err)
	}
}
func (app *app) runLorca() {
	// Create UI with basic HTML passed via data URI
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 480, 320, args...)
	if err != nil {
		app.shutdown.Kill(err)
		return
	}
	defer ui.Close()

	// A simple way to know when UI is ready (uses body.onload event in JS)
	ui.Bind("lorcaStarted", func() {
		log.Println("UI is ready")
	})


	// Wait until the http server is running, then try and load
	// the main page.
	select {
	case addr := <-app.ready:
		ui.Load(fmt.Sprintf("http://%s/static/index.html", addr))
		close(app.ready)

		/* noop */
	case <-app.shutdown.Dying():
		return
	}

	// Start our exec command goroutine.
	go func() {
		for {
			select {
			case exec := <-app.exec:
				log.Println("app: ", exec)
				ui.Eval(exec)

			case <-app.shutdown.Dying():
				return
			}
		}
	}()
	
	// Run our main application processing loop.
	for {
		select {
		case e := <-app.events:
			app.processEvent(e)

		case <-ui.Done():
			app.shutdown.Kill(ErrClosed)
			return

		case <-app.shutdown.Dying():
			return
		}
	}
}

func (app *app) processEvent(e interface{}) {
	switch v := e.(type) {
	case *journal.NavRoute:
		app.setNavRoute(v)
	default:
	}
}

func (app *app) setNavRoute(v *journal.NavRoute) {
	if len(v.Route) == 0 {
		app.navroute = "{}"
		return
	}

	systems := []string{}
	routes := []string{}

	// not exactly sure which coordinates I should use here.
	for _, v := range v.Route {
		systems = append(systems,
			fmt.Sprintf(`{ "name": "%s", "coords": { "x": %f, "y": %f, "z": %f }}`,
				v.StarSystem, v.StarPos[0], v.StarPos[1], v.StarPos[2]))
		routes = append(routes,
			fmt.Sprintf(`{"s":"%s"}`, v.StarSystem))
	}

	navroute := fmt.Sprintf(`
	{
		"systems": [
			%s
		],
		"routes": [{
			"title":"NavRoute",
			"points": [
				%s
			]
		}]
	}`,
		strings.Join(systems, ",\n			"),
		strings.Join(routes, ",\n				"))
	app.navroute = navroute

	log.Println("Updating navroute: ", navroute)
	// Maybe trigger a reload
	select {
	case app.exec <- `Ed3d.rebuild();`:
		/**/
	default:
	}
}

func AppMain(events chan interface{}, shutdown watch.Shutdown) {
	a := &app{
		events:   events,
		ready:    make(chan net.Addr, 1),
		exec:     make(chan string, 1),
		shutdown: shutdown,
		navroute: "{}",
	}
	go a.runHttpServer()
	a.runLorca()
}
