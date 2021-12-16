package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"
)

const expectedResult = 908

/*
 *   --- Day 16: Packet Decoder ---
 *          --- Part One ---
 *
 *   https://adventofcode.com/2021/day/16
 */

const input = "00537390040124EB240B3EDD36B68014D4C9ECCCE7BDA54E62522A300525813003560004223BC3F834200CC108710E98031C94C8B4BFFF42398309DDD30EEE00BCE63F03499D665AE57B698F9802F800824DB0CE1CC23100323610069D8010ECD4A5CE5B326098419C319AA2FCC44C0004B79DADB1EB48CE5EB7B2F4A42D9DF0AA74E66468C0139341F005A7BBEA5CA65F3976200D4BC01091A7E155991A7E155B9B4830056C01593829CC1FCD16C5C2011A340129496A7EFB3CA4B53F7D92675A947AB8A016CD631BE15CD5A17CB3CEF236DBAC93C4F4A735385E401804AA86802D291ED19A523DA310006832F07C97F57BC4C9BBD0764EE88800A54D5FB3E60267B8ED1C26AB4AAC0009D8400854138450C4C018855056109803D11E224112004DE4DB616C493005E461BBDC8A80350000432204248EA200F4148FD06C804EE1006618419896200FC1884F0A00010A8B315A129009256009CFE61DBE48A7F30EDF24F31FCE677A9FB018F6005E500163E600508012404A72801A4040688010A00418012002D51009FAA0051801CC01959801AC00F520027A20074EC1CE6400802A9A004A67C3E5EA0D3D5FAD3801118E75C0C00A97663004F0017B9BD8CCA4E2A7030C0179C6799555005E5CEA55BC8025F8352A4B2EC92ADF244128C44014649F52BC01793499EA4CBD402697BEBD18D713D35C9344E92CB67D7DFF05A60086001610E21A4DD67EED60A8402415802400087C108DB068001088670CA0DCC2E10056B282D6009CFC719DB0CD3980026F3EEF07A29900957801AB8803310A0943200042E3646789F37E33700BE7C527EECD13266505C95A50F0C017B004272DCE573FBB9CE5B9CAE7F77097EC830401382B105C0189C1D92E9CCE7F758B91802560084D06CC7DD679BC8048AF00400010884F18209080310FE0D47C94AA00"

var versions []string

func main() {
	defer func(t time.Time) {
		log.Println("took", time.Since(t))
	}(time.Now())

	versionSum := doPart1()
	log.Println("the sum of the version numbers from all packets is", versionSum)
}

func doPart1() int {
	bytes, err := hex.DecodeString(input)
	if err != nil {
		log.Panicln(err)
	}
	binary := ""
	for _, b := range bytes {
		B := fmt.Sprintf("%b", b)
		for i := len(B); i < 8; i++ {
			B = "0" + B
		}
		binary += B
	}

	versions = make([]string, 0)
	_ = parsePacket(binary)

	versionSum := 0
	for _, version := range versions {
		v, err := strconv.ParseInt(version, 2, 64)
		if err != nil {
			log.Panicln(err)
		}
		versionSum += int(v)
	}
	return versionSum
}

func parsePacket(binary string) string {
	version := binary[:3]
	versions = append(versions, version)

	typeID := binary[3:6]
	binary = binary[6:]

	switch typeID {
	case "100":
		binary = extractLiteral(binary)

	default:
		lengthTypeID := binary[0]
		binary = binary[1:]

		if lengthTypeID == '0' {
			binary = extractOperatorSubpacketBits(binary, 15)
		} else {
			binary = extractOperatorNSubpackets(binary, 11)
		}
	}

	return binary
}

func extractOperatorSubpacketBits(binary string, nSubpacketBits int) string {
	lenSubpacket, err := strconv.ParseInt(binary[:nSubpacketBits], 2, 64)
	if err != nil {
		log.Panicln(err)
	}
	binary = binary[nSubpacketBits:]

	for lenSubpacket > 0 {
		lenBefore := len(binary)
		binary = parsePacket(binary)
		nBits := lenBefore - len(binary)
		lenSubpacket -= int64(nBits)
	}

	return binary
}

func extractOperatorNSubpackets(binary string, nSubpacketBits int) string {
	nSubpackets, err := strconv.ParseInt(binary[:nSubpacketBits], 2, 64)
	if err != nil {
		log.Panicln(err)
	}
	binary = binary[nSubpacketBits:]

	for nSubpackets > 0 {
		binary = parsePacket(binary)
		nSubpackets--
	}

	return binary
}

func extractLiteral(binary string) string {
	var i int
	var continueBit byte
	for i = 5; i <= len(binary) && continueBit != '0'; i += 5 {
		continueBit = binary[i-5]
	}

	return binary[i-5:]
}
