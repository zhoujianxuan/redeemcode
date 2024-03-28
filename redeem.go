package redeemcode

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
)

// 定义兑换码的有效字符集，去掉O和0，1和i
var alphaMap = map[int64]string{
	0:  "A",
	1:  "B",
	2:  "C",
	3:  "D",
	4:  "E",
	5:  "F",
	6:  "G",
	7:  "H",
	8:  "J",
	9:  "K",
	10: "L",
	11: "M",
	12: "N",
	13: "P",
	14: "Q",
	15: "R",
	16: "X",
	17: "T",
	18: "U",
	19: "V",
	20: "W",
	21: "S",
	22: "Y",
	23: "Z",
	24: "2",
	25: "3",
	26: "4",
	27: "5",
	28: "6",
	29: "7",
	30: "8",
	31: "9",
}

// Gen 根据num生成兑换码，10位长度
func Gen(num int) (string, string) {
	// 自增id，占32位，1byte=4bit
	incrBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(incrBytes, uint32(num))

	sign := Sign(num)
	fresh := BytesToBinaryString([]byte{byte(rand.Intn(16))})[4:]

	codeList := sign + BytesToBinaryString(incrBytes) + fresh
	return ToCode(codeList), codeList
}

// ToCode 通过5位bit转0~31
func ToCode(codeList string) string {
	code := ""
	// 5个二进制位就是0~31
	for i := 0; i < len(codeList)/5; i++ {
		bs := codeList[i*5 : (i+1)*5]
		c, _ := strconv.ParseInt(bs, 2, 64)
		code += alphaMap[c]
	}
	return code
}

func BytesToBinaryString(bs []byte) string {
	buf := bytes.NewBuffer([]byte{})
	for _, v := range bs {
		buf.WriteString(fmt.Sprintf("%08b", v))
	}
	return buf.String()
}

func Sign(num int) string {
	bl := make([]byte, 4)
	binary.BigEndian.PutUint32(bl, uint32(num))
	str := BytesToBinaryString(bl)
	length := len(str) / 8
	var sum int
	var strList []string
	for i := 0; i < 8; i++ {
		c, _ := strconv.ParseInt(str[i*length:(i+1)*length], 2, 64)
		strList = append(strList, str[i*length:(i+1)*length])

		weight := rand.Intn(6555 / 8)
		sum += int(c) + weight
	}

	// 采集数据进行sign
	signBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(signBytes, uint16(sum))
	return BytesToBinaryString(signBytes)[2:]
}
