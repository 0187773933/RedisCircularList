# Redis Circular List

```
package main

import (
	"fmt"
	redis_lib "github.com/go-redis/redis/v8"
	circular_list "github.com/0187773933/RedisCircularList/list"
)

func main() {

	redis_connection := redis_lib.NewClient( &redis_lib.Options{
		Addr: "localhost:6379" ,
		DB: 3 ,
		Password: "" ,
	})

	fmt.Printf( "Current = %s\n"  , circular_list.Current( redis_connection , "test" ) )
	
	fmt.Printf( "Previous = %s\n"  , circular_list.Previous( redis_connection , "test" ) )
	fmt.Printf( "Previous = %s\n"  , circular_list.Previous( redis_connection , "test" ) )
	fmt.Printf( "Previous = %s\n"  , circular_list.Previous( redis_connection , "test" ) )
	
	fmt.Printf( "Next = %s\n"  , circular_list.Next( redis_connection , "test" ) )
	fmt.Printf( "Next = %s\n"  , circular_list.Next( redis_connection , "test" ) )
	fmt.Printf( "Next = %s\n"  , circular_list.Next( redis_connection , "test" ) )

	fmt.Printf( "Current = %s\n"  , circular_list.Current( redis_connection , "test" ) )

}
```