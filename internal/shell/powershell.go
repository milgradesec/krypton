package shell

import (
	"os/exec"
	"syscall"
	"unsafe"
)

func Run(flags, workingDir string) error {
	var powershellPath string
	wow64, err := isWoW64()
	if err != nil {
		return err
	}

	if wow64 {
		powershellPath = "c:/windows/sysnative/WindowsPowerShell/v1.0/powershell.exe"
	} else {
		powershellPath = "powershell.exe"
	}
	cmd := exec.Command(powershellPath, "-ExecutionPolicy", "Bypass", flags)
	if workingDir != "" {
		cmd.Dir = workingDir
	}

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func isWoW64() (bool, error) {
	dll, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return false, err
	}
	defer dll.Release() //nolint

	proc, err := dll.FindProc("IsWow64Process")
	if err != nil {
		return false, err
	}

	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		return false, err
	}

	var result bool
	_, _, _ = proc.Call(uintptr(handle), uintptr(unsafe.Pointer(&result))) //nolint
	return result, nil
}
