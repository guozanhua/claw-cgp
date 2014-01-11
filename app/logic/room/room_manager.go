// Author: sheppard(ysf1026@gmail.com) 2014-01-08

package room

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/yangsf5/claw-cgp/util"
)

type RoomConfig struct {
	XMLName xml.Name `xml:"room"`
	Games []Game `xml:"game"`
}

type Game struct {
	Name string `xml:"name,attr"`
	Items []Item `xml:"item"`
}

type Item struct {
	Name string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

var (
	config RoomConfig
)


func ReadConfig(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	util.CheckFatal(err)

	err = xml.Unmarshal(content, &config)
	util.CheckFatal(err)

	fmt.Println(config)
}

