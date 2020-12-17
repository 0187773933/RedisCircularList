package main

import (
	"fmt"
	redis_lib "github.com/go-redis/redis/v8"
	circular_list "circularlist/list"
)

// https://github.com/andymccurdy/redis-py/blob/1f857f0053606c23cb3f1abd794e3efbf6981e09/tests/test_commands.py
// https://github.com/ceberous/redis-manager-utils/blob/master/BaseClass.js
// https://github.com/48723247842/RedisCirclularList/blob/master/redis_circular_list/__init__.py
// https://redis.io/commands/sadd
// https://pkg.go.dev/builtin#error
// https://pkg.go.dev/github.com/go-redis/redis/v8#BoolCmd.Err

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