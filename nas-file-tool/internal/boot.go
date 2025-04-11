package internal

import "nas-file-tool/internal/menu"

func Registe() {
	menu.GetInstance().AddMenuItem("文件移动", filesMove)
	menu.GetInstance().AddMenuItem("子目录复制", copyFilesWithAllChildren)
}
