package main
import (
	"fmt"
	"log"
    "github.com/gomodule/redigo/redis"
)




func main(){

	type Mahasiswa struct {
		Nama     string  `redis:"nama"`
		NIM      string  `redis:"nim"`
		IPK      float64 `redis:"ipk"`
		Semester int     `redis:"semester"`
	}



	conn, err := redis.Dial("tcp", "localhost:6380")
	if err != nil {
		log.Panic(err)
	}

	_, err = conn.Do("HSET", "mahasiswa:1", "nama", "Redha Juanda", "nim", "12345", "ipk", 3.34, "semester", 4)
	if err != nil {
		log.Panic(err)
	}

	// mengambil data dengan tipe data string 
	nama, err := redis.String(conn.Do("HGET", "mahasiswa:1", "nama"))
	if err != nil {
		log.Panic(err)
	}

	// mengambil data dengan tipe data string
	nim, err := redis.String(conn.Do("HGET", "mahasiswa:1", "nim"))
	if err != nil {
		log.Panic(err)
	}

	// mengambil data dengan tipe integer
	ipk, err := redis.Float64(conn.Do("HGET", "mahasiswa:1", "ipk"))
	if err != nil {
		log.Panic(err)
	}

	// mengambil data dengan tipe integer 
	semester, err := redis.Int(conn.Do("HGET", "mahasiswa:1", "semester"))
	if err != nil {
		log.Panic(err)
	}

	// Mengambil semua data 
	rep, err := redis.StringMap(conn.Do("HGETALL", "mahasiswa:1"))
	if err != nil {
		log.Panic(err)
	}

	rep1, err := redis.Values(conn.Do("HGETALL", "mahasiswa:1"))
	if err != nil {
		log.Panic(err)
	}
	var mahasiswa Mahasiswa
	// ScanStruct akan mempopulate semua data reply ke struct
	// mahasiswa, ScanStruct juga otomatis mencasting tipe data
	// sesuai dengan yang didefinisikan di sruct
	err = redis.ScanStruct(rep1, &mahasiswa)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%+v", mahasiswa)

	fmt.Println("1", nama, nim, ipk, semester)
	fmt.Println("2", rep)

}