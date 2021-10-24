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
				Length: 0x7f,
				Data:   []byte{0x01, 0x02},
			},
			want: []byte{
				0x7f, 0x00, 0x00,
				0x01, 0x02,
			},
		},
		{
			name: "success - length 254",
			fields: fields{
				Length: 0b11111110,
				Data:   []byte{0x01, 0x02},
			},
			want: []byte{
				0b11111110, 0x01, 0x00,
				0x01, 0x02,
			},
		},
		{
			name: "success - length 255",
			fields: fields{
				Length: 0b11111111,
				Data:   []byte{0x01, 0x02},
			},
			want: []byte{
				0b11111111, 0x01, 0x00,
				0x01, 0x02,
			},
		},
		{
			name: "success - length 256",
			fields: fields{
				Length: 256,
				Data:   []byte{0x01, 0x02},
			},
			want: []byte{
				0b10000000, 0x02, 0x00,
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
