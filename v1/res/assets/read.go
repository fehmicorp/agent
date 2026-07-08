package assets

func Read(name string) ([]byte, error) {
	return FS.ReadFile("assets/" + name)
}
