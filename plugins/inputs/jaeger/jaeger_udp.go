package jaeger

import (
	"context"
	"net"
	"time"

	"github.com/uber/jaeger-client-go/thrift"
	"github.com/uber/jaeger-client-go/thrift-gen/agent"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	itrace "gitlab.jiagouyun.com/cloudcare-tools/datakit/io/trace"
)

func StartUDPAgent(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}

	log.Debugf("%s(UDP): listen on path: %s", inputName, addr)

	// receiving loop
	buf := make([]byte, 65535)
	for {
		select {
		case <-datakit.Exit.Wait():
			if err := udpConn.Close(); err != nil {
				log.Warnf("Close: %s", err)
			}
			log.Infof("jaeger udp agent closed")

			return nil
		default:
		}

		err := udpConn.SetDeadline(time.Now().Add(time.Second))
		if err != nil {
			log.Errorf("SetDeadline failed: %v", err)
			continue
		}

		n, addr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Debug(err.Error())
			continue
		}
		log.Debugf("Read from udp server:%s %d bytes", addr, n)

		if n <= 0 {
			continue
		}

		dktrace, err := parseJaegerUDP(buf[:n])
		if err != nil {
			continue
		}
		if len(dktrace) != 0 {
			itrace.StatTracingInfo(dktrace)
			itrace.BuildPointsBatch(inputName, dktrace, false)
		} else {
			log.Debug("empty batch")
		}
	}
}

func parseJaegerUDP(data []byte) (itrace.DatakitTrace, error) {
	thriftBuffer := thrift.NewTMemoryBufferLen(len(data))
	if _, err := thriftBuffer.Write(data); err != nil {
		log.Error("buffer write failed :%v,", err)

		return nil, err
	}

	protocolFactory := thrift.NewTCompactProtocolFactoryConf(&thrift.TConfiguration{})
	thriftProtocol := protocolFactory.GetProtocol(thriftBuffer)
	_, _, _, err := thriftProtocol.ReadMessageBegin(context.TODO()) //nolint:dogsled
	if err != nil {
		log.Error("read message begin failed :%v,", err)

		return nil, err
	}

	batch := agent.AgentEmitBatchArgs{}
	err = batch.Read(context.TODO(), thriftProtocol)
	if err != nil {
		log.Error("read batch failed :%v,", err)

		return nil, err
	}

	groups, err := batchToAdapters(batch.Batch)
	if err != nil {
		log.Error("process batch failed :%v,", err)

		return nil, err
	}

	err = thriftProtocol.ReadMessageEnd(context.TODO())
	if err != nil {
		log.Error("read message end failed :%v,", err)

		return nil, err
	}

	return groups, nil
}
