package main


import(
    "fmt"
    "time"
    "github.com/fatih/color"
    "net"
)


func Check(destination string , port string) string{
    address := destination+":"+port

    timeout := time.Duration(5 * time.Second)
    conn,err := net.DialTimeout("tcp",address,timeout)
    var status string
    
    red:= color.New(color.BgRed).Sprint("[DOWN]")
    green := color.New(color.BgGreen).Sprint("[UP]")
    if err != nil{
        status = fmt.Sprintf("%s %v is unreachable,\nError :%v ",red,destination , err)

    }else{
        status = fmt.Sprintf("%s %v is reachable,\nFrom: %v \nTo: %v ",green,destination , conn.LocalAddr() , conn.RemoteAddr())
    }
    return status 
}
