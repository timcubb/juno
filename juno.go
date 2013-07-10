package main

import(
	"fmt"
	"github.com/timcubb/juno/system"
    "github.com/timcubb/juno/deploy"
	)


func main() {
    
    node1 := system.Node{"node1.pfy.com", 8001, "App1003", "new"}
    node2 := system.Node{"node2.pfy.com", 8002, "App1043", "new"}
    node3 := system.Node{"node3.pfy.com", 8003, "App1553", "new"}
    
    sy := system.NewSystem()
    
    sy.AddNode("node1", node1)
    sy.AddNode("node2", node2)
    sy.AddNode("node3", node3)
    
    fmt.Println(sy);

    gn1 := sy.GetNode("node1")
    gn2 := sy.GetNode("node3")
    gn3 := sy.GetNode("node2")

    fmt.Println(gn1);
    fmt.Println(gn2);
    fmt.Println(gn3);
}
