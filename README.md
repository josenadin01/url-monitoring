# URL Monitoring
My first time using Golang on a personal project

You can run this project by writing: go run url-monitoring.go on your terminal inside the URL-MONITORING folder

After you have been introduced a menu to select which option you desire.

1 - Start Monitoring 
Will access all the urls from the urls.txt and go one by one checking
if the StatusCode is equal to 200 if it is then it will return a successful message
if it is not 200 then it will return the StatusCode it got from the http.Get(url)
It will also create a log.txt for the option number 2

2 - Show Logs 
Will read the log.txt file that is created after running option 1 for the first time, showing if the urls were online: true or false

0 - Quit
Will exit the program with a zero code that your OS will recognize as an safe exit.

Any other input will result in the program to close with a non zero code, your OS will recognize as an error.
