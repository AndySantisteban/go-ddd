package web

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Execute(directory string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("osascript", "-e", fmt.Sprintf(`tell app "Terminal" to do script "cd %s && wails dev"`, directory))
	case "linux":
		cmd = exec.Command("x-terminal-emulator", "-e", fmt.Sprintf("cd %s && wails dev", directory))
	case "windows":
		cmd = exec.Command("cmd", "/C", fmt.Sprintf("cd %s && start cmd /K wails dev", directory))
	default:
		fmt.Println("Sistema operativo no compatible.")
		return fmt.Errorf("sistema operativo no compatible")
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error al ejecutar 'wails dev':", err)
		return err
	}

	fmt.Println("'wails dev' ejecutado en una nueva terminal.")
	return nil
}
