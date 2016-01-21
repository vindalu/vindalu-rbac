package rbac

import (
	"encoding/json"
	"regexp"
)

type RbacAction string

type RbacOp string

const (
	RBAC_ACTION_ALLOW RbacAction = "allow"
	RBAC_ACTION_DENY  RbacAction = "deny"

	RBAC_OP_CREATE RbacOp = "create"
	RBAC_OP_READ   RbacOp = "read"
	RBAC_OP_UPDATE RbacOp = "update"
	RBAC_OP_DELETE RbacOp = "delete"
	RBAC_OP_ALL    RbacOp = "all"
)

type Policy struct {
	Name string `json:"name"`

	// url path
	Path string `json:"path"`

	pathRe *regexp.Regexp

	// allow, deny
	Action RbacAction `json:"action"`

	// read, create, update, delete, all
	Op RbacOp `json:"op"`
}

// Temp for unmarshal to avoid recursion
type policy Policy

func (p *Policy) UnmarshalJSON(data []byte) (err error) {
	var t policy

	if err = json.Unmarshal(data, &t); err == nil {

		if t.pathRe, err = regexp.Compile(t.Path); err == nil {
			*p = Policy(t)
		}
	}

	return
}

func (p *Policy) IsGranted(path string, op RbacOp) bool {
	return p.pathRe.MatchString(path) && (p.Op == RBAC_OP_ALL || p.Op == op) &&
		p.Action == RBAC_ACTION_ALLOW
}

type Role struct {
	Name     string
	Policies []Policy
}

func (r *Role) IsGranted(path string, op RbacOp) bool {
	for _, p := range r.Policies {
		if p.IsGranted(path, op) {
			return true
		}
	}
	return false
}

type RoleMapping struct {
	// Name of user or group
	Name string `json:"name"`

	// Assigned roles
	Roles []Role `json:"roles"`
}

func (r *RoleMapping) GetRole(name string) *Role {
	for _, r := range r.Roles {
		if r.Name == name {
			return &r
		}
	}
	return nil
}
