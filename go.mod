module github.com/pierdipi/duckeye

go 1.14

require (
	k8s.io/apimachinery v0.18.7-rc.0
	knative.dev/discovery v0.0.0-20200814171806-3dc3c066826b
	knative.dev/pkg v0.0.0-20200813155605-c9f9284521f1
)

replace (
	k8s.io/api => k8s.io/api v0.17.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
)
