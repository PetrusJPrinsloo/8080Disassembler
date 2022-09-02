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

	read, err := decode(rom, currentByte)
	check(err)
	currentByte += read

	for err == nil {
		read, err = decode(rom, currentByte)
		currentByte += read
	}

}

func decode(rom []byte, currentByte int) (int, error) {
	bytesRead := 1

	//convert byte to hex
	if len(rom) < currentByte+1 {
		return 0, fmt.Errorf("Out of bounds")
	}

	switch rom[currentByte] {
	case 0x00:
		fmt.Println("NOP")
		break
	case 0x01:
		fmt.Printf("LXI\tB,#$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x02:
		fmt.Println("STAX\tB")
		break
	case 0x03:
		fmt.Println("INX\tB")
		break
	case 0x04:
		fmt.Println("INR\tB")
		break
	case 0x05:
		fmt.Println("DCR\tB")
		break
	case 0x06:
		fmt.Printf("MVI\tB,#$%02x\n", rom[currentByte+1])
		bytesRead = 2
		break
	case 0x07:
		fmt.Println("RLC")
		break
	case 0x08:
		fmt.Println("NOP")
		break
	case 0x09:
		fmt.Println("DAD\tB")
		break
	case 0x0a:
		fmt.Println("LDAX\tB")
		break
	case 0x0b:
		fmt.Println("DCX\tB")
		break
	case 0x0c:
		fmt.Println("INR\tC")
		break
	case 0x0d:
		fmt.Println("DCR\tC")
		break
	case 0x0e:
		fmt.Printf("MVI\tC,#$%02x\n", rom[currentByte+1])
		bytesRead = 2
		break
	case 0x0f:
		fmt.Println("RRC")
		break
	case 0x10:
		fmt.Println("NOP")
		break
	case 0x11:
		fmt.Printf("LXI\tD,#$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x12:
		fmt.Println("STAX\tD")
		break
	case 0x13:
		fmt.Println("INX\tD")
		break
	case 0x14:
		fmt.Println("INR\tD")
		break
	case 0x15:
		fmt.Println("DCR\tD")
		break
	case 0x16:
		fmt.Printf("MVI\tD,#$%02x\n", rom[currentByte+1])
		bytesRead = 2
		break
	case 0x17:
		fmt.Println("RAL")
		break
	case 0x18:
		fmt.Println("NOP")
		break
	case 0x19:
		fmt.Println("DAD\tD")
		break
	case 0x1a:
		fmt.Println("LDAX\tD")
		break
	case 0x1b:
		fmt.Println("DCX\tD")
		break
	case 0x1c:
		fmt.Println("INR\tE")
		break
	case 0x1d:
		fmt.Println("DCR\tE")
		break
	case 0x1e:
		fmt.Printf("MVI\tE,#$%02x\n", rom[currentByte+1])
		bytesRead = 2
		break
	case 0x1f:
		fmt.Println("RAR")
		break
	case 0x20:
		fmt.Println("NOP")
		break
	case 0x21:
		fmt.Printf("LXI\tH,#$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x22:
		fmt.Printf("SHLD\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x23:
		fmt.Println("INX\tH")
		break
	case 0x24:
		fmt.Println("INR\tH")
		break
	case 0x25:
		fmt.Println("DCR\tH")
		break
	case 0x26:
		fmt.Printf("MVI\tH,#$%02x\n", rom[currentByte+1])
		bytesRead = 2
		break
	case 0x27:
		fmt.Println("DAA")
		break
	case 0x28:
		fmt.Println("NOP")
		break
	case 0x29:
		fmt.Println("DAD\tH")
		break
	case 0x2a:
		fmt.Printf("LHLD\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x2b:
		fmt.Println("DCX\tH")
		break
	case 0x2c:
		fmt.Println("INR\tL")
		break
	case 0x2d:
		fmt.Println("DCR\tL")
		break
	case 0x2e:
		fmt.Printf("MVI\tL,#$%02x\n", rom[currentByte+1])
		bytesRead = 2
		break
	case 0x2f:
		fmt.Println("CMA")
		break
	case 0x30:
		fmt.Println("NOP")
		break
	case 0x31:
		fmt.Printf("LXI\tSP,#$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x32:
		fmt.Printf("STA\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x33:
		fmt.Println("INX\tSP")
		break
	case 0x34:
		fmt.Println("INR\tM")
		break
	case 0x35:
		fmt.Println("DCR\tM")
		break
	case 0x36:
		fmt.Printf("MVI\tM,#$%02x\n", rom[currentByte+1])
		bytesRead = 2
		break
	case 0x37:
		fmt.Println("STC")
		break
	case 0x38:
		fmt.Println("NOP")
		break
	case 0x39:
		fmt.Println("DAD\tSP")
		break
	case 0x3a:
		fmt.Printf("LDA\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		bytesRead = 3
		break
	case 0x3b:
		fmt.Println("DCX\tSP")
		break
	case 0x3c:
		fmt.Println("INR\tA")
		break
	case 0x3d:
		fmt.Println("DCR\tA")
		break
	case 0x3e:
		fmt.Printf("MVI\tA,#$%02x\n", rom[currentByte+1])
		bytesRead = 2
		break
	case 0x3f:
		fmt.Println("CMC")
		break
	case 0x40:
		fmt.Println("MOV\tB,B")
		break
	case 0x41:
		fmt.Println("MOV\tB,C")
		break
	case 0x42:
		fmt.Println("MOV\tB,D")
		break
	case 0x43:
		fmt.Println("MOV\tB,E")
		break
	case 0x44:
		fmt.Println("MOV\tB,H")
		break
	case 0x45:
		fmt.Println("MOV\tB,L")
		break
	case 0x46:
		fmt.Println("MOV\tB,M")
		break
	case 0x47:
		fmt.Println("MOV\tB,A")
		break
	case 0x48:
		fmt.Println("MOV\tC,B")
		break
	case 0x49:
		fmt.Println("MOV\tC,C")
		break
	case 0x4a:
		fmt.Println("MOV\tC,D")
		break
	case 0x4b:
		fmt.Println("MOV\tC,E")
		break
	case 0x4c:
		fmt.Println("MOV\tC,H")
		break
	case 0x4d:
		fmt.Println("MOV\tC,L")
		break
	case 0x4e:
		fmt.Println("MOV\tC,M")
		break
	case 0x4f:
		fmt.Println("MOV\tC,A")
		break
	case 0x50:
		fmt.Println("MOV\tD,B")
		break
	case 0x51:
		fmt.Println("MOV\tD,C")
		break
	case 0x52:
		fmt.Println("MOV\tD,D")
		break
	case 0x53:
		fmt.Println("MOV\tD,E")
		break
	case 0x54:
		fmt.Println("MOV\tD,H")
		break
	case 0x55:
		fmt.Println("MOV\tD,L")
		break
	case 0x56:
		fmt.Println("MOV\tD,M")
		break
	case 0x57:
		fmt.Println("MOV\tD,A")
		break
	case 0x58:
		fmt.Println("MOV\tE,B")
		break
	case 0x59:
		fmt.Println("MOV\tE,C")
		break
	case 0x5a:
		fmt.Println("MOV\tE,D")
		break
	case 0x5b:
		fmt.Println("MOV\tE,E")
		break
	case 0x5c:
		fmt.Println("MOV\tE,H")
		break
	case 0x5d:
		fmt.Println("MOV\tE,L")
		break
	case 0x5e:
		fmt.Println("MOV\tE,M")
		break
	case 0x5f:
		fmt.Println("MOV\tE,A")
		break
	case 0x60:
		fmt.Println("MOV\tH,B")
		break
	case 0x61:
		fmt.Println("MOV\tH,C")
		break
	case 0x62:
		fmt.Println("MOV\tH,D")
		break
	case 0x63:
		fmt.Println("MOV\tH,E")
		break
	case 0x64:
		fmt.Println("MOV\tH,H")
		break
	case 0x65:
		fmt.Println("MOV\tH,L")
		break
	case 0x66:
		fmt.Println("MOV\tH,M")
		break
	case 0x67:
		fmt.Println("MOV\tH,A")
		break
	case 0x68:
		fmt.Println("MOV\tL,B")
		break
	case 0x69:
		fmt.Println("MOV\tL,C")
		break
	case 0x6a:
		fmt.Println("MOV\tL,D")
		break
	case 0x6b:
		fmt.Println("MOV\tL,E")
		break
	case 0x6c:
		fmt.Println("MOV\tL,H")
		break
	case 0x6d:
		fmt.Println("MOV\tL,L")
		break
	case 0x6e:
		fmt.Println("MOV\tL,M")
		break
	case 0x6f:
		fmt.Println("MOV\tL,A")
		break
	case 0x70:
		fmt.Println("MOV\tM,B")
		break
	case 0x71:
		fmt.Println("MOV\tM,C")
		break
	case 0x72:
		fmt.Println("MOV\tM,D")
		break
	case 0x73:
		fmt.Println("MOV\tM,E")
		break
	case 0x74:
		fmt.Println("MOV\tM,H")
		break
	case 0x75:
		fmt.Println("MOV\tM,L")
		break
	case 0x76:
		fmt.Println("HLT")
		break
	case 0x77:
		fmt.Println("MOV\tM,A")
		break
	case 0x78:
		fmt.Println("MOV\tA,B")
		break
	case 0x79:
		fmt.Println("MOV\tA,C")
		break
	case 0x7a:
		fmt.Println("MOV\tA,D")
		break
	case 0x7b:
		fmt.Println("MOV\tA,E")
		break
	case 0x7c:
		fmt.Println("MOV\tA,H")
		break
	case 0x7d:
		fmt.Println("MOV\tA,L")
		break
	case 0x7e:
		fmt.Println("MOV\tA,M")
		break
	case 0x7f:
		fmt.Println("MOV\tA,A")
		break
	case 0x80:
		fmt.Println("ADD\tB")
		break
	case 0x81:
		fmt.Println("ADD\tC")
		break
	case 0x82:
		fmt.Println("ADD\tD")
		break
	case 0x83:
		fmt.Println("ADD\tE")
		break
	case 0x84:
		fmt.Println("ADD\tH")
		break
	case 0x85:
		fmt.Println("ADD\tL")
		break
	case 0x86:
		fmt.Println("ADD\tM")
		break
	case 0x87:
		fmt.Println("ADD\tA")
		break
	case 0x88:
		fmt.Println("ADC\tB")
		break
	case 0x89:
		fmt.Println("ADC\tC")
		break
	case 0x8a:
		fmt.Println("ADC\tD")
		break
	case 0x8b:
		fmt.Println("ADC\tE")
		break
	case 0x8c:
		fmt.Println("ADC\tH")
		break
	case 0x8d:
		fmt.Println("ADC\tL")
		break
	case 0x8e:
		fmt.Println("ADC\tM")
		break
	case 0x8f:
		fmt.Println("ADC\tA")
		break
	case 0x90:
		fmt.Println("SUB\tB")
		break
	case 0x91:
		fmt.Println("SUB\tC")
		break
	case 0x92:
		fmt.Println("SUB\tD")
		break
	case 0x93:
		fmt.Println("SUB\tE")
		break
	case 0x94:
		fmt.Println("SUB\tH")
		break
	case 0x95:
		fmt.Println("SUB\tL")
		break
	case 0x96:
		fmt.Println("SUB\tM")
		break
	case 0x97:
		fmt.Println("SUB\tA")
		break
	case 0x98:
		fmt.Println("SBB\tB")
		break
	case 0x99:
		fmt.Println("SBB\tC")
		break
	case 0x9a:
		fmt.Println("SBB\tD")
		break
	case 0x9b:
		fmt.Println("SBB\tE")
		break
	case 0x9c:
		fmt.Println("SBB\tH")
		break
	case 0x9d:
		fmt.Println("SBB\tL")
		break
	case 0x9e:
		fmt.Println("SBB\tM")
		break
	case 0x9f:
		fmt.Println("SBB\tA")
		break
	case 0xa0:
		fmt.Println("ANA\tB")
		break
	case 0xa1:
		fmt.Println("ANA\tC")
		break
	case 0xa2:
		fmt.Println("ANA\tD")
		break
	case 0xa3:
		fmt.Println("ANA\tE")
		break
	case 0xa4:
		fmt.Println("ANA\tH")
		break
	case 0xa5:
		fmt.Println("ANA\tL")
		break
	case 0xa6:
		fmt.Println("ANA\tM")
		break
	case 0xa7:
		fmt.Println("ANA\tA")
		break
	case 0xa8:
		fmt.Println("XRA\tB")
		break
	case 0xa9:
		fmt.Println("XRA\tC")
		break
	case 0xaa:
		fmt.Println("XRA\tD")
		break
	case 0xab:
		fmt.Println("XRA\tE")
		break
	case 0xac:
		fmt.Println("XRA\tH")
		break
	case 0xad:
		fmt.Println("XRA\tL")
		break
	case 0xae:
		fmt.Println("XRA\tM")
		break
	case 0xaf:
		fmt.Println("XRA\tA")
		break
	case 0xb0:
		fmt.Println("ORA\tB")
		break
	case 0xb1:
		fmt.Println("ORA\tC")
		break
	case 0xb2:
		fmt.Println("ORA\tD")
		break
	case 0xb3:
		fmt.Println("ORA\tE")
		break
	case 0xb4:
		fmt.Println("ORA\tH")
		break
	case 0xb5:
		fmt.Println("ORA\tL")
		break
	case 0xb6:
		fmt.Println("ORA\tM")
		break
	case 0xb7:
		fmt.Println("ORA\tA")
		break
	case 0xb8:
		fmt.Println("CMP\tB")
		break
	case 0xb9:
		fmt.Println("CMP\tC")
		break
	case 0xba:
		fmt.Println("CMP\tD")
		break
	case 0xbb:
		fmt.Println("CMP\tE")
		break
	case 0xbc:
		fmt.Println("CMP\tH")
		break
	case 0xbd:
		fmt.Println("CMP\tL")
		break
	case 0xbe:
		fmt.Println("CMP\tM")
		break
	case 0xbf:
		fmt.Println("CMP\tA")
		break
	case 0xc0:
		fmt.Println("RNZ")
		break
	case 0xc1:
		fmt.Println("POP\tB")
		break
	case 0xc2:
		fmt.Printf("JNZ\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xc3:
		fmt.Printf("JMP\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xc4:
		fmt.Printf("CNZ\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xc5:
		fmt.Println("PUSH\tB")
		break
	case 0xc6:
		fmt.Printf("ADI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xc7:
		fmt.Println("RST\t0")
		break
	case 0xc8:
		fmt.Println("RZ")
		break
	case 0xc9:
		fmt.Println("RET")
		break
	case 0xca:
		fmt.Printf("JZ\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xcb:
		fmt.Println("NOP")
		break
	case 0xcc:
		fmt.Printf("CZ\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xcd:
		fmt.Printf("CALL\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xce:
		fmt.Printf("ACI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xcf:
		fmt.Println("RST\t1")
		break
	case 0xd0:
		fmt.Println("RNC")
		break
	case 0xd1:
		fmt.Println("POP\tD")
		break
	case 0xd2:
		fmt.Printf("JNC\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xd3:
		fmt.Printf("OUT\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xd4:
		fmt.Printf("CNC\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xd5:
		fmt.Println("PUSH\tD")
		break
	case 0xd6:
		fmt.Printf("SUI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xd7:
		fmt.Println("RST\t2")
		break
	case 0xd8:
		fmt.Println("RC")
		break
	case 0xd9:
		fmt.Println("NOP")
		break
	case 0xda:
		fmt.Printf("JC\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xdb:
		fmt.Printf("IN\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xdc:
		fmt.Printf("CC\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xdd:
		fmt.Println("NOP")
		break
	case 0xde:
		fmt.Printf("SBI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xdf:
		fmt.Println("RST\t3")
		break
	case 0xe0:
		fmt.Println("RPO")
		break
	case 0xe1:
		fmt.Println("POP\tH")
		break
	case 0xe2:
		fmt.Printf("JPO\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xe3:
		fmt.Println("XTHL")
		break
	case 0xe4:
		fmt.Printf("CPO\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xe5:
		fmt.Println("PUSH\tH")
		break
	case 0xe6:
		fmt.Printf("ANI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xe7:
		fmt.Println("RST\t4")
		break
	case 0xe8:
		fmt.Println("RPE")
		break
	case 0xe9:
		fmt.Println("PCHL")
		break
	case 0xea:
		fmt.Printf("JPE\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xeb:
		fmt.Println("XCHG")
		break
	case 0xec:
		fmt.Printf("CPE\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xed:
		fmt.Println("NOP")
		break
	case 0xee:
		fmt.Printf("XRI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xef:
		fmt.Println("RST\t5")
		break
	case 0xf0:
		fmt.Println("RP")
		break
	case 0xf1:
		fmt.Println("POP\tPSW")
		break
	case 0xf2:
		fmt.Printf("JP\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xf3:
		fmt.Println("DI")
		break
	case 0xf4:
		fmt.Printf("CP\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xf5:
		fmt.Println("PUSH\tPSW")
		break
	case 0xf6:
		fmt.Printf("ORI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xf7:
		fmt.Println("RST\t6")
		break
	case 0xf8:
		fmt.Println("RM")
		break
	case 0xf9:
		fmt.Println("SPHL")
		break
	case 0xfa:
		fmt.Printf("JM\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xfb:
		fmt.Println("EI")
		break
	case 0xfc:
		fmt.Printf("CM\t$%02x%02x\n", rom[currentByte+2], rom[currentByte+1])
		currentByte = 3
		break
	case 0xfd:
		fmt.Println("NOP")
		break
	case 0xfe:
		fmt.Printf("CPI\t$%02x\n", rom[currentByte+1])
		currentByte = 2
		break
	case 0xff:
		fmt.Println("RST\t7")
		break
	default:
		fmt.Println("Unknown opcode")
	}

	return bytesRead, nil
}
