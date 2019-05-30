package main

import (
	"fmt"

	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// Run jaeger
// docker run -d --name jaeger \
//     -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
//     -p 5775:5775/udp \
//     -p 6831:6831/udp \
//     -p 6832:6832/udp \
//     -p 5778:5778 \
//     -p 16686:16686 \
//     -p 14268:14268 \
//     -p 9411:9411 \
//     jaegertracing/all-in-one:1.8

func main() {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, closer, err := cfg.New("test", config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Failed to init jaeger, err= %v", err))
	}
	defer func() {
		_ = closer.Close()
	}()

	name := "Sadlil"
	str := fmt.Sprintf("Hello %s", name)

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", name)

	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", name),
	)
	println(str)
	span.LogKV("event", "println")
	span.Finish()
}
