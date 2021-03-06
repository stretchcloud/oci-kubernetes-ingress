package config

import (
	"testing"
)

var validConfig = `
auth:
  region: uk-london-1
  tenancy: ocid1.tenancy.oc1.
  user: ocid1.user.oc1.
  key: |
    -----BEGIN RSA PRIVATE KEY-----
    -----END RSA PRIVATE KEY-----
  fingerprint: a4:bb:34:...
useInstancePrincipals: false
loadbalancer:
  # Specifies which compartment to create a load balancer in (by default your cluster compartment should be set)
  compartment: ocid1.compartment.oc1..aaaaaaaaob4ckouj3cjmf36ifjkff33wvln5fnnarumafqzpqq7tmbig2n5q
  # Controls the subnets in which load balancers are provisioned.
  # A minimum of two must be declared.
  subnets:
    - ocid1.subnet.oc1.uk-london-1.
    - ocid1.subnet.oc1.uk-london-1.
`

func TestParseConfigInvalidConfig(t *testing.T) {
	_, err := Parse([]byte("Invalid config"))
	if err == nil {
		t.Fatal("expected error when given invalid config")
	}
}
func TestParseConfigValidConfig(t *testing.T) {
	_, err := Parse([]byte(validConfig))
	if err != nil {
		t.Fatalf("expected no error but got '%s'", err)
	}
}
