package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhouweigithub/goutil/cacheutil"
	"github.com/zhouweigithub/goutil/cmdutil"
	"github.com/zhouweigithub/goutil/compressutil"
	configutil "github.com/zhouweigithub/goutil/configUtil"
	"github.com/zhouweigithub/goutil/encryptutil"
	"github.com/zhouweigithub/goutil/excelutil"
	"github.com/zhouweigithub/goutil/fileutil"
	"github.com/zhouweigithub/goutil/guidutil"
	"github.com/zhouweigithub/goutil/iputil"
	"github.com/zhouweigithub/goutil/qrcodutil"
	"github.com/zhouweigithub/goutil/randutil"
	"github.com/zhouweigithub/goutil/setutil"
	sliceutil "github.com/zhouweigithub/goutil/sliceUtil"
	"github.com/zhouweigithub/goutil/stringutil"
	"github.com/zhouweigithub/goutil/threadutil"
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

var ints = []int{1, 4, 2, 1, 5, 3, 0, 4}

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
	fmt.Println(qrcodutil.CreatePngFile("http://promotion.79yougame.com/char.html", 200, "pngfile.png"))
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
	fileutil.CreateFolderIfNotExists("\\los")
}

func TestCopyFile(t *testing.T) {
	//fmt.Println(fileutil.CopyFile("C:/backup/notebook_202301031618.zip", "d:/zzz/x.zip"))
	fmt.Println(fileutil.CopyFolder("C:/backup", "d:/backup"))
}
