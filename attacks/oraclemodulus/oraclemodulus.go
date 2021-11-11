// Package oraclemodulus implements recoving an RSA modulus given the ciphertext of 2,3,4 and 9.
package oraclemodulus

import (
	"fmt"

	"github.com/sourcekris/goRsaTool/keys"
	fmp "github.com/sourcekris/goflint"
)

// name is the name of this attack.
const name = "modulus recovery via encryption oracle"

// Attack calculates an RSA modulus when we know the ciphertext of 2, 3, 4 and 9.
func Attack(ts []*keys.RSA) error {
	var (
		e2, e3, e4, e9 *fmp.Fmpz
		ok             bool
	)

	t := ts[0]
	if t.OracleCiphertexts == nil {
		return fmt.Errorf("%s failed, input requires the ciphertext of integers 2, 3, 4, and 9", name)
	}

	if len(t.OracleCiphertexts) != 4 {
		return fmt.Errorf("%s failed, this attack requires the ciphertext of integers 2, 3, 4, and 9", name)
	}

	if e2, ok = t.OracleCiphertexts[2]; !ok {
		return fmt.Errorf("%s failed, missing the ciphertext of 2", name)
	}

	if e3, ok = t.OracleCiphertexts[3]; !ok {
		return fmt.Errorf("%s failed, missing the ciphertext of 3", name)
	}

	if e4, ok = t.OracleCiphertexts[4]; !ok {
		return fmt.Errorf("%s failed, missing the ciphertext of 4", name)
	}

	if e9, ok = t.OracleCiphertexts[9]; !ok {
		return fmt.Errorf("%s failed, missing the ciphertext of 9", name)
	}

	// n = GCD(e2**2 - e4, e3**2 - e9)
	n := new(fmp.Fmpz).GCD(new(fmp.Fmpz).Sub(new(fmp.Fmpz).ExpXI(e2, 2), e4), new(fmp.Fmpz).Sub(new(fmp.Fmpz).ExpXI(e3, 2), e9))
	t.Key.PublicKey.N = new(fmp.Fmpz).Set(n)

	if t.Key.PublicKey.E != nil {
		fmt.Println("Recovered public key:")
		fmt.Println(keys.EncodeFMPPublicKey(t.Key.PublicKey))
	} else {
		fmt.Println("Recovered modulus:")
		fmt.Printf("n = %v", n)
	}

	return nil
}
