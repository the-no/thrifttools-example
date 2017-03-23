package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"git.oschina.net/theno/thrifttools-example/demo"

	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	HOST = "127.0.0.1"
	PORT = "10086"
)

func main() {
	startTime := currentTimeMillis()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := protocol.NewTBinaryProtocolFactory()
	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := demo.NewDemoThriftClientFactory(useTransport, protocolFactory)
	/*inpor := client.InputProtocol.(*protocol.Protocol) //
	inpor.ReadEndAction = func() { fmt.Println("Read End") }*/
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	/*
		for i := 0; i < 3; i++ {
			paramMap := make(map[int64]string)
			paramMap[1] = "idoall"
			paramMap[2] = "org" + strconv.Itoa(i+1)
			tmp := time.Now().UnixNano() / 1000000
			r1, _ := client.CallBack(tmp, "go client", paramMap)
			fmt.Println("GOClient Call->", r1)
			fmt.Println(client.GetM())
		}
	*/
	/*model := demo.Student{11, "student-idoall-go", true, 20}
	client.Put(&model)
	fmt.Println(client.GetM())*/

	tst := &demo.Ttest{

		Isbool: false,
		Str:    "test",
		Ti8:    8,
		Ti16:   16,
		Ti32:   32,
		Ti64:   64,
		Tbl:    65.4,
		// unused field # 8
		Lstr: []string{"1", "2"},
		Mp:   map[int32]bool{1: true, 2: false},
	}
	mp := map[int32]bool{}
	lt := []int32{}
	fmt.Println(client.TestSet(tst, mp, lt))
	endTime := currentTimeMillis()
	fmt.Printf("本次调用用时:%d-%d=%d毫秒\n", endTime, startTime, (endTime - startTime))

	fmt.Print("TestBool: ")
	fmt.Println(client.TestBool(true))

	fmt.Print("TestByte: ")
	fmt.Println(client.TestByte(8))

	fmt.Print("TestI16: ")
	fmt.Println(client.TestI16(16))

	fmt.Print("TestI32: ")
	fmt.Println(client.TestI32(32))

	fmt.Print("TestI64: ")
	fmt.Println(client.TestI64(64))

	fmt.Print("TestDouble: ")
	fmt.Println(client.TestDouble(64.64))
	fmt.Print("TestString: ")
	fmt.Println(client.TestString("TestString"))

	fl := []float64{1.1, 2.2, 3.3, 4.4}
	fmt.Print("TestListD: ")
	fmt.Println(client.TestListD(fl))

	sts := []*demo.Student{
		{1, "1", false, 1},
		{2, "2", true, 2},
		{3, "3", false, 3},
		{4, "4", false, 4},
	}
	fmt.Print("TestListS: ")
	fmt.Println(client.TestListS(sts))

	stsl := [][]*demo.Student{
		{{1, "1", false, 1}, {1, "1", true, 1}},
		{{2, "2", true, 2}, {2, "2", false, 2}},
		{{3, "3", false, 3}, {3, "3", true, 3}},
		{{4, "4", true, 4}, {4, "4", false, 4}},
	}
	fmt.Print("TestListSL: ")
	fmt.Println(client.TestListSL(stsl))

	fsl := [][]float64{
		{1.1, 2.2, 3.3, 4.4},
		{5.5, 6.6, 7.7, 8.8},
		{9.9, 10.10, 11.11, 12.12},
	}
	fmt.Print("TestListDL: ")
	fmt.Println(client.TestListDL(fsl))

	fmt.Println("TestVoid :", client.TestVoid())
	fmt.Println("TestVoidError: ", client.TestVoidError())
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
