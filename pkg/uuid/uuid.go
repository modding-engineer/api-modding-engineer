package uuid

import (
	"github.com/google/uuid"
)

type UUID uuid.UUID

const DNSNameSpace = "{6ba7b810-9dad-11d1-80b4-00c04fd430c8}"
const URLNameSpace = "{6ba7b811-9dad-11d1-80b4-00c04fd430c8}"

func New(nameSpace string) UUID {
	nameSpacedId := uuid.NewSHA1(uuid.New(), []byte(nameSpace))
	return UUID(nameSpacedId)
}

func FromAPIPath(apiPath string) UUID     { return New("https://api.modding.engineer" + apiPath) }
func FromURL(fullUrl string) UUID         { return New(URLNameSpace + fullUrl) }
func FromHostname(fullHost string) UUID   { return New(DNSNameSpace + fullHost) }
func FromSubDomain(shortHost string) UUID { return FromHostname(shortHost + ".modding.engineer") }
func (u UUID) String() string             { return uuid.UUID(u).String() }
