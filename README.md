# Echo-Template_0.1

![image](https://github.com/Abhimanyu-Barge/Echo-Template_0.1/assets/67216291/0434af7a-4530-41ec-ac96-2a8103ae2302)

High performance, extensible, minimalist Go web framework

Installation Requirements

To install Echo Go 1.13 or higher is required. Go 1.12 has limited support and some middlewares will not be available. Make sure your project folder is outside your $GOPATH.

- go get github.com/labstack/echo/v4
- go mod init server
- go mod tidy
- 
If you are working with Go v1.14 or earlier use:

$ GO111MODULE=on go get github.com/labstack/echo/v4

# Auto Reload Requirement 
    1)noode js
    2)npm
    3)nodemon 
      - npm -g i nodemon
        1. Open Windows PowerShell with Run as Administrator
        2. Run this command: Set-ExecutionPolicy Unrestricted
# Run server 
 - nodemon --exec go run server.go --signal SIGTERM

