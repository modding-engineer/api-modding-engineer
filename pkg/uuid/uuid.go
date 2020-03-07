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
	apiURLNameSpace = uuid.NewSHA1(uuid.NameSpaceURL, []byte("https://api.modding.engineer"))
)

func DNSNamespace() UUID { return UUID(dnsNameSpace) }
func URLNamespace() UUID { return UUID(apiURLNameSpace) }

func newFromString(v string) (UUID, error) {
	if u, err := url.Parse(v); err == nil {
		if strings.HasSuffix(u.Hostname(), ".api.modding.engineer") || u.Hostname() == "api.modding.engineer" {
			if u.String() == "https://api.modding.engineer" || u.String() == "https://api.modding.engineer/" {
				return UUID(apiURLNameSpace), nil
			}
			return newId(apiURLNameSpace, u.String()), nil
		}
	} else {
		return UUID(uuid.Nil), fmt.Errorf("could not resolve namespace for URL value: %v", v)
	}
	if strings.HasSuffix(v, ".modding.engineer") || v == "modding.engineer" {
		if v == "modding.engineer" {
			return UUID(dnsNameSpace), nil
		}
		return newId(dnsNameSpace, v), nil
	}
	return UUID(uuid.Nil), fmt.Errorf("could not resolve namespace for value: %v", v)
}

func New(value interface{}) UUID {
	switch v := value.(type) {
	case string:
		id, err := newFromString(v)
		if err != nil {
			panic(err)
		}
		return id
	case fmt.Stringer:
		id, err := newFromString(v.String())
		if err != nil {
			panic(err)
		}
		return id
	}
	panic(fmt.Errorf("could not resolve namespace for value: %v", value))
}

func newId(nameSpace uuid.UUID, value string) UUID {
	return UUID(uuid.NewSHA1(nameSpace, []byte(value)))
}

func FromAPIURL(apiUrl string) UUID {
	if apiUrl == "https://api.modding.engineer" || apiUrl == "https://api.modding.engineer/" {
		return UUID(apiURLNameSpace)
	}
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
	if strings.HasSuffix(newHost, ".modding.engineer") || newHost == "modding.engineer" {
		if newHost == "modding.engineer" {
			return UUID(dnsNameSpace)
		}
		return newId(dnsNameSpace, newHost)
	}
	panic(fmt.Errorf("refusing to create id for unknown host: %v", newHost))
}

func (u UUID) String() string { return uuid.UUID(u).String() }
func (u UUID) Validate(value string) bool {
	return u.String() == New(value).String()
}
