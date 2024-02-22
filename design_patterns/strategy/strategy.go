package main

import "fmt"

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}
type HashAlgorithm interface {
	Hash(h *PasswordProtector)
}

func NewPasswordProtector(user string, passwordName string, hash HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

func (pp *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	pp.hashAlgorithm = hash
}

func (pp *PasswordProtector) Hash() {
	pp.hashAlgorithm.Hash(pp)
}

type SHA struct{}

func (SHA) Hash(pp *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", pp.passwordName)
}

type MD5 struct{}

func (MD5) Hash(pp *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s\n", pp.passwordName)
}

func main() {
	sha := new(SHA)
	md5 := &MD5{}
	passwordProtector := NewPasswordProtector("Hugo", "Gmail Password", sha)
	passwordProtector.Hash()
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}
