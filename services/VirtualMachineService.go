package services

import (
	"context"
	"fmt"
	"github.com/beevik/etree"
	"github.com/onego-project/onego/blueprint"
	"github.com/onego-project/onego/requests"
	"github.com/onego-project/onego/resources"
	"github.com/onego-project/xmlrpc"
)

// VirtualMachineService struct
type VirtualMachineService struct {
	RPC *resources.RPC
}

const (
	vmDelete       = "terminate"
	vmForceDelete  = "terminate-hard"
	vmUndeploy     = "undeploy"
	vmUndeployHard = "undeploy-hard"
	vmPoweroff     = "poweroff"
	vmPoweroffHard = "poweroff-hard"
	vmReboot       = "reboot"
	vmRebootHard   = "reboot-hard"
	vmHold         = "hold"
	vmRelease      = "release"
	vmStop         = "stop"
	vmSuspend      = "suspend"
	vmResume       = "resume"
	vmReschedule   = "resched"
	vmUnreschedule = "unresched"
)

// UpdateType type
type UpdateType int

const (
	// Replace const
	Replace UpdateType = iota
	// Merge const
	Merge
)

// RecoverOperation type
type RecoverOperation int

const (
	// Failure const
	Failure RecoverOperation = iota
	// Success const
	Success
	// Retry const
	Retry
	// Delete const
	Delete
	// DeleteRecreate const
	DeleteRecreate
)

// OwnershipFilter type
type OwnershipFilter int

const (
	// PrimaryGroup const
	PrimaryGroup OwnershipFilter = iota - 4
	// User const
	User
	// All const
	All
	// UserGroups const
	UserGroups
)

// StateFilter type
type StateFilter int

const (
	// AnyStateIncludingDone const
	AnyStateIncludingDone StateFilter = iota - 2
	// AnyStateExceptDone const
	AnyStateExceptDone
	// Init const
	Init
	// Pending const
	Pending
	// Hold const
	Hold
	// Active const
	Active
	// Stopped const
	Stopped
	// Suspended const
	Suspended
	// Done const
	Done
	// Failed const
	Failed
	// PowerOff const
	PowerOff
	// Undeployed const
	Undeployed
	// Cloning const
	Cloning
	// CloningFailure const
	CloningFailure
)

func (s VirtualMachineService) call(methodName string, args ...interface{}) ([]*xmlrpc.Result, error) {
	ctx := context.TODO()

	result, err := s.RPC.Client.Call(ctx, methodName, args...)
	if err != nil {
		return nil, err
	}

	resArr := result.ResultArray()
	if !resArr[0].ResultBoolean() {
		return nil, fmt.Errorf("%s, code: %d", resArr[1].ResultString(), resArr[2].ResultInt())
	}

	return resArr, nil
}

// Deploy method
func (s VirtualMachineService) Deploy(vm resources.VirtualMachine, host resources.Host, overCommit bool, datastore resources.DataStore) error {
	args := []interface{}{s.RPC.Key, vm.ID(), host.ID(), overCommit, datastore.ID()}
	_, err := s.call("one.vm.deploy", args...)
	return err
}

func (s VirtualMachineService) actions(vm resources.VirtualMachine, action string) error {
	args := []interface{}{s.RPC.Key, action, vm.ID()}
	_, err := s.call("one.vm.action", args...)
	return err
}

// Terminate method
func (s VirtualMachineService) Terminate(vm resources.VirtualMachine, hard bool) error {
	if hard {
		return s.actions(vm, vmForceDelete)
	}
	return s.actions(vm, vmDelete)
}

// Undeploy method
func (s VirtualMachineService) Undeploy(vm resources.VirtualMachine, hard bool) error {
	if hard {
		return s.actions(vm, vmUndeployHard)
	}
	return s.actions(vm, vmUndeploy)
}

// Poweroff method
func (s VirtualMachineService) Poweroff(vm resources.VirtualMachine, hard bool) error {
	if hard {
		return s.actions(vm, vmPoweroffHard)
	}
	return s.actions(vm, vmPoweroff)
}

// Reboot method
func (s VirtualMachineService) Reboot(vm resources.VirtualMachine, hard bool) error {
	if hard {
		return s.actions(vm, vmRebootHard)
	}
	return s.actions(vm, vmReboot)
}

// Hold method
func (s VirtualMachineService) Hold(vm resources.VirtualMachine) error {
	return s.actions(vm, vmHold)
}

// Release method
func (s VirtualMachineService) Release(vm resources.VirtualMachine) error {
	return s.actions(vm, vmRelease)
}

// Stop method
func (s VirtualMachineService) Stop(vm resources.VirtualMachine) error {
	return s.actions(vm, vmStop)
}

// Suspend method
func (s VirtualMachineService) Suspend(vm resources.VirtualMachine) error {
	return s.actions(vm, vmSuspend)
}

// Resume method
func (s VirtualMachineService) Resume(vm resources.VirtualMachine) error {
	return s.actions(vm, vmResume)
}

// Reschedule method
func (s VirtualMachineService) Reschedule(vm resources.VirtualMachine) error {
	return s.actions(vm, vmReschedule)
}

// Unreschedule method
func (s VirtualMachineService) Unreschedule(vm resources.VirtualMachine) error {
	return s.actions(vm, vmUnreschedule)
}

// Migrate method
func (s VirtualMachineService) Migrate(vm resources.VirtualMachine, host resources.Host, datastore resources.DataStore, liveMigration bool, overcommit bool) error {
	args := []interface{}{s.RPC.Key, vm.ID(), host.ID(), liveMigration, overcommit, datastore.ID()}
	_, err := s.call("one.vm.migrate", args...)
	return err
}

// Chmod method
func (s VirtualMachineService) Chmod(vm resources.VirtualMachine, request requests.PermissionRequest) error {
	args := []interface{}{s.RPC.Key, vm.ID()}
	for pGroup := 0; pGroup < 3; pGroup++ {
		for pType := 0; pType < 3; pType++ {
			args = append(args, request.Permissions[pGroup][pType])
		}
	}

	_, err := s.call("one.vm.chmod", args...)
	return err
}

// Chown method
func (s VirtualMachineService) Chown(vm resources.VirtualMachine, request requests.OwnershipRequest) error {
	args := []interface{}{s.RPC.Key, vm.ID(), request.User, request.Group}
	_, err := s.call("one.vm.chown", args...)
	return err
}

// Rename method
func (s VirtualMachineService) Rename(vm resources.VirtualMachine, name string) error {
	args := []interface{}{s.RPC.Key, vm.ID(), name}
	_, err := s.call("one.vm.rename", args...)
	return err
}

// CreateSnapshot method
func (s VirtualMachineService) CreateSnapshot(vm resources.VirtualMachine, name string) (*resources.Snapshot, error) {
	args := []interface{}{s.RPC.Key, vm.ID(), name}

	resArr, err := s.call("one.vm.snapshotcreate", args...)
	if err != nil {
		return nil, err
	}

	snapshot := resources.Snapshot{SnapshotID: int(resArr[1].ResultInt())}

	return &snapshot, nil
}

// RevertSnapshot method
func (s VirtualMachineService) RevertSnapshot(vm resources.VirtualMachine, snapshot resources.Snapshot) error {
	args := []interface{}{s.RPC.Key, vm.ID(), snapshot.SnapshotID}
	_, err := s.call("one.vm.snapshotrevert", args...)
	return err
}

// DeleteSnapshot method
func (s VirtualMachineService) DeleteSnapshot(vm resources.VirtualMachine, snapshot resources.Snapshot) error {
	args := []interface{}{s.RPC.Key, vm.ID(), snapshot.SnapshotID}
	_, err := s.call("one.vm.snapshotdelete", args...)
	return err
}

// Resize method
func (s VirtualMachineService) Resize(vm resources.VirtualMachine, request blueprint.Interface, overCommit bool) error {
	args := []interface{}{s.RPC.Key, vm.ID(), request.Render(), overCommit}
	_, err := s.call("one.vm.resize", args...)
	return err
}

// UpdateUserTemplate method
func (s VirtualMachineService) UpdateUserTemplate(vm resources.VirtualMachine, blueprint blueprint.Interface, updateType UpdateType) error {
	args := []interface{}{s.RPC.Key, vm.ID(), blueprint.Render(), updateType}
	_, err := s.call("one.vm.update", args...)
	return err
}

// UpdateTemplate method
func (s VirtualMachineService) UpdateTemplate(vm resources.VirtualMachine, blueprint blueprint.Interface) error {
	args := []interface{}{s.RPC.Key, vm.ID(), blueprint.Render()}
	_, err := s.call("one.vm.updateconf", args...)
	return err
}

// Recover method
func (s VirtualMachineService) Recover(vm resources.VirtualMachine, operation RecoverOperation) error {
	args := []interface{}{s.RPC.Key, vm.ID(), operation}
	_, err := s.call("one.vm.recover", args...)
	return err
}

// RetrieveInfo method
func (s VirtualMachineService) RetrieveInfo(vm resources.VirtualMachine) (*resources.VirtualMachine, error) {
	args := []interface{}{s.RPC.Key, vm.ID()}
	resArr, err := s.call("one.vm.info", args...)
	if err != nil {
		return nil, err
	}

	doc := etree.NewDocument()
	if err = doc.ReadFromString(resArr[1].ResultString()); err != nil {
		return nil, err
	}

	vminfo := resources.VirtualMachine{XMLData: doc.Root()}

	return &vminfo, nil
}

// ListAll method
func (s VirtualMachineService) ListAll(ownershipFilter OwnershipFilter, stateFilter StateFilter) ([]*resources.VirtualMachine, error) {
	args := []interface{}{s.RPC.Key, ownershipFilter, -1, -1, stateFilter}
	resArr, err := s.call("one.vmpool.info", args...)
	if err != nil {
		return nil, err
	}

	doc := etree.NewDocument()
	if err = doc.ReadFromString(resArr[1].ResultString()); err != nil {
		return nil, err
	}

	elements := doc.FindElements("VM_POOL/VM")
	virtualMachines := make([]*resources.VirtualMachine, len(elements))
	for i, e := range elements {
		virtualMachines[i] = &resources.VirtualMachine{XMLData: e}
	}

	return virtualMachines, nil
}

// ListAllForUser method
func (s VirtualMachineService) ListAllForUser(user int, stateFilter StateFilter) ([]*resources.VirtualMachine, error) {
	args := []interface{}{s.RPC.Key, user, -1, -1, stateFilter}
	resArr, err := s.call("one.vmpool.info", args...)
	if err != nil {
		return nil, err
	}

	doc := etree.NewDocument()
	if err = doc.ReadFromString(resArr[1].ResultString()); err != nil {
		return nil, err
	}

	elements := doc.FindElements("VM_POOL/VM")
	virtualMachines := make([]*resources.VirtualMachine, len(elements))
	for i, e := range elements {
		virtualMachines[i] = &resources.VirtualMachine{XMLData: e}
	}

	return virtualMachines, nil
}

// List method
func (s VirtualMachineService) List(pageOffset int, pageSize int, ownershipFilter OwnershipFilter, stateFilter StateFilter) ([]*resources.VirtualMachine, error) {
	args := []interface{}{s.RPC.Key, ownershipFilter, -pageOffset, -pageSize, stateFilter}
	resArr, err := s.call("one.vmpool.info", args...)
	if err != nil {
		return nil, err
	}

	doc := etree.NewDocument()
	if err = doc.ReadFromString(resArr[1].ResultString()); err != nil {
		return nil, err
	}

	elements := doc.FindElements("VM_POOL/VM")
	virtualMachines := make([]*resources.VirtualMachine, len(elements))
	for i, e := range elements {
		virtualMachines[i] = &resources.VirtualMachine{XMLData: e}
	}

	return virtualMachines, nil
}

// ListForUser method
func (s VirtualMachineService) ListForUser(user int, pageOffset int, pageSize int, stateFilter StateFilter) ([]*resources.VirtualMachine, error) {
	args := []interface{}{s.RPC.Key, user, -pageOffset, -pageSize, stateFilter}
	resArr, err := s.call("one.vmpool.info", args...)
	if err != nil {
		return nil, err
	}

	doc := etree.NewDocument()
	if err = doc.ReadFromString(resArr[1].ResultString()); err != nil {
		return nil, err
	}

	elements := doc.FindElements("VM_POOL/VM")
	virtualMachines := make([]*resources.VirtualMachine, len(elements))
	for i, e := range elements {
		virtualMachines[i] = &resources.VirtualMachine{XMLData: e}
	}

	return virtualMachines, nil
}

// Allocate method
func (s VirtualMachineService) Allocate(blueprintInterface blueprint.Interface, onHold bool) (*resources.VirtualMachine, error) {
	args := []interface{}{s.RPC.Key, blueprintInterface.Render(), onHold}

	resArr, err := s.call("one.vm.allocate", args...)
	if err != nil {
		return nil, err
	}

	vm := resources.CreateVM(int(resArr[1].ResultInt()))

	return vm, nil
}
