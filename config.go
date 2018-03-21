package healer

import "errors"

type ProducerConfig struct {
	BootstrapServers string
	ClientID         string
	Acks             int16
	CompressionType  string
	BatchSize        int
	MessageMaxCount  int
	FlushIntervalMS  int
	MetadataMaxAgeMS int

	Retries          int
	RequestTimeoutMS int32
}

var DefaultProducerConfig *ProducerConfig = &ProducerConfig{
	ClientID:         "healer",
	Acks:             1,
	RequestTimeoutMS: 30000,
	CompressionType:  "none",
	Retries:          0,
	BatchSize:        16384,
	MessageMaxCount:  10,
	MetadataMaxAgeMS: 300000,
	FlushIntervalMS:  200,
}

var (
	messageMaxCountError   = errors.New("message.max.count must > 0")
	flushIntervalMSError   = errors.New("flush.interval.ms must > 0")
	unknownCompressionType = errors.New("unknown compression type")
	bootstrapServersNotSet = errors.New("bootstrap servers not set")
)

func (config *ProducerConfig) checkValid() error {
	if config.BootstrapServers == "" {
		return bootstrapServersNotSet
	}
	if config.MessageMaxCount <= 0 {
		return messageMaxCountError
	}
	if config.FlushIntervalMS <= 0 {
		return flushIntervalMSError
	}

	switch config.CompressionType {
	case "none":
	case "gzip":
	case "snappy":
	case "lz4":
	default:
		return unknownCompressionType
	}
	return nil
}
