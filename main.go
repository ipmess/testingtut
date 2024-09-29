package main

import ( "fmt"
         "crypto/sha256"
		 "math/bits"
)

// Function to be tested: adds two integers.
func Add(a, b int) int {
	return a + b
}

func DeltaBits(h1, h2 [32]uint8) int {
	var bc int
	var octetbc uint8
	for i := range h1 {
		octetbc = h1[i] ^ h2[i]
        bc = bc + bits.OnesCount8(octetbc)
	}
	return bc
}

func main() {

c1 := sha256.Sum256([]byte("x"))
c2 := sha256.Sum256([]byte("X"))
fmt.Printf("SHA256 of \"x\": %x\n", c1)
fmt.Printf("The type of the SHA256 result is: %T\n", c1)
fmt.Printf("SHA256 of \"X\": %x\n", c2)
i := DeltaBits(c1, c2)
fmt.Printf("Delta Bits: %d\n", i)

}