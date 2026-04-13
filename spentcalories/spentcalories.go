package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const stepLengthCoefficient = 0.65
const mInKm = 1000.0
const minInH = 60.0
const walkingCaloriesCoefficient = 0.5

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, "", 0, fmt.Errorf("invalid steps")
	}

	activity := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, err
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)
	hours := duration.Hours()

	return dist / hours
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("invalid input data")
	}

	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	calories := (weight * speed * minutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("invalid input data")
	}

	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	calories := (weight * speed * minutes) / minInH
	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	var calories float64

	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity,
		duration.Hours(),
		dist,
		speed,
		calories,
	), nil
}
