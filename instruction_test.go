package solana

import (
	"github.com/BRBussy/solgo/internal/pkg/encoding"
	"reflect"
	"testing"
)

func TestInstructions_Compact(t *testing.T) {
	tests := []struct {
		name    string
		i       Instructions
		want    encoding.CompactArray
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.Compact()
			if (err != nil) != tt.wantErr {
				t.Errorf("Compact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compact() got = %v, want %v", got, tt.want)
			}
		})
	}
}
