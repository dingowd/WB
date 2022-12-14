package models

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    int     `json:"temp_kf"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Clouds struct {
	All int `json:"all"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Sys struct {
	Pod string `json:"pod"`
}

type List struct {
	Dt         int       `json:"dt"`
	Main       Main      `json:"main"`
	Weather    []Weather `json:"weather"`
	Clouds     Clouds    `json:"clouds"`
	Wind       Wind      `json:"wind"`
	Visibility int       `json:"visibility"`
	Pop        int       `json:"pop"`
	Sys        Sys       `json:"sys"`
	DtTxt      string    `json:"dt_txt"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type CityProp struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Coord      Coord  `json:"coord"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int    `json:"sunrise"`
	Sunset     int    `json:"sunset"`
}

type Resp struct {
	CityId  int      `json:"-"`
	Cod     string   `json:"cod"`
	Message int      `json:"message"`
	Cnt     int      `json:"cnt"`
	List    []List   `json:"list"`
	City    CityProp `json:"city"`
}

type ShortWeather struct {
	Country string   `json:"country"`
	City    string   `json:"city"`
	Date    string   `json:"date"`
	AvTemp  float64  `json:"av_temp"`
	Dates   []string `json:"dates"`
}

type RespWithDate struct {
	CityId  int      `json:"-"`
	Date    string   `json:"date"`
	Cod     string   `json:"cod"`
	Message int      `json:"message"`
	Cnt     int      `json:"cnt"`
	List    []List   `json:"list"`
	City    CityProp `json:"city"`
}
