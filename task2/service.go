package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strconv"
	"time"
)

const BOOKING_API = "https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList"
const MASTER_PRICELIST_API = "https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi"

type MasterJenisPayload struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	MaxPrice  uint64    `json:"maxPrice"`
}

type BookingPayload struct {
	Id               string              `json:"id"`
	RoomName         string              `json:"roomName"`
	Participants     uint                `json:"participants"`
	ListConsumptions []map[string]string `json:"listConsumption"`
	BookingDate      time.Time           `json:"bookingDate"`
	StartTime        time.Time           `json:"startTime"`
	EndTime          time.Time           `json:"endTime"`
	OfficeName       string              `json:"officeName"`
}

type ResponsePayload struct {
	Payload interface{} `json:"payload"`
}

type OfficePayload struct {
	Name       string             `json:"name"`
	RoomUsages []RoomUsagePayload `json:"roomUsages"`
}

type RoomUsagePayload struct {
	Name            string          `json:"name"`
	UsagePercentage float64         `json:"usagePercentage"`
	FoodExpense     uint64          `json:"foodExpense"`
	Consumptions    map[string]uint `json:"consumptions"`
}

func proceedSummary(c *fiber.Ctx, month int, year int) error {

	//agent := fiber.Get(MASTER_PRICELIST_API)
	//statusCode, masterResponsePayload, errs := agent.Bytes()
	//
	//if errs != nil {
	//	log.Error("Error for with request", string(c.Request().RequestURI()), errs)
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "we're sorry, something went wrong"})
	//}
	//
	//if statusCode != 200 {
	//	responseBody := string(masterResponsePayload)
	//	log.Error("Error for with request", string(c.Request().RequestURI()), " Getting code, not 200", responseBody)
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "we're sorry, something went wrong"})
	//}
	//
	//agent = fiber.Get(BOOKING_API)
	//statusCode, bookingResponsePayload, errs := agent.Bytes()
	//
	//if errs != nil {
	//	log.Error("Error for with request", string(c.Request().RequestURI()), errs)
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "we're sorry, something went wrong"})
	//}
	//
	//if statusCode != 200 {
	//	responseBody := string(bookingResponsePayload)
	//	log.Error("Error for with request", string(c.Request().RequestURI()), " Getting code, not 200", responseBody)
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "we're sorry, something went wrong"})
	//}

	bookingResponsePayload := []byte("[{\"bookingDate\":\"2024-01-03T07:25:52.737Z\",\"officeName\":\"UID JAYA\",\"startTime\":\"2024-01-04T09:00:00.000Z\",\"endTime\":\"2024-01-04T16:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"},{\"name\":\"Snack Sore\"}],\"participants\":62,\"roomName\":\"Ruang Borobudur\",\"id\":\"1\"},{\"bookingDate\":\"2024-01-03T08:25:52.737Z\",\"officeName\":\"UID JAYA\",\"startTime\":\"2024-01-04T09:00:00.000Z\",\"endTime\":\"2024-01-04T14:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"}],\"participants\":42,\"roomName\":\"Ruang Prambanan\",\"id\":\"2\"},{\"bookingDate\":\"2024-01-03T08:25:52.737Z\",\"officeName\":\"UID JAYA\",\"startTime\":\"2024-01-04T09:00:00.000Z\",\"endTime\":\"2024-01-04T11:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"}],\"participants\":15,\"roomName\":\"Ruang Mendhut\",\"id\":\"3\"},{\"bookingDate\":\"2024-01-03T08:25:52.737Z\",\"officeName\":\"UID JAYA\",\"startTime\":\"2024-01-04T13:00:00.000Z\",\"endTime\":\"2024-01-04T16:00:00.000Z\",\"listConsumption\":[{\"name\":\"Makan Siang\"},{\"name\":\"Snack Sore\"}],\"participants\":27,\"roomName\":\"Ruang Mendhut\",\"id\":\"4\"},{\"bookingDate\":\"2024-01-06T08:25:52.737Z\",\"officeName\":\"UID JABAR\",\"startTime\":\"2024-01-08T13:00:00.000Z\",\"endTime\":\"2024-01-08T16:00:00.000Z\",\"listConsumption\":[{\"name\":\"Makan Siang\"},{\"name\":\"Snack Sore\"}],\"participants\":27,\"roomName\":\"Ruang Rinjani\",\"id\":\"5\"},{\"bookingDate\":\"2024-01-05T08:25:52.737Z\",\"officeName\":\"UID BALI\",\"startTime\":\"2024-01-07T09:00:00.000Z\",\"endTime\":\"2024-01-07T11:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"}],\"participants\":27,\"roomName\":\"Ruang Ganesha\",\"id\":\"6\"},{\"bookingDate\":\"2024-01-08T08:25:52.737Z\",\"officeName\":\"UID JAYA\",\"startTime\":\"2024-01-10T09:00:00.000Z\",\"endTime\":\"2024-01-10T16:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"},{\"name\":\"Snack Sore\"}],\"participants\":30,\"roomName\":\"Ruang Prambanan\",\"id\":\"7\"},{\"bookingDate\":\"2024-01-01T08:00:52.737Z\",\"officeName\":\"UID JATENG & DIY\",\"startTime\":\"2024-01-03T09:00:00.000Z\",\"endTime\":\"2024-01-03T14:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"}],\"participants\":35,\"roomName\":\"Ruang Arjuna\",\"id\":\"8\"},{\"bookingDate\":\"2024-01-05T09:25:52.737Z\",\"officeName\":\"UID JAYA\",\"startTime\":\"2024-01-06T14:00:00.000Z\",\"endTime\":\"2024-01-06T16:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Sore\"}],\"participants\":40,\"roomName\":\"Ruang Borobudur\",\"id\":\"9\"},{\"bookingDate\":\"2024-01-15T09:25:52.737Z\",\"officeName\":\"UID JAYA\",\"startTime\":\"2024-01-17T08:00:00.000Z\",\"endTime\":\"2024-01-17T11:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"}],\"participants\":40,\"roomName\":\"Ruang Borobudur\",\"id\":\"10\"},{\"bookingDate\":\"2024-01-02T08:30:52.737Z\",\"officeName\":\"UID JATIM\",\"startTime\":\"2024-01-05T08:00:00.000Z\",\"endTime\":\"2024-01-05T11:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"}],\"participants\":23,\"roomName\":\"Ruang Sudirman\",\"id\":\"11\"},{\"bookingDate\":\"2024-01-03T09:30:52.737Z\",\"officeName\":\"UID JATIM\",\"startTime\":\"2024-01-05T08:00:00.000Z\",\"endTime\":\"2024-01-05T11:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"}],\"participants\":30,\"roomName\":\"Ruang Fatmawati\",\"id\":\"12\"},{\"bookingDate\":\"2024-01-11T09:30:52.737Z\",\"officeName\":\"UID JATIM\",\"startTime\":\"2024-01-13T09:00:00.000Z\",\"endTime\":\"2024-01-13T16:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"},{\"name\":\"Snack Sore\"}],\"participants\":25,\"roomName\":\"Ruang Fatmawati\",\"id\":\"13\"},{\"bookingDate\":\"2024-01-15T09:30:52.737Z\",\"officeName\":\"UID JATIM\",\"startTime\":\"2024-01-18T09:00:00.000Z\",\"endTime\":\"2024-01-18T16:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"},{\"name\":\"Snack Sore\"}],\"participants\":23,\"roomName\":\"Ruang Sudirman\",\"id\":\"14\"},{\"bookingDate\":\"2024-01-04T09:30:52.737Z\",\"officeName\":\"UID JABAR\",\"startTime\":\"2024-01-07T09:00:00.000Z\",\"endTime\":\"2024-01-07T16:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"},{\"name\":\"Snack Sore\"}],\"participants\":23,\"roomName\":\"Ruang Galunggung\",\"id\":\"15\"},{\"bookingDate\":\"2024-01-04T10:30:52.737Z\",\"officeName\":\"UID JABAR\",\"startTime\":\"2024-01-08T09:00:00.000Z\",\"endTime\":\"2024-01-08T14:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"}],\"participants\":20,\"roomName\":\"Ruang Semeru\",\"id\":\"16\"},{\"bookingDate\":\"2024-01-04T10:30:52.737Z\",\"officeName\":\"UID JABAR\",\"startTime\":\"2024-01-08T15:00:00.000Z\",\"endTime\":\"2024-01-08T17:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Sore\"}],\"participants\":18,\"roomName\":\"Ruang Semeru\",\"id\":\"17\"},{\"bookingDate\":\"2024-01-10T08:30:52.737Z\",\"officeName\":\"UID JABAR\",\"startTime\":\"2024-01-13T15:00:00.000Z\",\"endTime\":\"2024-01-13T17:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Sore\"}],\"participants\":20,\"roomName\":\"Ruang Semeru\",\"id\":\"18\"},{\"bookingDate\":\"2024-01-15T08:30:52.737Z\",\"officeName\":\"UID JABAR\",\"startTime\":\"2024-01-17T09:00:00.000Z\",\"endTime\":\"2024-01-17T16:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"},{\"name\":\"Snack Sore\"}],\"participants\":20,\"roomName\":\"Ruang Bromo\",\"id\":\"19\"},{\"bookingDate\":\"2024-01-05T08:45:52.737Z\",\"officeName\":\"UID BALI\",\"startTime\":\"2024-01-08T09:00:00.000Z\",\"endTime\":\"2024-01-08T14:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"}],\"participants\":27,\"roomName\":\"Ruang Ganesha\",\"id\":\"20\"},{\"bookingDate\":\"2024-01-10T08:45:52.737Z\",\"officeName\":\"UID BALI\",\"startTime\":\"2024-01-12T09:00:00.000Z\",\"endTime\":\"2024-01-12T14:00:00.000Z\",\"listConsumption\":[{\"name\":\"Snack Siang\"},{\"name\":\"Makan Siang\"}],\"participants\":27,\"roomName\":\"Ruang Ganesha\",\"id\":\"21\"}]")
	masterResponsePayload := []byte("[{\"createdAt\":\"2024-07-04T04:03:28.947Z\",\"name\":\"Snack Siang\",\"maxPrice\":20000,\"id\":\"1\"},{\"createdAt\":\"2024-07-04T10:04:22.946Z\",\"name\":\"Makan Siang\",\"maxPrice\":30000,\"id\":\"2\"},{\"createdAt\":\"2024-07-03T21:57:46.929Z\",\"name\":\"Snack Sore\",\"maxPrice\":20000,\"id\":\"3\"}]")

	var bookings []BookingPayload

	var masterPrices []MasterJenisPayload

	e := json.Unmarshal(bookingResponsePayload, &bookings)
	if e != nil {
		log.Error("error ", e)
	}

	e = json.Unmarshal(masterResponsePayload, &masterPrices)
	if e != nil {
		log.Error("error ", e)
	}

	var masterPriceMaps = map[string]*MasterJenisPayload{}

	for _, masterPrice := range masterPrices {
		masterPriceMaps[masterPrice.Name] = &masterPrice
	}

	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0)

	filterStart := startDate.Unix()
	filterEnd := endDate.Unix()

	roomUsageGist := map[string]map[string]*RoomUsagePayload{}
	statistic := map[string]map[string]uint64{}

	for _, booking := range bookings {
		bookStartUnix := booking.StartTime.Unix()
		bookEndUnix := booking.EndTime.Unix()
		if (bookStartUnix >= filterStart && bookStartUnix <= filterEnd) || (bookEndUnix >= filterStart && bookEndUnix <= filterEnd) {
			officeContext := roomUsageGist[booking.OfficeName]
			statisticOfficeContext := statistic[booking.OfficeName]
			if officeContext == nil {
				officeContext = map[string]*RoomUsagePayload{}
				roomUsageGist[booking.OfficeName] = officeContext
				statisticOfficeContext = map[string]uint64{}
				statistic[booking.OfficeName] = statisticOfficeContext
			}

			roomPayload := officeContext[booking.RoomName]

			if roomPayload == nil {
				roomPayload = new(RoomUsagePayload)
				roomPayload.Consumptions = map[string]uint{}
				roomPayload.Name = booking.RoomName
				officeContext[booking.RoomName] = roomPayload
			}

			for _, value := range booking.ListConsumptions {
				consumptionType := value["name"]
				selectedMasterPrice := masterPriceMaps[consumptionType]
				prevConsumpt, exists := roomPayload.Consumptions[consumptionType]
				if !exists {
					roomPayload.Consumptions[consumptionType] = 0
					prevConsumpt = 0
				}
				roomPayload.Consumptions[consumptionType] = prevConsumpt + booking.Participants
				roomPayload.FoodExpense = roomPayload.FoodExpense + (uint64(booking.Participants) * selectedMasterPrice.MaxPrice)
			}
			prevTotalRoom, exists := statisticOfficeContext["total"+booking.RoomName]
			if !exists {
				prevTotalRoom = 0
				statisticOfficeContext["total"+booking.RoomName] = 0
			}
			statisticOfficeContext["total"+booking.RoomName] = prevTotalRoom + 1

			prevTotal, exists := statisticOfficeContext["total"]
			if !exists {
				prevTotal = 0
				statisticOfficeContext["total"] = 0
			}
			statisticOfficeContext["total"] = prevTotal + 1
		}
	}

	offices := make([]OfficePayload, len(roomUsageGist))
	oi := 0
	for officeName, roomUsages := range roomUsageGist {
		roomUsageContainers := make([]RoomUsagePayload, len(roomUsages))
		i := 0
		for roomName, roomUsage := range roomUsages {
			roomUsage.UsagePercentage = float64(statistic[officeName]["total"+roomName]*100) / float64(statistic[officeName]["total"])
			roomUsageContainers[i] = *roomUsage
			i++
		}
		offices[oi] = OfficePayload{
			Name:       officeName,
			RoomUsages: roomUsageContainers,
		}

		oi++
	}

	c.Status(fiber.StatusOK).JSON(ResponsePayload{Payload: offices})
	return nil
}

func getSummary(c *fiber.Ctx) error {
	qMonth, qYear := c.Query("month", ""), c.Query("year", "")

	if qMonth == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "month is require"})
	}

	if qYear == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "year is require"})
	}

	month, e := strconv.Atoi(qMonth)
	if e != nil || month > 11 || month < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid month, must be between 0(jan) - 11(dec)"})
	}

	year, e := strconv.Atoi(qYear)

	now := time.Now()
	if e != nil || year < 1970 || year > now.Year() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid year must be between 1970 and " + string(now.Year())})
	}

	return proceedSummary(c, month+1, year)
}
