package security

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"github.com/pkg/errors"
	"strings"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetSha512(text string) string {
	hasher := sha512.New512_256()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetKey(mac string) (string, error) {
	dmi := New()
	if err := dmi.Run(); err != nil {
		//fmt.Printf("Unable to get dmidecode information. Error: %v\n", err)
		return "", err
	}

	//You can search by record name
	byNameData, byNameErr := dmi.SearchByName("System Information")
	UUID := ""
	if byNameErr == nil {
		//fmt.Println(byNameData[0]["UUID"])
		for _, record := range byNameData {
			uuid, exists := record["UUID"]
			if exists {
				UUID = uuid
			}
		}
		if UUID == "" {
			return "", errors.New("Not Found :)")
		}
	}
	//uuid := "7E6B2E00-D7E0-11DD-A11E-D017C297B551"

	t := GetSha512(UUID) + GetSha512(mac)

	t1 := GetSha512(t)
	t2 := GetMD5Hash(t1)
	t3 := GetMD5Hash(t2)
	t3 = strings.ToUpper(t3)
	//fmt.Println(t3)
	key := fmt.Sprintf("%s-%s-%s-%s-%s-%s", t3[:4], t3[4:10], t3[10:16], t3[16:20], t3[20:24], t3[24:32]) // 4-6-6-4-4-8
	return key, nil
}

func getUSBKey() (string, error) {
	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        115200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		//fmt.Println(err)
		//log.Fatalf("serial.Open: %v", err)
		return "", err
	}

	// Make sure to close it later.
	defer port.Close()

	//buffer := make([]byte, 200)
	buffer := []byte{}
	hasStart := false
	hasEnd := false
	iStart := 0
	iEnd := 0
	sizeBuffer := 0

	nReads := 10
	iRead := 0

	for {
		if iRead == nReads {
			break
		}
		iRead++;
		buf := make([]byte, 100)
		n, err := port.Read(buf)
		if err != nil {
			//if err != io.EOF {
			//	fmt.Println("Error reading from serial port: ", err)
			//}
			return "", err
		} else {
			buffer = append(buffer, buf[:n]...)
			sizeBuffer = sizeBuffer + n
			strBuffer := fmt.Sprintf("%s", buffer)
			hasStart = false
			i := 0
			for {
				if i == sizeBuffer {
					// incomplete data
					hasStart = false
					buffer = []byte{}
					sizeBuffer = 0
					break
				}
				if strBuffer[i] == '@' {
					hasStart = true
					iStart = i
					break
				}
				i++
			}
			//fmt.Println("buffer2: ", strBuffer)
			hasEnd = false
			if hasStart {
				i = iStart
				for {
					if i == sizeBuffer {
						// incomplete data
						//hasStart = false
						//buffer = []byte{}
						break
					}
					if strBuffer[i] == '\n' {
						hasEnd = true
						iEnd = i
						break
					}
					i++
				}
			}
			if hasStart && hasEnd {
				buffer = []byte{}
				sizeBuffer = 0
				return strBuffer[iStart+1 : iEnd], nil // without @
			} else if !hasStart && hasEnd {
				// error on data
				buffer = []byte{}
				sizeBuffer = 0
			}
		}
	}
	return "", errors.New("Failed")
}

func ValidateSerial(serialOk string) (bool, error) {
	mac, err := getUSBKey()
	if err != nil {
		//fmt.Println(err)
		//fmt.Println("error usb key")
		return false, errors.New("error usb key")
	}
	//keyOk := "9728-9C9A13-E206E3-2B4A-BFF0-4A4E5F1D"
	//mac := "30:AE:A4:23:39:1C"
	key, err := GetKey(mac)
	if err == nil {
		//fmt.Println(key)
		if key == serialOk {
			//fmt.Println("ok serial")
			return true, nil
		} else {
			//fmt.Println("invalid serial")
			return false, errors.New("invalid serial")
		}
	} else {
		//fmt.Println("error get key")
		//fmt.Println(err)
		return false, errors.New("error get key")
	}
}
