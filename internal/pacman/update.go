package pacman

import (
	"fmt"
	"os"
	"os/exec"
)

func Update() {
	fmt.Println("🔄 Actualizando sistema con pacman...")
	cmd := exec.Command("sudo", "pacman", "-Syu", "--noconfirm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error actualizando pacman:", err)
	}
}

