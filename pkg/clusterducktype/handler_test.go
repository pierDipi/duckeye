package clusterducktype

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	discovery "knative.dev/discovery/pkg/apis/discovery/v1alpha1"
	. "knative.dev/pkg/reconciler/testing"

	discoveryclusterducktype "knative.dev/discovery/pkg/client/injection/informers/discovery/v1alpha1/clusterducktype/fake"
)

func TestServeHTTP(t *testing.T) {

	ctx, _ := SetupFakeContext(t)

	cdt := &discovery.ClusterDuckType{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterDuckType",
			APIVersion: discovery.SchemeGroupVersion.Version,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "duck",
			Namespace: "duck",
		},
		Spec: discovery.ClusterDuckTypeSpec{
			Group: "discovery.knative.dev",
			Names: discovery.DuckTypeNames{
				Name:     "duck",
				Plural:   "ducks",
				Singular: "duck",
			},
		},
		Status: discovery.ClusterDuckTypeStatus{
			Ducks: map[string][]discovery.ResourceMeta{
				"v1": {
					{
						APIVersion: "v1",
						Kind:       "addressable",
						Scope:      "Namespaced",
					},
				},
				"v2": {
					{
						APIVersion: "v1",
						Kind:       "podspecable",
						Scope:      "Namespaced",
					},
				},
			},
			DuckCount: 1,
		},
	}

	cdti := discoveryclusterducktype.Get(ctx)
	if err := cdti.Informer().GetStore().Add(cdt); err != nil {
		t.Fatal("Cannot add object", err)
	}

	h := New(cdti.Lister())

	r := httptest.NewRecorder()

	h.ServeHTTP(r, nil)

	b, err := ioutil.ReadAll(r.Result().Body)
	if err != nil {
		t.Fatal("failed to read response body", err)
	}
	_ = r.Result().Body.Close()

	expectedDucks := DucksList{
		{
			Name: "ducks",
			Status: map[string][]discovery.ResourceMeta{
				"v1": {
					{
						APIVersion: "v1",
						Kind:       "addressable",
						Scope:      "Namespaced",
					},
				},
				"v2": {
					{
						APIVersion: "v1",
						Kind:       "podspecable",
						Scope:      "Namespaced",
					},
				},
			},
		},
	}

	expected, err := json.Marshal(&expectedDucks)
	if err != nil {
		t.Fatal("Failed to marshal expected ducks", err)
	}

	if diff := cmp.Diff(expected, b); diff != "" {
		t.Fatal("Want, got (-, +)", diff)
	}
}
