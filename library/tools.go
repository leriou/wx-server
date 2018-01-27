package library

import (
	// "encoding/xml"
	"encoding/base64"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)
// tools
type Tools struct {

}

func NewTools() *Tools {
	return new(Tools)
}
// md5加密
func (t *Tools) Md5(s string) string{ 
	return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}
// base64_encode 
func (t *Tools) Base64_encode(s string) string{
	str := base64.StdEncoding.EncodeToString([]byte(s))
	return str
}
// base64_decode
func (t *Tools) Base64_decode(s string) string{
	str,_ := base64.StdEncoding.DecodeString(s)
	return string(str)
}
// sha1
func (t *Tools) Sha1(s string) string{
	n := sha1.New()
	io.WriteString(n,s)
	tmp := fmt.Sprintf("%x",n.Sum(nil))
	return tmp
}