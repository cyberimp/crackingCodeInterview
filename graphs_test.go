package main

import (
	"reflect"
	"testing"
)

func Test_isPassable(t *testing.T) {
	type args struct {
		edges [][]int
		s     int
		t     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"small cycled graph",
			args{
				[][]int{{0, 1}, {1, 2}, {2, 0}},
				0,
				2,
			},
			true,
		},
		{
			"disconnected graph",
			args{
				[][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
				0,
				5,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPassable(tt.args.edges, tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isPassable() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_makeBST(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"simple test",
			args{[]int{0, 1, 2}},
			`{"value":1,"left":{"value":0},"right":{"value":2}}`,
		},
		{
			"even test",
			args{[]int{0, 1, 2, 3}},
			`{"value":2,"left":{"value":1,"left":{"value":0}},"right":{"value":3}}`,
		},
		{
			"big one",
			args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 16, 17, 18, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 36, 38, 39, 40, 41, 42, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 56, 57, 59, 60, 61, 62, 63, 65, 66, 67, 68, 69, 70, 71, 72, 75, 76, 77, 80, 81, 82, 83, 84, 86, 87, 90, 92, 93, 94, 95, 96, 97, 98}},
			`{"value":47,"left":{"value":24,"left":{"value":11,"left":{"value":5,"left":{"value":2,"left":{"value":1,"left":{"value":0}},"right":{"value":4,"left":{"value":3}}},"right":{"value":9,"left":{"value":7,"left":{"value":6}},"right":{"value":10}}},"right":{"value":18,"left":{"value":16,"left":{"value":13,"left":{"value":12}},"right":{"value":17}},"right":{"value":22,"left":{"value":21,"left":{"value":20}},"right":{"value":23}}}},"right":{"value":34,"left":{"value":29,"left":{"value":27,"left":{"value":26,"left":{"value":25}},"right":{"value":28}},"right":{"value":32,"left":{"value":31,"left":{"value":30}},"right":{"value":33}}},"right":{"value":41,"left":{"value":39,"left":{"value":38,"left":{"value":36}},"right":{"value":40}},"right":{"value":45,"left":{"value":44,"left":{"value":42}},"right":{"value":46}}}}},"right":{"value":71,"left":{"value":60,"left":{"value":52,"left":{"value":50,"left":{"value":49,"left":{"value":48}},"right":{"value":51}},"right":{"value":57,"left":{"value":56,"left":{"value":53}},"right":{"value":59}}},"right":{"value":66,"left":{"value":63,"left":{"value":62,"left":{"value":61}},"right":{"value":65}},"right":{"value":69,"left":{"value":68,"left":{"value":67}},"right":{"value":70}}}},"right":{"value":86,"left":{"value":80,"left":{"value":76,"left":{"value":75,"left":{"value":72}},"right":{"value":77}},"right":{"value":83,"left":{"value":82,"left":{"value":81}},"right":{"value":84}}},"right":{"value":94,"left":{"value":92,"left":{"value":90,"left":{"value":87}},"right":{"value":93}},"right":{"value":97,"left":{"value":96,"left":{"value":95}},"right":{"value":98}}}}}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeBST(tt.args.values); got != tt.want {
				t.Errorf("makeBST() = %v, want %q", got, tt.want)
			}
		})
	}
}

func Test_binaryTreeIntoLists(t *testing.T) {
	type args struct {
		tree string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"simple test",
			args{`{"value":2,"left":{"value":1,"left":{"value":0}},"right":{"value":3}}`},
			[][]int{{2}, {1, 3}, {0}},
		},
		{
			"big test",
			args{`{"value":47,"left":{"value":24,"left":{"value":11,"left":{"value":5,"left":{"value":2,"left":{"value":1,"left":{"value":0}},"right":{"value":4,"left":{"value":3}}},"right":{"value":9,"left":{"value":7,"left":{"value":6}},"right":{"value":10}}},"right":{"value":18,"left":{"value":16,"left":{"value":13,"left":{"value":12}},"right":{"value":17}},"right":{"value":22,"left":{"value":21,"left":{"value":20}},"right":{"value":23}}}},"right":{"value":34,"left":{"value":29,"left":{"value":27,"left":{"value":26,"left":{"value":25}},"right":{"value":28}},"right":{"value":32,"left":{"value":31,"left":{"value":30}},"right":{"value":33}}},"right":{"value":41,"left":{"value":39,"left":{"value":38,"left":{"value":36}},"right":{"value":40}},"right":{"value":45,"left":{"value":44,"left":{"value":42}},"right":{"value":46}}}}},"right":{"value":71,"left":{"value":60,"left":{"value":52,"left":{"value":50,"left":{"value":49,"left":{"value":48}},"right":{"value":51}},"right":{"value":57,"left":{"value":56,"left":{"value":53}},"right":{"value":59}}},"right":{"value":66,"left":{"value":63,"left":{"value":62,"left":{"value":61}},"right":{"value":65}},"right":{"value":69,"left":{"value":68,"left":{"value":67}},"right":{"value":70}}}},"right":{"value":86,"left":{"value":80,"left":{"value":76,"left":{"value":75,"left":{"value":72}},"right":{"value":77}},"right":{"value":83,"left":{"value":82,"left":{"value":81}},"right":{"value":84}}},"right":{"value":94,"left":{"value":92,"left":{"value":90,"left":{"value":87}},"right":{"value":93}},"right":{"value":97,"left":{"value":96,"left":{"value":95}},"right":{"value":98}}}}}}`},
			[][]int{{47}, {24, 71}, {11, 34, 60, 86}, {5, 18, 29, 41, 52, 66, 80, 94}, {2, 9, 16, 22, 27, 32, 39, 45, 50, 57, 63, 69, 76, 83, 92, 97}, {1, 4, 7, 10, 13, 17, 21, 23, 26, 28, 31, 33, 38, 40, 44, 46, 49, 51, 56, 59, 62, 65, 68, 70, 75, 77, 82, 84, 90, 93, 96, 98}, {0, 3, 6, 12, 20, 25, 30, 36, 42, 48, 53, 61, 67, 72, 81, 87, 95}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := binaryTreeIntoLists(test.args.tree); !reflect.DeepEqual(got, test.want) {
				t.Errorf("binaryTreeIntoLists() = %v, want %v", got, test.want)
			}
		})
	}
}
