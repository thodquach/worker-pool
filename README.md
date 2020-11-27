# worker-pool
  Worker pool in Go programing language

# Reference:
https://brandur.org/go-worker-pool
You need to append all to-do tasks into a slice and then, you can start worker-pool to do all tasks.

# Improve:
  I split the task one by one, so you can use go routine to add task and run task
  The pool.Stop() function will wait all task to be done and close the channel `*Tasks`

  This is just a prototype and very simple worker-pool written in go lang.

```console
─ go run main.go
The nth: 0 element task run in workerpool.
The nth: 1 element task run in workerpool.
The nth: 2 element task run in workerpool.
The nth: 3 element task run in workerpool.
The nth: 5 element task run in workerpool.
The nth: 4 element task run in workerpool.
The nth: 6 element task run in workerpool.
The nth: 8 element task run in workerpool.
The nth: 7 element task run in workerpool.
The nth: 9 element task run in workerpool.
```
