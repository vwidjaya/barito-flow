package flow

import (
	"github.com/BaritoLog/go-boilerplate/errkit"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	stpb "github.com/golang/protobuf/ptypes/struct"
	lokipb "github.com/vwidjaya/barito-proto/loki"
	prodpb "github.com/vwidjaya/barito-proto/producer"
)

const (
	JsonParseError  = errkit.Error("JSON Parse Error")
	ProtoParseError = errkit.Error("Protobuf Parse Error")
)

func ConvertTimberToKafkaMessage(timber *prodpb.Timber, topic string) *sarama.ProducerMessage {
	b, _ := proto.Marshal(timber)

	return &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(b),
	}
}

func ConvertKafkaMessageToTimber(message *sarama.ConsumerMessage) (timber prodpb.Timber, err error) {
	err = proto.Unmarshal(message.Value, &timber)
	if err != nil {
		err = errkit.Concat(ProtoParseError, err)
		return
	}

	return
}

func ConvertTimberToEsDocumentString(timber prodpb.Timber, m *jsonpb.Marshaler) string {
	doc := ExtractTimberContents(timber)
	docStr, _ := m.MarshalToString(doc)
	return docStr
}

func ConvertTimberToLokiEntry(timber prodpb.Timber, m *jsonpb.Marshaler) *lokipb.Entry {
	doc := ExtractTimberContents(timber)
	line, _ := m.MarshalToString(doc)

	return &lokipb.Entry{
		Timestamp: ptypes.TimestampNow(),
		Line:      line,
	}
}

func ExtractTimberContents(timber prodpb.Timber) (doc *stpb.Struct) {
	doc = timber.GetContent()

	ts := &stpb.Value{
		Kind: &stpb.Value_StringValue{
			StringValue: timber.GetTimestamp(),
		},
	}
	doc.Fields["@timestamp"] = ts
	return
}
