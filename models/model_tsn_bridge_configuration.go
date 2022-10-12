package models

type TsnBridgeConfiguration struct {
	StreamPriority int `json:"streamPriority"`
	IngressPort    int `json:"ingressPort"`
}
