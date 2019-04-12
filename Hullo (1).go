package hulloratelimiter

import (
        "fmt"
        "net/http"
        "os"

        "github.com/gomodule/redigo/redis"

        
)
//establishing connection with redis
var redisPool *redis.Pool

func hulloratelimiter(limit int,ctr int) bool{
        redisAddr := os.Getenv("REDIS_ADDR")
        redisPassword := os.Getenv("REDIS_PASSWORD")

        redisPool = &redis.Pool{
                Dial: func() (redis.Conn, error) {
                        conn, err := redis.Dial("tcp", redisAddr)
                        if redisPassword == "" {
                                return conn, err
                        }
                        if err != nil {
                                return nil, err
                        }
                        if _, err := conn.Do("AUTH", redisPassword); err != nil {
                                conn.Close()
                                return nil, err
                        }
                        return conn, nil
                },
        }
	//cts ->current time stamp ,ts-> timestamp 
            value := []byte
			err := pool.Do(radix.Cmd(&value, "GET", redisAddr))
			currentTime := time.Now()
			cts := currentTime.Second() + (currentTime.Minute()*60) + (currentTime.Hour()*3600)
			l:=len(value)
			ts,_ := strconv.Atoi(string(value[:l-1]))
			cntr,_:= strconv.Atoi(string(value[:l-1]))
			if cts-ts > limit || value == " "{
				cntr = ctr
				newvalue := []byte(strconv.Itoa(cts))
				newvalue = append(newvalue,strconv.Itoa(cntr)...)
				err1 := pool.Do(radix.Cmd(&newvalue, "HSET ",redisAddr))
				
			} else if cts-ts <limit {
				if cntr != 0 {
					cntr--
					copy(value[l-1:],strconv.Itoa(cntr))
				err1 := pool.Do(radix.Cmd(&value, "HSET ",redisAddr))
				} else if cntr == 0 {
					return false				
				}
			} 
			return true
		}

