package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/abdulmeLINK/goNtrace/api"
	"github.com/abdulmeLINK/goNtrace/pkg"
	"github.com/gorilla/mux"
)

func ensureDirExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	servePtr := flag.Bool("serve", false, "start the server")
	mapPtr := flag.String("map", "", "generate a map image")
	portPtr := flag.String("port", "8080", "port to serve on")
	flag.Parse()

	// ensure output directory exists
	ensureDirExists("./output")

	if *servePtr {
		r := mux.NewRouter()
		r.HandleFunc("/", api.HomeHandler)
		fs := http.FileServer(http.Dir("resources/public"))
		r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
		http.Handle("/public", fs)
		r.HandleFunc("/trace", api.TracerouteHandler).Methods("GET")
		r.HandleFunc("/traceWithMTR", api.TracerouteWithMTRHandler).Methods("GET")
		r.HandleFunc("/map", api.MapHandler).Methods("GET")
		fmt.Printf("Serving on port %s...\n", *portPtr)
		http.ListenAndServe(":"+*portPtr, r)
	} else if *mapPtr != "" {
		hops, err := pkg.TraceRouteWithMTR(*mapPtr)
		if err != nil {
			fmt.Printf("Could not trace route: %v\n", err)
			return
		}

		var locations []pkg.GeoLocation
		for _, hop := range hops {
			location, err := pkg.GetGeoLocation(hop.IPAddr)
			if err != nil {
				fmt.Printf("Could not get geolocation: %v\n", err)
				return
			}
			locations = append(locations, location)
		}

		err = pkg.GenerateMapImage(locations)
		if err != nil {
			fmt.Printf("Could not generate map image: %v\n", err)
			return
		}

		fmt.Println("Map image generated successfully.")
	} else {
		fmt.Println("Please provide a command: --serve or --map")
	}
}
