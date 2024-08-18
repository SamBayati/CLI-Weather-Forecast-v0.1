
![image](https://github.com/user-attachments/assets/0f3560df-f033-4f68-9612-e2c72d41eb71)




## Dependencies 
### 1. Install Dependencies for adjusting the colours from the  ** color ** Repo on GitHub

ref : https://github.com/fatih/color.git 

```bash
go get github.com/fatih/color
```
### 2. Head to https://www.weatherapi.com/docs/#intro-location and get your API key
Base URL: http://api.weatherapi.com/v1  +  	/forecast.json or /forecast.xml + your api key + q = city's name


## Build

``` 
go build
```

### Default running

```Run the cli app
go run ./main.go
```



### Run with an specific city
// run this command if you like to look up an specific city other than the default one
```
go run ./main.go "name of the city"
go run ./main new york

```

### Adding the Cli application to your 

```
// mv Breadcrumbs CLI-Weather-Forecast-v0.1 /usr/local/bin 
```

### Run from the terminal 

```
CLI-Weather-Forecast-v0.1

#### or

CLI-Weather-Forecast-v0.1 new york 
```
