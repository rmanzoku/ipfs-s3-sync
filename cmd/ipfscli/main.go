package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	ipfss3 "github.com/rmanzoku/ipfs-s3-sync"
)

func copy(ctx context.Context, src, dst string) (err error) {
	fmt.Println("copy", src, dst)

	srcProtocol, srcPath := ipfss3.ParsePath(src)
	fmt.Println(srcProtocol, srcPath)
	body := []byte{}
	switch srcProtocol {
	case "local":
		body, err = ipfss3.LoadLocal(srcPath)
		if err != nil {
			return err
		}

	case "ipfs":
		body, err = ipfss3.LoadIpfs(srcPath)
		if err != nil {
			return err
		}

	default:
		fmt.Println("unsupported protocol:", srcProtocol)
	}

	dstProtocol, dstPath := ipfss3.ParsePath(dst)
	fmt.Println(dstProtocol, dstPath)
	switch dstProtocol {
	case "local":
		err = ipfss3.StoreLocal(dstPath, body)
		if err != nil {
			return err
		}

	case "ipfs":
		hash, err := ipfss3.StoreIpfs(body)
		if err != nil {
			return err
		}
		fmt.Println(hash)

	default:
		fmt.Println("unsupported protocol:", dstProtocol)
	}

	return nil
}

func sync() (err error) {
	return nil
}

func usage() {
	fmt.Println("Usage of ipfscli:")
	fmt.Println("")
	fmt.Println("   list     Show list of keys")
	fmt.Println("            --tags: with tags")
	fmt.Println("            --balance: with balance")
	fmt.Println("   new      Create key")
	fmt.Println("   add-tags [keyID] [name:value] [name:value]...")
	fmt.Println("            add tag to exist key")
}

func main() {
	var err error
	cpFlag := flag.NewFlagSet("cp", flag.ExitOnError)

	var flagRecursive bool
	cpFlag.BoolVar(&flagRecursive, "recursive", flagRecursive, "Recursive")

	if len(os.Args) == 1 {
		usage()
		return
	}

	err = cpFlag.Parse(os.Args[2:])
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	switch os.Args[1] {
	case "cp":
		src := os.Args[2]
		dst := os.Args[3]
		err = copy(ctx, src, dst)
	case "sync":
		err = sync()

	default:
		usage()
	}

	if err != nil {
		panic(err)
	}
}
