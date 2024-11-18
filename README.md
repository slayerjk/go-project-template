# go-project-template
Just a template for new Go project

<h2>Flags</h2>

* log-dir - path to logs dir; default is relative to exe - 'logs_http-param-to-db'
* keep-logs - number of logs to keep after rotation; default = 7

<h2>Helper packages</h2>

<h3>internal/logging</h3>

Logging to file using builtin "log" package and based on logdir, appname and date.

<h3>internal/mailing</h3>

Email using builtin "net/smtp" based on appname, message, date and mailing.json(found in "data" dir of program root) file with smtp data.

<h3>internal/vafswork</h3>

* Get programm's executable path.
* Rotate files in dir using va builtin packages based on dirname, number of files to keep(most recent).

<h3>internal/vawebwork</h3>

Create insecure http client