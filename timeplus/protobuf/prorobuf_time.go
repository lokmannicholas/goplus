package timeplus

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func ProtobufTimeStampToTime(tt *timestamp.Timestamp) (*time.Time, error) {

	if t, err := ptypes.Timestamp(tt); err == nil {
		return &t, nil
	} else {
		return nil, err
	}

}

func ProtobufTimeToTimeStamp(t *time.Time) (*timestamp.Timestamp, error) {
	if tt, err := ptypes.TimestampProto(*t); err == nil {
		return tt, nil
	} else {
		return nil, err
	}

}
