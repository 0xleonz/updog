/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"regexp"

	"github.com/0xleonz/updog/internal/aur"
	"github.com/0xleonz/updog/internal/cargo"
	"github.com/0xleonz/updog/internal/pacman"
	"github.com/0xleonz/updog/internal/pip"
	"github.com/spf13/cobra"
)

var (
	configBits string
	onlyPacma  bool
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "🔄 Ejecuta actualización del sistema (pacman, AUR, cargo, pip)",
	Long: `Permite controlar las fuentes a actualizar con --config (4 bits binarios) o --pacma para solo pacman.
				 Ejemplo: updog run --config 0101  # Solo actualiza AUR y pip`,
	Run: func(cmd *cobra.Command, args []string) {
		doPacman, doAUR, doCargo, doPip := true, true, false, false

		if onlyPacma {
			doAUR, doCargo, doPip = false, false, false
		}

		if configBits != "" {
			matched, _ := regexp.MatchString("^[01]{4}$", configBits)
		if !matched {
			fmt.Println("❌ Config inválida: debe tener 4 bits binarios (ej: 1010)")
			return
		}
			doPacman = configBits[0] == '1'
			doAUR    = configBits[1] == '1'
			doCargo  = configBits[2] == '1'
			doPip    = configBits[3] == '1'
		}

		if doPacman {
			pacman.Update()
		}
		if doAUR {
			aur.Update()
		}
		if doCargo {
			cargo.Update()
		}
		if doPip {
			pip.Update()
		}

		fmt.Println("✅ Operación completada.")
	},
}

func init() {
	runCmd.Flags().BoolVar(&onlyPacma, "pacma", false, "Solo ejecutar actualización con pacman")
	runCmd.Flags().StringVarP(&configBits, "config", "c", "", "Cadena binaria de 4 bits (pacman,AUR,cargo,pip)")
	rootCmd.AddCommand(runCmd)
}
