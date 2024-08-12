package storage

type storage interface {
	TargetPath()
	InitDir(args []string) bool //初始化对应的文件夹
}
