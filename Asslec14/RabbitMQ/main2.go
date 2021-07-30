package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"rmq/dataHandling"
	"rmq/fileHandling"
	"rmq/rabbitmq"
	"rmq/producer"
	"rmq/consumer"
)

func main() {
	//uri = "amqp://acc:password@URL"
	// uri := "amqp://aa:aa@http://127.0.0.1:15672/#/"
	uri := "amqp://tfs:tfs-ocg@174.138.40.239:5672/#/"

	rmq := rabbitmq.CreateNewRMQ(uri)

	//config exchange
	exchangeName := "Exc"
	exchangeType := "direct"

	//config queue
	queueName1 := "Queue1"
	queueName2 := "Queue2"

	//config bindingKey
	bindingKey1 := "b1"
	bindingKey2 := "b2"

	//config routingKey 
	routingKey1 := bindingKey1
	routingKey2 := bindingKey2

	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	// create 1 channel for producer 1 & 2
	pCh1, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel: ", err)
		return
	}
	// create 1 channel for producer2
	// pCh2, err := rmq.GetChannel()
	// if err != nil {
	// 	fmt.Println("Cannot get channel: ", err)
	// 	return
	// }
	// create 1 channel for consumer1
	cCh1, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel: ", err)
		return
	}
	// create 1 channel for consumer2
	cCh2, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel: ", err)
		return
	}
	// create Consumer 
	consumer1 := consumer.CreateNewConsumer(exchangeName, exchangeType, bindingKey1, queueName1, cCh1, ctx, &wg)
	consumer2 := consumer.CreateNewConsumer(exchangeName, exchangeType, bindingKey2, queueName2, cCh2, ctx, &wg)


	// create Producer
	producer1 := producer.CreateNewProducer(exchangeName, exchangeType, routingKey1, pCh1, ctx, &wg)
	producer2 := producer.CreateNewProducer(exchangeName, exchangeType, routingKey2, pCh1, ctx, &wg)

	
	/////////////////////////////////////////////////
	// graceful shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c // waits for the termination signal
		fmt.Printf("Got %s signal. Exiting...\n", sig)
		producer1.Close() // stop scheduler at the end
		consumer1.Close()
		producer2.Close() // stop scheduler at the end
		consumer2.Close()
		cancelFunc()
	}()
	/////////////////////////////////////////////////

	wg.Add(4)

	//read file and send1
	var lineSend1 = make(chan string)
	go fileHandling.ReadFileLineByLine("./test.txt", lineSend1)
	go producer1.Send(lineSend1)

	//receive1, handle1 and send2
	var lineReceive1 = make(chan string)
	var mapStringSend2 = make(chan string)
	go consumer1.StartReceiveData(lineReceive1)
	go dataHandling.SplitLineToMap(lineReceive1, mapStringSend2)
	go producer2.Send(mapStringSend2)

	//receive2, handle2 and print
	var ResultMap = make(map[string]int)
	var mapReceiveString = make(chan string)
	go consumer2.StartReceiveData(mapReceiveString)
	go dataHandling.UpdateMapResult(&ResultMap, mapReceiveString)
	go func() {
		for {
			fmt.Printf("ResultMap: %v \n", ResultMap)
			time.Sleep(time.Second * 5)
		}
	}()

	//stop
	time.Sleep(time.Second * 20)
	fmt.Println("calling cancelFunc")
	cancelFunc()

	wg.Wait()

	

}