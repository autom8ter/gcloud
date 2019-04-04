package auth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
)

type IAM struct {
	svc *iam.Service
}

func NewIAM(ctx context.Context, opts ...option.ClientOption) (*IAM, error) {
	cli, err := iam.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &IAM{
		svc: cli,
	}, nil
}

//No-OP
func (i *IAM) Close() {}

func (i *IAM) Client() *iam.Service {
	return i.svc
}

// CreateRole creates a custom role.
func (i *IAM) CreateRole(projectID, name, title, description, stage, roleId string, permissions []string) (*iam.Role, error) {
	request := &iam.CreateRoleRequest{
		Role: &iam.Role{
			Deleted:             false,
			Description:         description,
			Etag:                "",
			IncludedPermissions: permissions,
			Name:                name,
			Stage:               stage,
			Title:               title,
			ServerResponse:      googleapi.ServerResponse{},
			ForceSendFields:     nil,
			NullFields:          nil,
		},
		RoleId: roleId,
	}
	role, err := i.svc.Projects.Roles.Create("projects/"+projectID, request).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.Roles.Create: %v", err)
	}
	return role, nil
}

// DeleteRole deletes a custom role.
func (i *IAM) DeleteRole(projectID, name string) error {
	_, err := i.svc.Projects.Roles.Delete("projects/" + projectID + "/roles/" + name).Do()
	if err != nil {
		return fmt.Errorf("Projects.Roles.Delete: %v", err)
	}
	return nil
}

// EditRole modifies a custom role.
func (i *IAM) EditRole(projectID, name, newTitle, newDescription, newStage string, newPermissions []string) (*iam.Role, error) {
	resource := "projects/" + projectID + "/roles/" + name
	role, err := i.svc.Projects.Roles.Get(resource).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.Roles.Get: %v", err)
	}
	role.Title = newTitle
	role.Description = newDescription
	role.IncludedPermissions = newPermissions
	role.Stage = newStage
	role, err = i.svc.Projects.Roles.Patch(resource, role).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.Roles.Patch: %v", err)
	}
	return role, nil
}

func (i *IAM) GetRole(name string) (*iam.Role, error) {
	return i.svc.Roles.Get(name).Do()
}

// CreateKey creates a service account key.
func (i *IAM) CreateKey(serviceAccountEmail string) (*iam.ServiceAccountKey, error) {
	resource := "projects/-/serviceAccounts/" + serviceAccountEmail
	request := &iam.CreateServiceAccountKeyRequest{}
	return i.svc.Projects.ServiceAccounts.Keys.Create(resource, request).Do()
}

// DeleteKey deletes a service account key.
func (i *IAM) DeleteKey(fullKeyName string) error {
	_, err := i.svc.Projects.ServiceAccounts.Keys.Delete(fullKeyName).Do()
	if err != nil {
		return fmt.Errorf("Projects.ServiceAccounts.Keys.Delete: %v", err)
	}
	return nil
}

// ListServiceAccounts lists a project's service accounts.
func (i *IAM) ListServiceAccounts(projectID string) (*iam.ListServiceAccountsResponse, error) {
	return i.svc.Projects.ServiceAccounts.List("projects/" + projectID).Do()
}

// ViewGrantableRoles lists roles grantable on a resource.
func (i *IAM) ViewGrantableRoles(fullResourceName string) ([]*iam.Role, error) {
	request := &iam.QueryGrantableRolesRequest{
		FullResourceName: fullResourceName,
	}
	resp, err := i.svc.Roles.QueryGrantableRoles(request).Do()
	if err != nil {
		return nil, err
	}
	return resp.Roles, nil
}

// ListKey lists a service account's keys.
func (i *IAM) ListKeys(serviceAccountEmail string) ([]*iam.ServiceAccountKey, error) {
	resource := "projects/-/serviceAccounts/" + serviceAccountEmail
	response, err := i.svc.Projects.ServiceAccounts.Keys.List(resource).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Keys.List: %v", err)
	}
	return response.Keys, nil
}

// queryTestablePermissions lists testable permissions on a resource.
func (i *IAM) AueryTestablePermissions(fullResourceName string) ([]*iam.Permission, error) {
	client, err := google.DefaultClient(context.Background(), iam.CloudPlatformScope)
	if err != nil {
		return nil, fmt.Errorf("google.DefaultClient: %v", err)
	}
	service, err := iam.New(client)
	if err != nil {
		return nil, fmt.Errorf("iam.New: %v", err)
	}

	request := &iam.QueryTestablePermissionsRequest{
		FullResourceName: fullResourceName,
	}
	response, err := service.Permissions.QueryTestablePermissions(request).Do()
	if err != nil {
		return nil, fmt.Errorf("Permissions.QueryTestablePermissions: %v", err)
	}
	return response.Permissions, nil
}
