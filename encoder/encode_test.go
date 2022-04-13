package encoder

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		encType int
		data    []string
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
				encType: 1,
				data:    []string{"res1", "res2", "res3"},
			},
			wantEncoded: []byte("{\"type\":1,\"result\":[\"res1\",\"res2\",\"res3\"]}"),
			wantErr:     false,
		},
		{
			name: "mapResultPositiveTest",
			args: args{
				encType: 2,
				data:    []string{"res1", "res2", "res3"},
			},
			wantEncoded: []byte("{\"type\":2,\"result\":{\"0\":\"res1\",\"1\":\"res2\",\"2\":\"res3\"}}"),
			wantErr:     false,
		},
		{
			name: "unsupportedResultNegativeTest",
			args: args{
				encType: 3,
				data:    []string{"res1", "res2", "res3"},
			},
			wantEncoded: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncoded, err := Encode(tt.args.encType, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEncoded, tt.wantEncoded) {
				t.Errorf("Encode() gotEncoded = %v, want %v", gotEncoded, tt.wantEncoded)
			}
		})
	}
}
