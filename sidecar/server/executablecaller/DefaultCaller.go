package executablecaller

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func ReadOutput(cmd string, out *bytes.Buffer, errOut *bytes.Buffer, args ...string) (int, error) {

	var command exec.Cmd

	cmdExec := &command
	exitStatus := 0

	var waitStatus syscall.WaitStatus

	if args != nil {
		cmdExec = exec.Command(cmd, args...)
	} else {
		cmdExec = exec.Command(cmd)
	}
	cmdExec.Stdout = out
	cmdExec.Stderr = errOut

	err := cmdExec.Start()

	if err != nil {
		return exitStatus, err
	}
	timer := time.AfterFunc(5*time.Second, func() {
		cmdExec.Process.Kill()
	})

	err = cmdExec.Wait()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			exitStatus = waitStatus.ExitStatus()
			//fmt.Println("err status: ", exitStatus)
		}
	} else {
		waitStatus = cmdExec.ProcessState.Sys().(syscall.WaitStatus)
		exitStatus = waitStatus.ExitStatus()
		//fmt.Println("ok status: ", exitStatus)
	}

	bOk := timer.Stop()

	if !bOk {
		return exitStatus, errors.New("Error -1: Timeout")
	}

	return exitStatus, err
}

type ErrorDesc struct {
	code int
	desc string
}

type DefaultCaller struct {
	status string
}

func (caller DefaultCaller) GetCreateWorkspaceExecutable() string {
	return "./CreateWorkspace"
}

func (caller DefaultCaller) CreateWorkspaceCaller(arg1 string) (string, int) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	execFile := caller.GetCreateWorkspaceExecutable()

	if _, err := os.Stat(execFile); os.IsNotExist(err) {
		return "Executable file does not exist", -1
	}

	exitStatus, err := ReadOutput(execFile, &out, &errOut, arg1)

	if exitStatus == -1 {

		return err.Error(), exitStatus
	}
	return out.String(), exitStatus
}

func (caller DefaultCaller) GetCreateConnectionExecutable() string {
	return "./CreateConnection"
}

func (caller DefaultCaller) CreateConnectionCaller(arg1 string, arg2 string) (string, int) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	execFile := caller.GetCreateConnectionExecutable()

	if _, err := os.Stat(execFile); os.IsNotExist(err) {
		return "Executable file does not exist", -1
	}

	exitStatus, err := ReadOutput(execFile, &out, &errOut, arg1, arg2)

	if exitStatus == -1 {

		return err.Error(), exitStatus
	}
	return out.String(), exitStatus
}

func (caller DefaultCaller) GetDeleteWorkspaceExecutable() string {
	return "./DeleteWorkspace"
}

func (caller DefaultCaller) DeleteWorkspaceCaller(arg1 string) (string, int) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	execFile := caller.GetDeleteWorkspaceExecutable()

	if _, err := os.Stat(execFile); os.IsNotExist(err) {
		return "Executable file does not exist", -1
	}

	exitStatus, _ := ReadOutput(execFile, &out, &errOut, arg1)

	if exitStatus == -1 {

		return errOut.String(), exitStatus
	}
	return out.String(), exitStatus
}

func (caller DefaultCaller) GetDeleteConnectionExecutable() string {
	return "./DeleteConnection"
}

func (caller DefaultCaller) DeleteConnectionCaller(arg1 string, arg2 string) (string, int) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	execFile := caller.GetDeleteConnectionExecutable()

	if _, err := os.Stat(execFile); os.IsNotExist(err) {
		return "Executable file does not exist", -1
	}

	exitStatus, _ := ReadOutput(execFile, &out, &errOut, arg1, arg2)

	if exitStatus == -1 {

		return errOut.String(), exitStatus
	}
	return out.String(), exitStatus
}

func (caller DefaultCaller) GetGetConnectionExecutable() string {
	return "./GetConnection"
}

func (caller DefaultCaller) GetConnectionCaller(arg1 string, arg2 string) (string, int) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	execFile := caller.GetGetConnectionExecutable()

	if _, err := os.Stat(execFile); os.IsNotExist(err) {
		return "Executable file does not exist", -1
	}

	exitStatus, _ := ReadOutput(execFile, &out, &errOut, arg1, arg2)

	if exitStatus == -1 {

		return errOut.String(), exitStatus
	}
	return out.String(), exitStatus
}

func (caller DefaultCaller) GetGetWorkspaceExecutable() string {
	return "./GetWorkspace"
}

func (caller DefaultCaller) GetWorkspaceCaller(arg1 string) (string, int) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	execFile := caller.GetGetWorkspaceExecutable()

	if _, err := os.Stat(execFile); os.IsNotExist(err) {
		return "Executable file does not exist", -1
	}

	exitStatus, _ := ReadOutput(execFile, &out, &errOut, arg1)

	if exitStatus == -1 {

		return errOut.String(), exitStatus
	}
	return out.String(), exitStatus
}