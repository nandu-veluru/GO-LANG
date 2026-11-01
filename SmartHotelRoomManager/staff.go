package main

import "errors"

type serviceRequest struct {
	requestId   int
	roomNumber  int
	requestType string
	status      string
	priority    string
}

const (
	OPEN        = "OPEN"
	IN_PROGRESS = "IN_PROGRESS"
	COMPLETED   = "COMPLETED"
	CANCELLED   = "CANCELLED"
)

const (
	LOW    = "LOW"
	NORMAL = "NORMAL"
	HIGH   = "HIGH"
)

type RoomServiceManager struct {
	request []serviceRequest
}

func (s *staff) addRoom(h *hotel, number int, roomType string) error {

	if number <= 0 {
		return errors.New("Invalid room number")
	}

	if roomType != "Deluxe" && roomType != "Standard" {
		return errors.New("Invalid room type")
	}

	for _, r := range h.roomDetails {

		if r.number == number {
			return errors.New("Room already exists")
		}
	}

	h.addRoom(room{number, roomType, false})

	return nil
}

func (m *RoomServiceManager) createRequest(h *hotel, id int, roomNumber int, requestType string, priority string) error {

	if id <= 0 {
		return errors.New("Invalid request ID")
	}

	if roomNumber <= 0 {
		return errors.New("Invalid room number")
	}

	for _, r := range m.request {

		if r.requestId == id {
			return errors.New("Request ID already exists")
		}
	}

	roomExists := false

	for _, r := range h.roomDetails {

		if r.number == roomNumber {
			roomExists = true
		}
	}

	if roomExists == false {
		return errors.New("Room does not exist")
	}

	if requestType != "food" && requestType != "towels" && requestType != "housekeeping" && requestType != "maintenance" {
		return errors.New("Invalid request type")
	}

	if priority != LOW && priority != NORMAL && priority != HIGH {
        return errors.New("Invalid priority")
    }

	req := serviceRequest{
		requestId:   id,
		roomNumber:  roomNumber,
		requestType: requestType,
		status:      OPEN,
		priority:    priority,
	}

	m.request = append(m.request, req)

	return nil
}

func (m *RoomServiceManager) findRequestById(id int) (serviceRequest, error) {

	for _, r := range m.request {

		if r.requestId == id {
			return r, nil
		}
	}

	return serviceRequest{}, errors.New("Request not found")
}

func (m *RoomServiceManager) listRequests() []serviceRequest {

	result := []serviceRequest{}

	for _, r := range m.request {
		result = append(result, r)
	}

	return result
}

func (m *RoomServiceManager) updateRequestStatus(requestId int, status string) error {

	if requestId <= 0 {
		return errors.New("Invalid request ID")
	}

	for i := range m.request {

		if m.request[i].requestId == requestId {

			if status == IN_PROGRESS {

				if m.request[i].status == OPEN {
					m.request[i].status = IN_PROGRESS
					return nil
				} else {
					return errors.New("Only OPEN requests can be started")
				}

			} else if status == COMPLETED {

				if m.request[i].status == IN_PROGRESS {
					m.request[i].status = COMPLETED
					return nil
				} else {
					return errors.New("Only IN_PROGRESS requests can be completed")
				}

			} else if status == CANCELLED {

				if m.request[i].status == OPEN || m.request[i].status == IN_PROGRESS {

					m.request[i].status = CANCELLED
					return nil

				} else {
					return errors.New("Request cannot be cancelled")
				}

			} else {
				return errors.New("Invalid request status")
			}
		}
	}

	return errors.New("Request not found")
}

func (m *RoomServiceManager) filterByStatus(status string) []serviceRequest {

	result := []serviceRequest{}

	for _, r := range m.request {

		if r.status == status {

			if r.status == OPEN || r.status == IN_PROGRESS {
				result = append(result, r)
			}
		}
	}

	return result
}

func (m *RoomServiceManager) filterByRoomNumber(number int) []serviceRequest {

	result := []serviceRequest{}

	for _, r := range m.request {

		if r.roomNumber == number {

			if r.status == OPEN || r.status == IN_PROGRESS {
				result = append(result, r)
			}
		}
	}

	return result
}

func (m *RoomServiceManager) filterByRequestType(requestType string) []serviceRequest {

	result := []serviceRequest{}

	for _, r := range m.request {

		if r.requestType == requestType {

			if r.status == OPEN || r.status == IN_PROGRESS {
				result = append(result, r)
			}
		}
	}

	return result
}
