package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"./edgo"
	"./edgo/watch"
)

type filterFlag []string

func (i *filterFlag) String() string {
	return fmt.Sprint([]string(*i))	
}

func (i *filterFlag) Set(value string) error {
    *i = append(*i, value)
    return nil
}

var (
	ErrInterrupted = errors.New("main: interrupted")
	filters filterFlag
)

func waitForInterrupt(shutdown watch.Shutdown) {
	sigc := make(chan os.Signal)
	defer close(sigc)

	signal.Notify(sigc, os.Interrupt)
	defer signal.Stop(sigc)

	select {
	case <-sigc:
		shutdown.Kill(ErrInterrupted)
	case <-shutdown.Dying():
		/*noop*/
	}
}

func main() {
	flag.Var(&filters, "f", "Filtered events.")
	flag.Parse()

	var directory string
	if flag.NFlag() > 1 {
		directory = flag.Arg(1)
	} else if homedir, err := os.UserHomeDir(); err == nil {
		directory = filepath.Join(homedir, "Saved Games", "Frontier Developments", "Elite Dangerous")
	}
	if directory == "" {
		fmt.Printf("Usage: %s <path to elite dangerous journal>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	fmt.Printf("Using: %s %s\n", filepath.Base(os.Args[0]), directory)

	shutdown := watch.NewShutdown()
	w := edgolang.NewEliteWatcher(directory, shutdown)
	defer w.Close()

	if len(filters) > 0 {
		// Only add event filters if they have been specified on the comand line.
		w.EventFilter = make(map[string]struct{})
		for _, v := range filters {
			w.EventFilter[v] = struct{}{}	
			log.Println("main: filter ", v)
		}
		w.EventFilter["NavRoute"] = struct{}{}	
		log.Println("main: filter NavRoute")
	}

	go w.Main()
	go AppMain(w.Journals, shutdown)

	waitForInterrupt(shutdown)
	log.Println("done...")
}

/*
go get github.com/zserge/lorca
go get gopkg.in/tomb.v2
go get github.com/fsnotify/fsnotify

go generate
go run gen.go
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui"
*/
