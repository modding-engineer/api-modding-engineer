package uuid

import (
	"fmt"
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
		{
			"Creates expected id from sub-domain url",
			args{"https://other.api.modding.engineer/path"},
			UUID(uuid.MustParse("8517e550-1af3-5518-b1dd-bfe60343e7c4")),
		},
		{
			"Creates expected path url",
			args{"https://api.modding.engineer/objects"},
			UUID(uuid.MustParse("c9fece94-43f4-5d16-9591-c83ed963da62")),
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
	t.Run("panics with external URLs", func(t *testing.T) {
		panicUrl := "https://api.example.com/panic"
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("using URL %v should have generated a panic", panicUrl)
			} else {
				fmt.Println("\trecovered from panic:", r)
			}
		}()
		_ = FromAPIURL(panicUrl)
	})
}

func TestFromDomain(t *testing.T) {
	type args struct {
		id      uuid.UUID
		newHost string
	}
	tests := []struct {
		name string
		args args
		want UUID
	}{
		{
			"Root domain uuid",
			args{DNSNameSpace, "modding.engineer"},
			UUID(uuid.MustParse("5d987235-de1d-5db2-9c14-ec2034a49528")),
		},
		{
			"Sub-domain uuid",
			args{DNSNameSpace, "sub.modding.engineer"},
			UUID(uuid.MustParse("be3f8955-ef1d-5f8c-8591-6c2215f53024")),
		},
		{
			"API sub-domain uuid",
			args{APIDNSNameSpace, "sub.api.modding.engineer"},
			UUID(uuid.MustParse("542588a7-38ea-56fe-92ac-b85ae42ed951")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromDomain(tt.args.id, tt.args.newHost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromDomain() = %v, want %v", got, tt.want)
			} else {
				t.Run("validated", func(t *testing.T) {
					if !got.Validate(tt.args.id, tt.args.newHost) {
						t.Errorf("FromDomain() did not create a validated uuid; got: %v", got.String())
					}
				})
			}
		})
	}
	t.Run("panics with invalid namespace id", func(t *testing.T) {
		panicId := uuid.New()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("using namespace %v should have generated a panic", panicId.String())
			} else {
				fmt.Println("\trecovered from panic:", r)
			}
		}()
		_ = FromDomain(panicId, "modding.engineer")
	})
	t.Run("panics with invalid host", func(t *testing.T) {
		panicHost := "example.com"
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("using hostname %v should have generated a panic", panicHost)
			} else {
				fmt.Println("\trecovered from panic:", r)
			}
		}()
		_ = FromDomain(DNSNameSpace, panicHost)
	})
	t.Run("panics with invalid api host", func(t *testing.T) {
		panicHost := "blert.modding.engineer"
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("using hostname %v should have generated a panic", panicHost)
			} else {
				fmt.Println("\trecovered from panic:", r)
			}
		}()
		_ = FromDomain(APIDNSNameSpace, panicHost)
	})
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
