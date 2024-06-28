package luhn

import (
	"errors"
	"strconv"
)

func CheckNumber(cardNumber int) (bool, error) {
	var sum int = 0
	var isDouble bool = false
	payload := cardNumber / 10
	for i := len(strconv.Itoa(payload)) - 1; i >= 0; i-- {
		number, err := strconv.Atoi(string(strconv.Itoa(payload)[i]))
		if err != nil {
			return false, errors.New("error happened in strconv.atoi")
		}
		if isDouble {
			sum += number
		} else {

			if number*2 > 9 {
				sum += number*2/10 + number*2%10
			} else {
				sum += number * 2
			}
		}
		isDouble = !isDouble
	}
	if 10-sum%10 == cardNumber%10 {
		return true, nil
	}
	return false, nil
}
