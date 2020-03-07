package api

import (
	uuid2 "github.com/google/uuid"
	"github.com/modding-engineer/api-modding-engineer/pkg/uuid"
	"net/url"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	u, _ := url.Parse("https://test.api.modding.engineer/")
	u2, _ := url.Parse("https://test.api.modding.engineer/v2")
	type args struct {
		name     string
		uri      string
		testSign string
	}
	tests := []struct {
		name string
		args args
		want *API
	}{
		{
			name: "handles a new root api",
			args: args{
				name:     "Test API",
				uri:      "https://test.api.modding.engineer/",
				testSign: "https://test.api.modding.engineer/flaky",
			},
			want: &API{
				Name: "Test API",
				UUID: uuid.UUID(uuid2.MustParse("2ff12d50-3c93-521a-bd0b-6ac4477e0a0c")),
				URL:  u,
			},
		},
		{
			name: "handles a non-root api",
			args: args{
				name:     "Test API 2",
				uri:      "https://test.api.modding.engineer/v2",
				testSign: "https://test.api.modding.engineer/v2/flaky",
			},
			want: &API{
				Name: "Test API 2",
				UUID: uuid.UUID(uuid2.MustParse("47bf3691-7244-52f0-adeb-a4a8efa0e175")),
				URL:  u2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.name, tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			} else {
				t.Run("can sign", func(t *testing.T) {
					if sign, err := got.Sign(tt.args.testSign); err != nil {
						t.Errorf("could not sign: %v", err)
					} else {
						t.Run("validate", func(t *testing.T) {
							if !got.Validate(sign.String(), tt.args.testSign) {
								t.Errorf("signature did not validate: %v want %v", got, tt.want)
							}
						})
					}
				})
			}
		})
	}
}
