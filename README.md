[![go](https://img.shields.io/badge/Go-1.17.3-blue)](https://github.com/workcha/goScanner)

## ChineseCommandExec
解决go语言在中文windows上面执行命令出现乱码的问题

demo代码
    
    package main
    
    import "github.com/workcha/ChineseCommandExec/command"
    
    func main() {
        result := command.CommonExecute("ipconfig /all")
        println(result)
    }


![demo](img/1.png)