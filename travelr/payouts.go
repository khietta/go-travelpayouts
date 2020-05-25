package travelr

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// NewTraverPayouts ...
func NewTraverPayouts(token, marker string) *bodySetting {
	return &bodySetting{
		Token:       token,
		Host:        host,
		Locale:      locale,
		Marker:      marker,
		TripClass:   tripClass,
		UserIP:      userIP,
		KnowEnglish: "true",
		Passenger:   &Passengers{convertInt(1), convertInt(0), convertInt(0)},
	}
}

// NewSetting ...
func (st *bodySetting) NewSetting(host, userIP, locale, tripClass string) {
	*st = bodySetting{
		Host:      host,
		UserIP:    userIP,
		Locale:    locale,
		TripClass: tripClass,
	}
}

// NewPassengers ...
func (st *bodySetting) NewPassengers(ps *Passengers) {
	if ps != nil {
		*st = bodySetting{
			Passenger: ps,
		}
	}
}

// RoundTrip ...
func (st *bodySetting) RoundTrip(rq []SegmentRequest) (interface{}, error) {
	sg := make([]Segment, 0)
	for _, v := range rq {
		ti := v.Date.Format(timeFormat)
		sg = append(sg, Segment{
			Destination: &v.Destination,
			Date:        &ti,
			Origin:      &v.Origin,
		})
	}

	rt := RoundTripBody{
		Host:        &st.Host,
		Locale:      &st.Locale,
		Marker:      &st.Marker,
		UserIP:      &st.UserIP,
		TripClass:   &st.TripClass,
		Passengers:  st.Passenger,
		Segments:    &sg,
		KnowEnglish: &st.KnowEnglish,
	}

	params := convertStructToMap(rt)

	sig, err := MakeSignatureToken(st.Token, params)
	if err != nil {
		return nil, err
	}

	rt.Signature = &sig

	rs, err := json.Marshal(rt)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.Post(endpoint("v1/flight_search"), "application/json", bytes.NewReader(rs))
	if err != nil {
		return nil, err
	}

	resp, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
