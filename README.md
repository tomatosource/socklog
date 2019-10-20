# socklog

## Usage

```
	socklogger := NewLogger()
	defer socklogger.Close()
	log.SetOutput(socklogger)
```

## TODOs

 - take server addr as arg
 - return error when creating
 - maybe return close function?
