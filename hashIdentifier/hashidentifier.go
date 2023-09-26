package main

import (
	"fmt"
	"encoding/hex"
)

func main() {
	fmt.Print("Enter the hash value: ")
	var hashInput string
	fmt.Scanln(&hashInput)

	if isMD5(hashInput) {
		fmt.Println("This hash is generated using MD5.")
	} else if isSHA1(hashInput) {
		fmt.Println("This hash is generated using SHA-1.")
	} else if isSHA256(hashInput) {
		fmt.Println("This hash is generated using SHA-256.")
	} else if isSHA512(hashInput) {
		fmt.Println("This hash is generated using SHA-512.")
	} else if isSHA3_256(hashInput) {
		fmt.Println("This hash is generated using SHA3-256.")
	} else if isSHA3_512(hashInput) {
		fmt.Println("This hash is generated using SHA3-512.")
	} else if isBlake2b512(hashInput) {
		fmt.Println("This hash is generated using Blake2b-512.")
	} else if isBlake2s256(hashInput) {
		fmt.Println("This hash is generated using Blake2s-256.")
	} else if isRIPEMD160(hashInput) {
		fmt.Println("This hash is generated using RIPEMD-160.")
	} else if isWhirlpool(hashInput) {
		fmt.Println("This hash is generated using Whirlpool.")
	} else if isTiger192(hashInput) {
		fmt.Println("This hash is generated using Tiger-192.")
	} else if isGOST3411_94(hashInput) {
		fmt.Println("This hash is generated using GOST3411-94.")
	} else if isSkein256_128(hashInput) {
		fmt.Println("This hash is generated using Skein-256-128.")
	} else if isSkein512_128(hashInput) {
		fmt.Println("This hash is generated using Skein-512-128.")
	} else if isSkein512_256(hashInput) {
		fmt.Println("This hash is generated using Skein-512-256.")
	} else if isSkein512_512(hashInput) {
		fmt.Println("This hash is generated using Skein-512-512.")
	} else if isSHA224(hashInput) {
		fmt.Println("This hash is generated using SHA-224.")
	} else if isSHA384(hashInput) {
		fmt.Println("This hash is generated using SHA-384.")
	} else if isKeccak256(hashInput) {
		fmt.Println("This hash is generated using Keccak-256.")
	} else if isKeccak512(hashInput) {
		fmt.Println("This hash is generated using Keccak-512.")
	} else if isBLAKE3_256(hashInput) {
		fmt.Println("This hash is generated using BLAKE3-256.")
	} else {
		fmt.Println("This hash does not match any defined algorithm.")
	}
}

func isMD5(input string) bool {
	return len(input) == 32 && isHex(input)
}

func isSHA1(input string) bool {
	return len(input) == 40 && isHex(input)
}

func isSHA256(input string) bool {
	return len(input) == 64 && isHex(input)
}

func isSHA512(input string) bool {
	return len(input) == 128 && isHex(input)
}

func isSHA3_256(input string) bool {
	return len(input) == 64 && isHex(input)
}

func isSHA3_512(input string) bool {
	return len(input) == 128 && isHex(input)
}

func isBlake2b512(input string) bool {
	return len(input) == 128 && isHex(input)
}

func isBlake2s256(input string) bool {
	return len(input) == 64 && isHex(input)
}

func isRIPEMD160(input string) bool {
	return len(input) == 40 && isHex(input)
}

func isWhirlpool(input string) bool {
	return len(input) == 128 && isHex(input)
}

func isTiger192(input string) bool {
	return len(input) == 48 && isHex(input)
}

func isGOST3411_94(input string) bool {
	return len(input) == 64 && isHex(input)
}

func isSkein256_128(input string) bool {
	return len(input) == 32 && isHex(input)
}

func isSkein512_128(input string) bool {
	return len(input) == 32 && isHex(input)
}

func isSkein512_256(input string) bool {
	return len(input) == 64 && isHex(input)
}

func isSkein512_512(input string) bool {
	return len(input) == 128 && isHex(input)
}

func isSHA224(input string) bool {
	return len(input) == 56 && isHex(input)
}

func isSHA384(input string) bool {
	return len(input) == 96 && isHex(input)
}

func isKeccak256(input string) bool {
	return len(input) == 64 && isHex(input)
}

func isKeccak512(input string) bool {
	return len(input) == 128 && isHex(input)
}

func isBLAKE3_256(input string) bool {
	return len(input) == 64 && isHex(input)
}

func isHex(input string) bool {
	_, err := hex.DecodeString(input)
	return err == nil
}
