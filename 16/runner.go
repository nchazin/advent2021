package main

import (
	"fmt"
	"strconv"
)

var LITERAL_ID = int(4)

type Packet struct {
	literal int
	id      int
	version int
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
	next_i := lasti + (4 - lasti%4)
	literal_int, _ := strconv.ParseInt(string(literal), 2, 32)
	//fmt.Println("Next i:", next_i)
	return int(literal_int), next_i
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
	length, _ := strconv.ParseInt(string(input[7:21]), 2, 32)
	fmt.Println("%s", string(input))
	fmt.Println("Length: ", length)
	i := 22
	for i < int(22+length) {
		newpackets, n := parse_packet(input[i:])
		packets = append(packets, newpackets...)
		i += n

	}
	return packets, int(i + 22 + int(length))

}

func extract_count_packet(input []byte) ([]Packet, int) {
	packets := make([]Packet, 0)
	count, _ := strconv.ParseInt(string(input[7:18]), 2, 32)
	fmt.Println("Count: ", count)
	return packets, 0
}

func extract_operator(input []byte) ([]Packet, int) {
	length_type := length_type(input)
	if length_type == 0 {
		return extract_length_packet(input)
	} else {
		return extract_count_packet(input)
	}
}

func parse_packet(input []byte) ([]Packet, int) {
	//	fmt.Println(string(input))
	version := version(input)
	next_index := 0
	id := packet_id(input)
	packet := Packet{}
	packet.version = version
	packet.id = id
	packets := make([]Packet, 0)
	if id == LITERAL_ID {
		val, n := extract_literal(input)
		packet.literal = val
		next_index = n
		packets = append(packets, packet)
	} else {
		command_packets, n := extract_operator(input)
		packets = append(packets, packet)
		packets = append(packets, command_packets...)
		next_index = n
	}
	return packets, next_index
}

func main() {
	/*for _, c := range []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'} {
		fmt.Printf("%c -> %s\n", c, codewheel[c])
	}*/
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
	ip := input_to_bits([]byte("EE00D40C823060"))
	fmt.Println("Len type: ", length_type(ip))
	a, _ := extract_length_packet(ip)

	ip = input_to_bits([]byte("38006F45291200"))
	fmt.Println("Len type: ", length_type(ip))
	x, _ := extract_length_packet(ip)
	fmt.Println(x)

	ip = input_to_bits([]byte("020D64AEE52E55B"))
	fmt.Printf("Version: %d\n", version(ip))
	fmt.Printf("Id: %d\n", packet_id(ip))
	fmt.Printf("Lentght Type: %d\n", length_type((ip)))
	//y, _ := extract_length_packet(ip)
	fmt.Println("-", a)

	/*
		output := input_to_bits([]byte("ABC123"))
		fmt.Printf("%s\n", string(output))
		fmt.Printf("Version: %d\n", version(input_to_bits([]byte("D2FE28"))))
		fmt.Printf("Id: %d\n", packet_id(input_to_bits([]byte("D2FE28"))))
		fmt.Printf("literal: %d\n", extract_literal(input_to_bits([]byte("D2FE28"))))
	*/

}
