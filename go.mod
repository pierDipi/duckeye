module github.com/pierdipi/duckeye

go 1.14

require (
	k8s.io/apimachinery v0.18.12
	knative.dev/discovery v0.19.1-0.20201118215652-dd754f388592
	knative.dev/hack v0.0.0-20201120192952-353db687ec5b
	knative.dev/pkg v0.0.0-20201120183152-6a0e731e251a
)

replace (
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
)
