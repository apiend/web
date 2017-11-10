/**
	用户注册时，
	用户输入【账号】和【密码】（以及其他用户信息）；
	系统为用户生成【Salt值】；
	系统将【Salt值】和【用户密码】连接到一起；
	对连接后的值进行散列，得到【Hash值】；
	将【Hash值1】和【Salt值】分别放到数据库中。

	用户登录时，
	用户输入【账号】和【密码】；
	系统通过用户名找到与之对应的【Hash值】和【Salt值】；
	系统将【Salt值】和【用户输入的密码】连接到一起；
	对连接后的值进行散列，得到【Hash值2】（注意是即时运算出来的值）；
	比较【Hash值1】和【Hash值2】是否相等，相等则表示密码正确，否则表示密码错误。\

**/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	salt := time.Now().Unix()
	m5 := md5.New()
	pwd, _ := m5.Write([]byte("Mi Ma"))
	fmt.Println(pwd)
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)
	fmt.Println(st, hex.EncodeToString(st))

	fmt.Println("dd")

}
