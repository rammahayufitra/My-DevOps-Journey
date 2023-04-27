package main
import (
	"fmt"
	"log"
    "github.com/gomodule/redigo/redis"
)


func main(){
	pool := redis.NewPool(
		func() (redis.Conn, error){
			return redis.Dial("tcp", "localhost:6380")
		}, 
		10,
	)
	pool.MaxActive = 10
	// mengambil satu koneksi dari pool 
	conn := pool.Get() 
	defer conn.Close() 

	rep, err := conn.Do("SET", "test", "oke")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(rep)
}