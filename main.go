package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
	"hash"
	"os"
	"strings"
)

type fn func([]string) error

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: ch <op> [arguments...] <algorithm>")
		os.Exit(1)
	}
	op := os.Args[1]
	if op == "list" {
		algorithms := []string{"MD4", "MD5", "SHA1", "SHA224", "SHA256", "SHA384", "SHA512", "SHA3-224", "SHA3-256",
			"SHA3-384", "SHA3-512"}
		fmt.Println("Algorithms:")
		for _, algorithm := range algorithms {
			fmt.Printf("	%s\n", algorithm)
		}
		os.Exit(0)
	}
	opTypes := make(map[string]fn, 4)
	opTypes["check"] = check
	opTypes["checks"] = checks
	opTypes["create"] = create
	opTypes["creates"] = creates
	if _, ok := opTypes[op]; !ok {
		fmt.Printf("unknown op: %s\n", op)
		fmt.Println("possible values:")
		fmt.Println("	check	- check a file against a hash")
		fmt.Println("	checks	- check a string against a hash")
		fmt.Println("	create	- create a hash from a file")
		fmt.Println("	creates	- create a hash from a string")
		os.Exit(1)
	}
	opFunction := opTypes[op]
	err := opFunction(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func check(args []string) error {
	if len(args) != 4 {
		return errors.New("usage: ch check <file> <hash> <algorithm>")
	}
	file := args[1]
	exists, err := checkIfPathIsValid(file)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("file %s does not exist", file)
	}
	fileContent, err := readFile(file)
	if err != nil {
		return err
	}
	err = checkHash(fileContent, args[2], args[3])
	if err != nil {
		return err
	}
	return nil
}

func checks(args []string) error {
	if len(args) != 4 {
		return errors.New("usage: ch checks <string> <hash> <algorithm>")
	}
	err := checkHash([]byte(args[1]), args[2], args[3])
	if err != nil {
		return err
	}
	return nil
}

func create(args []string) error {
	if len(args) != 3 {
		return errors.New("usage: ch create <file> <algorithm>")
	}
	file := args[1]
	exists, err := checkIfPathIsValid(file)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("file %s does not exist", file)
	}
	fileContent, err := readFile(file)
	if err != nil {
		return err
	}
	err = createHash(fileContent, &file, args[2], true)
	if err != nil {
		return err
	}
	return nil
}

func creates(args []string) error {
	if len(args) != 3 {
		return errors.New("usage: ch creates <string> <algorithm>")
	}
	err := createHash([]byte(args[1]), nil, args[2], false)
	if err != nil {
		return err
	}
	return nil
}

func checkIfPathIsValid(path string) (bool, error) {
	if fileInfo, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		if fileInfo.IsDir() {
			return false, fmt.Errorf("path %s is a directory, not a file", path)
		}
		return true, nil
	}
}

func readFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func checkHash(content []byte, hash, algorithm string) error {
	sum, err := hashBytes(content, algorithm)
	if err != nil {
		return err
	}
	if hash != *sum {
		color.Red("Not ok!")
	} else {
		color.Green("Ok!")
	}
	return nil
}

func createHash(content []byte, file *string, algorithm string, isFile bool) error {
	sum, err := hashBytes(content, algorithm)
	if err != nil {
		return err
	}
	if isFile {
		fmt.Printf("%s of %s: %s\n", color.CyanString(strings.ToUpper(algorithm)), *file, color.GreenString(*sum))
	} else {
		fmt.Printf("%s of \"%s\": %s\n", color.CyanString(strings.ToUpper(algorithm)), string(content), color.GreenString(*sum))
	}
	return nil
}

func hashBytes(content []byte, algorithm string) (*string, error) {
	algorithmLower := strings.ToLower(algorithm)
	var hashCtx hash.Hash
	switch algorithmLower {
	case "md4":
		hashCtx = md4.New()
	case "md5":
		hashCtx = md5.New()
	case "sha1":
		hashCtx = sha1.New()
	case "sha224":
		hashCtx = sha256.New224()
	case "sha256":
		hashCtx = sha256.New()
	case "sha384":
		hashCtx = sha512.New384()
	case "sha512":
		hashCtx = sha512.New()
	case "sha3-224":
		hashCtx = sha3.New224()
	case "sha3-256":
		hashCtx = sha3.New256()
	case "sha3-384":
		hashCtx = sha3.New384()
	case "sha3-512":
		hashCtx = sha3.New512()
	default:
		return nil, fmt.Errorf("unknown algorithm %s", algorithm)
	}
	_, err := hashCtx.Write(content)
	if err != nil {
		return nil, err
	}
	sum := hex.EncodeToString(hashCtx.Sum(nil))
	return &sum, nil
}
