package encoder

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	ArrType = 1 + iota
	MapType
)

var (
	ErrInterfaceCast   = errors.New("cannot cast")
	ErrUnsupportedType = errors.New("unsupported type")
)

type Template struct {
	Type   int         `json:"type"`
	Result interface{} `json:"result"`
}

func Encode(v interface{}) (encoded []byte, err error) {
	template, ok := v.(Template)
	if !ok {
		return nil, fmt.Errorf("%v %T to %T type", ErrInterfaceCast, v, &Template{})
	}

	switch template.Type {
	case ArrType:
		encoded, err = arrTemplateEncoder(template)
	case MapType:
		encoded, err = mapTemplateEncoder(template)
	default:
		return nil, ErrUnsupportedType
	}
	return
}

func arrTemplateEncoder(t Template) (encoded []byte, err error) {
	arr, ok := t.Result.([]string)
	if !ok {
		return nil, fmt.Errorf("arrTemplateEncoder: %v %T to %T type",
			ErrInterfaceCast,
			t.Result,
			[]string{})
	}
	t.Result = arr

	return json.Marshal(&t)
}

func mapTemplateEncoder(t Template) (encoded []byte, err error) {
	m, ok := t.Result.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("mapTemplateEncoder: %v %T to %T type",
			ErrInterfaceCast,
			t.Result,
			[]string{},
		)
	}
	t.Result = m

	return json.Marshal(&t)
}
