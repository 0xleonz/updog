package cargo

import (
	"fmt"
	"os/exec"
	"strings"
)

func Update() {
	fmt.Println("🦀 Actualizando herramientas instaladas con cargo...")

	if _, err := exec.LookPath("cargo-binstall"); err == nil {
		fmt.Println("⚡ Usando cargo-binstall para bins precompilados...")
		cmd := exec.Command("cargo", "install", "--list")
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("❌ Error listando cargo installs:", err)
			return
		}
		lines := string(out)
		for _, line := range splitInstalledBins(lines) {
			exec.Command("cargo", "binstall", "--no-confirm", "--force", line).Run()
		}
		return
	}

	if _, err := exec.LookPath("cargo-install-update"); err == nil {
		fmt.Println("📦 Usando cargo-update...")
		exec.Command("cargo", "install-update", "-a").Run()
		return
	}

	fmt.Println("🛠️  Ninguna opción encontrada, instalando cargo-binstall...")
	exec.Command("cargo", "install", "cargo-binstall").Run()
	exec.Command("cargo", "binstall", "--no-confirm", "--force").Run()
}

// helper para parsear nombres de binarios instalados
func splitInstalledBins(out string) []string {
	var bins []string
	for _, line := range splitLines(out) {
		if len(line) > 0 && line[0] != ' ' {
			name := line
			if i := indexOf(line, " v"); i != -1 {
				name = line[:i]
			}
			bins = append(bins, name)
		}
	}
	return bins
}

func splitLines(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}

func indexOf(s, substr string) int {
	return strings.Index(s, substr)
}

