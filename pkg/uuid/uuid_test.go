package uuid

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestFromAPIURL(t *testing.T) {
	type args struct {
		apiUrl string
	}
	tests := []struct {
		name string
		args args
		want UUID
	}{
		{
			"Creates expected root url",
			args{"https://api.modding.engineer/"},
			UUID(uuid.MustParse("510213cc-19ec-54bd-b226-7bd180548e3e")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromAPIURL(tt.args.apiUrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromAPIURL() = %v, want %v", got, tt.want)
			} else {
				t.Run("validated", func(t *testing.T) {
					if !got.Validate(APIURLNameSpace, tt.args.apiUrl) {
						t.Errorf("FromAPIURL() did not create a validated uuid; got: %v", got.String())
					}
				})
			}
		})
	}
}

func TestFromSubDomain(t *testing.T) {
	type args struct {
		id      uuid.UUID
		newHost string
	}
	tests := []struct {
		name string
		args args
		want UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromSubDomain(tt.args.id, tt.args.newHost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromSubDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		nameSpace uuid.UUID
		value     string
	}
	tests := []struct {
		name string
		args args
		want UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.nameSpace, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_String(t *testing.T) {
	tests := []struct {
		name string
		u    UUID
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_Validate(t *testing.T) {
	type args struct {
		nameSpace uuid.UUID
		value     string
	}
	tests := []struct {
		name string
		u    UUID
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Validate(tt.args.nameSpace, tt.args.value); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
