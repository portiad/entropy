package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

type Handler func(c *context, w http.ResponseWriter, r *http.Request)
type context struct{}

func ping(c *context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{'O', 'K'})
}

func version(c *context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(VERSION))
}

func info(c *context, w http.ResponseWriter, r *http.Request) {
	i := struct {
		Version string
	}{
		Version: VERSION,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(i); err != nil {
		panic(err)
	}
}

func list(c *context, w http.ResponseWriter, r *http.Request) {
	// TODO: connect to the Docker endpoint (injected via context)
	//       use ps and --filter for the entropy agent label
	//       Foreach result, extract an injector - use append
	//       FromContainerInfo(r)

	l := []injector{
		{
			Name:        "hopping_kenedy",
			Frequency:   "30",
			Probability: "10",
			Image:       "qualmente/gremlins",
			Faults: []string{
				"pause",
				"partition",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(l); err != nil {
		panic(err)
	}
}

func handlerNotYetImplemented(c *context, w http.ResponseWriter, r *http.Request) {
	httpError(w, "Not implemented yet.", http.StatusNotImplemented)
}

func notImplementedHandler(c *context, w http.ResponseWriter, r *http.Request) {
	httpError(w, "Not supported in clustering mode.", http.StatusNotImplemented)
}

func httpError(w http.ResponseWriter, err string, status int) {
	log.WithField("status", status).Errorf("HTTP error: %v", err)
	http.Error(w, err, status)
}
