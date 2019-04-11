package main

import (
	"syscall"
	"unsafe"
)

func getArgument() *uint16 {

	cmd := syscall.GetCommandLine()

	dll := syscall.MustLoadDLL("Shlwapi.dll")
	defer dll.Release()

	arg, _, err := dll.MustFindProc("PathGetArgsW").Call(uintptr(unsafe.Pointer(cmd)))
	if arg == 0 {
		err = syscall.GetLastError()
		panic(err)
	}

	return (*uint16)(unsafe.Pointer(arg))
}

func winExec(argv *uint16) {

	var si syscall.StartupInfo
	si.Cb = uint32(unsafe.Sizeof(si))
	si.Flags = syscall.STARTF_USESHOWWINDOW
	si.ShowWindow = syscall.SW_SHOWDEFAULT

	var pi syscall.ProcessInformation

	err := syscall.CreateProcess(
		nil,
		argv,
		nil,
		nil,
		false,
		0,
		nil,
		nil,
		&si,
		&pi)

	if err != nil {
		panic(err)
	}

	defer syscall.CloseHandle(pi.Process)
	syscall.CloseHandle(pi.Thread)

	syscall.WaitForSingleObject(pi.Process, syscall.INFINITE)
}

func main() {

	winExec(getArgument())
}
