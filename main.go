package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/decred/dcrd/chaincfg"
	"github.com/decred/dcrd/hdkeychain"
)

var (
	keyFlag  = flag.String("key", "", "hex private key, xpriv, or xpub of root key")
	pathFlag = flag.String("path", "", "derivation path from key")
	netFlag  = flag.String("net", "main", "decred network (main|test|sim)")
)

func main() {
	flag.Parse()

	var params *chaincfg.Params
	switch *netFlag {
	case "main", "mainnet":
		params = &chaincfg.MainNetParams
	case "test", "testnet":
		params = &chaincfg.TestNet3Params
	case "sim", "simnet":
		params = &chaincfg.SimNetParams
	default:
		fmt.Fprintf(os.Stderr, "unknown network %q\n", *netFlag)
		return
	}

	k, err := hdkeychain.NewKeyFromString(*keyFlag)
	if err != nil {
		if len(*keyFlag) != 64 {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		h, err := hex.DecodeString(*keyFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		k, err = hdkeychain.NewMaster(h, params)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	var path []uint32
	for i, s := range strings.Split(*pathFlag, "/") {
		if s == "" || (i == 0 && s == "m") {
			continue
		}
		hardened := s[len(s)-1] == '\''
		if hardened {
			s = s[:len(s)-1]
		}
		child, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if hardened {
			child += hdkeychain.HardenedKeyStart
		}
		path = append(path, uint32(child))
	}

	for _, child := range path {
		k, err = k.Child(child)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	b := new(bytes.Buffer)
	if k.IsPrivate() {
		fmt.Fprintf(b, "xpriv: %v\n", k)
		k, err = k.Neuter()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
	fmt.Fprintf(b, "xpub:  %v\n", k)
	addr, err := k.Address(params)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprintf(b, "addr:  %v\n", addr)
	io.Copy(os.Stdout, b)
}
