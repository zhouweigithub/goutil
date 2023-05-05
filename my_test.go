package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/zhouweigithub/goutil/cacheutil"
	"github.com/zhouweigithub/goutil/cmdutil"
	"github.com/zhouweigithub/goutil/compressutil"
	"github.com/zhouweigithub/goutil/configutil"
	"github.com/zhouweigithub/goutil/encryptutil"
	"github.com/zhouweigithub/goutil/excelutil"
	"github.com/zhouweigithub/goutil/fileutil"
	"github.com/zhouweigithub/goutil/guidutil"
	"github.com/zhouweigithub/goutil/iputil"
	"github.com/zhouweigithub/goutil/jsutil"
	"github.com/zhouweigithub/goutil/qrcodeutil"
	"github.com/zhouweigithub/goutil/randutil"
	"github.com/zhouweigithub/goutil/screenshotutil"
	"github.com/zhouweigithub/goutil/setutil"
	"github.com/zhouweigithub/goutil/sliceutil"
	"github.com/zhouweigithub/goutil/socketutil/udputil"
	"github.com/zhouweigithub/goutil/stringutil"
	"github.com/zhouweigithub/goutil/threadutil"
	"github.com/zhouweigithub/goutil/webutil"
)

type ConModel struct {
	Name string
	Sex  string
	Age  int
}

func TestConfig(t *testing.T) {
	var no = ConModel{}
	var err = configutil.ToModel(&no)
	t.Log(no, err)
}

func TestThreading(t *testing.T) {
	var sources = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("start")
	fmt.Println(sources)
	threadutil.Threading(sources, 2, func(item *int) {
		//fmt.Println(*item)
		*item = *item + 10
	})
	fmt.Println(sources)
	fmt.Println("over")
}

type Model struct {
	Name string
	Age  int
}

var sources = []Model{}

func init() {
	sources = append(sources, Model{Name: "liming1", Age: 10})
	sources = append(sources, Model{Name: "liming3", Age: 12})
	sources = append(sources, Model{Name: "liming4", Age: 13})
	sources = append(sources, Model{Name: "liming5", Age: 11})
	sources = append(sources, Model{Name: "liming6", Age: 15})
	sources = append(sources, Model{Name: "liming7", Age: 16})
	sources = append(sources, Model{Name: "liming1", Age: 10})
}

//var ints = []int{1, 4, 2, 1, 5, 3, 0, 4}

func TestFilter(t *testing.T) {
	// var a = sliceutil.First(sources, func(item *Model) bool { return item.Age > 13 })
	// a.Name = "hello world"
	// fmt.Println(*a)
	// fmt.Println(sources[4])
	// var b = sliceutil.Last(sources, func(item *Model) bool { return item.Name == "liming5" })
	// fmt.Println(*b)
	// var c = sliceutil.Contains(sources, func(item *Model) bool { return item.Age == 18 })
	// fmt.Println(c)
	// var d = sliceutil.Where(sources, func(item *Model) bool { return item.Age < 15 })
	// d[0].Name = "hello world"
	// fmt.Println(*d[0])
	// fmt.Println(sources[0])
	// var e = sliceutil.Select(sources, func(item *Model) string { return item.Name })
	// fmt.Println(e)

	// var x = sliceutil.Distinct(sources, func(item *Model) int { return item.Age })
	// fmt.Println(x)

	//sliceutil.OrderBy(sources, func(i, j int) bool { return sources[i].Age < sources[j].Age })

	//var a = sliceutil.Distinct(sources)
	var a = sliceutil.Remove(sources, func(item *Model) bool { return item.Age == 10 })
	fmt.Println(a)
}

func TestSimilarity(t *testing.T) {
	fmt.Println(stringutil.Similarity("hello world", "hello owrdl"))
}

func TestGetRandChar(t *testing.T) {
	fmt.Println(randutil.GetRandChars(20))
	fmt.Println(randutil.GetRandChars(20))
	fmt.Println(randutil.GetRandChars(20))
	fmt.Println(randutil.GetRandChars(20))
	fmt.Println(randutil.GetRandChars(20))
}

func TestSliceAppend(t *testing.T) {
	// s := []int{5}
	// s = append(s, 7)
	// fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])
	// fmt.Println(s)
	// s = append(s, 9)
	// fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])
	// fmt.Println(s)
	// x := append(s, 11)
	// fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(x) =", &x[0])
	// fmt.Println(s)
	// fmt.Println(x)
	// y := append(s, 12)
	// fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(y) =", &y[0])
	// fmt.Println(s)
	// fmt.Println(x)
	// fmt.Println(y)

	a := []int{1, 2, 3}
	fmt.Println(len(a), cap(a))
	b := append(a, 4)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
}

func TestEncp(t *testing.T) {
	// var a, b = encryptutil.AESEncodeStr("123456", "hello123hello123")
	// fmt.Println(a, b)
	// fmt.Println(encryptutil.AESDecodeStr(a, "hello123hello123"))
	// fmt.Println(encryptutil.HMAC_SHA1("HELLO", "WORLD"))
	// fmt.Println(encryptutil.HMAC_SHA256("HELLO", "WORLD"))
	// fmt.Println(encryptutil.HMAC_SHA512("HELLO", "WORLD"))
	// fmt.Println(encryptutil.SHA256Str("HELLO"))
	// fmt.Println(encryptutil.SHA512Str("HELLO"))

	var a = encryptutil.ToUnicode("支持ASCII编码与字符的相互转换。ABCD,./!")
	fmt.Println(a)
	fmt.Println(encryptutil.FromUnicode(a))
}

func TestSlice(t *testing.T) {
	// var a = []int{} // !=nil
	// var b []int     // ==nil
	// fmt.Println(a == nil, b == nil)
	// fmt.Println(a, b)
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// var b = []int{3, 5, 1, 9, 2, 100}
	//fmt.Println(sliceutil.Exclude(a, b))
	// var c = sliceutil.Union(a, b)
	// var d = sliceutil.Distinct(c)
	// fmt.Println(c, d)

	sliceutil.ForEach(a, func(item *int) { *item = *item + 10 })
	fmt.Println(a)

	// var a1 = sliceutil.Where(sources, func(item *Model) bool { return item.Age > 13 })
	// var a2 = sliceutil.WhereReference(sources, func(item *Model) bool { return item.Age > 13 })

	// a1[0].Name = "chen"
	// a2[1].Name = "where"
}
func TestZip(t *testing.T) {
	fmt.Println(compressutil.Zip("zip.zip", "goutil.exe", "git_tag.txt"))
	compressutil.Unzip("zip.zip", "zipfolder")
}

func TestGuid(t *testing.T) {
	for i := 0; i < 5; i++ {
		// var a = guidutil.GetGUID()
		// fmt.Println(a.Hex())
		fmt.Println(guidutil.NewGUID())
	}
}

func TestQrcode(t *testing.T) {
	fmt.Println(qrcodeutil.CreatePngFile("http://promotion.79yougame.com/char.html", 200, "pngfile.png"))
	//qrcodutil.CreateQrcodePngBytes()
}

func TestExcel(t *testing.T) {
	var a, err = excelutil.ReadFromExcel(`C:\Users\juscc\Desktop\hello.xlsx`, "")
	//fmt.Println(a)
	fmt.Println(err)

	fmt.Println(excelutil.WriteToExcel(`C:\Users\juscc\Desktop\777.xlsx`, a, "hello"))
}

func TestCache(t *testing.T) {
	var c = cacheutil.New()
	c.Set("hello", "world", time.Second*5)
	c.Set("gogogo", "alalala", time.Second*10)
	fmt.Println(c.Get("hello"), c.Len())
	time.Sleep(time.Second * 3)
	c.Delete("gogogo")
	fmt.Println(c.Get("hello"), c.Len())
	time.Sleep(time.Second * 3)
	fmt.Println(c.Get("hello"), c.Len())
}

func TestIP(t *testing.T) {
	var a, err = iputil.GetLocalIp()
	fmt.Println(a, err)
}

func TestSet(t *testing.T) {
	var set = setutil.NewSet[int]()
	set.Add(3, 4, 5, 6, 3, 2, 5, 1, 4, 9)
	fmt.Println(set.Exists(3))
	fmt.Println(set.Exists(4))
	fmt.Println(set.Exists(5))
	set.Delete(3, 7, 6, 1)
	fmt.Println(set.Exists(3))
	fmt.Println(set.Exists(4))
	fmt.Println(set.Exists(5))
	fmt.Println(set.GetAll())
	fmt.Println(set.GetAllSorted())
}

func TestCmd(t *testing.T) {
	//fmt.Println(cmdutil.ExecuteCmd("ping", "www.baidu.com"))
	//fmt.Println(cmdutil.ExecuteCmdInFolder("d:\\", "ipconfig"))
	//fmt.Println(cmdutil.ExecuteCmdInFolderToShow("", "ping", "www.qq.com"))
	//fmt.Println(cmdutil.ExecuteCmdInFolderToFile("d:\\", "logs/mylog.txt", "ping", "www.sina.com"))

	fmt.Println(cmdutil.ExecuteCmd("cmd.exe", "/c", "dir"))
	fmt.Println(cmdutil.ExecuteCmdInFolder("d:/", "cmd.exe", "/c", "dir"))

	// cmd := exec.Command("cmd.exe", "/c", "dir")
	// //cmd.Dir = "E:/svn"
	// r, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// r, _ = simplifiedchinese.GBK.NewDecoder().Bytes(r)
	// fmt.Println(string(r))
}

func TestByte(t *testing.T) {
	var str = "你好w"
	var b = []byte(str)
	var c = []rune(str)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(string(b))
	fmt.Println(string(c))
}

func TestFile(t *testing.T) {
	//fileutil.WriteTextFile("d:\\hello.txt", "hello world6")
	fileutil.AppendTextFile("d:/hello.txt", "hello world3")
}

func TestCopyFile(t *testing.T) {
	var maps map[string]string
	for k, v := range maps {
		fmt.Println(k, v)
	}
}

func TestRequest(t *testing.T) {
	var urls = "http://baidu.com/DayLog/Index?a=b&b=c"
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["header1"] = "one header"
	var cookies = make(map[string]string)
	cookies["cookit1"] = "oh,this a cookie"
	cookies["cookit2"] = "oh,this t cookie"
	var _, bbb, ccc, ddd = webutil.GetWeb(urls, headers, cookies, "http://10.254.0.191:8888", 2)
	fmt.Println(bbb, ccc, ddd)
}

func TestRSA(t *testing.T) {
	applyPubEPriD()
}

var publicKey = `-----BEGIN Public key-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAk+89V7vpOj1rG6bTAKYM
56qmFLwNCBVDJ3MltVVtxVUUByqc5b6u909MmmrLBqS//PWC6zc3wZzU1+ayh8xb
UAEZuA3EjlPHIaFIVIz04RaW10+1xnby/RQE23tDqsv9a2jv/axjE/27b62nzvCW
eItu1kNQ3MGdcuqKjke+LKhQ7nWPRCOd/ffVqSuRvG0YfUEkOz/6UpsPr6vrI331
hWRB4DlYy8qFUmDsyvvExe4NjZWblXCqkEXRRAhi2SQRCl3teGuIHtDUxCskRIDi
aMD+Qt2Yp+Vvbz6hUiqIWSIH1BoHJer/JOq2/O6X3cmuppU4AdVNgy8Bq236iXvr
MQIDAQAB
-----END Public key-----`

var privateKey = `-----BEGIN Private key-----
MIIEpAIBAAKCAQEAk+89V7vpOj1rG6bTAKYM56qmFLwNCBVDJ3MltVVtxVUUByqc
5b6u909MmmrLBqS//PWC6zc3wZzU1+ayh8xbUAEZuA3EjlPHIaFIVIz04RaW10+1
xnby/RQE23tDqsv9a2jv/axjE/27b62nzvCWeItu1kNQ3MGdcuqKjke+LKhQ7nWP
RCOd/ffVqSuRvG0YfUEkOz/6UpsPr6vrI331hWRB4DlYy8qFUmDsyvvExe4NjZWb
lXCqkEXRRAhi2SQRCl3teGuIHtDUxCskRIDiaMD+Qt2Yp+Vvbz6hUiqIWSIH1BoH
Jer/JOq2/O6X3cmuppU4AdVNgy8Bq236iXvrMQIDAQABAoIBAQCCbxZvHMfvCeg+
YUD5+W63dMcq0QPMdLLZPbWpxMEclH8sMm5UQ2SRueGY5UBNg0WkC/R64BzRIS6p
jkcrZQu95rp+heUgeM3C4SmdIwtmyzwEa8uiSY7Fhbkiq/Rly6aN5eB0kmJpZfa1
6S9kTszdTFNVp9TMUAo7IIE6IheT1x0WcX7aOWVqp9MDXBHV5T0Tvt8vFrPTldFg
IuK45t3tr83tDcx53uC8cL5Ui8leWQjPh4BgdhJ3/MGTDWg+LW2vlAb4x+aLcDJM
CH6Rcb1b8hs9iLTDkdVw9KirYQH5mbACXZyDEaqj1I2KamJIU2qDuTnKxNoc96HY
2XMuSndhAoGBAMPwJuPuZqioJfNyS99x++ZTcVVwGRAbEvTvh6jPSGA0k3cYKgWR
NnssMkHBzZa0p3/NmSwWc7LiL8whEFUDAp2ntvfPVJ19Xvm71gNUyCQ/hojqIAXy
tsNT1gBUTCMtFZmAkUsjqdM/hUnJMM9zH+w4lt5QM2y/YkCThoI65BVbAoGBAMFI
GsIbnJDNhVap7HfWcYmGOlWgEEEchG6Uq6Lbai9T8c7xMSFc6DQiNMmQUAlgDaMV
b6izPK4KGQaXMFt5h7hekZgkbxCKBd9xsLM72bWhM/nd/HkZdHQqrNAPFhY6/S8C
IjRnRfdhsjBIA8K73yiUCsQlHAauGfPzdHET8ktjAoGAQdxeZi1DapuirhMUN9Zr
kr8nkE1uz0AafiRpmC+cp2Hk05pWvapTAtIXTo0jWu38g3QLcYtWdqGa6WWPxNOP
NIkkcmXJjmqO2yjtRg9gevazdSAlhXpRPpTWkSPEt+o2oXNa40PomK54UhYDhyeu
akuXQsD4mCw4jXZJN0suUZMCgYAgzpBcKjulCH19fFI69RdIdJQqPIUFyEViT7Hi
bsPTTLham+3u78oqLzQukmRDcx5ddCIDzIicMfKVf8whertivAqSfHytnf/pMW8A
vUPy5G3iF5/nHj76CNRUbHsfQtv+wqnzoyPpHZgVQeQBhcoXJSm+qV3cdGjLU6OM
HgqeaQKBgQCnmL5SX7GSAeB0rSNugPp2GezAQj0H4OCc8kNrHK8RUvXIU9B2zKA2
z/QUKFb1gIGcKxYr+LqQ25/+TGvINjuf6P3fVkHL0U8jOG0IqpPJXO3Vl9B8ewWL
cFQVB/nQfmaMa4ChK0QEUe+Mqi++MwgYbRHx1lIOXEfUJO+PXrMekw==
-----END Private key-----`

// Public key encryption private key decryption
func applyPubEPriD() {
	var s = encryptutil.Rsa()
	{
		fmt.Println("公钥加密，私钥解密")
		var a, _ = s.PublicEncrypt("hello", publicKey)
		fmt.Println(a)
		var b, _ = s.PriKeyDecrypt(a, privateKey)
		fmt.Println(b)
	}
	{
		fmt.Println("私钥加密，公钥解密")
		var a, _ = s.PriKeyEncrypt("hello", privateKey)
		fmt.Println(a)
		var b, _ = s.PublicDecrypt(a, publicKey)
		fmt.Println(b)
	}
	{
		fmt.Println("使用RSAWithMD5算法签名")
		var a, _ = s.SignMd5WithRsa("hello", privateKey)
		fmt.Println(a)
		var b = s.VerifySignMd5WithRsa("hello", a, publicKey)
		fmt.Println(b)
	}
	{
		fmt.Println("使用RSAWithSHA1算法签名")
		var a, _ = s.SignSha1WithRsa("hello", privateKey)
		fmt.Println(a)
		var b = s.VerifySignSha1WithRsa("hello", a, publicKey)
		fmt.Println(b)
	}
	{
		fmt.Println("使用RSAWithSHA256算法签名")
		var a, _ = s.SignSha256WithRsa("hello", privateKey)
		fmt.Println(a)
		var b = s.VerifySignSha256WithRsa("hello", a, publicKey)
		fmt.Println(b)
	}
}

func TestJs(t *testing.T) {
	var js jsMap
	if err := js.LoadFile("test.js"); err != nil {
		fmt.Println(err.Error())
	} else {
		var a, err = js.AddInt()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(a(3, 5))
		}
	}
}

type jsMap struct {
	jsutil.JsHelper
}

func (j *jsMap) AddInt() (fn func(int, int) int, err error) {
	var tmp = j.Vm.Get("addInt")
	if tmp == nil {
		err = errors.New("Js函数 addInt 映射到 Go 函数失败！js中未找到该函数 addInt")
	} else {
		err = j.Vm.ExportTo(tmp, &fn)
		if err != nil {
			err = errors.New("Js函数 addInt 映射到 Go 函数失败！\n" + err.Error())
		}
	}
	return
}

func TestUdpClient(t *testing.T) {
	var cli = udputil.CreateClient("10.254.0.191", 3387)

	//for i := 0; i < 5; i++ {
	var msg = "hello , boy [" + strconv.Itoa(1) + "]"
	if err := cli.SendUdp(msg); err != nil {
		fmt.Println("send err:" + err.Error())
	} else {
		fmt.Println("client send data: " + msg)
	}
	time.Sleep(time.Second)
	//}

	cli.ReceiveUdp(func(data []byte, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("client receive data from %s:%s", cli.RemoteIp, string(data))
		}
	})
}
func TestUdpServer(t *testing.T) {
	var ser = udputil.CreateServer(3387)
	var lis = ser.ListenUdp(func(data []byte, remoteAddr *net.UDPAddr, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("server receive data from %s:%s\n", remoteAddr.String(), string(data))
		}
		var msg = string(data) + " // back to you"
		if err := ser.SendUdp(remoteAddr.String(), msg); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("server send data: " + msg)
		}
	})

	if lis == nil {
		fmt.Printf("listing %s\n", ser.LocalAddress)
	} else {
		fmt.Println(lis)
	}

	if err := ser.SendToAnyUdp("10.254.0.19", 6654, "hello myself"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("server send data over ")
	}
}

func TestSceenShot(t *testing.T) {
	var sc = screenshotutil.New(`C:\Users\Me\AppData\Local\Google\Chrome\Application\chrome.exe`, "imgs")
	sc.ScreenShot(`https://www.baidu.com`)
}
