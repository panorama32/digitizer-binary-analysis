package service

import "github.com/panorama32/digitizer-binary-analysis/domain/rawdata"

type RawDataService struct {
	rawDatas []*rawdata.RawData
}

func NewRawDataService(rawDatas []*rawdata.RawData) *RawDataService {
	return &RawDataService{
		rawDatas: rawDatas,
	}
}

func (s *RawDataService) Join() *rawdata.RawData {
	data := []byte{}
	for _, rd := range s.rawDatas {
		data = append(data, rd.Bytes...)
	}

	return &rawdata.RawData{
		Bytes: data,
	}
}
