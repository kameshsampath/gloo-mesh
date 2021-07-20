package secrets

const (
	/*
		TODO(ilackarms): document the expected structure of secrets (required for VirtualMeshes  using a user-provided root CA)
	*/
	// CaCertID is the CA certificate chain file.
	CaCertID = "ca-cert.pem"
	// CaPrivateKeyID is the private key file of CA.
	CaPrivateKeyID = "ca-key.pem"
	// CertChainID is the ID/name for the certificate chain file.
	CertChainID = "cert-chain.pem"
	// RootCertID is the ID/name for the CA root certificate file.
	RootCertID = "root-cert.pem"
)

// The intermediate CA derived from the root CA of the MeshGroup
type CAData struct {
	RootCert     []byte
	CertChain    []byte
	CaCert       []byte
	CaPrivateKey []byte
}

func (d CAData) ToSecretData() map[string][]byte {
	return map[string][]byte{
		CertChainID:    d.CertChain,
		RootCertID:     d.RootCert,
		CaCertID:       d.CaCert,
		CaPrivateKeyID: d.CaPrivateKey,
	}
}

func CADataFromSecretData(data map[string][]byte) CAData {
	caKey := data[CaPrivateKeyID]
	caCert := data[CaCertID]
	certChain := data[CertChainID]
	rootCert := data[RootCertID]
	return CAData{
		RootCert:     rootCert,
		CertChain:    certChain,
		CaCert:       caCert,
		CaPrivateKey: caKey,
	}
}
