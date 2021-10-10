package encoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToCompactU16(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToCompactU16()
			assert.Equalf(t, tt.wantErr, err != nil, "error response not as expected")
			assert.Equalf(t, tt.want, got, "result not as expected")
		})
	}
}
