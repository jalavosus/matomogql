package model

import (
	"encoding/json"
)

//go:generate stringer -type LogAction -linecomment
type LogAction uint8

const (
	LogActionAction         LogAction = iota // action
	LogActionEvent                           // event
	LogActionGoal                            // goal
	LogActionEcommerceOrder                  // ecommerceOrder
	logActionUnknown                         // unknown
)

var validLogActions = []LogAction{
	LogActionAction,
	LogActionEvent,
	LogActionGoal,
	LogActionEcommerceOrder,
}

func LogActionFromString(s string) LogAction {
	var l = logActionUnknown
	for _, v := range validLogActions {
		if s == v.String() {
			l = v
			break
		}
	}

	return l
}

func (l LogAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func (l *LogAction) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*l = LogActionFromString(raw)

	return nil
}
