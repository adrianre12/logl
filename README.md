# logl

Package logl (loglevel) implements logging at levels from NONE to DEBUG, The output is filtered by a globally set logging level.
The objective was to write a logging package that is small and simple. The alternatives I looked at were either larger and more complex than my applications, and/or did not do what I wanted.
To do this logl wraps log and calls the relevant functions in log to write the messages. It uses the writer instance in log so it is possible to also use log if required.

The output looks like:
```
2020/12/02 09:35:11 logl_test.go:40: [ERR] Line 1
2020/12/02 09:35:11 logl_test.go:41: [WRN] Line 2
2020/12/02 09:35:11 logl_test.go:42: [INF] Line 3
2020/12/02 09:35:11 logl_test.go:43: [DBG] Line 4
2020/12/02 09:35:11 logl_test.go:44: [ERR] Line 5
2020/12/02 09:35:11 logl_test.go:45: [WRN] Line 6
2020/12/02 09:35:11 logl_test.go:46: [INF] Line 7
2020/12/02 09:35:11 logl_test.go:47: [DBG] Line 8
```

## ToDo
~~* Look at switching to log's instance of log.Logger.  So logl and log would write to the same output.~~
* Log rotation on size and day with ageing off.
* Multiple outputs e.g. log and file (see io.MultiWriter)
* Look at changing to using multiple instances of log each logging with a diffrent prefix to the same writer. Uisng a dummy log instance to disable writing for levels. I think a lot of coding for little gain.


