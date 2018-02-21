package main

import (
	"github.com/hazelcast/hazelcast-go-client"
	"fmt"
	"time"
)

func getKey(client hazelcast.IHazelcastInstance, i int,pi int) {
	start := time.Now()
	mp, _ := client.GetMap("simpleExample")
	key := fmt.Sprintf("%s%d-%d", "key_", i,pi)
	result1, _ := mp.Get(key)
	fmt.Printf("\n %s has the value of %s", key,result1)
	fmt.Printf("%s%d-%d", "get Key End:", i,pi)

	//result1, _ := mp.Get("key1")
	//fmt.Println("key1 has the value of ", result1)

	//result2, _ := mp.Get("key2")
	//fmt.Println("key2 has the value of ", result2)

	//mp.Clear()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("%d%s%d\n", i, " Duraction : ", elapsed.Seconds())
}

func setKey(client hazelcast.IHazelcastInstance, pi int) {
	start := time.Now()
	mp, _ := client.GetMap("simpleExample")
	for i := 0; i < 10000; i++ {
		mp.SetWithTtl(fmt.Sprintf("%s%d-%d", "key_", i,pi), fmt.Sprintf("%s%d-%d", "value_", i,pi), 10, time.Second)
		go getKey(client,i,pi)
	}
	fmt.Printf("%s%d", "Set Key End:", pi)

	//result1, _ := mp.Get("key1")
	//fmt.Println("key1 has the value of ", result1)

	//result2, _ := mp.Get("key2")
	//fmt.Println("key2 has the value of ", result2)

	//mp.Clear()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("%d%s%d\n", pi, " Duraction : ", elapsed.Seconds())
}

func main() {
	config := hazelcast.NewHazelcastConfig()
	config.GroupConfig().SetName("pttkep")
	config.GroupConfig().SetPassword("pttkep")
	config.ClientNetworkConfig().AddAddress("10.145.172.13:5701")

	client, err := hazelcast.NewHazelcastClientWithConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 100; i++ {
		go setKey(client, i)
	}
	fmt.Print(" End Main ")
	fmt.Scanln()
	//client.Shutdown()
}
