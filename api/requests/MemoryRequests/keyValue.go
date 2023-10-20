package MemoryRequests

type KeyValue struct {
	Key   string `json:"key" example:"ben" validate:"required"`
	Value string `json:"value" example:"samet" validate:"required"`
}
