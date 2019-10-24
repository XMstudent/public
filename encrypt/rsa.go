package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	errors "smart4s.com/public/perror"
)

// encrypt rsa-sha1 with pkcs8 private key
func RsaEncryptSha1WithPKCS8(originalData []byte, privKey string) (string, error) {
	var strKey string
	split(privKey, &strKey)
	blockKeyStr := "-----BEGIN RSA PRIVATE KEY-----\n" +
		strKey +
		"-----END RSA PRIVATE KEY-----"
	block, _ := pem.Decode([]byte(blockKeyStr))

	var encryptedData []byte
	var err error
	h := sha1.New()
	h.Write(originalData)
	privateKey, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
	if privateKey == nil {
		//秘钥解析失败
		return "", errors.New( "private key parsing with pkcs8 failed")
	}
	encryptedData, err = rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, h.Sum(nil))
	return base64.StdEncoding.EncodeToString(encryptedData), err
}

// encrypt rsa-sha1 with pkcs1 private key
func RsaEncryptSha1WithPKCS1(originalData []byte, privKey string) (string, error) {
	var strKey string
	split(privKey, &strKey)
	blockKeyStr := "-----BEGIN RSA PRIVATE KEY-----\n" +
		strKey +
		"-----END RSA PRIVATE KEY-----"
	block, _ := pem.Decode([]byte(blockKeyStr))

	var encryptedData []byte
	var err error
	h := sha1.New()
	h.Write(originalData)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	if privateKey == nil {
		//秘钥解析失败
		return "", errors.New( "private key parsing with pkcs1 failed")
	}
	encryptedData, err = rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, h.Sum(nil))
	return base64.StdEncoding.EncodeToString(encryptedData), err
}

func split(key string, temp *string) {
	if len(key) <= 64 {
		*temp = *temp + key + "\n"
	}
	for i := 0; i < len(key); i++ {
		if (i+1)%64 == 0 {
			*temp = *temp + key[:i+1] + "\n"
			key = key[i+1:]
			split(key, temp)
			break
		}
	}
}

// verify sign
func RsaVerifySignWithSha1(originalData []byte, signData, pubKey string) error {
	var strKey string
	split(pubKey, &strKey)
	blockKeyStr := "-----BEGIN PUBLIC KEY-----\n" +
		strKey +
		"-----END PUBLIC KEY-----"
	block, _ := pem.Decode([]byte(blockKeyStr))
	publicKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	if publicKey == nil {
		return errors.New("there is something wrong with public key")
	}
	pub := publicKey.(*rsa.PublicKey)
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write(originalData)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hash.Sum(nil), sign)
}
