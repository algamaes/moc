// Copyright (c) Microsoft Corporation.
// Licensed under the Apache v2.0 license.
package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/microsoft/moc/pkg/certs"
	"github.com/microsoft/moc/pkg/errors"
)

func init() {
	os.MkdirAll("/tmp/auth", os.ModePerm)
	key, _ = rsa.GenerateKey(rand.Reader, 2048)
}

func Test_GetCertRenewRequired(t *testing.T) {
	now := time.Now().UTC()

	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName: "test",
		},
		NotBefore: now,
		NotAfter:  now.Add(time.Second * 10),
	}

	b, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		t.Errorf("Failed creating certificate %v", err)
	}

	x509Cert, err := x509.ParseCertificate(b)
	if err != nil {
		t.Errorf("Failed parsing certificate %v", err)
	}
	if renewRequired(x509Cert) {
		t.Errorf("RenewRequired Expected:false Actual:true")
	}
}

func Test_GetCertRenewRequiredExpired(t *testing.T) {
	now := time.Now().UTC()

	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName: "test",
		},
		NotBefore: now.Add(-(time.Second * 10)),
		NotAfter:  now,
	}

	b, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		t.Errorf("Failed creating certificate %v", err)
	}

	x509Cert, err := x509.ParseCertificate(b)
	if err != nil {
		t.Errorf("Failed parsing certificate %v", err)
	}
	if !renewRequired(x509Cert) {
		t.Errorf("RenewRequired Expected:true Actual:false")
	}
}

func Test_GetCertRenewRequiredBeforeThreshold(t *testing.T) {
	now := time.Now().UTC()

	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName: "test",
		},
		NotBefore: now.Add(-(time.Second * 6)),
		NotAfter:  now.Add(time.Second * 4),
	}

	b, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		t.Errorf("Failed creating certificate %v", err)
	}

	x509Cert, err := x509.ParseCertificate(b)
	if err != nil {
		t.Errorf("Failed parsing certificate %v", err)
	}
	if renewRequired(x509Cert) {
		t.Errorf("RenewRequired Expected:false Actual:true")
	}
}

func Test_GetCertRenewRequiredAfterThreshold(t *testing.T) {
	now := time.Now().UTC()

	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName: "test",
		},
		NotBefore: now.Add(-(time.Second * 8)),
		NotAfter:  now.Add(time.Second * 2),
	}

	b, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		t.Errorf("Failed creating certificate %v", err)
	}

	x509Cert, err := x509.ParseCertificate(b)
	if err != nil {
		t.Errorf("Failed parsing certificate %v", err)
	}
	if !renewRequired(x509Cert) {
		t.Errorf("RenewRequired Expected:true Actual:false")
	}
}

func Test_GetCertRenewRequiredDelay(t *testing.T) {
	now := time.Now().UTC()

	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName: "test",
		},
		NotBefore: now.Add(-(time.Second * 6)),
		NotAfter:  now.Add(time.Second * 4),
	}

	b, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		t.Errorf("Failed creating certificate %v", err)
	}

	x509Cert, err := x509.ParseCertificate(b)
	if err != nil {
		t.Errorf("Failed parsing certificate %v", err)
	}
	if renewRequired(x509Cert) {
		t.Errorf("RenewRequired Expected:false Actual:true")
	}
	time.Sleep(time.Second * 2)
	if !renewRequired(x509Cert) {
		t.Errorf("RenewRequired Expected:true Actual:false")
	}
}

func Test_CertCheckNotExpired(t *testing.T) {
	now := time.Now().UTC()

	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName: "test",
		},
		NotBefore: now,
		NotAfter:  now.Add(time.Second * 30),
	}

	b, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		t.Errorf("Failed creating certificate %v", err)
	}

	x509Cert, err := x509.ParseCertificate(b)
	if err != nil {
		t.Errorf("Failed parsing certificate %v", err)
	}

	certPem := certs.EncodeCertPEM(x509Cert)
	if err = certCheck(certPem); err != nil {
		if errors.IsExpired(err) {
			t.Errorf("certCheck return certificate expired %v: Expected Valid Certificate", err)
		} else {
			t.Errorf("certCheck Expected:nil Actual:%v", err)
		}
	}
}

func Test_CertCheckExpired(t *testing.T) {
	now := time.Now().UTC()

	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName: "test",
		},
		NotBefore: now.Add(time.Second * -10),
		NotAfter:  now.Add(time.Second * -1),
	}

	b, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, key.Public(), key)
	if err != nil {
		t.Errorf("Failed creating certificate %v", err)
	}

	x509Cert, err := x509.ParseCertificate(b)
	if err != nil {
		t.Errorf("Failed parsing certificate %v", err)
	}

	certPem := certs.EncodeCertPEM(x509Cert)
	if err = certCheck(certPem); err == nil || !errors.IsExpired(err) {
		t.Errorf("certCheck Expected:Expired Actual:%v", err)
	}
}

func Test_CertCheckEmpty(t *testing.T) {

	if err := certCheck([]byte{}); err == nil || !errors.IsInvalidInput(err) {
		t.Errorf("certCheck Expected:InvalidInput Actual:%v", err)
	}
}
