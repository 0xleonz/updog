package pip

import (
	"fmt"
	"os/exec"
)

func Update() {
	fmt.Println("🐍 Actualizando paquetes de pip...")
	if _, err := exec.LookPath("pip"); err != nil {
		fmt.Println("❌ 'pip' no está instalado. Saltando pip.")
		return
	}

	cmd := exec.Command("bash", "-c", `
pip list --outdated --format=json --disable-pip-version-check |
python -c "import sys, json; print('\n'.join(p['name'] for p in json.load(sys.stdin)))" |
xargs -r -n1 pip install -U --break-system-packages
`)
	cmd.Stdout = cmd.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error actualizando pip:", err)
	}
}

