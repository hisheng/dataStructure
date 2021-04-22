go tool compile -S ch.go

go build

go tool objdump -s main.main ch0



>Go汇编语言提供了DATA命令用于初始化包变量，DATA命令的语法如下：
DATA symbol+offset(SB)/width, value
> 
var Id = 9527

我们采用以下命令可以给Id变量初始化为十六进制的0x2537，
对应十进制的9527（常量需要以美元符号$开头表示）：
> DATA	symbol+offset(SB)/width, value
>> 含义是：使用给定的value，
> 对相对于symbol偏移offset的地址，
> 宽度为width字节的内存区域进行初始化。
> 对于一个给定symbol上的一系列DATA指令，初始化时，偏移量必须是递增的。
> 不一定需要单步递增。


GLOBL symbol(SB), [flgs], width
GLOBL命令将一个符号声明为全局数据。
后面的参数是可选的标志，和符号的数据大小，
这个符号的初始值全为0，除非一个DATA指令对它进行了初始化。
GLOBL命令必须在对应的DATA命令后面。










