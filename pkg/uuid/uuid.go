package uuid

import (
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

func FromAPIURL(apiUrl string) UUID                   { return New(APIURLNameSpace, apiUrl) }
func FromSubDomain(id uuid.UUID, newHost string) UUID { return New(id, newHost) }
func (u UUID) String() string                         { return uuid.UUID(u).String() }
func (u UUID) Validate(nameSpace uuid.UUID, value string) bool {
	switch nameSpace {
	case DNSNameSpace:
		if strings.HasSuffix(value, ".modding.engineer") {
			subDomain := strings.TrimSuffix(value, ".modding.engineers")
			return u.String() == FromSubDomain(DNSNameSpace, subDomain).String()
		}
	case APIDNSNameSpace:
		if strings.HasSuffix(value, ".api.modding.engineer") {
			subDomain := strings.TrimSuffix(value, ".api.modding.engineer")
			return u.String() == FromSubDomain(APIDNSNameSpace, subDomain).String()
		}
	case APIURLNameSpace:
		u, err := url.Parse(value)
		if err != nil {
			return false
		}
		return u.String() == FromAPIURL(u.Path).String()
	}
	return false
}
