[![Go Version](https://img.shields.io/github/go-mod/go-version/Rob8150/gh)](https://tip.golang.org/doc/go1.18)
[![GoDoc](https://godoc.org/github.com/rwxrob/gh?status.svg)](https://godoc.org/github.com/Rob8150/gh)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rwxrob/gh)](https://goreportcard.com/report/github.com/rwxrob/gh)
[![GoInt](https://github.com/Rob8150/Rob8150/blob/main/golang.png)]
# Weather Utilities in Go 1.19+
🚧 *under construction* 🚧


## Weather Functions

* **Dew Point, Absolute Humidity, Relative Humidity, Vapour Pressure**

* **Weather functions Copyright 2022 Robert Clark under Open Source Apache2.**


### Work in Progress
* **99 % Complete**

##### **Inputs Temperature, Relative Humidity, Barometric Pressure all unints in SI**

* **Absolute Humidity**

Absolute humidity is the total amount of water vapor contained in a given volume of dry air. 
Water vapor is measured in units of weight (grams) and dry air in units of volume called cubic meters

* **Specific Humidity**

Specific humidity is the total amount of water vapor contained in a given volume of dry air. 
Water vapor is measured in units of volume (ml) and dry air in units of volume called cubic meters


* **Vapor Pressure**

Vapor pressure (or vapour pressure in English-speaking countries other than the US; see spelling differences)
or equilibrium vapor pressure is defined as the pressure exerted by a vapor in thermodynamic equilibrium
with its condensed phases (solid or liquid) at a given temperature in a closed system.The equilibrium
vapor pressure is an indication of a liquid's evaporation rate.

__________________________________________________________________________________________________

### Output

```go
Tell me the temp 24.7
Tell me humidity 33
Presure in hpa 1006
Alt in m 0
VAPOUR PARTIAL PRESSURE (HPa)
10.188213546228416
---------
Temp (c)
24.7
---------
Relative Humidity %
33
---------
Baro (HPa)
1006
---------
Vapour Saturation Pressure (HPa)
31.047343166324474
---------
Dew Point (C)
7.252448530556964
---------
Absolute Humidity (g/m^3)
7.411637136694928
---------
Specific Humidity (g/kg
6.323480581122353
---------
Cloud Base (m)
2180.9439336803794
{"T":24.7,"Rh":33,"Mb":1006,"Alt":0,"E":10.188213546228416,"Es":31.047343166324474,"Tdc":7.252448530556964,"Ah":7.411637136694928,"SH":6.323480581122353,"Cb":2180.9439336803794}

```
__________________________________________________________________________________________________


### Contributor License Agreements

Before we can accept your pull requests you'll need to sign a Contributor License Agreement (CLA):

- **If you are an individual writing original source code** and **you own the
intellectual property**, then you'll need to sign an client agreement.
- **If you work for a company that wants to allow you to contribute your
  work**, then you'll need to sign a corporate client agreement.

You can sign these via email by contacting the Author. After that,
we'll be able to accept your pull requests.

Note not all projects may need contributions so just email robert.clark.1967@gmail.com
for license see License.md in this REPO.
