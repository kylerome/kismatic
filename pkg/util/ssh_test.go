package util

import "testing"

func TestIsEncrypted(t *testing.T) {
	for _, data := range testData {
		if isEncrypted := isEncrypted(data.pemData); isEncrypted != data.ecrypted {
			t.Errorf("Expected: %t, got: %t", data.ecrypted, isEncrypted)
		}
	}
}

var testData = []struct {
	ecrypted bool
	pemData  []byte
}{
	{
		ecrypted: false,
		pemData: []byte(`
-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`),
	},
	{
		ecrypted: true,
		pemData: []byte(`
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-128-CBC,D0FFFDD404B4871ABBAE43ACE8F391C2

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
-----END RSA PRIVATE KEY-----`),
	},
}
