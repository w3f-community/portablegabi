package wasm

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/KILTprotocol/portablegabi/pkg/credentials"
	"github.com/privacybydesign/gabi"
)

// GenKey creates the private key for the claimer
func GenKey(this js.Value, inputs []js.Value) ([]interface{}, error) {
	claimer, err := credentials.NewUser(SysParams)
	if err != nil {
		return nil, err
	}
	return []interface{}{claimer}, nil
}

// RequestAttestation creates a session object and a message which request the
// attestation of specific attributes. The second object should be send to an
// attester. This method expects as inputs the private key of the claimer, a
// json encoded string containing the claim which should be attested, the
// handshake message from the attester and the public key of the attester.
func RequestAttestation(this js.Value, inputs []js.Value) ([]interface{}, error) {
	claimer := &credentials.Claimer{}
	claim := &credentials.Claim{}
	handshakeMsg := &credentials.StartSessionMsg{}
	issuerPubKey := &gabi.PublicKey{}

	if err := json.Unmarshal([]byte(inputs[0].String()), claimer); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(inputs[1].String()), claim); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(inputs[2].String()), handshakeMsg); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(inputs[3].String()), issuerPubKey); err != nil {
		return nil, err
	}
	fmt.Printf("test %+v \n", claim)
	session, msg, err := claimer.RequestSignatureForClaim(issuerPubKey, handshakeMsg, claim)
	if err != nil {
		return nil, err
	}
	return []interface{}{
		session,
		msg,
	}, nil
}

// BuildCredential creates a credential which can be used to create proofs that
// the claimer posseses certain attributes. This method expects as input the
// private key of the claimer, the session object created by requestAttestation
// and the signature message send the attester.
func BuildCredential(this js.Value, inputs []js.Value) ([]interface{}, error) {
	claimer := &credentials.Claimer{}
	session := &credentials.UserIssuanceSession{}
	signature := &gabi.IssueSignatureMessage{}

	if err := json.Unmarshal([]byte(inputs[0].String()), claimer); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(inputs[1].String()), session); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(inputs[2].String()), signature); err != nil {
		return nil, err
	}

	credential, err := claimer.BuildAttestedClaim(signature, session)
	if err != nil {
		return nil, err
	}
	return []interface{}{credential}, nil
}

// RevealAttributes creates a proof that the claimer posseses the requested
// attributes. This method takes as input the private key of the claimer, the
// credential which contains the requested attributes, a json encoded list
// containing the requested attributes and the public key of the attester.
// It returns a proof containing the values of the requested attributes.
func RevealAttributes(this js.Value, inputs []js.Value) ([]interface{}, error) {
	claimer := &credentials.Claimer{}
	credential := &credentials.AttestedClaim{}
	request := &credentials.RequestDiscloseAttributes{}
	issuerPubKey := &gabi.PublicKey{}

	if err := json.Unmarshal([]byte(inputs[0].String()), claimer); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(inputs[1].String()), credential); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(inputs[2].String()), request); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(inputs[3].String()), issuerPubKey); err != nil {
		return nil, err
	}

	disclosedAttr, err := claimer.RevealAttributes(issuerPubKey, credential, request)
	if err != nil {
		return nil, err
	}
	return []interface{}{disclosedAttr}, nil
}
