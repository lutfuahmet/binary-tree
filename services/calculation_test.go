package services

import (
	"binary-tree/types"
	"encoding/json"
	"testing"
)

func TestMaxPathSum(t *testing.T) {

	type args struct {
		payload string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{payload: `	{
"tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "3", "left": "6", "right": "7", "value": 3},
      {"id": "7", "left": null, "right": null, "value": 7},
      {"id": "6", "left": null, "right": null, "value": 6},
      {"id": "2", "left": "4", "right": "5", "value": 2},
      {"id": "5", "left": null, "right": null, "value": 5},
      {"id": "4", "left": null, "right": null, "value": 4}
],
"root": "1" }
}`},
			want: 18,
		},
		{
			name: "test2",
			args: args{payload: `{
"tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "3", "left": null, "right": null, "value": 3},
      {"id": "2", "left": null, "right": null, "value": 2}
],
"root": "1" }
}`},
			want: 6,
		},
		{
			name: "test3",
			want: 154,
			args: args{payload: `{
"tree": {
    "nodes": [
      {"id": "1", "left": "-10", "right": "-5", "value": 1},
      {"id": "-5", "left": "-20", "right": "-21", "value": -5},
      {"id": "-21", "left": "100-2", "right": "1-3", "value": -21},
      {"id": "1-3", "left": null, "right": null, "value": 1},
      {"id": "100-2", "left": null, "right": null, "value": 100},
      {"id": "-20", "left": "100", "right": "2", "value": -20},
      {"id": "2", "left": null, "right": null, "value": 2},
      {"id": "100", "left": null, "right": null, "value": 100},
      {"id": "-10", "left": "30", "right": "45", "value": -10},
      {"id": "45", "left": "3", "right": "-3", "value": 45},
      {"id": "-3", "left": null, "right": null, "value": -3},
      {"id": "3", "left": null, "right": null, "value": 3},
      {"id": "30", "left": "5", "right": "1-2", "value": 30},
      {"id": "1-2", "left": null, "right": null, "value": 1},
      {"id": "5", "left": null, "right": null, "value": 5}
],
"root": "1" }
}`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := types.CalculationRequest{}
			err := json.Unmarshal([]byte(tt.args.payload), &req)
			if err != nil {
				t.Fatal(err)
			}

			if got := MaxPathSum(req); got != tt.want {
				t.Errorf("MaxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
