package main

import (
	"github.com/gin-gonic/gin"
	"errors"
	"net/http"
	"strconv"
)

type lawnmower struct {
	ID			int `json:"id"`
	Header 		string `json:"header"`
	Blade 		string `json:"blade"`
	Battery 	string `json:"battery"`
	Text		string `json:"text"`
	Price 		int `json:"price"`
	Img			string `json:"img"`
}

var lawnmowers = []lawnmower{
  {
      ID:1,
      Header: "Lawnmower V1",
      Blade: "nylon",
      Battery: "10000",
      Text: "MTL Lawnmower V1",
      Price: 700,
      Img: "https://s.wsj.net/public/resources/images/BN-NN210_mowweb_J_20160412151533.jpg",
  },
  {
      ID:2,
      Header: "Lawnmower V2",
      Blade: "metal",
      Battery: "20000",
      Text: "MTL Lawnmower V2",
      Price: 1400,
      Img: "https://images.immediate.co.uk/production/volatile/sites/18/2022/01/Mower-Driving-2-7ab851b.jpg?quality=90&webp=true&fit=975,649",
  },

}

func getLawnmowers(context *gin.Context){
	context.IndentedJSON(http.StatusOK, lawnmowers)
}

func addLawnmower(context *gin.Context){
	var newLawnmower lawnmower

	if err := context.BindJSON(&newLawnmower); err !=nil{
		return
	}

	lawnmowers = append(lawnmowers, newLawnmower)

	context.IndentedJSON(http.StatusCreated, newLawnmower)
}

func getLawnmowerById(id int) (*lawnmower, error){
	for i, t:= range lawnmowers{
		if t.ID == id {
			return &lawnmowers[i], nil
		}
	}

	return nil, errors.New("lawnmower not found")
}

func getLawnmower(context *gin.Context){
	id, err := strconv.Atoi(context.Param("id"))
	lawnmower, err := getLawnmowerById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"To do not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, lawnmower)
}

// func toggleLawnmowerStatus(context *gin.Context){
// 	id := context.Param("id")
// 	todo, err := getTodoById(id)

// 	if err != nil{
// 		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"To do not found"})
// 		return
// 	}

// 	todo.Completed = !todo.Completed

// 	context.IndentedJSON(http.StatusOK, todo)
// }
func filteredLawnmowers(blade string, battery string , c *gin.Context){
	var newLawnmowers []lawnmower;
	for i, t:= range lawnmowers{
		if t.Blade == blade && t.Battery == battery{
			newLawnmowers = append(newLawnmowers, lawnmowers[i])
		}
	}
	c.IndentedJSON(http.StatusOK, newLawnmowers)
}
func filteredLawnmowersByBlade(blade string, c *gin.Context){
	var newLawnmowers []lawnmower;
	for i, t:= range lawnmowers{
		if t.Blade == blade {
			newLawnmowers = append(newLawnmowers, lawnmowers[i])
		}
	}
	c.IndentedJSON(http.StatusOK, newLawnmowers)
}
func filteredLawnmowersByBattery(battery string , c *gin.Context){
	var newLawnmowers []lawnmower;
	for i, t:= range lawnmowers{
		if t.Battery == battery{
			newLawnmowers = append(newLawnmowers, lawnmowers[i])
		}
	}
	c.IndentedJSON(http.StatusOK, newLawnmowers)
}
func main() {
	router := gin.Default()
	router.GET("/lawnmowers", getLawnmowers)
	router.GET("/lawnmowers/filters", func(c *gin.Context){
		blade := c.Query("blade")
		battery := c.Query("battery")
		if( blade == "None" && battery == "None"){
			getLawnmowers(c)
		} else if blade == "None"{
			filteredLawnmowersByBattery(battery, c)
		}else if battery == "None"{
			filteredLawnmowersByBlade(blade, c)
		}else {
			filteredLawnmowers(blade, battery, c)
		}
		
		
	})
	router.GET("/lawnmowers/:id", getLawnmower)
	
	//router.PATCH("/lawnmowers/:id", toggleToDoStatus)
	router.POST("/lawnmowers", addLawnmower)
	router.Run("localhost:8090")
}