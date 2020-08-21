package main

import (
	"log"
	"net/http"
	"os"

	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"

	"github.com/pierdipi/duckeye/pkg/clusterducktype"

	discoveryclusterducktype "knative.dev/discovery/pkg/client/injection/informers/discovery/v1alpha1/clusterducktype"
)

func main() {

	ctx := signals.NewContext()
	config := sharedmain.ParseAndGetConfigOrDie()
	ctx, informers := injection.Default.SetupInformers(ctx, config)

	if err := controller.StartInformers(ctx.Done(), informers...); err != nil {
		log.Fatal("Failed to start informers", err)
	}

	clusterDuckTypeHandler := clusterducktype.New(discoveryclusterducktype.Get(ctx).Lister())

	s := http.NewServeMux()

	// Serve website
	s.Handle("/", http.FileServer(http.Dir("/var/run/ko/")))

	// Serve ClusterDuckType
	s.Handle("/clusterducktypes", clusterDuckTypeHandler)
	s.Handle("/clusterducktypes/", clusterDuckTypeHandler)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), s); err != nil {
		log.Fatal(err)
	}
}
