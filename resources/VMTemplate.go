package resources

import (
	"fmt"
	"github.com/beevik/etree"
	"strconv"
)

// VMTemplate struct
type VMTemplate struct {
	XMLData *etree.Element
}

// CreateVMTemplate constructs VMTemplate
func CreateVMTemplate(id int) *Host {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0"`)

	el := doc.CreateElement("VMTEMPLATE").CreateElement("ID")
	el.SetText(fmt.Sprintf("%d", id))

	return &Host{doc.Root()}
}

// Attribute method
func (t VMTemplate) Attribute(path string) string {
	elements := t.XMLData.FindElements(path)
	if elements == nil {
		return ""
	}
	return elements[0].Text()
}

// ID method
func (t VMTemplate) ID() int {
	i, err := strconv.Atoi(t.Attribute("ID"))
	if err != nil {
		return -1
	}
	return i
}

// Name method
func (t VMTemplate) Name() string {
	return t.Attribute("NAME")
}

// User method
func (t VMTemplate) User() *User {
	i, err := strconv.Atoi(t.Attribute("UID"))
	if err != nil {
		return nil
	}
	return CreateUser(i)
}

// Group method
func (t VMTemplate) Group() *Group {
	i, err := strconv.Atoi(t.Attribute("GID"))
	if err != nil {
		return nil
	}
	return CreateGroup(i)
}

// PermissionUserUse method
func (t VMTemplate) PermissionUserUse() int {
	i, err := strconv.Atoi(t.Attribute("PERMISSIONS/OWNER_U"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionUserManage method
func (t VMTemplate) PermissionUserManage() int {
	i, err := strconv.Atoi(t.Attribute("PERMISSIONS/OWNER_M"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionUserAdmin method
func (t VMTemplate) PermissionUserAdmin() int {
	i, err := strconv.Atoi(t.Attribute("PERMISSIONS/OWNER_A"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionGroupUse method
func (t VMTemplate) PermissionGroupUse() int {
	i, err := strconv.Atoi(t.Attribute("PERMISSIONS/GROUP_U"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionGroupManage method
func (t VMTemplate) PermissionGroupManage() int {
	i, err := strconv.Atoi(t.Attribute("PERMISSIONS/GROUP_M"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionGroupAdmin method
func (t VMTemplate) PermissionGroupAdmin() int {
	i, err := strconv.Atoi(t.Attribute("PERMISSIONS/GROUP_A"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionOtherUse method
func (t VMTemplate) PermissionOtherUse() int {
	i, err := strconv.Atoi(t.Attribute("PERMISSIONS/OTHER_U"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionOtherManage method
func (t VMTemplate) PermissionOtherManage() int {
	i, err := strconv.Atoi(t.Attribute("PERMISSIONS/OTHER_M"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionOtherAdmin method
func (t VMTemplate) PermissionOtherAdmin() int {
	i, err := strconv.Atoi(t.Attribute("PERMISSIONS/OTHER_A"))
	if err != nil {
		return -1
	}
	return i
}
