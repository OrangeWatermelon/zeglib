package zeglib

func Die(err interface{}) {
	if err == nil {
		return
	}
	panic(err)
}
