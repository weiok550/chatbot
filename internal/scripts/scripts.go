package scripts

import "time"

func RegisterCronJob() {

	go func() {
		for {
			
			time.Sleep(2 * time.Second)
		}
	}()

}
