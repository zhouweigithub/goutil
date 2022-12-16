package guidutil

import (
	uuid "github.com/satori/go.uuid"
)

// var objectIdCounter uint32 = 0

// var machineId = readMachineId()

// type ObjectId string

// func readMachineId() []byte {
// 	var sum [3]byte
// 	id := sum[:]
// 	hostname, err1 := os.Hostname()
// 	if err1 != nil {
// 		_, err2 := io.ReadFull(rand.Reader, id)
// 		if err2 != nil {
// 			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
// 		}
// 		return id
// 	}
// 	hw := md5.New()
// 	hw.Write([]byte(hostname))
// 	copy(id, hw.Sum(nil))
// 	fmt.Println("readMachineId:" + string(id))
// 	return id
// }

// // GUID returns a new unique ObjectId.
// // 4byte 时间，
// // 3byte 机器ID
// // 2byte pid
// // 3byte 自增ID
// func GetGUID() ObjectId {
// 	var b [12]byte
// 	// Timestamp, 4 bytes, big endian
// 	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
// 	// Machine, first 3 bytes of md5(hostname)
// 	b[4] = machineId[0]
// 	b[5] = machineId[1]
// 	b[6] = machineId[2]
// 	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
// 	pid := os.Getpid()
// 	b[7] = byte(pid >> 8)
// 	b[8] = byte(pid)
// 	// Increment, 3 bytes, big endian
// 	i := atomic.AddUint32(&objectIdCounter, 1)
// 	b[9] = byte(i >> 16)
// 	b[10] = byte(i >> 8)
// 	b[11] = byte(i)
// 	return ObjectId(b[:])
// }

// // Hex returns a hex representation of the ObjectId.
// // 返回16进制对应的字符串
// func (id ObjectId) Hex() string {
// 	return hex.EncodeToString([]byte(id))
// }

// // 生成随机GUID，会重复
// func GetRandGUID(hasSplitChar bool) string {
// 	b := make([]byte, 16)
// 	_, err := rand.Read(b)
// 	if err != nil {
// 		return ""
// 	}
// 	if hasSplitChar {
// 		return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
// 	} else {
// 		return fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
// 	}
// }

// Returns canonical string representation of UUID:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func NewGUID() string {
	return uuid.NewV4().String()
}
