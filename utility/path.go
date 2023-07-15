package utility

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func GetProjectAbsRootPath() (rootPath string) {
    var(
        selfPath string
        pathArr []string
    )

    selfPath = gfile.Pwd()
    pathArr = gstr.Split(selfPath, "guoshao-fm-web")
    rootPath = pathArr[0] + "guoshao-fm-web/"

    return
}
