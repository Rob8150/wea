package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var Tc float64 = 26.2
var RH float64 = 27
var Mb float64 = 1013
var Alt float64 = 0

var TestScenarios []testScenario

type testScenario = struct {
	T   float64 `json:"T"`
	RH  float64 `json:"Rh"`
	Mb  float64 `json:"Mb"`
	Alt float64 `json:"Alt"`
	E   float64 `json:"E"`
	Es  float64 `json:"Es"`
	Tdc float64 `json:"Tdc"`
	Ah  float64 `json:"Ah"`
	SH  float64 `json:"SH"`
	Cb  float64 `json:"Cb"`
}

func main() {
	//Calculate the Dew Point from Temperature and relative humidity.
	//Calculate the Absolute Humidity from Temperature and relative humidity.
	text2 := getInput("Tell me the temp ")
	text3 := getInput("Tell me humidity ")
	text4 := getInput("Presure in hpa ")
	text5 := getInput("Alt in m ")

	Tc = getNum(text2)
	RH = getNum(text3)
	Mb = getNum(text4)
	Alt = getNum(text5)

	E := VapourPartial(Tc, RH)

	Es := VapourSat(Tc)

	Tdc := DewPoint(Tc, RH)

	Ah := AbsoluteHum(Tc, Tdc)

	SH := SpecificHum(Mb, E)

	Cb := CloudBase(Tc, Tdc, Alt)

	fmt.Println("VAPOUR PARTIAL PRESSURE (HPa)")
	fmt.Println(E)
	fmt.Println("---------")

	fmt.Println("Temp (c)")
	fmt.Println(Tc)
	fmt.Println("---------")

	fmt.Println("Relative Humidity %")
	fmt.Println(RH)
	fmt.Println("---------")

	fmt.Println("Baro (HPa)")
	fmt.Println(Mb)
	fmt.Println("---------")

	fmt.Println("Vapour Saturation Pressure (HPa)")
	fmt.Println(Es)
	fmt.Println("---------")

	fmt.Println("Dew Point (C)")
	fmt.Println(Tdc)
	fmt.Println("---------")

	fmt.Println("Absolute Humidity (g/m^3)")
	fmt.Println(Ah)
	fmt.Println("---------")

	fmt.Println("Specific Humidity (g/kg")
	fmt.Println(SH)
	fmt.Println("---------")

	fmt.Println("Cloud Base (m)")
	fmt.Println(Cb)

	//TestScenarios = append(TestScenarios, testScenario{T: Tc, RH: RH, Mb: Mb, Alt: Alt, E: E, Es:Es, Tdc: Tdc, Ah: Ah, SH: SH, Cb: Cb})
	TestScenarios = append(TestScenarios, testScenario{T: Tc, RH: RH, Mb: Mb, Alt: Alt, E: E, Es: Es, Tdc: Tdc, Ah: Ah, SH: SH, Cb: Cb})

	b := ObjToJson(TestScenarios[0])
	fmt.Println(string(b))

	//fmt.Println(TestScenarios[0])

}

func VapourPartial(Tc float64, RH float64) float64 {
	E := (6.11 * math.Pow(10, (7.5*(Tc-(14.55+0.114*Tc)*(1-(0.01*RH))-math.Pow(((2.5+0.007*Tc)*(1-(0.01*RH))), 3)-(15.9+0.117*Tc)*math.Pow((1-(0.01*RH)), 14))/(237.7+(Tc-(14.55+0.114*Tc)*(1-(0.01*RH))-math.Pow(((2.5+0.007*Tc)*(1-(0.01*RH))), 3)-(15.9+0.117*Tc)*math.Pow((1-(0.01*RH)), 14))))))
	return E
}

func VapourSat(Tc float64) float64 {
	var nu float64 = 6.11
	var Po float64 = (7.5 * Tc / (237.7 + Tc))
	A := math.Pow(10, Po)
	Es := nu * A
	return Es
}

func DewPoint(Tc float64, RH float64) float64 {
	Tdc := (Tc - (14.55+0.114*Tc)*(1-(0.01*RH)) - math.Pow(((2.5+0.007*Tc)*(1-(0.01*RH))), 3) - (15.9+0.117*Tc)*math.Pow((1-(0.01*RH)), 14))
	return Tdc
}

func CloudBase(Tc float64, Tdc float64, Alt float64) float64 {
	Cb := ((Tc-Tdc)/8)*1000 + Alt
	return Cb
}

func AbsoluteHum(Tc float64, Tdc float64) float64 {
	Ah := ((6.11 * math.Pow(10.0, (7.5*Tdc/(237.7+Tdc)))) * 100) / ((Tc + 273.16) * 461.5)
	return Ah * 1000
}

func SpecificHum(Mb float64, E float64) float64 {
	SH := (0.622 * E) / (Mb - (0.378 * E))
	return SH * 1000
}

func ObjToJson(obj testScenario) []byte {
	b, err := json.Marshal(TestScenarios[0])
	if err != nil {
		fmt.Println(err)
	}
	return b
}

// getInput asks for a string from keyboard
func getInput(askfor string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Print(askfor)
	scanner.Scan()
	text = scanner.Text()
	return text
}

// getNum convert string to float64
func getNum(given string) float64 {
	var clean string
	var gotnum float64
	clean = strings.TrimSpace(given)
	gotnum, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		log.Fatal(err)
	}
	return gotnum
}
