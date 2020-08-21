package clusterducktype

import (
	"encoding/json"
	"net/http"

	"k8s.io/apimachinery/pkg/labels"

	discovery "knative.dev/discovery/pkg/apis/discovery/v1alpha1"
	listers "knative.dev/discovery/pkg/client/listers/discovery/v1alpha1"
)

type Handler struct {
	lister listers.ClusterDuckTypeLister
}

func New(lister listers.ClusterDuckTypeLister) *Handler {
	return &Handler{
		lister: lister,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ducks, err := h.lister.List(labels.Everything())
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	var ducksList DucksList
	for _, d := range ducks {
		ducksList = append(ducksList, NewDucks(*d))
	}

	b, err := json.Marshal(ducksList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(b)
}

type DucksList []*Ducks

type Ducks struct {
	Name   string                              `json:"name"`
	Status map[string][]discovery.ResourceMeta `json:"status"`
}

func NewDucks(duckType discovery.ClusterDuckType) *Ducks {
	return &Ducks{
		Name:   duckType.Spec.Names.Plural,
		Status: duckType.Status.Ducks,
	}
}
