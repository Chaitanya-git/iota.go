package curl_examples_test

import (
	"fmt"
	"github.com/iotaledger/iota.go/curl"
	"github.com/iotaledger/iota.go/trinary"
)

// o: *Curl, The new Curl instance.
func ExampleNewCurl() {}

// i req: length, The length of the trits to squeeze out.
// o: Trits, The Trits representation of the hash.
// o: error, Returned for invalid lengths.
func ExampleSqueeze() {}

// i req: length, The length of the trits to squeeze out.
// o: Trits, The Trits representation of the hash.
func ExampleMustSqueeze() {}

// i req: length, The length of the trits to squeeze out.
// o: Trytes, The Trytes representation of the hash.
// o: error, Returned for invalid lengths.
func ExampleSqueezeTrytes() {}

// i req: length, The length of the trits to squeeze out.
// o: Trytes, The Trytes representation of the hash.
func ExampleMustSqueezeTrytes() {}

// i req: in, The Trits to absorb.
// o: error, Returned for internal errors.
func ExampleAbsorb() {}

// i req: inn, The Trytes to absorb.
// o: error, Returned for internal errors.
func ExampleAbsorbTrytes() {}

// i req: inn, The Trytes to absorb.
func ExampleMustAbsorbTrytes() {}

func ExampleTransform() {}

func ExampleReset() {}

// i req: trits, The Trits of which to compute the hash of.
// o: Trits, The Trits representation of the hash.
// o: error, Returned for internal errors.
func ExampleHashTrits() {
	trytes := "PDFIDVWRXONZSPJJQVZVVMLGSVB"
	trits := trinary.MustTrytesToTrits(trytes)
	tritsHash, err := curl.HashTrits(trits)
	if err != nil {
		// handle error
		return
	}
	fmt.Println(tritsHash) // output: [0 1 -1 0 -1 0 -1 1 ...]
}

// i req: t, The Trytes of which to compute the hash of.
// o: Trytes, The Trytes representation of the hash.
// o: error, Returned for internal errors.
func ExampleHashTrytes() {
	trytes := "PDFIDVWRXONZSPJJQVZVVMLGSVB"
	hash, err := curl.HashTrytes(trytes)
	if err != nil {
		// handle error
		return
	}
	fmt.Println(hash) // output: UXBXSI9LHCPYFFZXOWALCBTUIVXYKMCEDDIFXXGXJ9ZLEWKOTXSGYHPEAD9SXSRAWM9TPPXWZMZSIEKGX
}

// i req: t, The Trytes of which to compute the hash of.
// o: Trytes, The Trytes representation of the hash.
func ExampleMustHashTrytes() {}
