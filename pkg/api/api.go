package api

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	"github.com/modding-engineer/api-modding-engineer/pkg/uuid"
	"net/url"
	"strings"
)

type API struct {
	Name string
	UUID uuid.UUID
	URL  *url.URL
}

func (A API) Validate(id, uri string) bool {
	if signature, err := A.Sign(uri); err == nil && id == signature.String() {
		return true
	}
	return false
}

var InvalidSigningURIProvided = fmt.Errorf("invalid signing uri provided")

func (A API) Sign(uri string) (uuid.UUID, error) {
	// first try full URL
	u, err := url.Parse(uri)
	if err == nil {
		if strings.HasPrefix(u.String(), A.URL.String()) {
			return uuid.UUID(uuid2.NewSHA1(uuid2.UUID(A.UUID), []byte(u.String()))), nil
		}
	}
	if strings.HasPrefix(uri, "/") {
		signUri := fmt.Sprintf("%v/%v", A.URL.String(), strings.TrimPrefix(uri, "/"))
		if _, err := url.Parse(signUri); err != nil {
			return uuid.Nil, fmt.Errorf("%w: %v", InvalidSigningURIProvided, uri)
		}
		return uuid.UUID(uuid2.NewSHA1(uuid2.UUID(A.UUID), []byte(signUri))), nil
	}
	return uuid.Nil, fmt.Errorf("%w: %v", InvalidSigningURIProvided, uri)
}

func New(name, uri string) *API {
	var A = &API{
		Name: name,
	}
	if strings.HasSuffix(uri, "/") {
		uri = strings.TrimSuffix(uri, "/")
	}
	A.URL, _ = url.Parse(uri)
	A.UUID = uuid.New(A.URL.String())
	return A
}
