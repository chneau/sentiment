package sentiment

import (
	"testing"

	_ "github.com/chneau/sentiment/pkg/statik"
)

func TestEvaluate(t *testing.T) {
	type args struct {
		wwww []string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "happy",
			args: args{
				wwww: []string{"happy"},
			},
			want: 1,
		},
		{
			name: "bad",
			args: args{
				wwww: []string{"bad"},
			},
			want: -1,
		},
		{
			name: "neutral",
			args: args{
				wwww: []string{"neutral"},
			},
			want: 0,
		},
		{
			name: "happy#bad",
			args: args{
				wwww: []string{"happy#bad"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Evaluate(tt.args.wwww...); got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
