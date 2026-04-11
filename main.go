package main

import (
    "fmt"
    "github.com/Yandex-Practicum/tracker/daysteps"
    "github.com/Yandex-Practicum/tracker/spentcalories"
)

func main() {
	fmt.Println(daysteps.DayActionInfo("678,0h50m", 75, 1.75))

	info, err := spentcalories.TrainingInfo("3456,Бег,1h00m", 75, 1.75)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info)
}
