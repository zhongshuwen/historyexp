package migrator

import (
zsw "github.com/zhongshuwen/zswchain-go"
)

type PermissionObject struct {
	// Parent of this permission object
	Parent zsw.PermissionName `json:"parent,omitempty"`
	// Owner is the account for which this permission belongs to
	Owner zsw.AccountName `json:"owner,omitempty"`
	// Name is the permission's name this permission object is known as (human-readable name for the permission)
	Name zsw.PermissionName `json:"name,omitempty"`
	// Authority required to execute this permission
	Authority *zsw.Authority `json:"authority,omitempty"`
}

type LinkAuth struct {
	Permission string `json:"permission"`
	Contract   string `json:"contract"`
	Action     string `json:"action"`
}

type AccountInfo struct {
	Permissions []*PermissionObject `json:"permissions"`
	LinkAuths   []*LinkAuth         `json:"link_auths"`

	nameToPerm map[zsw.PermissionName]*PermissionObject
}

func newAccountInfo(permissions []*PermissionObject, linkAuths []*LinkAuth) *AccountInfo {
	info := &AccountInfo{
		Permissions: permissions,
		LinkAuths:   linkAuths,
	}
	info.setupIDtoPerm()
	return info
}

func (a *AccountInfo) setupIDtoPerm() {
	a.nameToPerm = make(map[zsw.PermissionName]*PermissionObject, len(a.Permissions))
	for _, perm := range a.Permissions {
		a.nameToPerm[perm.Name] = perm
	}
}

func (a *AccountInfo) sortPermissions() (out []*PermissionObject) {
	var roots []*PermissionObject
	parentToChildren := map[zsw.PermissionName][]*PermissionObject{}
	for _, perm := range a.Permissions {
		if perm.Parent == "" {
			roots = append(roots, perm)
			continue
		}

		parentToChildren[perm.Parent] = append(parentToChildren[perm.Parent], perm)
	}

	var walk func(roots []*PermissionObject, index int)
	walk = func(roots []*PermissionObject, index int) {
		if index >= len(roots) {
			return
		}
		ele := roots[index]
		out = append(out, ele)

		for _, child := range parentToChildren[ele.Name] {
			roots = append(roots, child)
		}
		index = index + 1
		walk(roots, index)
	}

	walk(roots, 0)

	return out
}
