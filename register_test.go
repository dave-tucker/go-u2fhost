package u2fhost

import (
	"encoding/hex"
	"testing"
)

// Expected inputs and outputs are taken from the examples from the
// specification.
// https://fidoalliance.org/specs/fido-u2f-v1.0-nfc-bt-amendment-20150514/fido-u2f-raw-message-formats.html#examples
func TestRegisterRequest(t *testing.T) {
	cidPubKey := JSONWebKey{
		Kty: "EC",
		Crv: "P-256",
		X:   "HzQwlfXX7Q4S5MtCCnZUNBw3RMzPO9tOyWjBqRl4tJ8",
		Y:   "XVguGFLIZx1fXg3wNqfdbn75hi4-_7-BxhMljw42Ht4",
	}
	clientDataHash := "4142d21c00d94ffb9d504ada8f99b721f4b191ae4e37ca0140f696b6983cfacb"
	appIdHash := "f0e6a6a97042a4f1f1c87f5f7d44315b2d852c2df5c7991cc66241bf7072d1c4"
	regRequest := &RegisterRequest{
		Challenge:  "vqrS6WXDe1JUs5_c3i4-LkKIHRr-3XVb3azuA5TifHo",
		Facet:      "http://example.com",
		AppId:      "http://example.com",
		JSONWebKey: &cidPubKey,
	}
	expectedRequest := clientDataHash + appIdHash
	expectedJson := "{\"typ\":\"navigator.id.finishEnrollment\",\"challenge\":\"vqrS6WXDe1JUs5_c3i4-LkKIHRr-3XVb3azuA5TifHo\",\"cid_pubkey\":{\"kty\":\"EC\",\"crv\":\"P-256\",\"x\":\"HzQwlfXX7Q4S5MtCCnZUNBw3RMzPO9tOyWjBqRl4tJ8\",\"y\":\"XVguGFLIZx1fXg3wNqfdbn75hi4-_7-BxhMljw42Ht4\"},\"origin\":\"http://example.com\"}"
	clientJson, request, err := registerRequest(regRequest)
	if err != nil {
		t.Errorf("Error constructing authenticate request: %s", err)
	}
	if string(clientJson) != expectedJson {
		t.Errorf("Expected client json to be %s but got %s", expectedJson, string(clientJson))
	}
	requestString := hex.EncodeToString(request)
	if requestString != expectedRequest {
		t.Errorf("Expected %s but got %s", expectedRequest, requestString)
	}
}