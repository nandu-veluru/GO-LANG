package main

import "errors"

type room struct {
	number   int
	roomType string
	occupied bool
}

type guest struct {
	id   int
	name string
}

type staff struct {
	id   int
	name string
}

type hotel struct {
	roomDetails  []room
	guestDetails []guest
	staffDetails []staff
}

type serviceRequest struct {
	requestId   int
	roomNumber  int
	requestType string
}

type RoomServiceManager struct {
	request []serviceRequest
}

func (staff *staff) addRoom(h *hotel, number int, roomType string) error {

	if number <= 0 {
		return errors.New("Invalid room number")
	}

	if roomType != "Deluxe" && roomType != "Standard" {
		return errors.New("Room type dosent exist")
	}

	for _, r := range h.roomDetails {
		if r.number == number {
			return errors.New("Room is already exist")
		}
	}

	r := room{number, roomType, false}
	h.addRoom(r)
	return nil
}

func (hotel *hotel) addRoom(r room) room {

	hotel.roomDetails = append(hotel.roomDetails, r)
	return r
}

func (hotel *hotel) addStaff(id int, name string) staff {

	s := staff{id, name}
	hotel.staffDetails = append(hotel.staffDetails, s)
	return s
}

func (staff *staff) addGuest(h *hotel, id int, name string) guest {

	g := guest{id, name}
	h.guestDetails = append(h.guestDetails, g)
	return g
}

func (staff *staff) checkOccupancyStatus(h *hotel, number int) bool {

	for _, r := range h.roomDetails {
		if r.number == number {
			return r.occupied
		}
	}
	return false
}

func (hotel *hotel) viewStaffRecords(h *hotel) []staff {

	return h.staffDetails
}

func (staff *staff) viewGuestRecordes(h *hotel) []guest {

	return h.guestDetails
}

func (m *RoomServiceManager) createRequest(h *hotel, id int, roomNumber int, requestType string) error {

	roomExists := false

	for _, r := range h.roomDetails {
		if r.number == roomNumber {
			roomExists = true
		}
	}

	if !roomExists {
		return errors.New("Room does not exist")
	}

	if requestType != "food" && requestType != "towels" &&
		requestType != "housekeeping" && requestType != "maintenance" {
		return errors.New("Invalid request type")
	}

	if requestType == " " {
		return errors.New("request type cant be blank")
	}

	req := serviceRequest{
		id,
		roomNumber,
		requestType,
	}
	m.request = append(m.request, req)

	return nil
}
