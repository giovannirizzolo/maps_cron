# Maps cronjob 🕐
This is a cronjob written in <img width="32" height="32" src="https://github.com/giovannirizzolo/maps_cron/assets/34986490/d050618f-c094-4a03-9547-22ca0f05dc80" alt="Golang logo"/> to periodically fetch traffic data from my home to office.

The cronjob is scheduled to be running:
- Every 2 minutes
- from 6 a.m. to 11 a.m. (included)
- On working weekdays (Monday-Friday)
- from March to June

The schedule is set in ```main.go``` using:

```go
  c.AddFunc("*/2 6-10 * 3-6 1-5", func() {
  ...
  }
```

## Why ❓

I was bored of having to check every morning the 🏠home-to-office🏢 ETA cause of my busy morning routing, so I needed a tool to perform it for me.

## How to run
The recommended option is to run the application using Docker.

## Setup ⚙️

To run the application, you must provide a ```.env``` file with the same structure provided in ```.env.example```.
```MAPS_API_KEY```  is an API key that you can get by using the free tier of [MapBox Navigation service](https://www.mapbox.com/navigation).
Setting ```ROUTING_PROFILE='driving-traffic'``` will make the application to fetch data including traffic jams delays


## Docker <img width="24" height="24" src="https://github.com/giovannirizzolo/maps_cron/assets/34986490/e2d755d8-3a49-4022-b798-bfc4b44b0853" alt="Docker logo"/>
cd into the root of the project, run ```docker compose up``` and wait for the application to launch


## Output

The cron will update according to its schedule a ✍️ ```report.txt``` file including info about the ETA and corresponding timestamp
