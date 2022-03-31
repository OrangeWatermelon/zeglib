package zeglib

func die(err interface{}) {
	if err == nil {
		return
	}
	panic(err)
}
