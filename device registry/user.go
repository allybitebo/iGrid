package registry

import (
	"fmt"
	"github.com/piusalfred/registry/pkg/errors"
	"golang.org/x/net/idna"
	"regexp"
	"strings"
	"time"
)

var (
	ErrInvalidEmail  = errors.New("invalid email format")
	ErrShortPassword = errors.New("password length is short")
)

type UserGroup int

const (
	Admin UserGroup = iota + 1 //owner of the network
	RegionAdmin
	RegionUser
)

const (
	minPassLen   = 8
	maxLocalLen  = 64
	maxDomainLen = 255
	maxTLDLen    = 24 // longest TLD currently in existence

	atSeparator  = "@"
	dotSeparator = "."
)

var (
	userRegexp    = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	hostRegexp    = regexp.MustCompile("^[^\\s]+\\.[^\\s]+$")
	userDotRegexp = regexp.MustCompile("(^[.]{1})|([.]{1}$)|([.]{2,})")
)

type User struct {
	ID       string `json:"id,omitempty"`       //id or user token | uuid
	Name     string `json:"name"`               //fullname
	Email    string `json:"email"`              //email
	Password string `json:"password,omitempty"` //password of user
	Group    int    `json:"group,omitempty"`    //user group
	Region   string `json:"region,omitempty"`   //operating region in case of multi cloud
	Created  string `json:"created,omitempty"`  //when was this user added
}

func CreateUser(hasher Hasher, provider UUIDProvider, name, email, password, region string) (User, error) {
	id, err := provider.ID()
	if err != nil {
		return User{}, err
	}

	if !isEmail(email) {
		return User{}, ErrInvalidEmail
	}

	hash, err := hasher.Hash(password)

	if !isEmail(email) {
		return User{}, err
	}

	now := time.Now().Format(time.RFC3339)

	return User{
		ID:       id,
		Name:     name,
		Email:    email,
		Group:    3,
		Password: hash,
		Region:   region,
		Created:  now,
	}, nil

}

// Validate returns an error if user representation is invalid.
func (u User) Validate() error {
	if !isEmail(u.Email) {
		return ErrInvalidEmail
	}

	if len(u.Password) < minPassLen {
		return ErrShortPassword
	}

	return nil
}

func isEmail(email string) bool {
	if email == "" {
		return false
	}

	es := strings.Split(email, atSeparator)
	if len(es) != 2 {
		return false
	}
	local, host := es[0], es[1]

	if local == "" || len(local) > maxLocalLen {
		return false
	}

	hs := strings.Split(host, dotSeparator)
	if len(hs) < 2 {
		return false
	}
	domain, ext := hs[0], hs[1]

	// Check subdomain and validate
	if len(hs) > 2 {
		if domain == "" {
			return false
		}

		for i := 1; i < len(hs)-1; i++ {
			sub := hs[i]
			if sub == "" {
				return false
			}
			domain = fmt.Sprintf("%s.%s", domain, sub)
		}

		ext = hs[len(hs)-1]
	}

	if domain == "" || len(domain) > maxDomainLen {
		return false
	}
	if ext == "" || len(ext) > maxTLDLen {
		return false
	}

	punyLocal, err := idna.ToASCII(local)
	if err != nil {
		return false
	}
	punyHost, err := idna.ToASCII(host)
	if err != nil {
		return false
	}

	if userDotRegexp.MatchString(punyLocal) || !userRegexp.MatchString(punyLocal) || !hostRegexp.MatchString(punyHost) {
		return false
	}

	return true
}
