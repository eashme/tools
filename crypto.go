package tools

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
)

// 生成 私钥和公钥
func GenRsaPairs() (pub string, pri string, err error) {
	prv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}
	pub = hex.EncodeToString(x509.MarshalPKCS1PublicKey(&prv.PublicKey))
	pri = hex.EncodeToString(x509.MarshalPKCS1PrivateKey(prv))
	return pub, pri, nil
}

// RSA加密
func RsaPubEncode(raw string, pub string) (en string, err error) {
	pb, err := hex.DecodeString(pub)
	if err != nil {
		return "", err
	}
	pubK, err := x509.ParsePKCS1PublicKey(pb)
	if err != nil {
		return "", err
	}

	// 公钥加密
	b, err := rsa.EncryptPKCS1v15(rand.Reader, pubK, []byte(raw))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(reverse(b)), nil
}

// rsa解密
func RsaPriDecode(en string, pri string) (raw string, err error) {
	pb, err := hex.DecodeString(pri)
	if err != nil {
		return "", err
	}

	priK, err := x509.ParsePKCS1PrivateKey(pb)
	if err != nil {
		return "", nil
	}

	b, err := hex.DecodeString(en)
	if err != nil {
		return "", err
	}

	b, err = rsa.DecryptPKCS1v15(rand.Reader, priK, reverse(b))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func reverse(b []byte) []byte {
	n := len(b)
	for i := 0; i < n-1; i += 2 {
		b[i], b[i+1] = b[i+1]+4, b[i]-4
	}
	return b
}

func Md5(raw string, salt ...string) string {
	e := md5.New()
	e.Write([]byte(raw))
	var s []byte
	if len(salt) > 0 {
		s = []byte(salt[0])
	}
	return hex.EncodeToString(e.Sum(s))
}

// 多串组合成一个md5值
func Md5NoSalt(s ...string) string {
	e := md5.New()
	for _, raw := range s {
		e.Write([]byte(raw))
	}
	return hex.EncodeToString(e.Sum(nil))
}
