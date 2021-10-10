package encoding

import (
	"reflect"
	"testing"
)

func TestIntToCompactU16(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToCompactU16()
			if (err != nil) != tt.wantErr {
				t.Errorf("IntToCompactU16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToCompactU16() got = %v, want %v", got, tt.want)
			}
		})
	}
}
