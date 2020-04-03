package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SensorData is
type SensorData struct {
	SensorType string `json:"dataID"`
	DataSize   int    `json:"dataSize"`
}

//AccelerometerSensor is
type AccelerometerSensor struct {
	SensorData
	Data [][]int `json:"data"`
}

//GyroSensor is
type GyroSensor struct {
	SensorData
	Data []int `json:"data"`
}

// ErrorMessage is
type ErrorMessage struct {
	ErrorCode   int    `json:"errorType"`
	Description string `json:"description"`
}

// Error404 is
func Error404() ErrorMessage {
	return ErrorMessage{ErrorCode: 404, Description: "Page Not Found"}
}

// ErrorMessageTemplate is
func ErrorMessageTemplate(code int, message string) ErrorMessage {
	return ErrorMessage{ErrorCode: code, Description: message}
}

// QueryBody is
type QueryBody struct {
	DataSize int `json:"dataSize"`
}

func getAccelerometerSensorData(body QueryBody) AccelerometerSensor {
	var accelerometerSensorData AccelerometerSensor
	accelerometerSensorData.Data = [][]int{{1, 2, 3}, {4, 5, 6}}
	accelerometerSensorData.SensorData = SensorData{SensorType: "AccelerometerSensor"}
	accelerometerSensorData.DataSize = body.DataSize
	return accelerometerSensorData
}

func getGyroSensorData(body QueryBody) GyroSensor {
	var gyroSensorData GyroSensor
	gyroSensorData.SensorData = SensorData{SensorType: "GyroSensor"}
	gyroSensorData.Data = []int{1, 2, 3}
	gyroSensorData.DataSize = body.DataSize
	return gyroSensorData
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var queryBody QueryBody

	if err := json.NewDecoder(r.Body).Decode(&queryBody); err != nil {
		log.Println(err)
	}

	if queryBody.DataSize <= 0 {
		log.Println("Error Bad Request Data Size : ", queryBody)
		json.NewEncoder(w).Encode(ErrorMessageTemplate(400, "Bad Request"))
		return
	}

	params := mux.Vars(r)

	switch params["op"] {
	case "accelerometersensor":
		json.NewEncoder(w).Encode(getAccelerometerSensorData(queryBody))
	case "gyrosensor":
		json.NewEncoder(w).Encode(getGyroSensorData(queryBody))
	default:
		json.NewEncoder(w).Encode(Error404())
		log.Println("default")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sensor/{op}", getData).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
