package model

import (
"gorm.io/gorm"
)

type  projectCode2  struct {
UserID  string `json:"UserID"`
AnswerTime  string `json:"AnswerTime"`
EndTime  string `json:"EndTime"`
StartTime  string `json:"StartTime"`
Direction  string `json:"Direction"`
Disposition  string `json:"Disposition"`
Duration  string `json:"Duration"`
Callername  string `json:"Callername"`
Callernumber  string `json:"Callernumber"`
Calleename  string `json:"Calleename"`
Calleenumber  string `json:"Calleenumber"`
LegId  string `json:"LegId"`
RecordingIds  string `json:"RecordingIds"`
Queue  string `json:"Queue"`
Queueid  string `json:"Queueid"`
Queuename  string `json:"Queuename"`
ID  string `json:"ID"`
gorm.Model
}
