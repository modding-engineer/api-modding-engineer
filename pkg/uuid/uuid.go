package uuid

import (
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"strings"
)

type UUID uuid.UUID

var (
	DNSNameSpace    = uuid.NewSHA1(uuid.NameSpaceDNS, []byte("modding.engineer"))
	APIDNSNameSpace = uuid.NewSHA1(uuid.NameSpaceDNS, []byte("api.modding.engineer"))
	APIURLNameSpace = uuid.NewSHA1(uuid.NameSpaceURL, []byte("https://api.modding.engineer"))
)

func New(nameSpace uuid.UUID, value string) UUID {
	nameSpacedId := uuid.NewSHA1(nameSpace, []byte(value))
	return UUID(nameSpacedId)
}

func FromAPIURL(apiUrl string) UUID {
	u, err := url.Parse(apiUrl)
	if err != nil {
		panic(err)
	}
	if u.Hostname() != "api.modding.engineer" && !strings.HasSuffix(u.Hostname(), ".api.modding.engineer") {
		panic(fmt.Errorf("refusing to create id for non-API url: %v", u.String()))
	}
	return New(APIURLNameSpace, u.String())
}

func FromDomain(id uuid.UUID, newHost string) UUID {
	var U UUID
	switch id {
	case DNSNameSpace:
		if newHost != "modding.engineer" && !strings.HasSuffix(newHost, ".modding.engineer") {
			panic(fmt.Errorf("refusing to create id for invalid sub-domain: %v", newHost))
		}
		if strings.HasSuffix(newHost, ".modding.engineer") {
			U = New(id, strings.TrimSuffix(newHost, ".modding.engineer"))
			break
		}
		U = New(id, newHost)
	case APIDNSNameSpace:
		if newHost != "api.modding.engineer" && !strings.HasSuffix(newHost, ".api.modding.engineer") {
			panic(fmt.Errorf("refusing to create id for invalid sub-domain: %v", newHost))
		}
		if strings.HasSuffix(newHost, ".api.modding.engineer") {
			U = New(id, strings.TrimSuffix(newHost, ".api.modding.engineer"))
			break
		}
		U = New(id, newHost)
	default:
		panic(fmt.Errorf("refusing to create id in unknown namespace: %v", id.String()))
	}
	return U
}

func (u UUID) String() string { return uuid.UUID(u).String() }
func (u UUID) Validate(nameSpace uuid.UUID, value string) bool {
	switch nameSpace {
	case DNSNameSpace:
		if strings.HasSuffix(value, ".modding.engineer") {
			return u.String() == FromDomain(DNSNameSpace, value).String()
		} else if value == "modding.engineer" {
			return u.String() == FromDomain(DNSNameSpace, value).String()
		}
	case APIDNSNameSpace:
		if strings.HasSuffix(value, ".api.modding.engineer") {
			return u.String() == FromDomain(APIDNSNameSpace, value).String()
		} else if value == "api.modding.engineer" {
			return u.String() == FromDomain(APIDNSNameSpace, value).String()
		}
	case APIURLNameSpace:
		apiUrl, err := url.Parse(value)
		if err != nil {
			return false
		}
		return u.String() == FromAPIURL(apiUrl.String()).String()
	}
	return false
}
