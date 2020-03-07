package uuid

import (
	"github.com/google/uuid"
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

func FromAPIPath(apiPath string) UUID   { return New(APIURLNameSpace, apiPath) }
func FromSubDomain(newHost string) UUID { return New(DNSNameSpace, newHost) }
func (u UUID) String() string           { return uuid.UUID(u).String() }
func (u UUID) Validate(nameSpace uuid.UUID, value string) bool {
	switch nameSpace {
	case DNSNameSpace:
		if splits := strings.Split(value, "."); len(splits) > 2 {
			stub := strings.Join(splits[len(splits)-2:], ".")
			if stub == "modding.engineer" {
				return u.String() == FromSubDomain(value).String()
			}
		}
	case APIDNSNameSpace:
		return APIDNSNameSpace.String() == value
	case APIURLNameSpace:
	}
	return false
}
