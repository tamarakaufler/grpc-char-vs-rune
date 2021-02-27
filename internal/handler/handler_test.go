package handler_test

import (
	"reflect"
	"testing"

	"github.com/tamarakaufler/grpc-char-vs-rune/internal/handler"
)

func TestConvertToRune(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantR []uint32
		wantM map[string]uint32
	}{
		{
			name: "one letter conversion",
			args: args{
				s: "a",
			},
			wantR: []uint32{97},
			wantM: map[string]uint32{
				"a": 97,
			},
		},
		{
			name: "string conversion",
			args: args{
				s: "acb",
			},
			wantR: []uint32{97, 99, 98},
			wantM: map[string]uint32{
				"a": 97,
				"c": 99,
				"b": 98,
			},
		},
		{
			name: "string conversion",
			args: args{
				s: "世界 日本語",
			},
			wantR: []uint32{19990, 30028, 32, 26085, 26412, 35486},
			wantM: map[string]uint32{
				" ": 32,
				"世": 19990,
				"日": 26085,
				"本": 26412,
				"界": 30028,
				"語": 35486,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := handler.ConvertToRune(tt.args.s)
			if !reflect.DeepEqual(got, tt.wantR) {
				t.Errorf("ConvertToRune() got = %v, want %v", got, tt.wantR)
			}
			if !reflect.DeepEqual(got1, tt.wantM) {
				t.Errorf("ConvertToRune() got1 = %v, want %v", got1, tt.wantM)
			}
		})
	}
}
