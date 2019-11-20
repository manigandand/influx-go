package main

import (
	"fmt"
	"time"
)

// IP1 ...
const (
	IP1 = `incident,number=%d,incident_id=inc%d,orgnization_id=org%d,service_id=svc%d,alert_source_id=as%d,status=triggered,event_type=triggered fieldKey1="fieldVal1" %+v`
	IP2 = `incident,number=%d,incident_id=inc%d,orgnization_id=org%d,service_id=svc%d,alert_source_id=as%d,status=acknowledged,event_type=acknowledged fieldKey1="fieldVal1",tta=%d %+v`
	IP3 = `incident,number=%d,incident_id=inc%d,orgnization_id=org%d,service_id=svc%d,alert_source_id=as%d,status=resolved,event_type=resolved fieldKey1="fieldVal1",ttr=%d %+v`
)

func main() {
	set1()
	set2()
	set3()
}

func set1() {
	back15 := time.Now().UnixNano() - 1800000000000
	for i := 1; i <= 10; i++ {
		fmt.Printf(IP1+"\n", i, i, i, i, i, back15+int64(i*60000000000))
		fmt.Printf(IP2+"\n", i, i, i, i, i, i*5, back15+int64(i*60000000000))
		fmt.Printf(IP3+"\n", i, i, i, i, i, i*300, back15+int64(i*60000000000))
	}
}

func set2() {
	back15 := time.Now().UnixNano() - 900000000000
	for i := 1; i <= 10; i++ {
		fmt.Printf(IP1+"\n", i+10, i+10, i, i, i, back15+int64(i*60000000000))
		fmt.Printf(IP2+"\n", i+10, i+10, i, i, i, i+2*5, back15+int64(i*60000000000))
		fmt.Printf(IP3+"\n", i+10, i+10, i, i, i, i+3*300, back15+int64(i*60000000000))
	}
}

func set3() {
	back15 := time.Now().UnixNano() - 500000000000
	for i := 1; i <= 10; i++ {
		fmt.Printf(IP1+"\n", i+20, i+20, i, i, i, back15+int64(i*60000000000))
		fmt.Printf(IP2+"\n", i+20, i+20, i, i, i, i+4*5, back15+int64(i*60000000000))
		fmt.Printf(IP3+"\n", i+20, i+20, i, i, i, i+5*300, back15+int64(i*60000000000))
	}
}
