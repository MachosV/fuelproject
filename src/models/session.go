package models

import "time"

type Session struct{
	id string
	timestamp time.Time
	values map[string]string
}

func (s *Session) SetId(sessionid string){
	s.id = sessionid
}

func (s *Session) SetTimestamp(){
	s.timestamp = time.Now()
}

func (s *Session) Init(){
	s.values = make(map[string]string)
}

func (s *Session) SetValue(key string, value string){
	s.values[key] = value
}

func (s *Session) GetTimestamp() time.Time{
	return s.timestamp
}
