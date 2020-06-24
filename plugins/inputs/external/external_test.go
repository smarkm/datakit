package external

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestRCPServer(t *testing.T) {
	s := &Server{
		Listen: "/tmp/dk.sock",
	}

	s.Start(nil)
}

func TestRPC(t *testing.T) {
	wg := sync.WaitGroup{}

	s := &Server{
		Listen: "/tmp/dk.sock",
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Start(nil)
	}()

	time.Sleep(time.Second)

	conn, err := grpc.Dial("unix://"+s.Listen, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := NewDataKitClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Feed(ctx, &Request{
		Lines: []byte(strings.Join([]string{
			`test_a,tag1=val1,tag2=val2 f1=1i,f2=3,f3="abc",f4=T ` + fmt.Sprintf("%d", time.Now().UnixNano()),
			`test_b,tag1=val1,tag2=val2 f1=1i,f2=3,f3="abc",f4=T ` + fmt.Sprintf("%d", time.Now().UnixNano()),
			`test_c,tag1=val1,tag2=val2 f1=1i,f2=3,f3="abc",f4=T ` + fmt.Sprintf("%d", time.Now().UnixNano()),
		}, "\n"))})

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("[C] sending %d points ok, err: %s", r.GetPoints(), r.GetErr())

	r, err = c.Feed(ctx, &Request{
		Lines: []byte(strings.Join([]string{ // bad body
			`test_a tag1=val1,tag2=val2 f1=1i,f2=3,f3="abc",f4=T ` + fmt.Sprintf("%d", time.Now().UnixNano()),
			`test_b tag1=val1,tag2=val2 f1=1i,f2=3,f3="abc",f4=T ` + fmt.Sprintf("%d", time.Now().UnixNano()),
			`test_c tag1=val1,tag2=val2 f1=1i,f2=3,f3="abc",f4=T ` + fmt.Sprintf("%d", time.Now().UnixNano()),
		}, "\n"))})

	if err != nil {
		t.Fatal("should not been here")
	}

	log.Printf("[C] sending points: %d, err: %s", r.GetPoints(), r.GetErr())

	log.Printf("stopping server...")
	s.Stop()

	wg.Wait()
}
