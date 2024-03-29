package shoot

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// {
		// 	name: "regular JPG",
		// 	args: args{
		// 		filename: "data/image1.JPG",
		// 	},
		// 	want:    "2022-08-17T12:01:26+02:00",
		// 	wantErr: false,
		// },
		{
			name: "cr3",
			args: args{
				filename: "data/image2.CR3",
			},
			want:    "2022-08-17T12:01:26+02:00",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Date(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Date() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			d := got.Format(time.RFC3339)
			if d != tt.want {
				t.Errorf("Date() = %v, want %v", d, tt.want)
			}
		})
	}
}
