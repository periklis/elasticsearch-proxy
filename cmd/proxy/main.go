package main

import (
	"os"

	"github.com/openshift/elasticsearch-proxy/pkg/config"

	"github.com/openshift/elasticsearch-proxy/pkg/proxy"
	log "github.com/sirupsen/logrus"
)

func main() {
	// initLogging()
	// log.SetLevel(log.InfoLevel)

	opts, err := config.Init(os.Args[1:])
	if err != nil {
		log.Errorf("%s", err)
		os.Exit(1)
	}

	// proxyServer := proxy.NewProxyServer(opts)

	// log.Debugf("Registering Handlers....")
	// proxyServer.RegisterRequestHandlers(auth.NewHandlers(opts))

	// var h http.Handler = proxyServer
	// if opts.RequestLogging {
	//	h = logging.NewHandler(os.Stdout, h, true)
	// }

	h := &proxy.TinkerHandler{}

	s := &proxy.Server{
		Handler: h,
		Opts:    opts,
	}
	go s.ListenAndServe()

	// if opts.MetricsListeningAddress != "" {
	//	m := proxy.MetricsServer{
	//		Handler: h,
	//		Opts:    opts,
	//	}
	//	go m.ListenAndServe()
	// }
	select {}
}

// func initLogging() {
//	logLevel := os.Getenv("LOG_LEVEL")
//	if strings.TrimSpace(logLevel) == "" {
//		logLevel = "info"
//	}
//	level, err := log.ParseLevel(logLevel)
//	if err != nil {
//		level = log.WarnLevel
//		log.Infof("Setting loglevel to 'warn' as unable to parse %s", logLevel)
//	}
//	log.SetLevel(level)
// }
