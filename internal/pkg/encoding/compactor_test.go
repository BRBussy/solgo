package encoding

import (
	"reflect"
	"testing"
)

func TestCompactArray_ToBytes(t *testing.T) {
	type fields struct {
		Length uint64
		Data   []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "success - length 127",
			fields: fields{
				Length: 127,
				Data:   []byte{0x01, 0x02},
			},
			want: []byte{
				0b01111111, 0x00, 0x00,
				0x01, 0x02,
			},
		},
		{
			name: "success - length 300",
			fields: fields{
				Length: 127,
				Data:   []byte{0x01, 0x02},
			},
			want: []byte{
				0x7f, 0x00, 0x00,
				0x01, 0x02,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CompactArray{
				Length: tt.fields.Length,
				Data:   tt.fields.Data,
			}
			if got := c.ToBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
