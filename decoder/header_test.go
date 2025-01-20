package decoder

import (
	"github.com/jonahlewis4/bmp/bmp/headers"
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
		want    *headers.Header
		wantErr bool
	}{
		{
			name: "Get Header from rgb-triangle.bmp",
			args: args{
				fileName: "../test/bmps/original/rgb-triangle.bmp",
			},
			want: &headers.Header{
				BITMAPFILEHEADER: &headers.BITMAPFILEHEADER{
					Signature: headers.BitmapSignature,
					FileSize:  120054,
					Reserved:  headers.BitmapReserved,
					DataSize:  headers.ExpectedInfoHeaderSize + headers.FileHeaderSize,
				},
				InfoHeader: headers.InfoHeader(&headers.BITMAPINFOHEADER{
					Size:               headers.ExpectedInfoHeaderSize,
					Width:              200,
					Height:             200,
					Planes:             1,
					BitsPerPixel:       24,
					Compression:        headers.BI_RGB,
					ImageSize:          0,
					HorizontalRes:      0,
					VerticalRes:        0,
					NumColors:          0,
					NumImportantColors: 0,
				}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHeaderFromFileName(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHeaderFromFileName() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHeaderFromFileName() got = \n%+v\nwant\n%+v", got, tt.want)
			}
		})
	}
}
