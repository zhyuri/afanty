package core

import (
	"reflect"
	"testing"
)

var (
	passState = []byte(`{
  "Type": "Pass",
  "ResultPath": "$.coords",
  "Next": "End"
}`)
)

func TestBuildState(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Build PassState from Json",
			args: args{
				data: passState,
			},
			want: &PassState{
				State: State{
					BaseState: BaseState{
						Type: NamePassState,
					},
					Next: "End",
				},
				// cannot test Result field here, because it has different memory address
				// and cannot pass the reflect.DeepEqual
				ResultPath: "$.coords",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildState(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildState(%v) error = %v, wantErr %v", string(tt.args.data), err, tt.wantErr)
				return
			}
			// Be ware you need to use the pointer type !
			if _, ok := got.(*PassState); !ok {
				t.Errorf("BuildState type not match, want PassState.")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildState(%v) = %v, want %v", string(tt.args.data), got, tt.want)
			}
		})
	}
}

func TestParseStateType(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    reflect.Type
		wantErr bool
	}{
		{
			name: "PassState parse test",
			args: args{
				data: passState,
			},
			want:    reflect.TypeOf(PassState{}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStateType(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStateType(%v) error = %v, wantErr %v", string(tt.args.data), err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStateType(%v) = %v, want %v", string(tt.args.data), got, tt.want)
			}
		})
	}
}
