package external

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/lacion/mygolangproject/models"
	"github.com/lacion/mygolangproject/service"
	"go.uber.org/fx"
)

func worker(id int, jobs <-chan models.User, wg *sync.WaitGroup, service *service.Service) {
	defer wg.Done()
	for user := range jobs {
		err := service.UserReposiotry.CreateUser(&user)
		if err != nil {
			fmt.Printf("Worker %d failed to create user %s: %s\n", id, user.ID, err)
		} else {
			fmt.Printf("Worker %d created user %s\n", id, user.ID)
		}
	}
}

func InvokeSeed(lc fx.Lifecycle, service *service.Service) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			count, err := service.UserReposiotry.GetCount()
			if err != nil {
				return err
			}
			if count > 0 {
				fmt.Println("Users already seeded and we have ", count, " users")
				return nil
			}			

			file, err := os.Open("users_data.json")
			if err != nil {
				return err
			}
			defer file.Close()

			fmt.Print(file)

			jobs := make(chan models.User, 100)

			var wg sync.WaitGroup
			numOfWorkers := 10
			for i := 0; i < numOfWorkers; i++ {
				wg.Add(1)
				go worker(i, jobs, &wg, service)
			}

			decoder := json.NewDecoder(file)
			_, err = decoder.Token()
			if err != nil {
				return err
			}

			for decoder.More() {
				var user models.User
				err := decoder.Decode(&user)
				if err != nil {
					return err
				}
				jobs <- user
			}

			close(jobs)

			wg.Wait()

			return nil
		},
	})
}
