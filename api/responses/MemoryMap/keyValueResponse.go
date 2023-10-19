package MemoryMap

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func CreateKeyValueResponse(key, value string) KeyValue {
	return KeyValue{Key: key, Value: value}
}
