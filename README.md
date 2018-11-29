# rabbitmqmonitorgo

### Simple way to get queues info which are present rabbitmq

Replace the values of the url , username,password 
```$xslt
rmqc, _ := rabbithole.NewClient("URL", username, password)
```

```$xslt
rabbitmqmonitorgo:$ go get
rabbitmqmonitorgo:$ go run monitor.go
```

### Running Locally
* go to the browser
    * http://localhost:8081/details/{queue_name}
    

#### Reference

https://github.com/michaelklishin/rabbit-hole
