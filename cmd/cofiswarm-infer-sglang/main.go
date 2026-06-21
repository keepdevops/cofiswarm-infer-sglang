package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/keepdevops/cofiswarm-infer-sglang/internal/bus"
	"github.com/keepdevops/cofiswarm-observer-sdk/pkg/servicecomponent"
)

func main() {
	addr := flag.String("listen", ":8091", "health/metadata port (HTTP mode)")
	busMode := flag.Bool("bus", false, "announce + serve .infer.sglang.* on the NATS observer bus instead of HTTP")
	natsURL := flag.String("nats", "nats://127.0.0.1:4222", "NATS URL (bus mode)")
	flag.Parse()

	if *busMode {
		serveBus(*natsURL)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/v1/info", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"engine":"sglang","stub":true,"note":"run deploy/Dockerfile for full sglang-metal"}`))
	})
	log.Printf("infer-sglang metadata on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}

// serveBus announces infer-sglang on the observer bus and serves its .infer.sglang.* capability
// subjects until SIGINT/SIGTERM, when it says goodbye so presence flips offline cleanly.
func serveBus(url string) {
	nc, err := servicecomponent.Connect(url, "cofiswarm-infer-sglang")
	if err != nil {
		log.Fatalf("bus connect %s: %v", url, err)
	}
	defer nc.Close()
	comp := servicecomponent.New(nc, "infer-sglang", "infer-sglang", bus.Routes("sglang"))
	if err := comp.Start(); err != nil {
		log.Fatalf("bus start: %v", err)
	}
	defer comp.Shutdown()
	log.Printf("infer-sglang on bus %s (.infer.sglang.info/.health)", url)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Print("infer-sglang bus stopping")
}
