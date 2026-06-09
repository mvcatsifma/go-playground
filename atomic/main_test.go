package atomic

import "testing"

func Test_concurrentIncrement(t *testing.T) {
	type args struct {
		nr int
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				nr: 100,
			},
			want: 100,
		},
		{
			args: args{
				nr: 1000,
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concurrentIncrement(tt.args.nr); got != tt.want {
				t.Errorf("concurrentIncrement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runCleanExit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runCleanExit()
		})
	}
}

func Test_updateAndReadConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateAndReadConfig()
		})
	}
}
