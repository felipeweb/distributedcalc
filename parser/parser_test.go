package parser

import (
	"context"
	"testing"
)

func Test_replaceVars(t *testing.T) {
	type args struct {
		input Input
	}
	tests := []struct {
		name    string
		args    args
		wantStr string
		wantErr bool
	}{
		{
			"replace variable with success",
			args{
				Input{
					Expression: "X+(2Y+(X/Y))",
					Variables: map[string]float64{
						"X": 8,
						"Y": 4,
					},
				},
			},
			"8+(2*4+(8/4))",
			false,
		},
		{
			"replace variable after )",
			args{
				Input{
					Expression: "(2Y+(X/Y))X",
					Variables: map[string]float64{
						"X": 8,
						"Y": 4,
					},
				},
			},
			"(2*4+(8/4))*8",
			false,
		},
		{
			"undeclared variable error",
			args{
				Input{
					Expression: "(2Y+(X/Y))3",
					Variables: map[string]float64{
						"X": 8,
						"Y": 4,
					},
				},
			},
			"(2*4+(8/4))*3",
			false,
		},
		{
			"replace variable after )",
			args{
				Input{
					Expression: "(2Y+(X/Y))X",
					Variables: map[string]float64{
						"Y": 4,
					},
				},
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := replaceVars(context.Background(), tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("replaceVars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStr != tt.wantStr {
				t.Errorf("replaceVars() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}
