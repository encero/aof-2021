package solutions

import (
	"fmt"

	"github.com/encero/advent-of-code-2021/helpers"
)

const (
	PacketTypeLiteralValue = 4

	PacketSubTypeLen   = false
	PacketSubTypeCount = true
)

func Day16PacketDecoder() error {
	packet := NewPacket(HexToBITS(helpers.ReadLine("inputs/day16.txt")))

	version := SumPacketVersions(packet)

	fmt.Println("Day 16, Part 1", version)

	result := packet.Compute()

	fmt.Println("Day 16, Part 2", result)

	return nil
}

func SumPacketVersions(root Packet) int {
	sum := root.Version
	for _, v := range root.SubPackets() {
		sum += SumPacketVersions(v)
	}
	return sum
}

type Packet struct {
	bits    []bool
	Version int
	Type    int
}

func NewPacket(bits []bool) Packet {
	return Packet{
		bits:    bits,
		Version: BITStoInt(bits[0:3]),
		Type:    BITStoInt(bits[3:6]),
	}
}

func (p Packet) Compute() int {
	subs := p.SubPackets()

	switch p.Type {
	case 0:
		sum := 0
		for _, v := range subs {
			sum += v.Compute()
		}
		return sum
	case 1:
		product := 1
		for _, p := range subs {
			product *= p.Compute()
		}
		return product
	case 2:
		min := 2147483647
		for _, p := range subs {
			val := p.Compute()
			if val < min {
				min = val
			}
		}
		return min
	case 3:
		max := 0
		for _, p := range subs {
			val := p.Compute()
			if val > max {
				max = val
			}
		}
		return max
	case 4:
		val, _ := p.LiteralValue()
		return val
	case 5:
		if subs[0].Compute() > subs[1].Compute() {
			return 1
		}
		return 0
	case 6:
		if subs[0].Compute() < subs[1].Compute() {
			return 1
		}
		return 0
	case 7:
		if subs[0].Compute() == subs[1].Compute() {
			return 1
		}
		return 0
	}

	panic("unknown packet")
}

func (p Packet) LiteralValue() (int, int) {
	offset := 6
	valueBITS := []bool{}
	len := 0

	for {
		len += 5
		part := p.bits[offset : offset+5]

		valueBITS = append(valueBITS, part[1:5]...)

		if !part[0] {
			break
		}
		offset += 5
	}

	return BITStoInt(valueBITS), len
}

func (p Packet) SubPackets() []Packet {
	var packets []Packet
	if p.Type == PacketTypeLiteralValue {
		return nil
	}

	offset := 3 + 3 + 1
	var checkFn func() bool
	loadedBITS := 0

	if p.SubPacketType() == PacketSubTypeLen {
		offset += 15
		checkFn = func() bool {
			return loadedBITS < p.SubPacketDataLen()
		}
	} else {
		offset += 11
		checkFn = func() bool {
			return len(packets) < p.SubPacketCount()
		}
	}

	for checkFn() {
		subPacket := NewPacket(p.bits[offset+loadedBITS:])
		loadedBITS += subPacket.Len()

		packets = append(packets, subPacket)
	}

	return packets
}

func (p Packet) SubPacketType() bool {
	return p.bits[6]
}

func (p Packet) SubPacketDataLen() int {
	return BITStoInt(p.bits[7:22])
}

func (p Packet) SubPacketCount() int {
	return BITStoInt(p.bits[7:18])
}

func (p Packet) Len() int {
	if p.Type == PacketTypeLiteralValue {
		_, len := p.LiteralValue()
		return 3 + 3 + len
	}

	if p.SubPacketType() == PacketSubTypeLen {
		return 3 + 3 + 1 + 15 + p.SubPacketDataLen()
	}

	size := 3 + 3 + 1 + 11
	for _, v := range p.SubPackets() {
		size += v.Len()
	}

	return size
}

func BITStoInt(bits []bool) int {
	var out int

	for i := len(bits) - 1; i >= 0; i-- {
		if bits[i] {
			out |= 1 << (len(bits) - i - 1)
		}
	}

	return out
}

func HexToBITS(in string) []bool {
	var out []bool
	for _, v := range in {
		letter := string(v)
		switch letter {
		case "0":
			out = append(out, []bool{false, false, false, false}...)
		case "1":
			out = append(out, []bool{false, false, false, true}...)
		case "2":
			out = append(out, []bool{false, false, true, false}...)
		case "3":
			out = append(out, []bool{false, false, true, true}...)
		case "4":
			out = append(out, []bool{false, true, false, false}...)
		case "5":
			out = append(out, []bool{false, true, false, true}...)
		case "6":
			out = append(out, []bool{false, true, true, false}...)
		case "7":
			out = append(out, []bool{false, true, true, true}...)
		case "8":
			out = append(out, []bool{true, false, false, false}...)
		case "9":
			out = append(out, []bool{true, false, false, true}...)
		case "A":
			out = append(out, []bool{true, false, true, false}...)
		case "B":
			out = append(out, []bool{true, false, true, true}...)
		case "C":
			out = append(out, []bool{true, true, false, false}...)
		case "D":
			out = append(out, []bool{true, true, false, true}...)
		case "E":
			out = append(out, []bool{true, true, true, false}...)
		case "F":
			out = append(out, []bool{true, true, true, true}...)
		}
	}

	return out
}

func DumpBITS(bits []bool) {
	for _, v := range bits {
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
}
