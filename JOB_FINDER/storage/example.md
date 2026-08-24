//to use in main
//TODO - config add STORAGE_DIR for consistency
func main() {
	logger := loggersystem.Init()
	s, err := storage.NewStore("../interview", logger)
	if err != nil {
		panic(err)
	}

	err = s.Scan()
	s.Show()
}
