/*
 * Created by goland
 * User: Hasan UCAK <hasan.ucak@gmail.com>
 * Date: 2/21/2018
 * Time: 3:00 PM

load test hazelcast for cache mechanism.

*/


package main

import (
	"github.com/hazelcast/hazelcast-go-client"
	"fmt"
	"time"
)


//set hazelcast username
const username string = "pttkep"
//set hazelcast password of username
const password string = "pttkep"
//set hazelcast host ip and port number
const hazelcast_url string = "10.145.172.13:5701"
//thread number meanly Goroutines
const totalRoutines int = 1000
//number of insert key per thread
const maxKeyNumber int = 10000
//ttl value for key .key/value life time in second
const keyLifeTime int = 10
//map name for test
const mapName string = "testMap"



/*
 * get value of key from hazelcast with Goroutines
 */
func GetKey(client hazelcast.IHazelcastInstance, index int,routineIndex int) {
	startTime := time.Now()
	mp, _ := client.GetMap(mapName)
	key := fmt.Sprintf("%s%d-%d", "key_", routineIndex,index)
	result1, _ := mp.Get(key)
	fmt.Printf("\n %s has the value of %s", key,result1)
	fmt.Printf("--%s%d-%d", "get Key End:", routineIndex,index)
	//calculate duraction time
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%d%s%d\n", i, " Duraction : ", elapsed.Seconds())
}

/*
 * set value of key from hazelcast with Goroutines
 */
func SetKey(client hazelcast.IHazelcastInstance, routineIndex int) {
	startTime := time.Now()
	mp, _ := client.GetMap(mapName)
	for index := 0; index < maxKeyNumber; index++ {
		mp.SetWithTtl(fmt.Sprintf("%s%d-%d", "key_", routineIndex,index), fmt.Sprintf("%s%d-%d", "value_", routineIndex,index), keyLifeTime, time.Second)
		go GetKey(client,index,routineIndex)
	}
	fmt.Printf("%s%d", "Set Key End:", routineIndex)
	//calculate duraction time
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%d%s%d\n", routineIndex, " Duraction : ", elapsed.Seconds())
}

func main() {
	config := hazelcast.NewHazelcastConfig()
	//set hazelcast config
	config.GroupConfig().SetName(username)
	config.GroupConfig().SetPassword(password)
	config.ClientNetworkConfig().AddAddress(hazelcast_url)
	//starting test
	client, err := hazelcast.NewHazelcastClientWithConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	//tread number
	for routineIndex := 0; routineIndex < totalRoutines; routineIndex++ {
		go SetKey(client, routineIndex)
	}
	fmt.Print(" End Main ")
	//we have to wait for the process to end
	fmt.Scanln()
	client.Shutdown()
}
