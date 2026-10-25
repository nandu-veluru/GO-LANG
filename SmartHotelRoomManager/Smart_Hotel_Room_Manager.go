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
