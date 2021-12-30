package solutions

import (
	"testing"

	is_ "github.com/matryer/is"
)

func TestNewPacket(t *testing.T) {
	is := is_.New(t)

	pack := NewPacket(HexToBITS("D2FE28"))
	is.Equal(pack.Version, 6) // version
	is.Equal(pack.Type, 4)    // type

	val, ln := pack.LiteralValue()
	is.Equal(val, 2021) // literal value packet
	is.Equal(ln, 15)    // literal value BITS size

	ln = pack.Len()
	is.Equal(ln, 21)

	pack = NewPacket(HexToBITS("38006F45291200"))
	is.Equal(pack.Version, 1)                        // version
	is.Equal(pack.Type, 6)                           // type
	is.Equal(pack.SubPacketType(), PacketSubTypeLen) // sub packet type
	is.Equal(pack.SubPacketDataLen(), 27)            // sub packet data len

	packets := pack.SubPackets()
	is.Equal(len(packets), 2) // sub packet count

	val, _ = packets[0].LiteralValue() // packet 1 value
	is.Equal(val, 10)

	val, _ = packets[1].LiteralValue() // packet 2 value
	is.Equal(val, 20)

	pack = NewPacket(HexToBITS("EE00D40C823060"))
	is.Equal(pack.Version, 7)                          // version
	is.Equal(pack.Type, 3)                             // type
	is.Equal(pack.SubPacketType(), PacketSubTypeCount) // sub packet type
	is.Equal(pack.SubPacketCount(), 3)                 // sub packet count

	packets = pack.SubPackets()
	is.Equal(len(packets), 3) // sub packet count

	val, _ = packets[0].LiteralValue() // packet 1 value
	is.Equal(val, 1)

	val, _ = packets[1].LiteralValue() // packet 2 value
	is.Equal(val, 2)

	val, _ = packets[2].LiteralValue() // packet 2 value
	is.Equal(val, 3)
}

func TestCompute(t *testing.T) {
	is := is_.New(t)

	is.Equal(NewPacket(HexToBITS("C200B40A82")).Compute(), 3)
	is.Equal(NewPacket(HexToBITS("04005AC33890")).Compute(), 54)
	is.Equal(NewPacket(HexToBITS("880086C3E88112")).Compute(), 7)
	is.Equal(NewPacket(HexToBITS("CE00C43D881120")).Compute(), 9)
	is.Equal(NewPacket(HexToBITS("D8005AC2A8F0")).Compute(), 1)
	is.Equal(NewPacket(HexToBITS("F600BC2D8F")).Compute(), 0)
	is.Equal(NewPacket(HexToBITS("9C005AC2F8F0")).Compute(), 0)
	is.Equal(NewPacket(HexToBITS("9C0141080250320F1802104A08")).Compute(), 1)
}
func TestSumPacketVersions(t *testing.T) {
	is := is_.New(t)

	is.Equal(SumPacketVersions(NewPacket(HexToBITS("8A004A801A8002F478"))), 16)

	is.Equal(SumPacketVersions(NewPacket(HexToBITS("620080001611562C8802118E34"))), 12)
	is.Equal(SumPacketVersions(NewPacket(HexToBITS("C0015000016115A2E0802F182340"))), 23)
	is.Equal(SumPacketVersions(NewPacket(HexToBITS("A0016C880162017C3686B18A3D4780"))), 31)
}

func TestHextoBITS(t *testing.T) {
	is := is_.New(t)

	bits := HexToBITS("FF")
	is.Equal(bits, []bool{true, true, true, true, true, true, true, true})

	bits = HexToBITS("08")
	is.Equal(bits, []bool{false, false, false, false, true, false, false, false})
}

func TestBITStoInt(t *testing.T) {
	is := is_.New(t)

	result := BITStoInt([]bool{true})
	is.Equal(result, 1)

	result = BITStoInt([]bool{true, true})
	is.Equal(result, 3)

	result = BITStoInt([]bool{true, false, true})
	is.Equal(result, 5)

	result = BITStoInt([]bool{true, false, false, false, false, false, false, false})
	is.Equal(result, 128)

	result = BITStoInt([]bool{true, false, false, false, false, false, false, false, false})
	is.Equal(result, 256)

	result = BITStoInt([]bool{true, false, false, false, false, false, false, false, false, false})
	is.Equal(result, 512)
}
