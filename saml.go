package saml

import "github.com/RobotsAndPencils/go-saml/util"

// ServiceProviderSettings provides settings to configure server acting as a SAML Service Provider.
// Expect only one IDP per SP in this configuration. If you need to configure multipe IDPs for an SP
// then configure multiple instances of this module
type ServiceProviderSettings struct {
	PublicCertPath              string
	PrivateKeyPath              string
	IDPSSOURL                   string
	IDPSSODescriptorURL         string
	IDPPublicCertPath           string
	IDPPublicCertEncPath        string
	AssertionConsumerServiceURL string

	hasInit          bool
	publicCert       string
	privateKey       string
	iDPPublicCert    string
	iDPPublicCertEnc string
}

type IdentityProviderSettings struct {
}

func (s *ServiceProviderSettings) Init() (err error) {
	if s.hasInit {
		return nil
	}
	s.hasInit = true

	s.publicCert, err = util.LoadCertificate(s.PublicCertPath)
	if err != nil {
		panic(err)
	}

	s.privateKey, err = util.LoadCertificate(s.PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	s.iDPPublicCert, err = util.LoadCertificate(s.IDPPublicCertPath)
	if err != nil {
		panic(err)
	}

	s.iDPPublicCertEnc, err = util.LoadCertificate(s.IDPPublicCertEncPath)
	if err != nil {
		panic(err)
	}

	return nil
}

func (s *ServiceProviderSettings) PublicCert() string {
	if !s.hasInit {
		panic("Must call ServiceProviderSettings.Init() first")
	}
	return s.publicCert
}

func (s *ServiceProviderSettings) PrivateKey() string {
	if !s.hasInit {
		panic("Must call ServiceProviderSettings.Init() first")
	}
	return s.privateKey
}

func (s *ServiceProviderSettings) IDPPublicCertEnc() string {
	if !s.hasInit {
		panic("Must call ServiceProviderSettings.Init() first")
	}
	return s.iDPPublicCertEnc
}

func (s *ServiceProviderSettings) IDPPublicCert() string {
	if !s.hasInit {
		panic("Must call ServiceProviderSettings.Init() first")
	}
	return s.iDPPublicCert
}
