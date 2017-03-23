namespace go demo

struct Student{
	1: i32 sid 
	2: string sname
	3: bool ssex=0
	4: i16 sage
}

struct Ttest{
	1:bool isbool
	2:string str
	3:byte  ti8
	4:i16    ti16
	5:i32   ti32
	6:i64	ti64
	7:double tbl
	9:list<string>  lstr
	//10:optional set<i32>  sets
	10:map<i32,bool> mp
	11:list<byte> bs
 }

 exception SError {
    1:string code
    2:string message
    3:map<string, string> extra
}

const map<i64,string> MAPCONSTANT = {1:'world', 2:'moon'}

service DemoThrift {        
	        list<string> CallBack(1:i64 callTime, 2:string name, 3:map<i64, string> paramMap);
			void Put(1: Student s);
			void Set(1: required i64 test);
			string GetM();
			Ttest TestSet(1:Ttest tst,2:map<i32,bool> mp,3:list<i32> lt);

			bool TestBool(1:bool b);
			byte TestByte(1:byte b);
			i16 TestI16(1:i16 i);
			i32 TestI32(1:i32 i);
			i64 TestI64(1:i64 i);
			double TestDouble(1:double d);
			string TestString(1:string str);

			list<double> TestListD(1:list<double> ds);
			//list<byte> TestListB(1:list<byte> bs);
			list<Student> TestListS(1:list<Student> ss);
			list<list<Student>> TestListSL(1:list<list<Student>> ssl);
			list<list<double>> TestListDL(1:list<list<double>> dsl);
			//list<list<byte>> TestListBL(1:list<list<byte>> bsl);
			void  TestVoid();
			void  TestVoidError() throws(1:SError serr);

}
