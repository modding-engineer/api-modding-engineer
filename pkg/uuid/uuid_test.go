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
			UUID(uuid.MustParse("6109d8c9-6102-5d6c-bc6a-777248c7e3ce")),
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
					if !got.Validate(tt.args.apiUrl) {
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
			args{dnsNameSpace, "modding.engineer"},
			UUID(uuid.MustParse("3f23498c-b045-523a-94ab-dd0e7d7b973c")),
		},
		{
			"Sub-domain uuid",
			args{dnsNameSpace, "sub.modding.engineer"},
			UUID(uuid.MustParse("1323aeaf-e74d-5d50-9a73-d347568c49d9")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromDomain(tt.args.newHost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromDomain() = %v, want %v", got, tt.want)
			} else {
				t.Run("validated", func(t *testing.T) {
					if !got.Validate(tt.args.newHost) {
						t.Errorf("FromDomain() did not create a valid uuid for %v; got: %v", tt.args.newHost, got.String())
					}
				})
			}
		})
	}
	t.Run("panics with invalid host", func(t *testing.T) {
		panicHost := "example.com"
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("using hostname %v should have generated a panic", panicHost)
			} else {
				fmt.Println("\trecovered from panic:", r)
			}
		}()
		_ = FromDomain(panicHost)
	})
}

func TestNew(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want UUID
	}{
		{
			"gets default root host",
			args{"modding.engineer"},
			UUID(uuid.MustParse("3f23498c-b045-523a-94ab-dd0e7d7b973c")),
		},
		{
			"gets default root url",
			args{"https://api.modding.engineer/"},
			UUID(uuid.MustParse("6109d8c9-6102-5d6c-bc6a-777248c7e3ce")),
		},
		{
			"gets same value as host helper",
			args{"sub.modding.engineer"},
			FromDomain("sub.modding.engineer"),
		},
		{
			"gets same value as url helper",
			args{"https://sub.api.modding.engineer"},
			FromAPIURL("https://sub.api.modding.engineer"),
		},
		{
			"handles one of our namespace UUIDs",
			args{dnsNameSpace},
			UUID(uuid.MustParse("3f23498c-b045-523a-94ab-dd0e7d7b973c")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("panics with unknown uuid", func(t *testing.T) {
		panicValue := uuid.New()
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("using value %v should have generated a panic", panicValue)
			} else {
				fmt.Println("\trecovered from panic:", r)
			}
		}()
		_ = New(panicValue)
	})
	t.Run("panics with unresolved namespace", func(t *testing.T) {
		panicValue := struct{}{}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("using value %v should have generated a panic", panicValue)
			} else {
				fmt.Println("\trecovered from panic:", r)
			}
		}()
		_ = New(panicValue)
	})
}

var _ fmt.Stringer = new(UUID)
