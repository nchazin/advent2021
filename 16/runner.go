package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

var LITERAL_ID = int(4)

type Packet struct {
	value   uint64
	id      int
	version int
	length  int
}

var total_version int

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

var operators = map[int]string{
	0: "sum",
	1: "product",
	2: "min",
	3: "max",
	5: "gt",
	6: "lt",
	7: "equals",
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

func extract_literal(input []byte) (uint64, int) {
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
	literal_int, _ := strconv.ParseUint(string(literal), 2, 64)
	return uint64(literal_int), lasti
}

func length_type(input []byte) int {
	if input[6] == '1' {
		return 1
	} else {
		return 0
	}
}

func extract_length_packet(input []byte, packet *Packet) int {
	packets := make([]Packet, 0)
	length, _ := strconv.ParseInt(string(input[7:22]), 2, 32)
	i := 22
	for i < int(22+length) {
		newpacket, n := parse_packet(input[i:])
		packets = append(packets, newpacket)
		i += n
	}
	packet.value = do_operator(packet_id(input), packets)
	return int(i)

}

func extract_count_packet(input []byte, packet *Packet) int {
	packets := make([]Packet, 0)
	count, _ := strconv.ParseInt(string(input[7:18]), 2, 32)
	cur_point := 18
	for i := 0; i < int(count); i++ {
		packet, n := parse_packet(input[cur_point:])
		packets = append(packets, packet)
		cur_point += n
	}
	packet.value = do_operator(packet_id(input), packets)

	return cur_point
}

func extract_operator(input []byte, packet *Packet) int {
	length_type := length_type(input)
	if length_type == 0 {
		return extract_length_packet(input, packet)
	} else {
		return extract_count_packet(input, packet)
	}
}

func do_operator(operator int, packets []Packet) uint64 {
	//fmt.Printf("%d %s - %v\n", operator, operators[operator], packets)
	/*
		Packets with type ID 0 are sum packets - their value is the sum of the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
		Packets with type ID 1 are product packets - their value is the result of multiplying together the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
		Packets with type ID 2 are minimum packets - their value is the minimum of the values of their sub-packets.
		Packets with type ID 3 are maximum packets - their value is the maximum of the values of their sub-packets.
		Packets with type ID 5 are greater than packets - their value is 1 if the value of the first sub-packet is greater than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
		Packets with type ID 6 are less than packets - their value is 1 if the value of the first sub-packet is less than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
		Packets with type ID 7 are equal to packets - their value is 1 if the value of the first sub-packet is equal to the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
	*/
	switch operator {
	case 0:
		return do_sum(packets)
	case 1:
		return do_product(packets)
	case 2:
		return do_min(packets)
	case 3:
		return do_max(packets)
	case 5:
		return do_gt(packets)
	case 6:
		return do_lt(packets)
	case 7:
		return do_eq(packets)
	}
	return 0
}

func do_sum(packets []Packet) uint64 {
	sum := uint64(0)
	for _, p := range packets {
		sum += p.value
	}
	return sum
}

func do_product(packets []Packet) uint64 {
	product := uint64(1)
	for _, p := range packets {
		product *= p.value
	}
	return product
}

func do_min(packets []Packet) uint64 {
	min := uint64(math.MaxUint64)
	for _, p := range packets {
		if p.value < min {
			min = p.value
		}
	}
	return min
}

func do_max(packets []Packet) uint64 {
	max := uint64(0)
	for _, p := range packets {
		if p.value > max {
			max = p.value
		}
	}
	return max
}

func do_gt(packets []Packet) uint64 {
	if packets[0].value > packets[1].value {
		return 1
	}
	return 0
}

func do_lt(packets []Packet) uint64 {
	if packets[0].value < packets[1].value {
		return 1
	}
	return 0
}

func do_eq(packets []Packet) uint64 {
	if packets[0].value == packets[1].value {
		return 1
	}
	return 0
}

func parse_packet(input []byte) (Packet, int) {
	version := version(input)
	total_version += version
	next_index := 0
	id := packet_id(input)
	packet := Packet{}
	packet.version = version
	packet.id = id
	packets := make([]Packet, 0)
	if id == LITERAL_ID {
		val, n := extract_literal(input)
		packet.value = val
		packet.length = n
		next_index = n
		packets = append(packets, packet)
	} else {
		n := extract_operator(input, &packet)
		next_index = n
	}
	return packet, next_index
}

func do_main(data []byte) {
	op, _ := parse_packet(data)
	fmt.Println("value: ", op.value)
	fmt.Println("versions: ", total_version)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	do_main(input_to_bits(data))
}
