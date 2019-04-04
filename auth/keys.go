package auth

import (
	"cloud.google.com/go/iam"
	cloudkms "cloud.google.com/go/kms/apiv1"
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"google.golang.org/api/option"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	fieldmask "google.golang.org/genproto/protobuf/field_mask"
	"log"
	"math/big"
)

type Keys struct {
	cli *cloudkms.KeyManagementClient
}

func NewKeys(ctx context.Context, opts ...option.ClientOption) (*Keys, error) {
	client, err := cloudkms.NewKeyManagementClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &Keys{
		cli: client,
	}, nil
}

func (v *Keys) Close() {
	_ = v.cli.Close()
}

func (v *Keys) Client() *cloudkms.KeyManagementClient {
	return v.cli
}

// CreateKeyRing creates a new ring to store keys on KMS.
// example parentName: "projects/PROJECT_ID/locations/global/"
func (v *Keys) CreateKeyRing(ctx context.Context, parentName, keyRingId string) (*kmspb.KeyRing, error) {
	// Build the request.
	req := &kmspb.CreateKeyRingRequest{
		Parent:    parentName,
		KeyRingId: keyRingId,
	}
	// Call the API.
	return v.cli.CreateKeyRing(ctx, req)
}

// CreateCryptoKey creates a new symmetric encrypt/decrypt key on KMS.
// example keyRingName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID"
func (v *Keys) CreateCryptoKey(ctx context.Context, keyRingName, keyId string) (*kmspb.CryptoKey, error) {
	// Build the request.
	req := &kmspb.CreateCryptoKeyRequest{
		Parent:      keyRingName,
		CryptoKeyId: keyId,
		CryptoKey: &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ENCRYPT_DECRYPT,
			VersionTemplate: &kmspb.CryptoKeyVersionTemplate{
				Algorithm: kmspb.CryptoKeyVersion_GOOGLE_SYMMETRIC_ENCRYPTION,
			},
		},
	}
	// Call the API.
	return v.cli.CreateCryptoKey(ctx, req)
}

// DisableCryptoKeyVersion disables a specified key version on KMS.
// example keyVersionName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) DisableCryptoKeyVersion(ctx context.Context, keyVersionName string) error {
	// Build the request.
	req := &kmspb.UpdateCryptoKeyVersionRequest{
		CryptoKeyVersion: &kmspb.CryptoKeyVersion{
			Name:  keyVersionName,
			State: kmspb.CryptoKeyVersion_DISABLED,
		},
		UpdateMask: &fieldmask.FieldMask{
			Paths: []string{"state"},
		},
	}
	// Call the API.
	result, err := v.cli.UpdateCryptoKeyVersion(ctx, req)
	if err != nil {
		return err
	}
	log.Printf("Disabled crypto key version: %s", result)
	return nil
}

// EnableCryptoKeyVersion enables a previously disabled key version on KMS.
// example keyVersionName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) EnableCryptoKeyVersion(ctx context.Context, keyVersionName string) error {
	// Build the request.
	req := &kmspb.UpdateCryptoKeyVersionRequest{
		CryptoKeyVersion: &kmspb.CryptoKeyVersion{
			Name:  keyVersionName,
			State: kmspb.CryptoKeyVersion_ENABLED,
		},
		UpdateMask: &fieldmask.FieldMask{
			Paths: []string{"state"},
		},
	}
	// Call the API.
	result, err := v.cli.UpdateCryptoKeyVersion(ctx, req)
	if err != nil {
		return err
	}
	log.Printf("Enabled crypto key version: %s", result)
	return nil
}

// DestroyCryptoKeyVersion marks a specified key version for deletion. The key can be restored if requested within 24 hours.
// example keyVersionName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) DestroyCryptoKeyVersion(ctx context.Context, keyVersionName string) error {
	// Build the request.
	req := &kmspb.DestroyCryptoKeyVersionRequest{
		Name: keyVersionName,
	}
	// Call the API.
	result, err := v.cli.DestroyCryptoKeyVersion(ctx, req)
	if err != nil {
		return err
	}
	log.Printf("Destroyed crypto key version: %s", result)
	return nil
}

// RestoreCryptoKeyVersion attempts to recover a key that has been marked for destruction within the last 24 hours.
// example keyVersionName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) RestoreCryptoKeyVersion(ctx context.Context, keyVersionName string) error {
	// Build the request.
	req := &kmspb.RestoreCryptoKeyVersionRequest{
		Name: keyVersionName,
	}
	// Call the API.
	result, err := v.cli.RestoreCryptoKeyVersion(ctx, req)
	if err != nil {
		return err
	}
	log.Printf("Restored crypto key version: %s", result)
	return nil
}

// GetRingPolicy retrieves and prints the IAM policy associated with the key ring
// example keyRingName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID"
func (v *Keys) GetRingPolicy(ctx context.Context, keyRingName string) (*iam.Policy, error) {
	// Get the KeyRing.
	keyRingObj, err := v.cli.GetKeyRing(ctx, &kmspb.GetKeyRingRequest{Name: keyRingName})
	if err != nil {
		return nil, err
	}
	// Get IAM Policy.
	handle := v.cli.KeyRingIAM(keyRingObj)
	policy, err := handle.Policy(ctx)
	if err != nil {
		return nil, err
	}
	for _, role := range policy.Roles() {
		for _, member := range policy.Members(role) {
			log.Printf("Role: %s Member: %s\n", role, member)
		}
	}
	return policy, nil
}

// GetCryptoKeyPolicy retrieves and prints the IAM policy associated with the key
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
func (v *Keys) GetCryptoKeyPolicy(ctx context.Context, keyName string) (*iam.Policy, error) {
	// Get the KeyRing.
	keyObj, err := v.cli.GetCryptoKey(ctx, &kmspb.GetCryptoKeyRequest{Name: keyName})
	if err != nil {
		return nil, err
	}
	// Get IAM Policy.
	handle := v.cli.CryptoKeyIAM(keyObj)
	policy, err := handle.Policy(ctx)
	if err != nil {
		return nil, err
	}
	for _, role := range policy.Roles() {
		for _, member := range policy.Members(role) {
			log.Printf("Role: %s Member: %s\n", role, member)
		}
	}
	return policy, nil
}

// AddMemberRingPolicy adds a new member to a specified IAM role for the key ring
// example keyRingName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID"
func (v *Keys) AddMemberRingPolicy(ctx context.Context, keyRingName, member string, role iam.RoleName) error {
	// Get the KeyRing.
	keyRingObj, err := v.cli.GetKeyRing(ctx, &kmspb.GetKeyRingRequest{Name: keyRingName})
	if err != nil {
		return err
	}
	// Get IAM Policy.
	handle := v.cli.KeyRingIAM(keyRingObj)
	policy, err := handle.Policy(ctx)
	if err != nil {
		return err
	}
	// Add Member.
	policy.Add(member, role)
	err = handle.SetPolicy(ctx, policy)
	if err != nil {
		return err
	}
	log.Print("Added member to keyring policy.")
	return nil
}

// RemoveMemberRingPolicy removes a specified member from an IAM role for the key ring
// example keyRingName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID"
func (v *Keys) RemoveMemberRingPolicy(ctx context.Context, keyRingName, member string, role iam.RoleName) error {
	// Get the KeyRing.
	keyRingObj, err := v.cli.GetKeyRing(ctx, &kmspb.GetKeyRingRequest{Name: keyRingName})
	if err != nil {
		return err
	}
	// Get IAM Policy.
	handle := v.cli.KeyRingIAM(keyRingObj)
	policy, err := handle.Policy(ctx)
	if err != nil {
		return err
	}

	// Remove Member.
	policy.Remove(member, role)
	err = handle.SetPolicy(ctx, policy)
	if err != nil {
		return err
	}
	log.Print("Removed member from keyring policy.")
	return nil
}

// AddMemberCryptoKeyPolicy adds a new member to a specified IAM role for the key
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
func (v *Keys) AddMemberCryptoKeyPolicy(ctx context.Context, keyName, member string, role iam.RoleName) error {
	// Get the desired CryptoKey.
	keyObj, err := v.cli.GetCryptoKey(ctx, &kmspb.GetCryptoKeyRequest{Name: keyName})
	if err != nil {
		return err
	}
	// Get IAM Policy.
	handle := v.cli.CryptoKeyIAM(keyObj)
	policy, err := handle.Policy(ctx)
	if err != nil {
		return err
	}
	// Add Member.
	policy.Add(member, role)
	err = handle.SetPolicy(ctx, policy)
	if err != nil {
		return err
	}
	log.Print("Added member to cryptokey policy.")
	return nil
}

// RemoveMemberCryptoKeyPolicy removes a specified member from an IAM role for the key
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
func (v *Keys) RemoveMemberCryptoKeyPolicy(ctx context.Context, keyName, member string, role iam.RoleName) error {
	// Get the desired CryptoKey.
	keyObj, err := v.cli.GetCryptoKey(ctx, &kmspb.GetCryptoKeyRequest{Name: keyName})
	if err != nil {
		return err
	}
	// Get IAM Policy.
	handle := v.cli.CryptoKeyIAM(keyObj)
	policy, err := handle.Policy(ctx)
	if err != nil {
		return err
	}
	// Remove Member.
	policy.Remove(member, role)
	err = handle.SetPolicy(ctx, policy)
	if err != nil {
		return err
	}
	log.Print("Removed member from cryptokey policy.")
	return nil
}

// Encrypt will encrypt the input plaintext with the specified symmetric key
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
func (v *Keys) EncryptSymmetric(ctx context.Context, keyName string, plaintext []byte) ([]byte, error) {
	// Build the request.
	req := &kmspb.EncryptRequest{
		Name:      keyName,
		Plaintext: plaintext,
	}
	// Call the API.
	resp, err := v.cli.Encrypt(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Ciphertext, nil
}

// Decrypt will decrypt the input ciphertext bytes using the specified symmetric key
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"
func (v *Keys) DecryptSymmetric(ctx context.Context, keyName string, ciphertext []byte) ([]byte, error) {
	req := &kmspb.DecryptRequest{
		Name:       keyName,
		Ciphertext: ciphertext,
	}
	// Call the API.
	resp, err := v.cli.Decrypt(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Plaintext, nil
}

// CreateAsymmetricKey creates a new RSA encrypt/decrypt key pair on KMS.
// example keyRingName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID"
func (v *Keys) CreateAsymmetricKey(ctx context.Context, keyRingName, keyId string) error {

	// Build the request.
	req := &kmspb.CreateCryptoKeyRequest{
		Parent:      keyRingName,
		CryptoKeyId: keyId,
		CryptoKey: &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ASYMMETRIC_DECRYPT,
			VersionTemplate: &kmspb.CryptoKeyVersionTemplate{
				Algorithm: kmspb.CryptoKeyVersion_RSA_DECRYPT_OAEP_2048_SHA256,
			},
		},
	}
	// Call the API.
	result, err := v.cli.CreateCryptoKey(ctx, req)
	if err != nil {
		return err
	}
	log.Printf("Created crypto key. %s", result)
	return nil
}

// GetAsymmetricPublicKey retrieves the public key from a saved asymmetric key pair on KMS.
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) GetAsymmetricPublicKey(ctx context.Context, keyName string) (interface{}, error) {
	// Build the request.
	req := &kmspb.GetPublicKeyRequest{
		Name: keyName,
	}
	// Call the API.
	response, err := v.cli.GetPublicKey(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch public key: %+v", err)
	}
	// Parse the key.
	keyBytes := []byte(response.Pem)
	block, _ := pem.Decode(keyBytes)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %+v", err)
	}
	return publicKey, nil
}

// DecryptRSA will attempt to decrypt a given ciphertext with an 'RSA_DECRYPT_OAEP_2048_SHA256' private key.stored on Cloud KMS
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) DecryptRSA(ctx context.Context, keyName string, ciphertext []byte) ([]byte, error) {

	// Build the request.
	req := &kmspb.AsymmetricDecryptRequest{
		Name:       keyName,
		Ciphertext: ciphertext,
	}
	// Call the API.
	response, err := v.cli.AsymmetricDecrypt(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("decryption request failed: %+v", err)
	}
	return response.Plaintext, nil
}

// [END kms_decrypt_rsa]

// [START kms_encrypt_rsa]

// encryptRSA will encrypt data locally using an 'RSA_DECRYPT_OAEP_2048_SHA256' public key retrieved from Cloud KMS
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) EncryptRSA(ctx context.Context, keyName string, plaintext []byte) ([]byte, error) {

	// Retrieve the public key from KMS.
	response, err := v.cli.GetPublicKey(ctx, &kmspb.GetPublicKeyRequest{Name: keyName})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch public key: %+v", err)
	}
	// Parse the key.
	block, _ := pem.Decode([]byte(response.Pem))
	abstractKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %+v", err)
	}
	rsaKey, ok := abstractKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key '%s' is not RSA", keyName)
	}
	// Encrypt data using the RSA public key.
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaKey, plaintext, nil)
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %+v", err)
	}
	return ciphertext, nil
}

// SignAsymmetric will sign a plaintext message using a saved asymmetric private key.
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) SignAsymmetric(ctx context.Context, keyName string, message []byte) ([]byte, error) {
	// Find the digest of the message.
	digest := sha256.New()
	digest.Write(message)
	// Build the signing request.
	req := &kmspb.AsymmetricSignRequest{
		Name: keyName,
		Digest: &kmspb.Digest{
			Digest: &kmspb.Digest_Sha256{
				Sha256: digest.Sum(nil),
			},
		},
	}
	// Call the API.
	response, err := v.cli.AsymmetricSign(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("asymmetric sign request failed: %+v", err)
	}
	return response.Signature, nil
}

// VerifySignatureRSA will verify that an 'RSA_SIGN_PSS_2048_SHA256' signature is valid for a given message.
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) VerifySignatureRSA(ctx context.Context, keyName string, signature, message []byte) error {

	// Retrieve the public key from KMS.
	response, err := v.cli.GetPublicKey(ctx, &kmspb.GetPublicKeyRequest{Name: keyName})
	if err != nil {
		return fmt.Errorf("failed to fetch public key: %+v", err)
	}
	// Parse the key.
	block, _ := pem.Decode([]byte(response.Pem))
	abstractKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse public key: %+v", err)
	}
	rsaKey, ok := abstractKey.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("key '%s' is not RSA", keyName)
	}
	// Verify RSA signature.
	hash := sha256.New()
	hash.Write(message)
	digest := hash.Sum(nil)
	pssOptions := rsa.PSSOptions{SaltLength: len(digest), Hash: crypto.SHA256}
	err = rsa.VerifyPSS(rsaKey, crypto.SHA256, digest, signature, &pssOptions)
	if err != nil {
		return fmt.Errorf("signature verification failed: %+v", err)
	}
	return nil
}

// VerifySignatureEC will verify that an 'EC_SIGN_P256_SHA256' signature is valid for a given message.
// example keyName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
func (v *Keys) VerifySignatureEC(ctx context.Context, keyName string, signature, message []byte) error {
	// Retrieve the public key from KMS.
	response, err := v.cli.GetPublicKey(ctx, &kmspb.GetPublicKeyRequest{Name: keyName})
	if err != nil {
		return fmt.Errorf("failed to fetch public key: %+v", err)
	}
	// Parse the key.
	block, _ := pem.Decode([]byte(response.Pem))
	abstractKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse public key: %+v", err)
	}
	ecKey, ok := abstractKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("key '%s' is not EC", keyName)
	}
	// Verify Elliptic Curve signature.
	var parsedSig struct{ R, S *big.Int }
	_, err = asn1.Unmarshal(signature, &parsedSig)
	if err != nil {
		return fmt.Errorf("failed to parse signature bytes: %+v", err)
	}
	hash := sha256.New()
	hash.Write(message)
	digest := hash.Sum(nil)
	if !ecdsa.Verify(ecKey, digest, parsedSig.R, parsedSig.S) {
		return errors.New("signature verification failed")
	}
	return nil
}
