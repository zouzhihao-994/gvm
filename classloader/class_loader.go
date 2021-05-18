package classloader

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/config"
	"github.com/zouzhihao-994/gvm/exception"
	"sync"
)

var once sync.Once
var BSCLoader *BootStrapLoader
var EXCLoader *ExtensionLoader
var GSCLoader *ApplicationLoader
var APPLoader *ApplicationLoader

// InitClassLoader 初始化类加载器
func InitClassLoader() {
	once.Do(func() {
		BSCLoader = newBootStrapLoader()
		EXCLoader = newExtensionLoader()
		GSCLoader = newApplicationLoader(config.VMNativePathDefault)
		APPLoader = newApplicationLoader(config.ClassPath)
	})
}

// Loading 加载字节码文件到方法区 Perm 中
// 加载顺序依次为 BootStrapLoader 、 ExtensionLoader 、 ApplicationLoader
// 《dynamic class loading in the java virtual machine》 url: https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.18.762&rep=rep1&type=pdf
// @param fileName 类名
func Loading(fileName string) []byte {
	fileName = fileName + ".class"
	fmt.Println("loadding calss file -> " + fileName)
	var data []byte

	// 从启动类加载器中加载
	if data = BSCLoader.Loading(fileName); data != nil {
		return data
	}

	// 从扩展类加载器中加载
	if data = EXCLoader.Loading(fileName); data != nil {
		return data
	}

	// 从Gvm系统库中加载
	if data = GSCLoader.Loading(fileName); data != nil {
		return data
	}

	// 从用户类加载器中加载
	if data = APPLoader.Loading(fileName); data != nil {
		return data
	}

	exception.GvmError{Msg: "classfile not found"}.Throw()
	return nil
}
