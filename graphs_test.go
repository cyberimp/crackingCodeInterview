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
			`{"left":{"value":0},"right":{"value":2},"value":1}`,
		},
		{
			"even test",
			args{[]int{0, 1, 2, 3}},
			`{"left":{"left":{"value":0},"value":1},"right":{"value":3},"value":2}`,
		},
		{
			"big one",
			args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 16, 17, 18, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 36, 38, 39, 40, 41, 42, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 56, 57, 59, 60, 61, 62, 63, 65, 66, 67, 68, 69, 70, 71, 72, 75, 76, 77, 80, 81, 82, 83, 84, 86, 87, 90, 92, 93, 94, 95, 96, 97, 98}},
			`{"left":{"left":{"left":{"left":{"left":{"left":{"value":0},"value":1},"right":{"left":{"value":3},"value":4},"value":2},"right":{"left":{"left":{"value":6},"value":7},"right":{"value":10},"value":9},"value":5},"right":{"left":{"left":{"left":{"value":12},"value":13},"right":{"value":17},"value":16},"right":{"left":{"left":{"value":20},"value":21},"right":{"value":23},"value":22},"value":18},"value":11},"right":{"left":{"left":{"left":{"left":{"value":25},"value":26},"right":{"value":28},"value":27},"right":{"left":{"left":{"value":30},"value":31},"right":{"value":33},"value":32},"value":29},"right":{"left":{"left":{"left":{"value":36},"value":38},"right":{"value":40},"value":39},"right":{"left":{"left":{"value":42},"value":44},"right":{"value":46},"value":45},"value":41},"value":34},"value":24},"right":{"left":{"left":{"left":{"left":{"left":{"value":48},"value":49},"right":{"value":51},"value":50},"right":{"left":{"left":{"value":53},"value":56},"right":{"value":59},"value":57},"value":52},"right":{"left":{"left":{"left":{"value":61},"value":62},"right":{"value":65},"value":63},"right":{"left":{"left":{"value":67},"value":68},"right":{"value":70},"value":69},"value":66},"value":60},"right":{"left":{"left":{"left":{"left":{"value":72},"value":75},"right":{"value":77},"value":76},"right":{"left":{"left":{"value":81},"value":82},"right":{"value":84},"value":83},"value":80},"right":{"left":{"left":{"left":{"value":87},"value":90},"right":{"value":93},"value":92},"right":{"left":{"left":{"value":95},"value":96},"right":{"value":98},"value":97},"value":94},"value":86},"value":71},"value":47}`,
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
			args{`{"left":{"left":{"value":0},"value":1},"right":{"value":3},"value":2}`},
			[][]int{{2}, {1, 3}, {0}},
		},
		{
			"big test",
			args{`{"left":{"left":{"left":{"left":{"left":{"left":{"value":0},"value":1},"right":{"left":{"value":3},"value":4},"value":2},"right":{"left":{"left":{"value":6},"value":7},"right":{"value":10},"value":9},"value":5},"right":{"left":{"left":{"left":{"value":12},"value":13},"right":{"value":17},"value":16},"right":{"left":{"left":{"value":20},"value":21},"right":{"value":23},"value":22},"value":18},"value":11},"right":{"left":{"left":{"left":{"left":{"value":25},"value":26},"right":{"value":28},"value":27},"right":{"left":{"left":{"value":30},"value":31},"right":{"value":33},"value":32},"value":29},"right":{"left":{"left":{"left":{"value":36},"value":38},"right":{"value":40},"value":39},"right":{"left":{"left":{"value":42},"value":44},"right":{"value":46},"value":45},"value":41},"value":34},"value":24},"right":{"left":{"left":{"left":{"left":{"left":{"value":48},"value":49},"right":{"value":51},"value":50},"right":{"left":{"left":{"value":53},"value":56},"right":{"value":59},"value":57},"value":52},"right":{"left":{"left":{"left":{"value":61},"value":62},"right":{"value":65},"value":63},"right":{"left":{"left":{"value":67},"value":68},"right":{"value":70},"value":69},"value":66},"value":60},"right":{"left":{"left":{"left":{"left":{"value":72},"value":75},"right":{"value":77},"value":76},"right":{"left":{"left":{"value":81},"value":82},"right":{"value":84},"value":83},"value":80},"right":{"left":{"left":{"left":{"value":87},"value":90},"right":{"value":93},"value":92},"right":{"left":{"left":{"value":95},"value":96},"right":{"value":98},"value":97},"value":94},"value":86},"value":71},"value":47}`},
			[][]int{{47}, {24, 71}, {11, 34, 60, 86}, {5, 18, 29, 41, 52, 66, 80, 94}, {2, 9, 16, 22, 27, 32, 39, 45, 50, 57, 63, 69, 76, 83, 92, 97}, {1, 4, 7, 10, 13, 17, 21, 23, 26, 28, 31, 33, 38, 40, 44, 46, 49, 51, 56, 59, 62, 65, 68, 70, 75, 77, 82, 84, 90, 93, 96, 98}, {0, 3, 6, 12, 20, 25, 30, 36, 42, 48, 53, 61, 67, 72, 81, 87, 95}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := binaryTreeIntoLists(test.args.tree); !reflect.DeepEqual(got, test.want) {
				t.Errorf("binaryTreeIntoLists() = %#v, want %v", got, test.want)
			}
		})
	}
}
