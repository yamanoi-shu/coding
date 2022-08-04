import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func step4(lines []string) error {
	var menuNum int

	n, err := fmt.Sscanf(lines[0], "%d", &menuNum)

	menuMap := make(map[int]*lib.MenuInfo)

	if n < 1 || err != nil {
		return errors.New("Bad input")
	}

	cookingMap := make(map[int]*lib.Queue)
	orderMap := make(map[int]int)

	lines = lines[1:]

	var lineCnt int

	for lineCnt = 0; lineCnt < menuNum; lineCnt++ {
		var id, stockNum, price int

		n, err := fmt.Sscanf(lines[lineCnt], "%d %d %d", &id, &stockNum, &price)

		if n < 3 || err != nil {
			return errors.New("Bad input")
		}

		menu := &lib.MenuInfo{
			ID:       id,
			StockNum: stockNum,
			Price:    price,
		}
		menuMap[id] = menu
	}
	lines = lines[lineCnt:]

	for _, line := range lines {
		lineElem := strings.Split(line, " ")

		switch lineElem[0] {
		case "received":
			seatNum, _ := strconv.Atoi(lineElem[2])

			menuId, _ := strconv.Atoi(lineElem[3])

			order := lib.Order{
				SeatNum: seatNum,
				MenuId:  menuId,
			}
			if cookingMap[menuId] == nil {
				cookingMap[menuId] = lib.NewQueue(menuNum)
			}

			orderMap[seatNum]++

		case "complete":
			completeMenuId, err := strconv.Atoi(lineElem[1])
			if err != nil {
				return errors.New("Bad input")
			}

			completeOrder := cookingMap[completeMenuId].Dequeue()

			orderMap[seatNum]--

			fmt.Printf("ready %d %d\n", completeOrder.SeatNum, completeOrder.MenuId)

		case "check":
			seatNum, err := strconv.Atoi(lineElem[2])

			if err != nil {
				return err
			}

			if orderMap[seatNum] > 0 {
				fmt.Println("wait ready")
			} else {
				fmt.Println("check")
			}
		}

	}
	return nil
}
