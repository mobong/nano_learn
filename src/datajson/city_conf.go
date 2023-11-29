package datajson

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

var CityData *City

type conditions struct {
	Type  int `json:"type"`
	Level int `json:"level"`
}

type City struct {
	Title      string       `json:"title"`
	Des        string       `json:"des"`
	Name       string       `json:"name"`
	Type       int8         `json:"type"`
	Additions  []int8       `json:"additions"`
	Conditions []conditions `json:"conditions"`
	Levels     []Levels     `json:"levels"`
}

type NeedRes struct {
	Decree int `json:"decree"`
	Grain  int `json:"grain"`
	Wood   int `json:"wood"`
	Iron   int `json:"iron"`
	Stone  int `json:"stone"`
	Gold   int `json:"gold"`
}

type Levels struct {
	Level  int     `json:"level"`
	Values []int   `json:"values"`
	Need   NeedRes `json:"need"`
	Time   int     `json:"time"` //升级需要的时间
}

func Load() {
	dir, _ := os.Getwd()
	fileName := path.Join(dir, "data", "city.json")
	jdata, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("load city data error:", err)
	}
	city := new(City)
	json.Unmarshal(jdata, city)
	CityData = city
	fmt.Println(CityData)
}
