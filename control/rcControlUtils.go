package control

import (
	_ "strconv"
	_ "time"
)

func get_bytepack_plmnId(lInPlmnId []byte) []byte {

var tmpPlmnId []byte
var bytePackedPlmnId []byte

tmpPlmnId = make([]byte, 6)
bytePackedPlmnId = make([]byte, 3)

num_mnc_digits := len(lInPlmnId)

tmpPlmnId[0] = lInPlmnId[0]
tmpPlmnId[1] = lInPlmnId[1]
tmpPlmnId[2] = lInPlmnId[2]

if num_mnc_digits == 5 {
tmpPlmnId[3] = 0x0f
tmpPlmnId[4] = lInPlmnId[3]
tmpPlmnId[5] = lInPlmnId[4]
} else {
tmpPlmnId[3] = lInPlmnId[3]
tmpPlmnId[4] = lInPlmnId[4]
tmpPlmnId[5] = lInPlmnId[5]
}
bytePackedPlmnId[0] = ((tmpPlmnId[1] & 0x0F) << 4) | (tmpPlmnId[0] & 0x0F)
bytePackedPlmnId[1] = ((tmpPlmnId[3] & 0x0F) << 4) | (tmpPlmnId[2] & 0x0F)
bytePackedPlmnId[2] = ((tmpPlmnId[5] & 0x0F) << 4) | (tmpPlmnId[4] & 0x0F)

return bytePackedPlmnId
}

func ConverIntegerToByte(lIntVar int) []byte {
var rightMost, tempIntVar int
var byteArray []byte
tempIntVar = lIntVar
for {
rightMost = tempIntVar % 10
byteArray = append(byteArray, byte(rightMost)) // convert single digit to byte
tempIntVar /= 10
if tempIntVar == 0 {
break
}
}
// need to reverse the order
fixByteArray := []byte{}
for i := range byteArray {
n := byteArray[len(byteArray)-1-i]
fixByteArray = append(fixByteArray, n)
}

return fixByteArray
}
