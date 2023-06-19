// Copyright 2023 NgSinCung <wuxiansong0125@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

// SecureServingOptions contains configuration items related to HTTPS server startup.
type SecureServingOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	// BindPort is ignored when Listener is set, will serve HTTPS even with 0.
	BindPort int `json:"bind-port"    mapstructure:"bind-port"`
	// Required set to true means that BindPort cannot be zero.
	Required bool
	// ServerCert is the TLS cert info for serving secure traffic
	ServerCert GeneratableKeyCert `json:"tls"          mapstructure:"tls"`
	// AdvertiseAddress net.IP
}

// NewSecureServingOptions creates a SecureServingOptions object with default parameters.
func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddress: "0.0.0.0",
		BindPort:    8443,
		Required:    true,
		ServerCert: GeneratableKeyCert{
			PairName:      "iam",
			CertDirectory: "/var/run/iam",
		},
	}
}

// GeneratableKeyCert contains configuration items related to certificate.
type GeneratableKeyCert struct {
	// CertKey allows setting an explicit cert/key file to use.
	CertKey CertKey `json:"cert-key" mapstructure:"cert-key"`

	// CertDirectory specifies a directory to write generated certificates to if CertFile/KeyFile aren't explicitly set.
	// PairName is used to determine the filenames within CertDirectory.
	// If CertDirectory and PairName are not set, an in-memory certificate will be generated.
	CertDirectory string `json:"cert-dir"  mapstructure:"cert-dir"`
	// PairName is the name which will be used with CertDirectory to make a cert and key filenames.
	// It becomes CertDirectory/PairName.crt and CertDirectory/PairName.key
	PairName string `json:"pair-name" mapstructure:"pair-name"`
}

// CertKey contains configuration items related to certificate.
type CertKey struct {
	// CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
	CertFile string `json:"cert-file"        mapstructure:"cert-file"`
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string `json:"private-key-file" mapstructure:"private-key-file"`
}
