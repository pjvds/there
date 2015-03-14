package main

import (
	"net/http"

	"github.com/pjvds/tidy"
)

var (
	locations map[string]string
)

func main() {
	log := tidy.GetLogger()

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		host := request.URL.Host
		location, ok := locations[host]

		if !ok {
			location = "http://github.com/pjvds/there/#README"
		}

		http.Redirect(response, request, location, http.StatusFound)

		log.WithFields(tidy.Fields{
			"host":     request.Host,
			"location": location,
			"found":    ok,
		}).Debug("request handled")
	})

	address := ":8080"
	log.WithField("address", address).Debug("listening")

	if err := http.ListenAndServe(address, nil); err != nil {
		log.WithField("error", err).Error("listening failed")
	}
}
