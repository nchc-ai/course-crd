package v1alpha1

type AccessType string

const (
	AccessTypeNodePort AccessType = "NodePort"
	AccessTypeIngress  AccessType = "Ingress"
)
