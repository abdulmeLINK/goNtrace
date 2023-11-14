package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/abdulmeLINK/goNtrace/pkg"
)

func traceHandler(hops []pkg.Hop, w http.ResponseWriter, r *http.Request) {

	var geoLocations []pkg.GeoLocation
	for _, hop := range hops {
		geoLocation, err := pkg.GetGeoLocation(hop.IPAddr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		geoLocations = append(geoLocations, geoLocation)
	}

	json.NewEncoder(w).Encode(geoLocations)
}
func TracerouteHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "Missing IP address", http.StatusBadRequest)
		return
	}
	hops, err := pkg.TraceRoute(ip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	traceHandler(hops, w, r)

}
func TracerouteWithMTRHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "Missing IP address", http.StatusBadRequest)
		return
	}
	hops, err := pkg.TraceRouteWithMTR(ip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	traceHandler(hops, w, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("resources/home.html")
	if err != nil {
		http.Error(w, "Could not open requested file", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/html")
	w.Write(content)
}

func MapHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "Missing IP address", http.StatusBadRequest)
		return
	}

	hops, err := pkg.TraceRouteWithMTR(ip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var locations []pkg.GeoLocation
	for _, hop := range hops {
		location, err := pkg.GetGeoLocation(hop.IPAddr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		locations = append(locations, location)
	}

	err = pkg.GenerateMapImage(locations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, "./output/map.png")
}
