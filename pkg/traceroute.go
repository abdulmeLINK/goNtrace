package pkg

import (
	"fmt"
	"net"
	"time"

	"github.com/abdulmeLINK/mtr/pkg/mtr"
	pj "github.com/hokaccha/go-prettyjson"
)

type Hop struct {
	IPAddr string
	RTT    time.Duration
}

var (
	version string
	date    string

	COUNT            = 5
	TIMEOUT          = 800 * time.Millisecond
	INTERVAL         = 100 * time.Millisecond
	HOP_SLEEP        = time.Nanosecond
	MAX_HOPS         = 64
	MAX_UNKNOWN_HOPS = 10
	RING_BUFFER_SIZE = 50
	PTR_LOOKUP       = false
	jsonFmt          = false
	srcAddr          = ""
	versionFlag      bool
)

func TraceRoute(dest string) ([]Hop, error) {
	var hops []Hop
	fmt.Println("IP: " + dest)
	host, err := net.ResolveIPAddr("ip", dest)
	if err != nil {
		return nil, fmt.Errorf("could not resolve IP: %v", err)
	}

	m, ch, err := mtr.NewMTR(host.String(), srcAddr, TIMEOUT, INTERVAL, HOP_SLEEP,
		MAX_HOPS, MAX_UNKNOWN_HOPS, RING_BUFFER_SIZE, PTR_LOOKUP)
	if err != nil {
		return nil, fmt.Errorf("could not run traceroute: %v", err)
	}

	go func(ch chan struct{}) {
		for {
			<-ch
		}
	}(ch)
	m.Run(ch, COUNT)
	s, err := pj.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("could not run traceroute: %v", err)
	}
	fmt.Println(string(s))
	for _, v := range m.Statistic {
		byt, _ := v.MarshalJSON()
		fmt.Println(string(byt))
		hops = append(hops, Hop{
			IPAddr: v.Targets[0],
			RTT:    v.SumElapsed,
		})
	}

	if err != nil {
		return nil, fmt.Errorf("could not run traceroute: %v", err)
	}

	return hops, nil
}
