package aur

import (
	"fmt"
	"os/exec"
)

func Update() {
	fmt.Println("🔄 Actualizando paquetes de AUR con yay...")
	if _, err := exec.LookPath("yay"); err != nil {
		fmt.Println("❌ 'yay' no está instalado. Saltando AUR.")
		return
	}
	cmd := exec.Command("yay", "-Syu", "--noconfirm")
	cmd.Stdout = cmd.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error actualizando AUR:", err)
	}
}

