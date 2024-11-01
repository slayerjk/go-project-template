# go-project-template
Just a template for new Go project

<h2>Helper packages</h2>

<h3>internal/logging</h3>

Logging to file using builtin "log" package and based on logdir, appname and date.

<h3>internal/mailing</h3>

Email using builtin "net/smtp" based on appname, message, date and mailing.json(found in "data" dir of program root) file with smtp data.

<h3>internal/vafswork</h3>

* Get programm's executable path.
* Rotate files in dir using va builtin packages based on dirname, number of files to keep(most recent).
