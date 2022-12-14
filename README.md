# How to run code

1. Download the dependencies
```bash
go mod download
```
2. Go to src/test and run main.go
```bash
go run main.go
```
3. The json file will automatically be generated. Filename format will be `{IMEI}-{NoOfData}.json`

- Data example
```json
{
  "IMEI": "866728060488102",
  "CodecID": 142,
  "NoOfData": 12,
  "Data": [
    {
      "UtimeMs": 1665667228597,
      "Utime": 1665667228,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": -3,
      "Angle": 0,
      "VisSat": 10,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 5
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9754
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 8
        },
        {
          "Name": "External Voltage",
          "Value": 12250
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665666928578,
      "Utime": 1665666928,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": -9,
      "Angle": 0,
      "VisSat": 11,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 5
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9762
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 8
        },
        {
          "Name": "External Voltage",
          "Value": 12256
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665666628562,
      "Utime": 1665666628,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 21,
      "Angle": 0,
      "VisSat": 11,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 4
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9760
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 8
        },
        {
          "Name": "External Voltage",
          "Value": 12207
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665666328542,
      "Utime": 1665666328,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 36,
      "Angle": 0,
      "VisSat": 10,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 4
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9758
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 10
        },
        {
          "Name": "External Voltage",
          "Value": 12263
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665666028527,
      "Utime": 1665666028,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 29,
      "Angle": 0,
      "VisSat": 10,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 5
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9753
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 8
        },
        {
          "Name": "External Voltage",
          "Value": 12233
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665665728511,
      "Utime": 1665665728,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 30,
      "Angle": 0,
      "VisSat": 9,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 4
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9749
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 8
        },
        {
          "Name": "External Voltage",
          "Value": 12189
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665665428492,
      "Utime": 1665665428,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 30,
      "Angle": 0,
      "VisSat": 9,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 4
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9750
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 10
        },
        {
          "Name": "External Voltage",
          "Value": 12258
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665665128476,
      "Utime": 1665665128,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 28,
      "Angle": 0,
      "VisSat": 9,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 4
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9748
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 10
        },
        {
          "Name": "External Voltage",
          "Value": 12258
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665664828460,
      "Utime": 1665664828,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 25,
      "Angle": 0,
      "VisSat": 10,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 4
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9742
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 8
        },
        {
          "Name": "External Voltage",
          "Value": 12206
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665664528441,
      "Utime": 1665664528,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 26,
      "Angle": 0,
      "VisSat": 10,
      "Speed": 0,
      "EventID": 0,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 4
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9737
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 13
        },
        {
          "Name": "GNSS HDOP",
          "Value": 12
        },
        {
          "Name": "External Voltage",
          "Value": 12320
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665664250320,
      "Utime": 1665664250,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 26,
      "Angle": 0,
      "VisSat": 10,
      "Speed": 0,
      "EventID": 239,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 0
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 5
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 0
        },
        {
          "Name": "Battery Voltage",
          "Value": 9737
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 8
        },
        {
          "Name": "External Voltage",
          "Value": 12271
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    },
    {
      "UtimeMs": 1665664249000,
      "Utime": 1665664249,
      "Priority": 0,
      "Lat": 215955040,
      "Lng": 391924742,
      "Altitude": 26,
      "Angle": 0,
      "VisSat": 9,
      "Speed": 0,
      "EventID": 240,
      "Elements": [
        {
          "Name": "BLE 3 Battery Voltage",
          "Value": 1
        },
        {
          "Name": "Dallas Temperature ID 4",
          "Value": "\u0003"
        },
        {
          "Name": "Movement",
          "Value": 0
        },
        {
          "Name": "GSM Signal",
          "Value": 5
        },
        {
          "Name": "Sleep Mode",
          "Value": 0
        },
        {
          "Name": "Ignition",
          "Value": 1
        },
        {
          "Name": "Battery Voltage",
          "Value": 9732
        },
        {
          "Name": "Battery Current",
          "Value": 51
        },
        {
          "Name": "GNSS PDOP",
          "Value": 12
        },
        {
          "Name": "GNSS HDOP",
          "Value": 10
        },
        {
          "Name": "External Voltage",
          "Value": 12269
        },
        {
          "Name": "Speed",
          "Value": 0
        }
      ]
    }
  ],
  "Response": "AAXK/gGWDA=="
}
```