package resources

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/onego-project/xmlrpc"
	"strconv"
)

// VirtualMachine struct
type VirtualMachine struct {
	XMLData *etree.Element
}

// RPC struct
type RPC struct {
	Client *xmlrpc.Client
	Key    string
}

// Snapshot struct
type Snapshot struct {
	//name       string
	SnapshotID int
}

// Monitoring struct
type Monitoring struct {
	VMid              int
	MonitoringRecords []MonitoringData
}

// MonitoringData struct
type MonitoringData struct {
	XMLData *etree.Element
}

// History struct
type History struct {
	XMLData *etree.Element
}

// CreateVM constructs VirtualMachine
func CreateVM(id int) *VirtualMachine {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0"`)

	el := doc.CreateElement("VM").CreateElement("ID")
	el.SetText(fmt.Sprintf("%d", id))

	return &VirtualMachine{doc.Root()}
}

// Attribute method
func (vm VirtualMachine) Attribute(path string) string {
	elements := vm.XMLData.FindElements(path)
	if elements == nil {
		return ""
	}
	return elements[0].Text()
}

// ID method
func (vm VirtualMachine) ID() int {
	i, err := strconv.Atoi(vm.Attribute("ID"))
	if err != nil {
		return -1
	}
	return i
}

// UID method
func (vm VirtualMachine) UID() int {
	i, err := strconv.Atoi(vm.Attribute("UID"))
	if err != nil {
		return -1
	}
	return i
}

// GID method
func (vm VirtualMachine) GID() int {
	i, err := strconv.Atoi(vm.Attribute("GID"))
	if err != nil {
		return -1
	}
	return i
}

// UName method
func (vm VirtualMachine) UName() string {
	return vm.Attribute("UNAME")
}

// GName method
func (vm VirtualMachine) GName() string {
	return vm.Attribute("GNAME")
}

// Name method
func (vm VirtualMachine) Name() string {
	return vm.Attribute("NAME")
}

// PermissionUserUse method
func (vm VirtualMachine) PermissionUserUse() int {
	i, err := strconv.Atoi(vm.Attribute("PERMISSIONS/OWNER_U"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionUserManage method
func (vm VirtualMachine) PermissionUserManage() int {
	i, err := strconv.Atoi(vm.Attribute("PERMISSIONS/OWNER_M"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionUserAdmin method
func (vm VirtualMachine) PermissionUserAdmin() int {
	i, err := strconv.Atoi(vm.Attribute("PERMISSIONS/OWNER_A"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionGroupUse method
func (vm VirtualMachine) PermissionGroupUse() int {
	i, err := strconv.Atoi(vm.Attribute("PERMISSIONS/GROUP_U"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionGroupManage method
func (vm VirtualMachine) PermissionGroupManage() int {
	i, err := strconv.Atoi(vm.Attribute("PERMISSIONS/GROUP_M"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionGroupAdmin method
func (vm VirtualMachine) PermissionGroupAdmin() int {
	i, err := strconv.Atoi(vm.Attribute("PERMISSIONS/GROUP_A"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionOtherUse method
func (vm VirtualMachine) PermissionOtherUse() int {
	i, err := strconv.Atoi(vm.Attribute("PERMISSIONS/OTHER_U"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionOtherManage method
func (vm VirtualMachine) PermissionOtherManage() int {
	i, err := strconv.Atoi(vm.Attribute("PERMISSIONS/OTHER_M"))
	if err != nil {
		return -1
	}
	return i
}

// PermissionOtherAdmin method
func (vm VirtualMachine) PermissionOtherAdmin() int {
	i, err := strconv.Atoi(vm.Attribute("PERMISSIONS/OTHER_A"))
	if err != nil {
		return -1
	}
	return i
}

// LastPoll method
func (vm VirtualMachine) LastPoll() int {
	i, err := strconv.Atoi(vm.Attribute("LAST_POLL"))
	if err != nil {
		return -1
	}
	return i
}

// State method
func (vm VirtualMachine) State() int {
	i, err := strconv.Atoi(vm.Attribute("STATE"))
	if err != nil {
		return -1
	}
	return i
}

// LCMState method
func (vm VirtualMachine) LCMState() int {
	i, err := strconv.Atoi(vm.Attribute("LCM_STATE"))
	if err != nil {
		return -1
	}
	return i
}

// PrevState method
func (vm VirtualMachine) PrevState() int {
	i, err := strconv.Atoi(vm.Attribute("PREV_STATE"))
	if err != nil {
		return -1
	}
	return i
}

// PrevLCMState method
func (vm VirtualMachine) PrevLCMState() int {
	i, err := strconv.Atoi(vm.Attribute("PREV_LCM_STATE"))
	if err != nil {
		return -1
	}
	return i
}

// Resched method
func (vm VirtualMachine) Resched() int {
	i, err := strconv.Atoi(vm.Attribute("RESCHED"))
	if err != nil {
		return -1
	}
	return i
}

// STime method
func (vm VirtualMachine) STime() int64 {
	i, err := strconv.ParseInt(vm.Attribute("STIME"), 10, 64)
	if err != nil {
		return -1
	}
	return i
}

// ETime method
func (vm VirtualMachine) ETime() int64 {
	i, err := strconv.ParseInt(vm.Attribute("ETIME"), 10, 64)
	if err != nil {
		return -1
	}
	return i
}

// DeployID method
func (vm VirtualMachine) DeployID() string {
	return vm.Attribute("DEPLOY_ID")
}

// Monitoring method
func (vm VirtualMachine) Monitoring() *Monitoring {
	return &Monitoring{
		VMid:              vm.ID(),
		MonitoringRecords: []MonitoringData{{vm.XMLData.FindElements("MONITORING")[0]}}}
}

// Template method
func (vm VirtualMachine) Template() *etree.Element {
	return vm.XMLData.FindElements("TEMPLATE")[0]
}

// UserTemplate method
func (vm VirtualMachine) UserTemplate() *etree.Element {
	return vm.XMLData.FindElements("USERTEMPLATE")[0]
}

// HistoryRecords method
func (vm VirtualMachine) HistoryRecords() []History {
	elements := vm.XMLData.FindElements("HISTORY_RECORDS/HISTORY")
	history := make([]History, len(elements))
	for i, e := range elements {
		history[i] = History{e}
	}
	return history
}
