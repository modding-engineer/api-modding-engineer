package uuid

import (
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"strings"
)

type UUID uuid.UUID

var (
	dnsNameSpace    = uuid.NewSHA1(uuid.NameSpaceDNS, []byte("modding.engineer"))
	apiDNSNameSpace = uuid.NewSHA1(uuid.NameSpaceDNS, []byte("api.modding.engineer"))
	apiURLNameSpace = uuid.NewSHA1(uuid.NameSpaceURL, []byte("https://api.modding.engineer"))
)

func New(value string) UUID {
	if u, err := url.Parse(value); err == nil {
		if strings.HasSuffix(u.Hostname(), ".api.modding.engineer") || u.Hostname() == "api.modding.engineer" {
			return newId(apiURLNameSpace, u.String())
		}
	}
	if strings.HasSuffix(value, ".api.modding.engineer") || value == "api.modding.engineer" {
		return newId(apiDNSNameSpace, value)
	}
	if strings.HasSuffix(value, ".modding.engineer") || value == "modding.engineer" {
		return newId(dnsNameSpace, value)
	}
	panic(fmt.Errorf("could not resolve namespace for value: %v", value))
}

func newId(nameSpace uuid.UUID, value string) UUID {
	return UUID(uuid.NewSHA1(nameSpace, []byte(value)))
}

func FromAPIURL(apiUrl string) UUID {
	u, err := url.Parse(apiUrl)
	if err != nil {
		panic(err)
	}
	if u.Hostname() != "api.modding.engineer" && !strings.HasSuffix(u.Hostname(), ".api.modding.engineer") {
		panic(fmt.Errorf("refusing to create id for non-API url: %v", u.String()))
	}
	return New(u.String())
}

func FromDomain(newHost string) UUID {
	if strings.HasSuffix(newHost, ".api.modding.engineer") || newHost == "api.modding.engineer" {
		return newId(apiDNSNameSpace, newHost)
	}
	if strings.HasSuffix(newHost, ".modding.engineer") || newHost == "modding.engineer" {
		return newId(dnsNameSpace, newHost)
	}
	panic(fmt.Errorf("refusing to create id for unknown host: %v", newHost))
}

func (u UUID) String() string { return uuid.UUID(u).String() }
func (u UUID) Validate(value string) bool {
	return u.String() == New(value).String()
}
