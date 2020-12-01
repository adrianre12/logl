# logl

Package logl (loglevel) implements logging at levels from NONE to DEBUG, The output is filtered by a globally set logging level.
The objective was to write a logging package that is small and simple. The alternatives I looked at were either larger and more complex than my applications, and/or did not do what I wanted.
To do this logl wraps an instance of log.Logger and calls the relevant functions in log to write the messages. I decided to not use the writer instance in log so that it is still possible to use log if required,  but I am not sure about this.


## ToDo
* Look at switching to log's instance of log.Logger.  So logl and log would write to the same output.
* Log rotation on size and day with ageing off.
* Multiple outputs e.g. log and file (see io.MultiWriter)


