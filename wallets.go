package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletFileDir = "wallet.dat"

// Keep map of address -> wallet
type Wallets struct {
	Wallets map[string]*Wallet
}

// Create Wallets object and load data from file if existed
func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile()

	return &wallets, err
}

// Create and add a Wallet to Wallets
func (wallets *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())

	wallets.Wallets[address] = wallet
	return address
}

// Given the address, return the wallet in map
func (wallets Wallets) GetWallet(address string) Wallet {
	return *wallets.Wallets[address]
}

// Return all addresses in this collection of wallet
func (wallets *Wallets) GetAddresses() []string {
	var result []string

	for address := range wallets.Wallets {
		result = append(result, address)
	}

	return result
}

// Load wallets object from file
func (wallets *Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFileDir); os.IsNotExist(err) {
		return err
	}

	walletFile, err := ioutil.ReadFile(walletFileDir)
	if err != nil {
		return err
	}

	var tempWallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(walletFile))
	err = decoder.Decode(&tempWallets)
	if err != nil {
		return err
	}

	wallets.Wallets = tempWallets.Wallets

	return nil
}

// Save wallets object to file
func (wallets Wallets) SaveToFile() {
	var data bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&data)
	err := encoder.Encode(wallets)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFileDir, data.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
