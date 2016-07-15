package main

type Response struct {
	WeatherInfo WeatherInfo `json:"weatherinfo"`
}

type WeatherInfo struct {
	City    string `json:"city"`
	CityId  string `json:"cityid"`
	Temp    string `json:"temp"`
	WD      string `json:"WD"`
	WS      string `json:"WS"`
	SD      string `json:"SD"`
	WSE     string `json:"WSE"`
	Time    string `json:"time"`
	IsRadar string `json:"isRadar"`
	Radar   string `json:"Radar"`
	NJD     string `json:"njd"`
	QY      string `json:"qy"`
	Rain    string `json:"rain"`
}
