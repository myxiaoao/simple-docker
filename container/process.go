package container

import (
	"os"
	"os/exec"
	"syscall"
)

// NewParentProcess 创建一个会隔离的 namespace 进程的 Command
func NewParentProcess(tty bool) (*exec.Cmd, *os.File) {
	readPipe, writePipe, _ := os.Pipe()
	// 调用自身，传入 init 参数，也就是执行 initCommand
	cmd := exec.Command("/proc/self/exe", "init")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET |
			syscall.CLONE_NEWIPC, // Cloneflags linux 特定参数
	}

	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	cmd.ExtraFiles = []*os.File{
		readPipe,
	}

	return cmd, writePipe
}