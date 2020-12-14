package classfile

// 验证，目的是确保类或者接口的二进制表示在结构上是正确的
// 验证过程可能导致某些额外的类或者接口被加载进来，但不一定会导致它们也需要验证或准备
// 具体的验证范围包括一下几种：
// 1. 方法的访问控制
// 2. 参数和静态类型检查
// 3. 堆栈是否滥用
// 4. 变量是否初始化
// 5. 变量是否赋予正确类型
// 6. 异常表必须引用类合法的指令
// 7. 验证局部变量表
// 8. 逐一验证每个字节码的合法性
func verification() {
	// gvm 暂时先不做验证，埋个坑在这
}