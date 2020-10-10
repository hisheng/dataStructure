package redis


type sds string
type  sdshdr struct{
	len int
	free int
	buf string
}

func sdsnewlen(init interface{},initlen int) sds {
	var sh *sdshdr = &sdshdr{}
	sh.len = initlen
	sh.free = 0
	sh.buf = init.(string)
	return sds(sh.buf)
}

func sdsenpty() sds {
	return sdsnewlen("",0)
}

func sdslen(s sds) int  {
	return len(s)
}

func sdsdup(s sds)  {
	return sdsnewlen(s, sdslen(s))
}




