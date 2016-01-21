package rbac

import (
	"encoding/json"
	"testing"
)

var (
	testPolicyBytes = []byte(`{"name": "test", "path": "^my/path", "op": "all", "action": "allow"}`)
	testPolicy      Policy
	testRole        Role
	testRoleMapping RoleMapping
)

func Test_Policy_UnmarshalJSON(t *testing.T) {

	err := json.Unmarshal(testPolicyBytes, &testPolicy)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", testPolicy)
}

func Test_Policy_IsGranted(t *testing.T) {
	if !testPolicy.IsGranted("my/path", RBAC_OP_CREATE) {
		t.Fatal("Should be granted")
	}
}

func Test_Role_IsGranted(t *testing.T) {
	testRole = Role{Name: "test-role", Policies: make([]Policy, 1)}
	testRole.Policies[0] = testPolicy
	if !testRole.IsGranted("my/path", RBAC_OP_READ) {
		t.Fatal("Should be granted")
	}
}

func Test_Role_IsGranted_not(t *testing.T) {
	if testRole.IsGranted("not_my/path", RBAC_OP_READ) {
		t.Fatal("Should not be granted")
	}
}

func Test_RoleMapping_GetRole(t *testing.T) {
	testRoleMapping = RoleMapping{Roles: []Role{testRole}}
	if testRoleMapping.GetRole("test-role") == nil {
		t.Fatal("Should return a role")
	}
}

func Test_RoleMapping_GetRole_Non_Existent(t *testing.T) {

	if testRoleMapping.GetRole("test-role1") != nil {
		t.Fatal("Should be nil")
	}
}
