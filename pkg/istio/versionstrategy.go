package istio

import v1 "maistra.io/istio-operator/api/v1alpha1"

type VersionStrategy interface {
	ApplyDefaults(ihi *v1.Istio) error
}