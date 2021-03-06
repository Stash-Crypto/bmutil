// Copyright (c) 2015 Monetas
// Copyright 2016 Daniel Krawisz.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package identity_test

import (
	"reflect"
	"testing"

	"github.com/DanielKrawisz/bmutil/identity"
	"github.com/DanielKrawisz/bmutil/pow"
)

func TestNewPublic(t *testing.T) {
	privId, _ := identity.ImportWIF("BM-2cXm1jokUVp9Nn1kBtkeMjpxaLJuP3FwET",
		"5K3oNuMzVEWdrtyBAZXrPQwQTSmCGrAZS1groRDQVGDeccLim15",
		"5HzhkuimkuizxJyw9b7qnFEMtUrAXD25Y5AV1sZ964dSSXReKnb",
		pow.DefaultNonceTrialsPerByte, pow.DefaultExtraBytes)
	id := privId.ToPublic()

	testId := identity.NewPublic(privId.SigningKey.PubKey(),
		privId.EncryptionKey.PubKey(), pow.DefaultNonceTrialsPerByte,
		pow.DefaultExtraBytes, 4, 1)

	if !reflect.DeepEqual(id, testId) {
		t.Errorf("Created public identity not equal to original.")
	}
}
