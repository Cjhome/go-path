package hash

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	cost int
}

//Make 加密方法
func (b *Bcrypt)Make(password []byte)([]byte, error) {
	return bcrypt.GenerateFromPassword(password,b.cost)
}
//Check 检查方法

func (b *Bcrypt) Check(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func (b *Bcrypt)NewHash() {
	hash := Bcrypt{
		cost: bcrypt.DefaultCost,
	}
	fmt.Println(hash)
}