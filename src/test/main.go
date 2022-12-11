package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/filipkroca/teltonikaparser"
	"io/ioutil"
	"log"
)

type JsonDecoded struct {
	IMEI     string               `json:"IMEI"`     // IMEI number, if len==15 also validated by checksum
	CodecID  byte                 `json:"CodecID"`  // 0x08 (codec 8) or 0x8E (codec 8 extended)
	NoOfData uint8                `json:"NoOfData"` // Number of Data
	Data     []*JsonDecodeAvlData `json:"Data"`     // Slice with avl data
	Response []byte               `json:"Response"` // Slice with a response
}

func (de *JsonDecoded) FromStruct(item *teltonikaparser.Decoded) (*JsonDecoded, error) {
	de.IMEI = item.IMEI
	de.CodecID = item.CodecID
	de.NoOfData = item.NoOfData
	de.Response = item.Response
	if len(item.Data) > 0 {
		datas := make([]*JsonDecodeAvlData, 0)
		for _, val := range item.Data {
			data, err := new(JsonDecodeAvlData).FromStruct(&val)
			if err != nil {
				return nil, err
			}
			datas = append(datas, data)
		}
		if len(datas) > 0 {
			de.Data = datas
		}
	}
	return de, nil
}

type JsonDecodeAvlData struct {
	UtimeMs     uint64                    `json:"UtimeMs"`  // Utime in mili seconds
	Utime       uint64                    `json:"Utime"`    // Utime in seconds
	Priority    uint8                     `json:"Priority"` // Priority, 	[0	Low, 1	High, 2	Panic]
	Lat         int32                     `json:"Lat"`      // Latitude (between 850000000 and -850000000), fit int32
	Lng         int32                     `json:"Lng"`      // Longitude (between 1800000000 and -1800000000), fit int32
	Altitude    int16                     `json:"Altitude"` // Altitude In meters above sea level, 2 bytes
	Angle       uint16                    `json:"Angle"`    // Angle In degrees, 0 is north, increasing clock-wise, 2 bytes
	VisSat      uint8                     `json:"VisSat"`   // Satellites Number of visible satellites
	Speed       uint16                    `json:"Speed"`    // Speed in km/h
	EventID     uint16                    `json:"EventID"`  // Event generated (0 – data generated not on event)
	FinalValues []*FinalValue             `json:"Elements"` // Human decoded for IO Elements
	Elements    []teltonikaparser.Element `json:"-"`        // Slice containing parsed IO Elements, no need to return this value
}

func (d *JsonDecodeAvlData) FromStruct(item *teltonikaparser.AvlData) (*JsonDecodeAvlData, error) {
	d.UtimeMs = item.UtimeMs
	d.Utime = item.Utime
	d.Priority = item.Priority
	d.Lat = item.Lat
	d.Lng = item.Lng
	d.Altitude = item.Altitude
	d.Angle = item.Angle
	d.VisSat = item.VisSat
	d.Speed = item.Speed
	d.EventID = item.EventID
	if len(item.Elements) > 0 {
		finalValues := make([]*FinalValue, 0)
		for _, ioel := range item.Elements {
			finalValue, err := new(FinalValue).FromStruct(&ioel)
			if err != nil {
				return nil, err
			}
			finalValues = append(finalValues, finalValue)
		}
		if len(finalValues) > 0 {
			d.FinalValues = finalValues
		}
	}
	return d, nil
}

type FinalValue struct {
	IOID  uint16      `json:"IOID"`
	Name  string      `json:"Name"`
	Value interface{} `json:"Value"`
}

func (v *FinalValue) FromStruct(item *teltonikaparser.Element) (*FinalValue, error) {
	var humanDecoder = teltonikaparser.HumanDecoder{}
	// decode to human readable format
	decoded, err := humanDecoder.Human(item, "FMBXY") // second parameter - device family type ["FMBXY", "FM64"]
	if err != nil {
		return nil, err
	}

	// get final decoded value to value which is specified in ./teltonikajson/ in paramether FinalConversion
	if val, err := (*decoded).GetFinalValue(); err != nil {
		return nil, err
	} else if val != nil {
		v.Name = decoded.AvlEncodeKey.PropertyName
		v.Value = val
	}
	return v, nil
}

func WriteJsonString(item interface{}) (string, error) {
	jsonData, err := json.MarshalIndent(item, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func WriteJsonFile(item interface{}, filename string) error {
	jsonData, err := json.MarshalIndent(item, "", "\t")
	if err != nil {
		return err
	}
	// write to json file
	return ioutil.WriteFile(filename, jsonData, 0644)
}

func main() {
	// test with Codec8 Extended packet
	//stringData := `0086cafe0101000f3335323039333038353639383230368e0100000167efa919800200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000710000fc00000900b5000000b600000042305600cd432a00ce6064001100090012ff22001303d1000f0000000200f1000059d900100000000000000000010086cafe0191000f3335323039333038353639383230368e0100000167efad92080200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000715800fc01000900b5000000b600000042039d00cd432a00ce60640011015f0012fd930013036f000f0000000200f1000059d900100000000000000000010086cafe01a0000f3335323039333038353639383230368e01000000f9cebaeac80200000000000000000000000000000000fc0013000800ef0000f00000150000c80000450200010000710000fc00000900b5000000b600000042305400cd000000ce0000001103570012fe8900130196000f0000000200f10000000000100000000000000000010083cafe0101000f3335323039333038353639383230368e0100000167f1aeec00000a750e8f1d43443100f800b210000000000012000700ef0000f00000150500c800004501000100007142000900b5000600b6000500422fb300cd432a00ce60640011000700120007001303ec000f0000000200f1000059d90010000000000000000001`
	stringData := `03D8CAFE0196000F3836363732383036303438383130328E0C00000183D180E3B500175C4C060CDF3660FFFD00000A00000000000C000600160000470300F00000150500C80000EF0000060043261A0044003300B5000C00B6000800422FDA0018000000000000000000000183D17C4FC200175C4C060CDF3660FFF700000B00000000000C000600160000470300F00000150500C80000EF000006004326220044003300B5000C00B6000800422FE00018000000000000000000000183D177BBD200175C4C060CDF3660001500000B00000000000C000600160000470300F00000150400C80000EF000006004326200044003300B5000C00B6000800422FAF0018000000000000000000000183D17327DE00175C4C060CDF3660002400000A00000000000C000600160000470300F00000150400C80000EF0000060043261E0044003300B5000C00B6000A00422FE70018000000000000000000000183D16E93EF00175C4C060CDF3660001D00000A00000000000C000600160000470300F00000150500C80000EF000006004326190044003300B5000C00B6000800422FC90018000000000000000000000183D169FFFF00175C4C060CDF3660001E00000900000000000C000600160000470300F00000150400C80000EF000006004326150044003300B5000C00B6000800422F9D0018000000000000000000000183D1656C0C00175C4C060CDF3660001E00000900000000000C000600160000470300F00000150400C80000EF000006004326160044003300B5000C00B6000A00422FE20018000000000000000000000183D160D81C00175C4C060CDF3660001C00000900000000000C000600160000470300F00000150400C80000EF000006004326140044003300B5000C00B6000A00422FE20018000000000000000000000183D15C442C00175C4C060CDF3660001900000A00000000000C000600160000470300F00000150400C80000EF0000060043260E0044003300B5000C00B6000800422FAE0018000000000000000000000183D157B03900175C4C060CDF3660001A00000A00000000000C000600160000470300F00000150400C80000EF000006004326090044003300B5000D00B6000C004230200018000000000000000000000183D15371D000175C4C060CDF3660001A00000A000000EF000C000600160000470300F00000150500C80000EF000006004326090044003300B5000C00B6000800422FEF0018000000000000000000000183D1536CA800175C4C060CDF3660001A000009000000F0000C000600160100470300F00000150500C80000EF010006004326040044003300B5000C00B6000A00422FED001800000000000000000C000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000`

	bs, _ := hex.DecodeString(stringData)

	// decode a raw data byte slice
	parsedData, err := teltonikaparser.Decode(&bs)
	if err != nil {
		log.Panicf("error when decoding bs, %v\n", err)
	}

	transformEntity, err := new(JsonDecoded).FromStruct(&parsedData)
	if err != nil {
		log.Panicf("error when transform entity item, %v\n", err)
	}

	// print json data
	data, _ := WriteJsonString(transformEntity)
	fmt.Println(data)

	// write json data to file
	folderPath := "./data"
	err = WriteJsonFile(transformEntity, fmt.Sprintf("%s/%s-%d.json", folderPath, transformEntity.IMEI, transformEntity.NoOfData))
	if err != nil {
		log.Panicf("error when write to json file, %v\n", err)
	}
	log.Println("success writing to json file")
}
