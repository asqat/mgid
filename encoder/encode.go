package encoder

import (
	"encoding/json"
	"errors"
	"strconv"
)

const (
	ArrType = 1 + iota
	MapType
)

var (
	ErrUnsupportedType = errors.New("unsupported type")
)

type Template struct {
	Type   int         `json:"type"`
	Result interface{} `json:"result"`
}

func Encode(encType int, data []string) (encoded []byte, err error) {
	t := &Template{}
	switch encType {
	case ArrType:
		t = arrTemplateEncoder(data)
	case MapType:
		t = mapTemplateEncoder(data)
	default:
		return nil, ErrUnsupportedType
	}
	return json.Marshal(t)
}

func arrTemplateEncoder(data []string) *Template {
	return &Template{
		Type:   ArrType,
		Result: data,
	}
}

func mapTemplateEncoder(data []string) *Template {
	m := make(map[string]string, len(data))
	for i, item := range data {
		m[strconv.Itoa(i)] = item
	}
	return &Template{
		Type:   MapType,
		Result: m,
	}
}
