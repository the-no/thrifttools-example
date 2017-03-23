package main

import (
	"fmt"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"

	"git.oschina.net/theno/thrifttools"
	"git.oschina.net/theno/thrifttools-example/demo"

	"time"
)

const networkAddr string = "0.0.0.0:10086"

type idoallThrift struct {
}

func (*idoallThrift) CallBack(callTime int64, name string, paramMap map[int64]string) (r []string, err error) {
	fmt.Println("-->from client Call:", callTime, name, paramMap)
	r = append(r, "key:"+paramMap[1]+"    value:"+paramMap[2])
	time.Sleep(2 * time.Second)
	return
}

func (*idoallThrift) Put(s *demo.Student) (err error) {
	fmt.Printf("Stduent--->id: %d\tname:%s\tsex:%t\tage:%d\n", s.Sid, s.Sname, s.Ssex, s.Sage)
	return
}
func (*idoallThrift) Set(test int64) (err error) {
	return nil
}

func (*idoallThrift) GetM() (r string, err error) {
	return "M", nil
}

func (*idoallThrift) TestSet(tst *demo.Ttest, mp map[int32]bool, lt []int32) (r *demo.Ttest, err error) {

	return tst, nil
}

func (*idoallThrift) TestBool(b bool) (r bool, err error) {
	fmt.Println("TestBool :", b)
	return b, nil
}

func (*idoallThrift) TestByte(b int8) (r int8, err error) {
	fmt.Println("TestByte :", b)
	return b, nil
}

func (*idoallThrift) TestI16(i int16) (r int16, err error) {
	fmt.Println("TestI16 :", i)
	return i, nil
}

func (*idoallThrift) TestI32(i int32) (r int32, err error) {
	fmt.Println("TestI32 :", i)
	return i, nil
}

func (*idoallThrift) TestI64(i int64) (r int64, err error) {
	fmt.Println("TestI64 :", i)
	return i, nil
}

func (*idoallThrift) TestDouble(d float64) (r float64, err error) {
	fmt.Println("TestDouble : ", d)
	return d, nil
}

func (*idoallThrift) TestString(str string) (r string, err error) {
	fmt.Println("TestString :", str)
	return str, nil
}

func (*idoallThrift) TestListD(ds []float64) (r []float64, err error) {
	fmt.Println("TestListD :", ds)
	return ds, nil
}

func (*idoallThrift) TestListS(ss []*demo.Student) (r []*demo.Student, err error) {
	fmt.Println("TestListS :", ss)
	return ss, nil
}

func (*idoallThrift) TestListSL(ssl [][]*demo.Student) (r [][]*demo.Student, err error) {
	fmt.Println("TestListSL :", ssl)
	return ssl, nil
}

func (*idoallThrift) TestListDL(dsl [][]float64) (r [][]float64, err error) {
	fmt.Println("TestListDL :", dsl)
	return dsl, nil
}

func (*idoallThrift) TestVoid() (err error) {
	return fmt.Errorf("test Error")
}
func (*idoallThrift) TestVoidError() (err error) {
	return &demo.SError{"1", "test", nil}
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(networkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &idoallThrift{}
	/* apache
	processor := demo.NewDemoThriftProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	*/

	/*   v1.0
	processor := demo.NewDemoThriftProcessor(handler)
	midware := thrifttools.NewThriftMidWare(processor)
	server := thrift.NewTSimpleServer4(midware, serverTransport, transportFactory, protocolFactory)
	*/
	// v2.0
	midware := thrifttools.NewThriftMidWare(handler)
	midware.Use(timeRecorder)
	server := thrift.NewTSimpleServer4(midware, serverTransport, transportFactory, protocolFactory)

	fmt.Println("thrift server in", networkAddr)
	//server.Serve()
	fmt.Println(server.Serve())
}

func timeRecorder(c *thrifttools.Context) {
	s := time.Now().UnixNano()
	c.Next()
	f := time.Now().UnixNano()
	fmt.Printf("TimeRecorder [%s] Time: %dms\n", c.Name, (f-s)/1000000)
}
