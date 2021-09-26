package openrtb

func intRef(v int) (vv *int) {
	if v != 0 {
		vv = new(int)
		*vv = v
	}
	return vv
}
