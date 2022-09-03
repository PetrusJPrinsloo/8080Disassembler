package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RetrieveROM(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	check(err)
	defer file.Close()

	fileInfo, err := file.Stat()
	check(err)

	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)
	check(err)

	return bytes, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	rom, err := RetrieveROM(filename)
	check(err)

	currentByte := 0
	lineNumber := int16(0)

	read, err := decode(rom, currentByte, lineNumber)
	check(err)
	currentByte += read

	for err == nil {
		lineNumber++
		read, err = decode(rom, currentByte, lineNumber)
		currentByte += read
	}

}

func decode(rom []byte, currentByte int, lineNumber int16) (int, error) {
	bytesRead := 1

	//convert byte to hex
	if len(rom) < currentByte+1 {
		return 0, fmt.Errorf("Out of bounds")
	}

	switch rom[currentByte] {
	case 0x00:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0x01:
		fmt.Printf("%04x\tLXI\tB,#$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x02:
		fmt.Printf("%04x\tSTAX\tB\n", lineNumber)
		break
	case 0x03:
		fmt.Printf("%04x\tINX\tB\n", lineNumber)
		break
	case 0x04:
		fmt.Printf("%04x\tINR\tB\n", lineNumber)
		break
	case 0x05:
		fmt.Printf("%04x\tDCR\tB\n", lineNumber)
		break
	case 0x06:
		fmt.Printf("%04x\tMVI\tB,#$%02x\n", lineNumber, rom[currentByte+1])
		bytesRead = 2
		break
	case 0x07:
		fmt.Printf("%04x\tRLC\n", lineNumber)
		break
	case 0x08:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0x09:
		fmt.Printf("%04x\tDAD\tB\n", lineNumber)
		break
	case 0x0a:
		fmt.Printf("%04x\tLDAX\tB\n", lineNumber)
		break
	case 0x0b:
		fmt.Printf("%04x\tDCX\tB\n", lineNumber)
		break
	case 0x0c:
		fmt.Printf("%04x\tINR\tC\n", lineNumber)
		break
	case 0x0d:
		fmt.Printf("%04x\tDCR\tC\n", lineNumber)
		break
	case 0x0e:
		fmt.Printf("%04x\tMVI\tC,#$%02x\n", lineNumber, rom[currentByte+1])
		bytesRead = 2
		break
	case 0x0f:
		fmt.Printf("%04x\tRRC\n", lineNumber)
		break
	case 0x10:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0x11:
		fmt.Printf("%04x\tLXI\tD,#$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x12:
		fmt.Printf("%04x\tSTAX\tD\n", lineNumber)
		break
	case 0x13:
		fmt.Printf("%04x\tINX\tD\n", lineNumber)
		break
	case 0x14:
		fmt.Printf("%04x\tINR\tD\n", lineNumber)
		break
	case 0x15:
		fmt.Printf("%04x\tDCR\tD\n", lineNumber)
		break
	case 0x16:
		fmt.Printf("%04x\tMVI\tD,#$%02x\n", lineNumber, rom[currentByte+1])
		bytesRead = 2
		break
	case 0x17:
		fmt.Printf("%04x\tRAL\n", lineNumber)
		break
	case 0x18:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0x19:
		fmt.Printf("%04x\tDAD\tD\n", lineNumber)
		break
	case 0x1a:
		fmt.Printf("%04x\tLDAX\tD\n", lineNumber)
		break
	case 0x1b:
		fmt.Printf("%04x\tDCX\tD\n", lineNumber)
		break
	case 0x1c:
		fmt.Printf("%04x\tINR\tE\n", lineNumber)
		break
	case 0x1d:
		fmt.Printf("%04x\tDCR\tE\n", lineNumber)
		break
	case 0x1e:
		fmt.Printf("%04x\tMVI\tE,#$%02x\n", lineNumber, rom[currentByte+1])
		bytesRead = 2
		break
	case 0x1f:
		fmt.Printf("%04x\tRAR\n", lineNumber)
		break
	case 0x20:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0x21:
		fmt.Printf("%04x\tLXI\tH,#$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x22:
		fmt.Printf("%04x\tSHLD\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x23:
		fmt.Printf("%04x\tINX\tH\n", lineNumber)
		break
	case 0x24:
		fmt.Printf("%04x\tINR\tH\n", lineNumber)
		break
	case 0x25:
		fmt.Printf("%04x\tDCR\tH\n", lineNumber)
		break
	case 0x26:
		fmt.Printf("%04x\tMVI\tH,#$%02x\n", rom[currentByte+1])
		bytesRead = 2
		break
	case 0x27:
		fmt.Printf("%04x\tDAA\n", lineNumber)
		break
	case 0x28:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0x29:
		fmt.Printf("%04x\tDAD\tH\n", lineNumber)
		break
	case 0x2a:
		fmt.Printf("%04x\tLHLD\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x2b:
		fmt.Printf("%04x\tDCX\tH\n", lineNumber)
		break
	case 0x2c:
		fmt.Printf("%04x\tINR\tL\n", lineNumber)
		break
	case 0x2d:
		fmt.Printf("%04x\tDCR\tL\n", lineNumber)
		break
	case 0x2e:
		fmt.Printf("%04x\tMVI\tL,#$%02x\n", lineNumber, rom[currentByte+1])
		bytesRead = 2
		break
	case 0x2f:
		fmt.Printf("%04x\tCMA\n", lineNumber)
		break
	case 0x30:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0x31:
		fmt.Printf("%04x\tLXI\tSP,#$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x32:
		fmt.Printf("%04x\tSTA\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x33:
		fmt.Printf("%04x\tINX\tSP\n", lineNumber)
		break
	case 0x34:
		fmt.Printf("%04x\tINR\tM\n", lineNumber)
		break
	case 0x35:
		fmt.Printf("%04x\tDCR\tM\n", lineNumber)
		break
	case 0x36:
		fmt.Printf("%04x\tMVI\tM,#$%02x\n", lineNumber, rom[currentByte+1])
		bytesRead = 2
		break
	case 0x37:
		fmt.Printf("%04x\tSTC\n", lineNumber)
		break
	case 0x38:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0x39:
		fmt.Printf("%04x\tDAD\tSP\n", lineNumber)
		break
	case 0x3a:
		fmt.Printf("%04x\tLDA\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x3b:
		fmt.Printf("%04x\tDCX\tSP\n", lineNumber)
		break
	case 0x3c:
		fmt.Printf("%04x\tINR\tA\n", lineNumber)
		break
	case 0x3d:
		fmt.Printf("%04x\tDCR\tA\n", lineNumber)
		break
	case 0x3e:
		fmt.Printf("%04x\tMVI\tA,#$%02x\n", lineNumber, rom[currentByte+1])
		bytesRead = 2
		break
	case 0x3f:
		fmt.Printf("%04x\tCMC\n", lineNumber)
		break
	case 0x40:
		fmt.Printf("%04x\tMOV\tB,B\n", lineNumber)
		break
	case 0x41:
		fmt.Printf("%04x\tMOV\tB,C\n", lineNumber)
		break
	case 0x42:
		fmt.Printf("%04x\tMOV\tB,D\n", lineNumber)
		break
	case 0x43:
		fmt.Printf("%04x\tMOV\tB,E\n", lineNumber)
		break
	case 0x44:
		fmt.Printf("%04x\tMOV\tB,H\n", lineNumber)
		break
	case 0x45:
		fmt.Printf("%04x\tMOV\tB,L\n", lineNumber)
		break
	case 0x46:
		fmt.Printf("%04x\tMOV\tB,M\n", lineNumber)
		break
	case 0x47:
		fmt.Printf("%04x\tMOV\tB,A\n", lineNumber)
		break
	case 0x48:
		fmt.Printf("%04x\tMOV\tC,B\n", lineNumber)
		break
	case 0x49:
		fmt.Printf("%04x\tMOV\tC,C\n", lineNumber)
		break
	case 0x4a:
		fmt.Printf("%04x\tMOV\tC,D\n", lineNumber)
		break
	case 0x4b:
		fmt.Printf("%04x\tMOV\tC,E\n", lineNumber)
		break
	case 0x4c:
		fmt.Printf("%04x\tMOV\tC,H\n", lineNumber)
		break
	case 0x4d:
		fmt.Printf("%04x\tMOV\tC,L\n", lineNumber)
		break
	case 0x4e:
		fmt.Printf("%04x\tMOV\tC,M\n", lineNumber)
		break
	case 0x4f:
		fmt.Printf("%04x\tMOV\tC,A\n", lineNumber)
		break
	case 0x50:
		fmt.Printf("%04x\tMOV\tD,B\n", lineNumber)
		break
	case 0x51:
		fmt.Printf("%04x\tMOV\tD,C\n", lineNumber)
		break
	case 0x52:
		fmt.Printf("%04x\tMOV\tD,D\n", lineNumber)
		break
	case 0x53:
		fmt.Printf("%04x\tMOV\tD,E\n", lineNumber)
		break
	case 0x54:
		fmt.Printf("%04x\tMOV\tD,H\n", lineNumber)
		break
	case 0x55:
		fmt.Printf("%04x\tMOV\tD,L\n", lineNumber)
		break
	case 0x56:
		fmt.Printf("%04x\tMOV\tD,M\n", lineNumber)
		break
	case 0x57:
		fmt.Printf("%04x\tMOV\tD,A\n", lineNumber)
		break
	case 0x58:
		fmt.Printf("%04x\tMOV\tE,B\n", lineNumber)
		break
	case 0x59:
		fmt.Printf("%04x\tMOV\tE,C\n", lineNumber)
		break
	case 0x5a:
		fmt.Printf("%04x\tMOV\tE,D\n", lineNumber)
		break
	case 0x5b:
		fmt.Printf("%04x\tMOV\tE,E\n", lineNumber)
		break
	case 0x5c:
		fmt.Printf("%04x\tMOV\tE,H\n", lineNumber)
		break
	case 0x5d:
		fmt.Printf("%04x\tMOV\tE,L\n", lineNumber)
		break
	case 0x5e:
		fmt.Printf("%04x\tMOV\tE,M\n", lineNumber)
		break
	case 0x5f:
		fmt.Printf("%04x\tMOV\tE,A\n", lineNumber)
		break
	case 0x60:
		fmt.Printf("%04x\tMOV\tH,B\n", lineNumber)
		break
	case 0x61:
		fmt.Printf("%04x\tMOV\tH,C\n", lineNumber)
		break
	case 0x62:
		fmt.Printf("%04x\tMOV\tH,D\n", lineNumber)
		break
	case 0x63:
		fmt.Printf("%04x\tMOV\tH,E\n", lineNumber)
		break
	case 0x64:
		fmt.Printf("%04x\tMOV\tH,H\n", lineNumber)
		break
	case 0x65:
		fmt.Printf("%04x\tMOV\tH,L\n", lineNumber)
		break
	case 0x66:
		fmt.Printf("%04x\tMOV\tH,M\n", lineNumber)
		break
	case 0x67:
		fmt.Printf("%04x\tMOV\tH,A\n", lineNumber)
		break
	case 0x68:
		fmt.Printf("%04x\tMOV\tL,B\n", lineNumber)
		break
	case 0x69:
		fmt.Printf("%04x\tMOV\tL,C\n", lineNumber)
		break
	case 0x6a:
		fmt.Printf("%04x\tMOV\tL,D\n", lineNumber)
		break
	case 0x6b:
		fmt.Printf("%04x\tMOV\tL,E\n", lineNumber)
		break
	case 0x6c:
		fmt.Printf("%04x\tMOV\tL,H\n", lineNumber)
		break
	case 0x6d:
		fmt.Printf("%04x\tMOV\tL,L\n", lineNumber)
		break
	case 0x6e:
		fmt.Printf("%04x\tMOV\tL,M\n", lineNumber)
		break
	case 0x6f:
		fmt.Printf("%04x\tMOV\tL,A\n", lineNumber)
		break
	case 0x70:
		fmt.Printf("%04x\tMOV\tM,B\n", lineNumber)
		break
	case 0x71:
		fmt.Printf("%04x\tMOV\tM,C\n", lineNumber)
		break
	case 0x72:
		fmt.Printf("%04x\tMOV\tM,D\n", lineNumber)
		break
	case 0x73:
		fmt.Printf("%04x\tMOV\tM,E\n", lineNumber)
		break
	case 0x74:
		fmt.Printf("%04x\tMOV\tM,H\n", lineNumber)
		break
	case 0x75:
		fmt.Printf("%04x\tMOV\tM,L\n", lineNumber)
		break
	case 0x76:
		fmt.Printf("%04x\tHLT\n", lineNumber)
		break
	case 0x77:
		fmt.Printf("%04x\tMOV\tM,A\n", lineNumber)
		break
	case 0x78:
		fmt.Printf("%04x\tMOV\tA,B\n", lineNumber)
		break
	case 0x79:
		fmt.Printf("%04x\tMOV\tA,C\n", lineNumber)
		break
	case 0x7a:
		fmt.Printf("%04x\tMOV\tA,D\n", lineNumber)
		break
	case 0x7b:
		fmt.Printf("%04x\tMOV\tA,E\n", lineNumber)
		break
	case 0x7c:
		fmt.Printf("%04x\tMOV\tA,H\n", lineNumber)
		break
	case 0x7d:
		fmt.Printf("%04x\tMOV\tA,L\n", lineNumber)
		break
	case 0x7e:
		fmt.Printf("%04x\tMOV\tA,M\n", lineNumber)
		break
	case 0x7f:
		fmt.Printf("%04x\tMOV\tA,A\n", lineNumber)
		break
	case 0x80:
		fmt.Printf("%04x\tADD\tB\n", lineNumber)
		break
	case 0x81:
		fmt.Printf("%04x\tADD\tC\n", lineNumber)
		break
	case 0x82:
		fmt.Printf("%04x\tADD\tD\n", lineNumber)
		break
	case 0x83:
		fmt.Printf("%04x\tADD\tE\n", lineNumber)
		break
	case 0x84:
		fmt.Printf("%04x\tADD\tH\n", lineNumber)
		break
	case 0x85:
		fmt.Printf("%04x\tADD\tL\n", lineNumber)
		break
	case 0x86:
		fmt.Printf("%04x\tADD\tM\n", lineNumber)
		break
	case 0x87:
		fmt.Printf("%04x\tADD\tA\n", lineNumber)
		break
	case 0x88:
		fmt.Printf("%04x\tADC\tB\n", lineNumber)
		break
	case 0x89:
		fmt.Printf("%04x\tADC\tC\n", lineNumber)
		break
	case 0x8a:
		fmt.Printf("%04x\tADC\tD\n", lineNumber)
		break
	case 0x8b:
		fmt.Printf("%04x\tADC\tE\n", lineNumber)
		break
	case 0x8c:
		fmt.Printf("%04x\tADC\tH\n", lineNumber)
		break
	case 0x8d:
		fmt.Printf("%04x\tADC\tL\n", lineNumber)
		break
	case 0x8e:
		fmt.Printf("%04x\tADC\tM\n", lineNumber)
		break
	case 0x8f:
		fmt.Printf("%04x\tADC\tA\n", lineNumber)
		break
	case 0x90:
		fmt.Printf("%04x\tSUB\tB\n", lineNumber)
		break
	case 0x91:
		fmt.Printf("%04x\tSUB\tC\n", lineNumber)
		break
	case 0x92:
		fmt.Printf("%04x\tSUB\tD\n", lineNumber)
		break
	case 0x93:
		fmt.Printf("%04x\tSUB\tE\n", lineNumber)
		break
	case 0x94:
		fmt.Printf("%04x\tSUB\tH\n", lineNumber)
		break
	case 0x95:
		fmt.Printf("%04x\tSUB\tL\n", lineNumber)
		break
	case 0x96:
		fmt.Printf("%04x\tSUB\tM\n", lineNumber)
		break
	case 0x97:
		fmt.Printf("%04x\tSUB\tA\n", lineNumber)
		break
	case 0x98:
		fmt.Printf("%04x\tSBB\tB\n", lineNumber)
		break
	case 0x99:
		fmt.Printf("%04x\tSBB\tC\n", lineNumber)
		break
	case 0x9a:
		fmt.Printf("%04x\tSBB\tD\n", lineNumber)
		break
	case 0x9b:
		fmt.Printf("%04x\tSBB\tE\n", lineNumber)
		break
	case 0x9c:
		fmt.Printf("%04x\tSBB\tH\n", lineNumber)
		break
	case 0x9d:
		fmt.Printf("%04x\tSBB\tL\n", lineNumber)
		break
	case 0x9e:
		fmt.Printf("%04x\tSBB\tM\n", lineNumber)
		break
	case 0x9f:
		fmt.Printf("%04x\tSBB\tA\n", lineNumber)
		break
	case 0xa0:
		fmt.Printf("%04x\tANA\tB\n", lineNumber)
		break
	case 0xa1:
		fmt.Printf("%04x\tANA\tC\n", lineNumber)
		break
	case 0xa2:
		fmt.Printf("%04x\tANA\tD\n", lineNumber)
		break
	case 0xa3:
		fmt.Printf("%04x\tANA\tE\n", lineNumber)
		break
	case 0xa4:
		fmt.Printf("%04x\tANA\tH\n", lineNumber)
		break
	case 0xa5:
		fmt.Printf("%04x\tANA\tL\n", lineNumber)
		break
	case 0xa6:
		fmt.Printf("%04x\tANA\tM\n", lineNumber)
		break
	case 0xa7:
		fmt.Printf("%04x\tANA\tA\n", lineNumber)
		break
	case 0xa8:
		fmt.Printf("%04x\tXRA\tB\n", lineNumber)
		break
	case 0xa9:
		fmt.Printf("%04x\tXRA\tC\n", lineNumber)
		break
	case 0xaa:
		fmt.Printf("%04x\tXRA\tD\n", lineNumber)
		break
	case 0xab:
		fmt.Printf("%04x\tXRA\tE\n", lineNumber)
		break
	case 0xac:
		fmt.Printf("%04x\tXRA\tH\n", lineNumber)
		break
	case 0xad:
		fmt.Printf("%04x\tXRA\tL\n", lineNumber)
		break
	case 0xae:
		fmt.Printf("%04x\tXRA\tM\n", lineNumber)
		break
	case 0xaf:
		fmt.Printf("%04x\tXRA\tA\n", lineNumber)
		break
	case 0xb0:
		fmt.Printf("%04x\tORA\tB\n", lineNumber)
		break
	case 0xb1:
		fmt.Printf("%04x\tORA\tC\n", lineNumber)
		break
	case 0xb2:
		fmt.Printf("%04x\tORA\tD\n", lineNumber)
		break
	case 0xb3:
		fmt.Printf("%04x\tORA\tE\n", lineNumber)
		break
	case 0xb4:
		fmt.Printf("%04x\tORA\tH\n", lineNumber)
		break
	case 0xb5:
		fmt.Printf("%04x\tORA\tL\n", lineNumber)
		break
	case 0xb6:
		fmt.Printf("%04x\tORA\tM\n", lineNumber)
		break
	case 0xb7:
		fmt.Printf("%04x\tORA\tA\n", lineNumber)
		break
	case 0xb8:
		fmt.Printf("%04x\tCMP\tB\n", lineNumber)
		break
	case 0xb9:
		fmt.Printf("%04x\tCMP\tC\n", lineNumber)
		break
	case 0xba:
		fmt.Printf("%04x\tCMP\tD\n", lineNumber)
		break
	case 0xbb:
		fmt.Printf("%04x\tCMP\tE\n", lineNumber)
		break
	case 0xbc:
		fmt.Printf("%04x\tCMP\tH\n", lineNumber)
		break
	case 0xbd:
		fmt.Printf("%04x\tCMP\tL\n", lineNumber)
		break
	case 0xbe:
		fmt.Printf("%04x\tCMP\tM\n", lineNumber)
		break
	case 0xbf:
		fmt.Printf("%04x\tCMP\tA\n", lineNumber)
		break
	case 0xc0:
		fmt.Printf("%04x\tRNZ\n", lineNumber)
		break
	case 0xc1:
		fmt.Printf("%04x\tPOP\tB\n", lineNumber)
		break
	case 0xc2:
		fmt.Printf("%04x\tJNZ\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xc3:
		fmt.Printf("%04x\tJMP\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xc4:
		fmt.Printf("%04x\tCNZ\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xc5:
		fmt.Printf("%04x\tPUSH\tB\n", lineNumber)
		break
	case 0xc6:
		fmt.Printf("%04x\tADI\t$%02x\n", lineNumber, rom[currentByte+1])
		currentByte = 2
		break
	case 0xc7:
		fmt.Printf("%04x\tRST\t0\n", lineNumber)
		break
	case 0xc8:
		fmt.Printf("%04x\tRZ\n", lineNumber)
		break
	case 0xc9:
		fmt.Printf("%04x\tRET\n", lineNumber)
		break
	case 0xca:
		fmt.Printf("%04x\tJZ\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xcb:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0xcc:
		fmt.Printf("%04x\tCZ\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xcd:
		fmt.Printf("%04x\tCALL\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xce:
		fmt.Printf("%04x\tACI\t$%02x\n", lineNumber, rom[currentByte+1])
		currentByte = 2
		break
	case 0xcf:
		fmt.Printf("%04x\tRST\t1\n", lineNumber)
		break
	case 0xd0:
		fmt.Printf("%04x\tRNC\n", lineNumber)
		break
	case 0xd1:
		fmt.Printf("%04x\tPOP\tD\n", lineNumber)
		break
	case 0xd2:
		fmt.Printf("%04x\tJNC\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xd3:
		fmt.Printf("OUT\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xd4:
		fmt.Printf("%04x\tCNC\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xd5:
		fmt.Printf("%04x\tPUSH\tD\n", lineNumber)
		break
	case 0xd6:
		fmt.Printf("SUI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xd7:
		fmt.Printf("%04x\tRST\t2\n", lineNumber)
		break
	case 0xd8:
		fmt.Printf("%04x\tRC\n", lineNumber)
		break
	case 0xd9:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0xda:
		fmt.Printf("%04x\tJC\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xdb:
		fmt.Printf("IN\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xdc:
		fmt.Printf("%04x\tCC\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xdd:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0xde:
		fmt.Printf("%04x\tSBI\t$%02x\n", lineNumber, rom[currentByte+1])
		currentByte = 2
		break
	case 0xdf:
		fmt.Printf("%04x\tRST\t3\n", lineNumber)
		break
	case 0xe0:
		fmt.Printf("%04x\tRPO\n", lineNumber)
		break
	case 0xe1:
		fmt.Printf("%04x\tPOP\tH\n", lineNumber)
		break
	case 0xe2:
		fmt.Printf("%04x\tJPO\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xe3:
		fmt.Printf("%04x\tXTHL\n", lineNumber)
		break
	case 0xe4:
		fmt.Printf("%04x\tCPO\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xe5:
		fmt.Printf("%04x\tPUSH\tH\n", lineNumber)
		break
	case 0xe6:
		fmt.Printf("ANI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xe7:
		fmt.Printf("%04x\tRST\t4\n", lineNumber)
		break
	case 0xe8:
		fmt.Printf("%04x\tRPE\n", lineNumber)
		break
	case 0xe9:
		fmt.Printf("%04x\tPCHL\n", lineNumber)
		break
	case 0xea:
		fmt.Printf("%04x\tJPE\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xeb:
		fmt.Printf("%04x\tXCHG\n", lineNumber)
		break
	case 0xec:
		fmt.Printf("%04x\tCPE\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xed:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0xee:
		fmt.Printf("%04x\tXRI\t$%02x\n", lineNumber, rom[currentByte+1])
		currentByte = 2
		break
	case 0xef:
		fmt.Printf("%04x\tRST\t5\n", lineNumber)
		break
	case 0xf0:
		fmt.Printf("%04x\tRP\n", lineNumber)
		break
	case 0xf1:
		fmt.Printf("%04x\tPOP\tPSW\n", lineNumber)
		break
	case 0xf2:
		fmt.Printf("%04x\tJP\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xf3:
		fmt.Printf("%04x\tDI\n", lineNumber)
		break
	case 0xf4:
		fmt.Printf("%04x\tCP\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xf5:
		fmt.Printf("%04x\tPUSH\tPSW\n", lineNumber)
		break
	case 0xf6:
		fmt.Printf("%04x\tORI\t$%02x\n", lineNumber, rom[currentByte+1])
		currentByte = 2
		break
	case 0xf7:
		fmt.Printf("%04x\tRST\t6\n", lineNumber)
		break
	case 0xf8:
		fmt.Printf("%04x\tRM\n", lineNumber)
		break
	case 0xf9:
		fmt.Printf("%04x\tSPHL\n", lineNumber)
		break
	case 0xfa:
		fmt.Printf("%04x\tJM\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xfb:
		fmt.Printf("%04x\tEI\n", lineNumber)
		break
	case 0xfc:
		fmt.Printf("%04x\tCM\t$%02x%02x\n", lineNumber, rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xfd:
		fmt.Printf("%04x\tNOP\n", lineNumber)
		break
	case 0xfe:
		fmt.Printf("%04x\tCPI\t$%02x\n", lineNumber, rom[currentByte+1])
		currentByte = 2
		break
	case 0xff:
		fmt.Printf("%04x\tRST\t7\n", lineNumber)
		break
	default:
		fmt.Printf("%04x\t#\tUnknown opcode: $%02x\n", lineNumber, rom[currentByte])
	}

	return bytesRead, nil
}
