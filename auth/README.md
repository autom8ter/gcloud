# auth
--
    import "github.com/autom8ter/gcloud/auth"


## Usage

#### type Auth

```go
type Auth struct {
}
```


#### func  New

```go
func New(ctx context.Context, opts ...option.ClientOption) (*Auth, error)
```

#### func (*Auth) Close

```go
func (a *Auth) Close()
```

#### func (*Auth) IAM

```go
func (a *Auth) IAM() *IAM
```

#### func (*Auth) Keys

```go
func (a *Auth) Keys() *Keys
```

#### type IAM

```go
type IAM struct {
}
```


#### func  NewIAM

```go
func NewIAM(ctx context.Context, opts ...option.ClientOption) (*IAM, error)
```

#### func (*IAM) AueryTestablePermissions

```go
func (i *IAM) AueryTestablePermissions(fullResourceName string) ([]*iam.Permission, error)
```
queryTestablePermissions lists testable permissions on a resource.

#### func (*IAM) Client

```go
func (i *IAM) Client() *iam.Service
```

#### func (*IAM) Close

```go
func (i *IAM) Close()
```
No-OP

#### func (*IAM) CreateKey

```go
func (i *IAM) CreateKey(serviceAccountEmail string) (*iam.ServiceAccountKey, error)
```
CreateKey creates a service account key.

#### func (*IAM) CreateRole

```go
func (i *IAM) CreateRole(projectID, name, title, description, stage, roleId string, permissions []string) (*iam.Role, error)
```
CreateRole creates a custom role.

#### func (*IAM) DeleteKey

```go
func (i *IAM) DeleteKey(fullKeyName string) error
```
DeleteKey deletes a service account key.

#### func (*IAM) DeleteRole

```go
func (i *IAM) DeleteRole(projectID, name string) error
```
DeleteRole deletes a custom role.

#### func (*IAM) EditRole

```go
func (i *IAM) EditRole(projectID, name, newTitle, newDescription, newStage string, newPermissions []string) (*iam.Role, error)
```
EditRole modifies a custom role.

#### func (*IAM) GetRole

```go
func (i *IAM) GetRole(name string) (*iam.Role, error)
```

#### func (*IAM) ListKeys

```go
func (i *IAM) ListKeys(serviceAccountEmail string) ([]*iam.ServiceAccountKey, error)
```
ListKey lists a service account's keys.

#### func (*IAM) ListServiceAccounts

```go
func (i *IAM) ListServiceAccounts(projectID string) (*iam.ListServiceAccountsResponse, error)
```
ListServiceAccounts lists a project's service accounts.

#### func (*IAM) ViewGrantableRoles

```go
func (i *IAM) ViewGrantableRoles(fullResourceName string) ([]*iam.Role, error)
```
ViewGrantableRoles lists roles grantable on a resource.

#### type Keys

```go
type Keys struct {
}
```


#### func  NewKeys

```go
func NewKeys(ctx context.Context, opts ...option.ClientOption) (*Keys, error)
```

#### func (*Keys) AddMemberCryptoKeyPolicy

```go
func (v *Keys) AddMemberCryptoKeyPolicy(ctx context.Context, keyName, member string, role iam.RoleName) error
```
AddMemberCryptoKeyPolicy adds a new member to a specified IAM role for the key
example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"

#### func (*Keys) AddMemberRingPolicy

```go
func (v *Keys) AddMemberRingPolicy(ctx context.Context, keyRingName, member string, role iam.RoleName) error
```
AddMemberRingPolicy adds a new member to a specified IAM role for the key ring
example keyRingName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID"

#### func (*Keys) Client

```go
func (v *Keys) Client() *cloudkms.KeyManagementClient
```

#### func (*Keys) Close

```go
func (v *Keys) Close()
```

#### func (*Keys) CreateAsymmetricKey

```go
func (v *Keys) CreateAsymmetricKey(ctx context.Context, keyRingName, keyId string) error
```
CreateAsymmetricKey creates a new RSA encrypt/decrypt key pair on KMS. example
keyRingName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID"

#### func (*Keys) CreateCryptoKey

```go
func (v *Keys) CreateCryptoKey(ctx context.Context, keyRingName, keyId string) (*kmspb.CryptoKey, error)
```
CreateCryptoKey creates a new symmetric encrypt/decrypt key on KMS. example
keyRingName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID"

#### func (*Keys) CreateKeyRing

```go
func (v *Keys) CreateKeyRing(ctx context.Context, parentName, keyRingId string) (*kmspb.KeyRing, error)
```
CreateKeyRing creates a new ring to store keys on KMS. example parentName:
"projects/PROJECT_ID/locations/global/"

#### func (*Keys) DecryptRSA

```go
func (v *Keys) DecryptRSA(ctx context.Context, keyName string, ciphertext []byte) ([]byte, error)
```
DecryptRSA will attempt to decrypt a given ciphertext with an
'RSA_DECRYPT_OAEP_2048_SHA256' private key.stored on Cloud KMS example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"

#### func (*Keys) DecryptSymmetric

```go
func (v *Keys) DecryptSymmetric(ctx context.Context, keyName string, ciphertext []byte) ([]byte, error)
```
Decrypt will decrypt the input ciphertext bytes using the specified symmetric
key example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"

#### func (*Keys) DestroyCryptoKeyVersion

```go
func (v *Keys) DestroyCryptoKeyVersion(ctx context.Context, keyVersionName string) error
```
DestroyCryptoKeyVersion marks a specified key version for deletion. The key can
be restored if requested within 24 hours. example keyVersionName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"

#### func (*Keys) DisableCryptoKeyVersion

```go
func (v *Keys) DisableCryptoKeyVersion(ctx context.Context, keyVersionName string) error
```
DisableCryptoKeyVersion disables a specified key version on KMS. example
keyVersionName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"

#### func (*Keys) EnableCryptoKeyVersion

```go
func (v *Keys) EnableCryptoKeyVersion(ctx context.Context, keyVersionName string) error
```
EnableCryptoKeyVersion enables a previously disabled key version on KMS. example
keyVersionName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"

#### func (*Keys) EncryptRSA

```go
func (v *Keys) EncryptRSA(ctx context.Context, keyName string, plaintext []byte) ([]byte, error)
```
encryptRSA will encrypt data locally using an 'RSA_DECRYPT_OAEP_2048_SHA256'
public key retrieved from Cloud KMS example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"

#### func (*Keys) EncryptSymmetric

```go
func (v *Keys) EncryptSymmetric(ctx context.Context, keyName string, plaintext []byte) ([]byte, error)
```
Encrypt will encrypt the input plaintext with the specified symmetric key
example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"

#### func (*Keys) GetAsymmetricPublicKey

```go
func (v *Keys) GetAsymmetricPublicKey(ctx context.Context, keyName string) (interface{}, error)
```
GetAsymmetricPublicKey retrieves the public key from a saved asymmetric key pair
on KMS. example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"

#### func (*Keys) GetCryptoKeyPolicy

```go
func (v *Keys) GetCryptoKeyPolicy(ctx context.Context, keyName string) (*iam.Policy, error)
```
GetCryptoKeyPolicy retrieves and prints the IAM policy associated with the key
example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"

#### func (*Keys) GetRingPolicy

```go
func (v *Keys) GetRingPolicy(ctx context.Context, keyRingName string) (*iam.Policy, error)
```
GetRingPolicy retrieves and prints the IAM policy associated with the key ring
example keyRingName: "projects/PROJECT_ID/locations/global/keyRings/RING_ID"

#### func (*Keys) RemoveMemberCryptoKeyPolicy

```go
func (v *Keys) RemoveMemberCryptoKeyPolicy(ctx context.Context, keyName, member string, role iam.RoleName) error
```
RemoveMemberCryptoKeyPolicy removes a specified member from an IAM role for the
key example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID"

#### func (*Keys) RemoveMemberRingPolicy

```go
func (v *Keys) RemoveMemberRingPolicy(ctx context.Context, keyRingName, member string, role iam.RoleName) error
```
RemoveMemberRingPolicy removes a specified member from an IAM role for the key
ring example keyRingName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID"

#### func (*Keys) RestoreCryptoKeyVersion

```go
func (v *Keys) RestoreCryptoKeyVersion(ctx context.Context, keyVersionName string) error
```
RestoreCryptoKeyVersion attempts to recover a key that has been marked for
destruction within the last 24 hours. example keyVersionName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"

#### func (*Keys) SignAsymmetric

```go
func (v *Keys) SignAsymmetric(ctx context.Context, keyName string, message []byte) ([]byte, error)
```
SignAsymmetric will sign a plaintext message using a saved asymmetric private
key. example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"

#### func (*Keys) VerifySignatureEC

```go
func (v *Keys) VerifySignatureEC(ctx context.Context, keyName string, signature, message []byte) error
```
VerifySignatureEC will verify that an 'EC_SIGN_P256_SHA256' signature is valid
for a given message. example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"

#### func (*Keys) VerifySignatureRSA

```go
func (v *Keys) VerifySignatureRSA(ctx context.Context, keyName string, signature, message []byte) error
```
VerifySignatureRSA will verify that an 'RSA_SIGN_PSS_2048_SHA256' signature is
valid for a given message. example keyName:
"projects/PROJECT_ID/locations/global/keyRings/RING_ID/cryptoKeys/KEY_ID/cryptoKeyVersions/1"
