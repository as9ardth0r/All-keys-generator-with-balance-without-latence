package main

import (
	"fmt"
	"math/big"
	"net/http"
        "io/ioutil"
        "os"
        //"crypto/sha256"
	"crypto/rand"
	"strings"

	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

func getbalance(address string) string {
  response, err := http.Get("https://blockchain.info/q/getreceivedbyaddress/" + address)
  if err != nil {
      fmt.Printf("%s", err)
      os.Exit(1)
  }
  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Printf("%s", err)
    os.Exit(1)
  }
  return string(contents)
}

func randInt(min, max *big.Int) (*big.Int, error) {
	if min.Cmp(max) > 0 {
		min, max = max, min
	}
	intvl := new(big.Int)
	intvl.Add(intvl.Sub(max, min), big.NewInt(1))
	r, err := rand.Int(rand.Reader, intvl)
	if err != nil {
		return nil, err
	}
	r.Add(r, min)
	return r, nil
}

func main() {
	// Print header
	fmt.Printf("%64s %34s %34s\n", "Private", "Public", "Public Compressed")

	// Initialise big numbers with small numbers
        

	// Create a slice to pad our count to 32 bytes
	padded := make([]byte, 32)
        //rand.Read(padded)

	// Loop forever because we're never going to hit the end anyway
	for {
	
	        min, ok := new(big.Int).SetString(
		"9223372036854775808",
		10,
	)
	if !ok {
		return
	}
	max, ok := new(big.Int).SetString(
		"18446744073709551615",
		10,
	)
	if !ok {
		return
	}
	r, err := randInt(min, max)
	if err != nil {
		return
	}
		// Increment our counter
		//r.Add(r, min)

		// Copy count value's bytes to padded slice
		copy(padded[32-len(r.Bytes()):], r.Bytes())

		// Get private and public keys
		//h := sha256.New()
	        //h.Write(padded[:])
                //_, public := btcec.PrivKeyFromBytes(btcec.S256(), h.Sum(nil))
                _, public := btcec.PrivKeyFromBytes(btcec.S256(), padded)

		// Get compressed and uncompressed addresses
		caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)
		
		// Print keys
		balanceu := getbalance(uaddr.EncodeAddress())
		balancec := getbalance(caddr.EncodeAddress())
		if strings.HasPrefix(caddr.EncodeAddress(), "16jY7") {
		fmt.Printf("%x %34s %1s %34s %1s\n", padded, uaddr.EncodeAddress(), balanceu, caddr.EncodeAddress(), balancec)
		}
	}
}
