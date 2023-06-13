package auth

import (
	"auth/driver/db/model"
	"auth/pkg/entity"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvertAuthEntity(t *testing.T) {

	type args struct {
		authModel *model.Auth
	}

	tests := []struct {
		name string
		args args
		want *entity.Auth
	}{
		{
			name: "test1",
			args: args{
				authModel: &model.Auth{
					UserID:   "test1",
					Password: "test1",
					Email:    "test1",
					Scope:    "test1",
				},
			},
			want: &entity.Auth{
				UserID:   "test1",
				Password: "test1",
				Email:    "test1",
				Scope:    "test1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, _ := ConvertAuthEntity(tt.args.authModel)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Diff() = %v", diff)
			}
		})
	}
}
