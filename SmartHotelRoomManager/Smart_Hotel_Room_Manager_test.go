package main

import (
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestHotelInitialization(t *testing.T) {
	h := hotel{}

	if len(h.guestDetails) != 0 {
		t.Error("Guest details is not empty")
	}

	if len(h.roomDetails) != 0 {
		t.Error("Room details is not empty")
	}

	if len(h.staffDetails) != 0 {
		t.Error("Staff details is not empty")
	}
}

func TestAddStaff(t *testing.T) {
	h := hotel{}

	h.addStaff(1, "Nandini")

	if len(h.staffDetails) != 1 {
		t.Error("Staff not added to the hotel")
	}

	if h.staffDetails[0].id != 1 {
		t.Error("Staff id mismatch")
	}

	if h.staffDetails[0].name != "Nandini" {
		t.Error("Staff name mismatch")
	}
}

func TestStaffShouldAddRoom(t *testing.T) {
	h := hotel{}
	s := h.addStaff(1, "Nandini")

	s.addRoom(&h, 101, "Deluxe")

	if len(h.roomDetails) != 1 {
		t.Error("Room not added to the hotel")
	}

	if h.roomDetails[0].number != 101 {
		t.Error("room number mismatch")
	}

	if h.roomDetails[0].roomType != "Deluxe" {
		t.Error("type mismatch")
	}
}

func TestStaffShouldAddGuest(t *testing.T) {
	h := hotel{}

	g := h.addStaff(1, "Nandini")

	g.addGuest(&h, 121, "Alice")

	if len(h.guestDetails) != 1 {
		t.Error("guest is not added to hotel")
	}

	if h.guestDetails[0].id != 121 {
		t.Error("guest id mismatched")
	}

	if h.guestDetails[0].name != "Alice" {
		t.Error("Guest name is mismatched")
	}
}

func TestStaffShouldCheckOccupancyStatus(t *testing.T) {
	h := hotel{}

	s := h.addStaff(1, "Nandini")

	s.addRoom(&h, 101, "Deluxe")

	occupied := s.checkOccupancyStatus(&h, 101)

	if occupied {
		t.Error("Room is occupied")
	}
}

func TestStaffShouldValidateRoom(t *testing.T) {

	h := hotel{}
	s := h.addStaff(1, "Nandini")

	err := s.addRoom(&h, 101, "Deluxe")

	if err != nil {
		t.Error("Valid room was not added")
	}

	err = s.addRoom(&h, 0, "Deluxe")

	if err == nil {
		t.Error("Invalid room number was accepted")
	}

	err = s.addRoom(&h, -1, "Deluxe")

	if err == nil {
		t.Error("Negitive room number accepted")
	}

	err = s.addRoom(&h, -1, "suite")

	if err == nil {
		t.Error("Inalid room type is accepted")
	}

	err = s.addRoom(&h, 101, "standard")

	if err == nil {
		t.Error("Dulicate number accepted")
	}

}
