package go_redis

import (
	"fmt"
)

func GetLock(user string) {
	lock := "mutex"
	res := client.SetNX(ctx, lock, user, 0)
	if res.Val() == false {
		owner := client.Get(ctx, lock)
		fmt.Printf("user %s get lock fail, lock owner is %s \n", user, owner.Val())
	} else {
		fmt.Printf("i get a lock, i am %s \n", user)
	}
}
