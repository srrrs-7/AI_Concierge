package redis

import (
	"auth/driver/cache"
	"auth/util/env"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_GetDomain(t *testing.T) {

	type args struct {
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rds, err := cache.New(&env.EnvParams[string]{})
			if err != nil {
				log.Fatal(err)
			}
			got := New(&env.EnvParams[string]{}, rds)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetDomain() = %v", diff)
			}
		})
	}

}
