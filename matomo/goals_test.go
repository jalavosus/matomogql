package matomo_test

import (
	"context"
	"testing"

	"github.com/jalavosus/matomogql/matomo"
)

const (
	testEndpoint = "https://demo.matomo.cloud/"
	testApiKey   = "anonymous"
)

func TestGetGoal(t *testing.T) {
	testCtx := context.TODO()

	type args struct {
		ctx    context.Context
		idSite int
		idGoal int
	}

	type testWant struct {
		name string
	}

	tests := []struct {
		name    string
		args    args
		want    testWant
		wantErr bool
	}{
		{
			name: "valid_idsite",
			args: args{
				ctx:    testCtx,
				idSite: 11,
				idGoal: 3,
			},
			want:    testWant{name: "Contact Us"},
			wantErr: false,
		},
		{
			name: "invalid_idsite",
			args: args{
				ctx:    testCtx,
				idSite: 15,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := matomo.GetGoal(tt.args.ctx, tt.args.idSite, tt.args.idGoal)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if tt.want.name != got.Name {
				t.Errorf("GetGoal() got.name = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGoals(t *testing.T) {
	testCtx := context.TODO()

	type args struct {
		ctx    context.Context
		idSite int
		idGoal int
	}

	type testWant struct {
		index  int
		name   string
		idGoal int
	}

	tests := []struct {
		name    string
		args    args
		want    []testWant
		wantErr bool
	}{
		{
			name: "valid_data",
			args: args{
				ctx:    testCtx,
				idSite: 11,
			},
			want: []testWant{
				{index: 2, name: "Contact Us", idGoal: 3},
				{index: 3, name: "Phone Number Clicks", idGoal: 4},
			},
			wantErr: false,
		},
		{
			name: "invalid_idsite",
			args: args{
				ctx:    testCtx,
				idSite: 15,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := matomo.GetGoals(tt.args.ctx, tt.args.idSite, nil)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			for _, want := range tt.want {
				if want.index > len(got) {
					t.Errorf("GetGoals() len(got) = %v, want >= %v", len(got), want.index)
				}

				item := got[want.index]
				if item.Name != want.name {
					t.Errorf("GetGoals() got[%d].Name = %v, want %v", want.index, item.Name, want.name)
				}
				if item.IDGoal != want.idGoal {
					t.Errorf("GetGoals() got[%d].IDGoal = %v, want %v", want.index, item.IDGoal, want.idGoal)
				}
			}
		})
	}
}
