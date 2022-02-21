# elon-musk-tweets

Twitter Data Visualization Web Application

## Brief Summary

This Golang application downloads a json file and inserts the entries in MongoDB. 
Several endpoints exist that provide the data from different query parameters.
Then a React application consumes this data and visualizes it in different charts using the Highcharts library.

## Commands

To start the backend use "go run ./main.go -importdb". ('importdb' flag is used to populate the database)
If the database is already populated use simply "go run ./main.go".

To start the frontend use "npm start" in 'elon-client' directory.

The frontend is available at 'localhost:3000' and the backend on 'localhost:9090.
