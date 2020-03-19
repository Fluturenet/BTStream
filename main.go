package main

import (
	"flag"
	_ "io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	_ "github.com/fluturenet/ed25519"
	log "github.com/sirupsen/logrus"
)

var (
	dataDir    string
	httpAddr   string
	httpServer http.Server
)

func main() {
	home, _ := os.UserHomeDir()
	defaultPath := filepath.Join(home, "BTStream")
	flag.StringVar(&dataDir, "datadir", defaultPath, "Path where store key and temporary data")
	flag.StringVar(&dataDir, "d", defaultPath, "Path where store key and temporary data")

	flag.StringVar(&httpAddr, "haddr", ":8080", "HTTP listening Addr")

	genkey := flag.Bool("genkey", false, "generates key identity. ex -genkey")
	help := flag.Bool("help", false, "prints help")
	flag.Parse()

	if *genkey {
		generateKey()
		os.Exit(0)
	}
	if *help {
		gentlydie()
		os.Exit(0)
	}

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		log.Infoln("Received an interrupt, stopping services...")
		cleanup()
		close(cleanupDone)
	}()
	log.Info("dataDir set to:", dataDir)
	log.Info("Starting Torrent Client")
	init_torrent()
	log.Info("Starting http server on: ", httpAddr)
	startHttp()
	log.Info("Http server: closed")

	<-cleanupDone
}

func gentlydie() {
	flag.Usage()
	os.Exit(0)
}

func cleanup() {
	httpServer.Close()
}

func startHttp() {
	httpServer.Addr = httpAddr
	httpServer.Handler = logHandler(http.DefaultServeMux)
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/", landingpage)
	http.HandleFunc("/tr", tr_handler)
	go log.Fatal(httpServer.ListenAndServe())
}
