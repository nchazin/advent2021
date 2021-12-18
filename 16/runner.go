package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

var LITERAL_ID = int(4)

type Packet struct {
	literal int
	id      int
	version int
	length  int
}

var codewheel = map[byte][]byte{
	'0': []byte{'0', '0', '0', '0'},
	'1': []byte{'0', '0', '0', '1'},
	'2': []byte{'0', '0', '1', '0'},
	'3': []byte{'0', '0', '1', '1'},
	'4': []byte{'0', '1', '0', '0'},
	'5': []byte{'0', '1', '0', '1'},
	'6': []byte{'0', '1', '1', '0'},
	'7': []byte{'0', '1', '1', '1'},
	'8': []byte{'1', '0', '0', '0'},
	'9': []byte{'1', '0', '0', '1'},
	'A': []byte{'1', '0', '1', '0'},
	'B': []byte{'1', '0', '1', '1'},
	'C': []byte{'1', '1', '0', '0'},
	'D': []byte{'1', '1', '0', '1'},
	'E': []byte{'1', '1', '1', '0'},
	'F': []byte{'1', '1', '1', '1'},
}

func input_to_bits(input []byte) []byte {
	output := make([]byte, 0)
	for _, c := range input {
		output = append(output, codewheel[c]...)
	}
	return output
}

func version(input []byte) int {
	version, _ := strconv.ParseInt(string(input[:3]), 2, 32)
	return int(version)
}

func packet_id(input []byte) int {
	id, _ := strconv.ParseInt(string(input[3:6]), 2, 32)
	return int(id)
}

func extract_literal(input []byte) (int, int) {
	fmt.Println("literal")
	literal := make([]byte, 0)
	lasti := 0
	for i := 6; i+4 < len(input); i += 5 {
		chunk := input[i : i+5]
		literal = append(literal, chunk[1:]...)
		if chunk[0] == '0' {
			lasti = i + 5
			break
		}
	}
	literal_int, _ := strconv.ParseInt(string(literal), 2, 32)
	//fmt.Println("Next i:", next_i)
	return int(literal_int), lasti
}

func length_type(input []byte) int {
	if input[6] == '1' {
		return 1
	} else {
		return 0
	}
}

func extract_length_packet(input []byte) ([]Packet, int) {
	packets := make([]Packet, 0)
	length, _ := strconv.ParseInt(string(input[7:22]), 2, 32)
	fmt.Println("Length: ", length)
	i := 22
	for i < int(22+length) {
		newpackets, n := parse_packet(input[i:])
		packets = append(packets, newpackets...)
		i += n
	}
	return packets, int(i)

}

func extract_count_packet(input []byte) ([]Packet, int) {
	packets := make([]Packet, 0)
	//	fmt.Println(string(input))
	//	fmt.Println(string(input[7:18]))
	count, _ := strconv.ParseInt(string(input[7:18]), 2, 32)
	cur_point := 18
	for i := 0; i < int(count); i++ {
		fmt.Printf("WOrking on %d of %d at %d\n", i, count, cur_point)
		packet, n := parse_packet(input[cur_point:])
		packets = append(packets, packet...)
		cur_point += n
	}
	fmt.Println("Curpoint is now: ", cur_point)
	return packets, cur_point
}

func extract_operator(input []byte) ([]Packet, int) {
	length_type := length_type(input)
	fmt.Println("length type: ", length_type)
	if length_type == 0 {
		return extract_length_packet(input)
	} else {
		return extract_count_packet(input)
	}
}

/*
000 000 100000110101 1001001010111011100101001011100101010110110100

100 100 1 0101 1 1011 1 0010 1 0010 1 1100 1 0101 0 1101
10100110000000001011110

101 001 1 00000000010

1111001001110001111101110111010110000101101011 (0 of 2)

111 100 1 0011 1 0001 1 1110 1 1101 1 1010 1 1000 0 1011

01011000000000001010100 1 of 2
010 110 0 000000000101010 010100 111101001000011111100111101001000011
*/

func parse_packet(input []byte) ([]Packet, int) {
	//	fmt.Println(string(input))
	//fmt.Println("Data length: ", len(input))
	fmt.Println("Data start: ", string(input[0:64]))
	version := version(input)
	next_index := 0
	id := packet_id(input)
	packet := Packet{}
	packet.version = version
	packet.id = id
	fmt.Println("packet version:", version)
	packets := make([]Packet, 0)
	if id == LITERAL_ID {
		val, n := extract_literal(input)
		packet.literal = val
		packet.length = n
		next_index = n
		packets = append(packets, packet)
	} else {
		fmt.Println("Operator on")
		command_packets, n := extract_operator(input)
		packets = append(packets, packet)
		packets = append(packets, command_packets...)
		next_index = n
	}
	return packets, next_index
}

func do_main(data []byte) {
	op, _ := parse_packet(data)
	osum := 0
	for _, pp := range op {
		fmt.Println("version: ", pp.version, " id:", pp.id, " value:", pp.literal)
		osum += pp.version
	}
	fmt.Println(osum)
}

func main() {
	/*for _, c := range []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'} {
		fmt.Printf("%c -> %s\n", c, codewheel[c])
	}*/

	/* SKIP FOR NOW...
	input := input_to_bits([]byte("D2FE28D2FE28"))
	cur_index := 0
	packets := make([]Packet, 0)
	i := 0
	for cur_index < len(input) {
		p, n := parse_packet(input[cur_index:])
		packets = append(packets, p...)
		cur_index += n
		if i > 5 {
			break
		}
		i++
	}
	for _, p := range packets {
		fmt.Println(p)
	}
	*/

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	sumbits := "00000110101"
	ss, _ := strconv.ParseInt(sumbits, 2, 32)
	fmt.Println(ss)

	fmt.Println()
	do_main(input_to_bits(data))
	return

	fmt.Println("EE00D40C823060")
	ip := input_to_bits([]byte("EE00D40C823060"))
	fmt.Println("Len type: ", length_type(ip))
	fmt.Println(string(ip))
	fmt.Println("-----")
	a, acount := extract_count_packet(ip)
	fmt.Println("-", a)
	fmt.Println("------> acount: ", acount, " ... ", len(ip))
	for _, pp := range a {
		fmt.Println("version: ", pp.version, " value:", pp.literal)
	}

	fmt.Println("38006F45291200")
	ip = input_to_bits([]byte("38006F45291200"))
	fmt.Println("Len type: ", length_type(ip))
	x, xcount := extract_length_packet(ip)
	fmt.Println("------> xcount: ", xcount, " ... ", len(ip))

	for _, pp := range x {
		fmt.Println("version: ", pp.version, " value:", pp.literal)
	}

	ip = input_to_bits([]byte("A0016C880162017C3686B18A3D4780"))
	xyz, _ := parse_packet(ip)
	sum := 0
	for _, pp := range xyz {
		fmt.Println("version: ", pp.version, " id:", pp.id, " value:", pp.literal)
		sum += pp.version
	}
	fmt.Println(sum)

	ip = input_to_bits([]byte("020D64AEE52E55B"))
	fmt.Printf("Version: %d\n", version(ip))
	fmt.Printf("Id: %d\n", packet_id(ip))
	fmt.Printf("Lentght Type: %d\n", length_type((ip)))
	//y, _ := extract_length_packet(ip)

	/*
		output := input_to_bits([]byte("ABC123"))
		fmt.Printf("%s\n", string(output))
		fmt.Printf("Version: %d\n", version(input_to_bits([]byte("D2FE28"))))
		fmt.Printf("Id: %d\n", packet_id(input_to_bits([]byte("D2FE28"))))
		fmt.Printf("literal: %d\n", extract_literal(input_to_bits([]byte("D2FE28"))))
	*/

}
