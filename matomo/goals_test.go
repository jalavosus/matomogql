package matomo_test

import (
	"context"
	"testing"

	"github.com/jalavosus/matomogql/matomo"
	"github.com/stretchr/testify/assert"

	_ "github.com/joho/godotenv/autoload"
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
				idSite: 3,
				idGoal: 3,
			},
			want:    testWant{name: "RFQ Form"},
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

	client := matomo.NewClient(matomo.GetEnv())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetGoal(tt.args.ctx, tt.args.idSite, tt.args.idGoal)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoErrorf(t, err, "GetGoal() returned an error")

			assert.Equal(t, tt.want.name, got.Name)
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
				idSite: 3,
			},
			want: []testWant{
				{index: 0, name: "Email clicks", idGoal: 2},
				{index: 1, name: "RFQ Form", idGoal: 3},
				{index: 2, name: "Phone Number Clicks", idGoal: 4},
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

	client := matomo.NewClient(matomo.GetEnv())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetAllGoals(tt.args.ctx, tt.args.idSite, nil)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoErrorf(t, err, "GetAllGoals() returned an error")

			for _, want := range tt.want {
				assert.GreaterOrEqual(t, len(got), want.index)

				item := got[want.index]
				assert.Equal(t, want.name, item.Name)
				assert.Equalf(t, want.idGoal, item.IDGoal, "expected goal %[1]s item.IDGoal to be %[2]d, got %[3]d", item.Name, want.idGoal, item.IDGoal)
			}
		})
	}
}
