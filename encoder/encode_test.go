package encoder

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name        string
		args        args
		wantEncoded []byte
		wantErr     bool
	}{
		{
			name: "arrayResultPositiveTest",
			args: args{
				v: Template{
					Type:   1,
					Result: []string{"res1", "res2", "res3"},
				},
			},
			wantEncoded: []byte("{\"type\":1,\"result\":[\"res1\",\"res2\",\"res3\"]}"),
			wantErr:     false,
		},
		{
			name: "mapResultPositiveTest",
			args: args{
				v: Template{
					Type:   2,
					Result: map[string]string{"0": "res1", "1": "res2", "2": "res3"},
				},
			},
			wantEncoded: []byte("{\"type\":2,\"result\":{\"0\":\"res1\",\"1\":\"res2\",\"2\":\"res3\"}}"),
			wantErr:     false,
		},
		{
			name: "unsupportedResultNegativeTest01",
			args: args{
				v: Template{
					Type:   3,
					Result: []string{"res1", "res2", "res3"},
				},
			},
			wantEncoded: nil,
			wantErr:     true,
		},
		{
			name: "unsupportedResultNegativeTest02",
			args: args{
				v: Template{
					Type:   1,
					Result: "unsupported_type",
				},
			},
			wantEncoded: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncoded, err := Encode(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEncoded, tt.wantEncoded) {
				t.Errorf("Encode() gotEncoded = %v, want %v", string(gotEncoded), string(tt.wantEncoded))
			}
		})
	}
}
