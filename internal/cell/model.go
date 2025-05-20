package cell

type CellSpec struct {
	Path       string `toml:"path"`
	Mountpoint string `toml:"mountpoint"`
	Boot       bool   `toml:"boot"`
}
