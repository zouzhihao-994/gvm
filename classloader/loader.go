package classloader

import (
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/exception"
	"os"
	"path/filepath"
	"strings"
)

type Loader interface {
	Loading(fileName string) []byte
	AddZip(string)
}

/*
创建一个通配符类型Entry.
将path下的所有jar包文件的路径作为string创建一个zipEntry.
最后返回一个path路径下所有jar包构成的zipEntry的数组.
后续可以通过readClass直接对数组中的所有zipEntry类型的jar包进行遍历，从而搜索class文件.
*/
func jars(path string) []string {
	// 获取路径(不包含通配符字符),remove *
	baseDir := path[:len(path)-1]
	var jars []string
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是目录而非文件
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		// 如果后缀是jar包
		if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") {
			// 创建jar包类型的文件到compositeEntry中
			jars = append(jars, path)
		}
		return nil

	}
	// 遍历每一个
	filepath.Walk(baseDir, walkFn)
	return jars
}

/*
查找jre目录的路径.
*/
func getJreDirPath() string {
	jreOption := config.JrePath

	// 如果用户输入了-Xjre 参数
	if exists(jreOption) {
		return jreOption
	}
	/*
		如果用户没有输入 -Xjre 参数
		在 './jre' 下找
	*/
	if exists("./jre") {
		return "./jre"
	}

	/**
	没有输入 -Xjre参数,而且当前目录下也没有找到
	在 JAVA_HOME中找
	*/
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	exception.GvmError{Msg: "jre folder does't exist"}.Throw()
	return ""
}

/*
判断path目录是否存在.
*/
func exists(path string) bool {
	// stat获取项目中path的文件信息
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}