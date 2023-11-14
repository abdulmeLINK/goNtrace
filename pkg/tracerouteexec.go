package pkg

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type MTRResult struct {
	Report struct {
		Hps []hubs `json:"hubs"`
	} `json:"report"`
}

type hubs struct {
	IPAddr string  `json:"host"`
	RTT    float64 `json:"Avg"`
}

func TraceRouteWithMTR(targetIP string) ([]Hop, error) {
	cmd := exec.Command("mtr", targetIP, "--json")

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not execute mtr command: %v", err)
	}

	var result MTRResult
	fmt.Println(string(output))
	err = json.Unmarshal(output, &result)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	var hops []Hop
	for _, hub := range result.Report.Hps {
		hops = append(hops, Hop{
			IPAddr: hub.IPAddr,
			RTT:    time.Duration(hub.RTT),
		})
	}
	return hops, nil
}
