package decoder

import (
	"reflect"
	"testing"
)

func Test_getHeader(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     *Header
	}{
		{
			name:     "rgb-triangle",
			filePath: "../test/bmps/original/rgb-traingle.bmp",
			want: &Header{
				height: 200,
				width:  200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHeader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHeader() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
