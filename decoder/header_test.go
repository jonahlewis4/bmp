package decoder

import (
	"reflect"
	"testing"
)

func TestGetHeaderFromFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    *Header
		wantErr bool
	}{
		{
			name: "Get Header from rgb-triangle.bmp",
			args: args{
				fileName: "../test/bmps/original/rgb-triangle.bmp",
			},
			want: &Header{
				BITMAPFILEHEADER{
					Signature: BitmapSignature,
					FileSize: 120054,
					Reserved: BitmapReserved,
					DataSize: 120054 -
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHeaderFromFileName(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHeaderFromFileName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHeaderFromFileName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
