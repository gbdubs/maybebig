package maybebig

func Min(fs ...*Float) *Float {
	m := fs[0]
	for i := 1; i < len(fs); i++ {
		if Lt(fs[i], m) {
			m = fs[i]
		}
	}
	return m
}

func Max(fs ...*Float) *Float {
	m := fs[0]
	for i := 1; i < len(fs); i++ {
		if Gt(fs[i], m) {
			m = fs[i]
		}
	}
	return m
}
