package builtin

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"time"

	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"math/rand"

	"github.com/tzmfreedom/land/ast"
)

var cryptoType = ast.CreateClass(
	"Crypto",
	[]*ast.Method{},
	ast.NewMethodMap(),
	ast.NewMethodMap(),
)

func init() {
	cryptoType.StaticMethods.Set(
		"decrypt",
		[]*ast.Method{
			ast.CreateMethod(
				"decrypt",
				BlobType,
				[]*ast.Parameter{
					stringTypeParameter,
					BlobTypeParameter,
					BlobTypeParameter,
					BlobTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					algorithmName := params[0].StringValue()
					privateKey := params[1].Value().([]byte)
					initializationVector := params[2].Value().([]byte)
					cipherText := params[3].Value().([]byte)

					keyLengthMap := map[string]int{
						"AES128": 128,
						"AES192": 192,
						"AES256": 256,
					}
					keyLength, ok := keyLengthMap[algorithmName]
					if !ok {
						panic("invalid algorithm: " + algorithmName)
					}
					if keyLength != len(privateKey) {
						panic(fmt.Sprintf("private key length does not match to algorithm: %d", len(privateKey)))
					}
					if keyLength != len(initializationVector) {
						panic(fmt.Sprintf("iv length does not match to algorithm: %d", len(privateKey)))
					}
					block, err := aes.NewCipher(privateKey)
					if err != nil {
						panic(err)
					}

					plain := make([]byte, len(cipherText))
					decrypter := cipher.NewCBCDecrypter(block, initializationVector)
					decrypter.CryptBlocks(plain, cipherText)
					padSize := int(plain[len(plain)-1])

					return NewBlob(plain[:len(plain)-padSize])
				},
			),
		},
	)

	cryptoType.StaticMethods.Set(
		"decryptWithManagedIV",
		[]*ast.Method{
			ast.CreateMethod(
				"decryptWithManagedIV",
				BlobType,
				[]*ast.Parameter{
					stringTypeParameter,
					BlobTypeParameter,
					BlobTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					algorithmName := params[0].StringValue()
					privateKey := params[1].Value().([]byte)
					cipherText := params[2].Value().([]byte)

					keyLengthMap := map[string]int{
						"AES128": 128,
						"AES192": 192,
						"AES256": 256,
					}
					keyLength, ok := keyLengthMap[algorithmName]
					if !ok {
						panic("invalid algorithm: " + algorithmName)
					}
					if keyLength != len(privateKey) {
						panic(fmt.Sprintf("private key length does not match to algorithm: %d", len(privateKey)))
					}
					block, err := aes.NewCipher(privateKey)
					if err != nil {
						panic(err)
					}

					plain := make([]byte, len(cipherText[aes.BlockSize:]))
					decrypter := cipher.NewCBCDecrypter(block, cipherText[:aes.BlockSize])
					decrypter.CryptBlocks(plain, cipherText[aes.BlockSize:])
					padSize := int(plain[len(plain)-1])

					return NewBlob(plain[:len(plain)-padSize])
				},
			),
		},
	)

	cryptoType.StaticMethods.Set(
		"encrypt",
		[]*ast.Method{
			ast.CreateMethod(
				"encrypt",
				BlobType,
				[]*ast.Parameter{
					stringTypeParameter,
					BlobTypeParameter,
					BlobTypeParameter,
					BlobTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					algorithmName := params[0].StringValue()
					privateKey := params[1].Value().([]byte)
					initializationVector := params[2].Value().([]byte)
					clearText := params[3].Value().([]byte)

					keyLengthMap := map[string]int{
						"AES128": 128,
						"AES192": 192,
						"AES256": 256,
					}
					keyLength, ok := keyLengthMap[algorithmName]
					if !ok {
						panic("invalid algorithm: " + algorithmName)
					}
					if keyLength != len(privateKey) {
						panic(fmt.Sprintf("private key length does not match to algorithm: %d", len(privateKey)))
					}
					if keyLength != len(initializationVector) {
						panic(fmt.Sprintf("iv length does not match to algorithm: %d", len(privateKey)))
					}
					block, err := aes.NewCipher(privateKey)
					if err != nil {
						panic(err)
					}

					paddedPlain := padPKCS7(clearText)
					cipherText := make([]byte, len(paddedPlain))

					encrypter := cipher.NewCBCEncrypter(block, initializationVector)
					encrypter.CryptBlocks(cipherText, paddedPlain)

					return NewBlob(cipherText)
				},
			),
		},
	)

	cryptoType.StaticMethods.Set(
		"encryptWithManagedIV",
		[]*ast.Method{
			ast.CreateMethod(
				"encryptWithManagedIV",
				BlobType,
				[]*ast.Parameter{
					stringTypeParameter,
					BlobTypeParameter,
					BlobTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					algorithmName := params[0].StringValue()
					privateKey := params[1].Value().([]byte)
					clearText := params[2].Value().([]byte)

					keyLengthMap := map[string]int{
						"AES128": 128,
						"AES192": 192,
						"AES256": 256,
					}
					keyLength, ok := keyLengthMap[algorithmName]
					if !ok {
						panic("invalid algorithm: " + algorithmName)
					}
					if keyLength != len(privateKey) {
						panic(fmt.Sprintf("private key length does not match to algorithm: %d", len(privateKey)))
					}
					block, err := aes.NewCipher(privateKey)
					if err != nil {
						panic(err)
					}

					paddedPlain := padPKCS7(clearText)
					cipherText := make([]byte, aes.BlockSize+len(paddedPlain))
					iv := cipherText[:aes.BlockSize]
					rand.Read(iv)

					encrypter := cipher.NewCBCEncrypter(block, iv)
					encrypter.CryptBlocks(cipherText[aes.BlockSize:], paddedPlain)

					return NewBlob(cipherText)
				},
			),
		},
	)

	cryptoType.StaticMethods.Set(
		"generateAesKey",
		[]*ast.Method{
			ast.CreateMethod(
				"generateAesKey",
				BlobType,
				[]*ast.Parameter{IntegerTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					size := params[0].IntegerValue()
					allowedSize := map[int]struct{}{
						128: {},
						192: {},
						256: {},
					}
					_, ok := allowedSize[size]
					if !ok {
						panic(fmt.Sprintf("invalid size: %d", size))
					}
					key := make([]byte, size/8)
					rand.Read(key)
					return NewBlob(key)
				},
			),
		},
	)

	cryptoType.StaticMethods.Set(
		"generateDigest",
		[]*ast.Method{
			ast.CreateMethod(
				"generateDigest",
				BlobType,
				[]*ast.Parameter{
					stringTypeParameter,
					BlobTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					algorithmName := params[0].StringValue()
					input := params[1].Value().([]byte)
					algorithmFuncMap := map[string]func() hash.Hash{
						"SHA1":   sha1.New,
						"SHA256": sha256.New,
						"SHA512": sha512.New,
						"MD5":    md5.New,
					}
					algorithmFunc, ok := algorithmFuncMap[algorithmName]
					if !ok {
						panic("invalid algorithm name: " + algorithmName)
					}
					h := algorithmFunc()
					h.Write(input)
					return NewBlob(h.Sum(nil))
				},
			),
		},
	)

	cryptoType.StaticMethods.Set(
		"generateMac",
		[]*ast.Method{
			ast.CreateMethod(
				"generateMac",
				BlobType,
				[]*ast.Parameter{
					stringTypeParameter,
					BlobTypeParameter,
					BlobTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					algorithmName := params[0].StringValue()
					input := params[1].Value().([]byte)
					privateKey := params[2].Value().([]byte)

					algorithmFuncMap := map[string]func() hash.Hash{
						"hmacSHA1":   sha1.New,
						"hmacSHA256": sha256.New,
						"hmacSHA512": sha512.New,
						"hmacMD5":    md5.New,
					}
					algorithmFunc, ok := algorithmFuncMap[algorithmName]
					if !ok {
						panic("invalid algorithm name: " + algorithmName)
					}
					mac := hmac.New(algorithmFunc, privateKey)
					_, err := mac.Write(input)
					if err != nil {
						panic(err)
					}
					return NewBlob(mac.Sum(nil))
				},
			),
		},
	)

	cryptoType.StaticMethods.Set(
		"getRandomInteger",
		[]*ast.Method{
			ast.CreateMethod(
				"getRandomInteger",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					rand.Seed(time.Now().UnixNano())
					return NewInteger(rand.Int())
				},
			),
		},
	)

	cryptoType.StaticMethods.Set(
		"getRandomLong",
		[]*ast.Method{
			ast.CreateMethod(
				"getRandomLong",
				IntegerType, // TODO: impl
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					rand.Seed(time.Now().UnixNano())
					return NewInteger(rand.Int())
				},
			),
		},
	)

	primitiveClassMap.Set("Date", DateType)
}

func padPKCS7(data []byte) []byte {
	padSize := aes.BlockSize - len(data)%aes.BlockSize
	appendChars := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(data, appendChars...)
}
