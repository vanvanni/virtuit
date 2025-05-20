package resources

type ResourceSpec struct {
    CPU int `toml:"cpu"`
    Mem int `toml:"mem"` // in MB
}