package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type key struct {
	firstParam  *big.Int
	secondParam *big.Int
}

type rsa struct {
	publicKey  key
	privateKey key
	mod        *big.Int
}

func (r *rsa) generatePublicKey(bitSize int) {
	p, err := rand.Prime(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}
	pSub := new(big.Int)
	pSub.Sub(p, big.NewInt(1))

	q, err := rand.Prime(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}
	qSub := new(big.Int)
	qSub.Sub(q, big.NewInt(1))

	r.mod = new(big.Int)
	r.mod.Mul(pSub, qSub)

	e, err := rand.Prime(rand.Reader, bitSize)
	if err != nil {
		panic(err)
	}
	for e.Cmp(r.mod) != -1 {
		e, err = rand.Prime(rand.Reader, bitSize)
		if err != nil {
			panic(err)
		}
	}

	r.publicKey = key{e, p.Mul(p, q)}
}

func (r *rsa) generatePrivateKey() {
	d := new(big.Int)
	d.ModInverse(r.publicKey.firstParam, r.mod)
	r.privateKey = key{d, r.publicKey.secondParam}
}

func (r *rsa) generateKeys(bitSize int) {
	r.generatePublicKey(bitSize)
	r.generatePrivateKey()
}

func rsaEncrypt(hash *big.Int, publicKey key) *big.Int {
	res := new(big.Int)
	return res.Exp(hash, publicKey.firstParam, publicKey.secondParam)
}

func rsaDecrypt(encryptText *big.Int, privateKey key) *big.Int {
	res := new(big.Int)
	return res.Exp(encryptText, privateKey.firstParam, privateKey.secondParam)
}

func main() {

	var r rsa
	r.generateKeys(1024)
	temp := rsaEncrypt(big.NewInt(20202), r.publicKey)
	fmt.Println(temp)
	fmt.Println(rsaDecrypt(temp, r.privateKey))
}
