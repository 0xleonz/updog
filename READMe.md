# updog

`updog` es una herramienta de línea de comandos para actualizar tu sistema desde múltiples fuentes:

- pacman (Arch Linux)
- yay (AUR)
- cargo (Rust)
- pip (Python)


## Instalación


```bash
go install gitlab.com/0xleonz/updog@latest
```

> Asegúrate de tener `$GOBIN` en tu `PATH` y Go ≥ 1.16.

## Uso

```bash
updog run
# Limpia huerfanos (-Qt)
updog clean
```

updog list

## Estructura del proyecto

* `cmd/` → comandos CLI (run, clean, list)
* `internal/` → lógica por fuente (pacman, aur, cargo, pip)
* `main.go` → entrada principal
