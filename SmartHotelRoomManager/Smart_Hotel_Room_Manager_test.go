
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

	h.addStaff(1, "Nandini")

	h.addGuest(121, "Alice")

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

	occupied := h.checkOccupancyStatus(101)

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
		t.Error("Negative room number accepted")
	}

	err = s.addRoom(&h, -1, "suite")

	if err == nil {
		t.Error("Invalid room type is accepted")
	}

	err = s.addRoom(&h, 101, "Deluxe")

	if err == nil {
		t.Error("Duplicate number accepted")
	}
}

func TestCreateRequest(t *testing.T) {

	h := hotel{}
	s := h.addStaff(1, "Nandini")

	s.addRoom(&h, 101, "Deluxe")

	m := RoomServiceManager{}

	err := m.createRequest(&h, 1, 101, "housekeeping", "HIGH")

	if err != nil {
		t.Error("Valid request was not created")
	}

	if len(m.request) != 1 {
		t.Error("Request not added")
	}

	if m.request[0].requestId != 1 {
		t.Error("Request ID mismatch")
	}

	if m.request[0].roomNumber != 101 {
		t.Error("Room number mismatch")
	}

	if m.request[0].requestType != "housekeeping" {
		t.Error("Request type mismatch")
	}

	if m.request[0].requestType == " " {
		t.Error("Request type cant be blank")
	}
}

func TestFindRequestById(t *testing.T) {

	h := hotel{}
	s := h.addStaff(1, "Nandini")
	s.addRoom(&h, 101, "Deluxe")

	m := RoomServiceManager{}

	m.createRequest(&h, 1, 101, "food", "HIGH")

	req, err := m.findRequestById(1)

	if err != nil {
		t.Error("Request not found")
	}

	if req.requestId != 1 {
		t.Error("Request ID mismatch")
	}

	if req.roomNumber != 101 {
		t.Error("Room number mismatch")
	}

	_, err = m.findRequestById(2)

	if err == nil {
		t.Error("Invalid request ID was accepted")
	}
}

func TestListRequests(t *testing.T) {

	h := hotel{}
	s := h.addStaff(1, "Nandini")
	s.addRoom(&h, 101, "Deluxe")
	s.addRoom(&h, 102, "Suite")

	m := RoomServiceManager{}

	m.createRequest(&h, 1, 101, "food", "HIGH")
	m.createRequest(&h, 2, 102, "towels", "NORMAL")

	result := m.listRequests()

	if len(result) != 2 {
		t.Error("Request list length mismatch")
	}

	if result[0].requestId != 1 {
		t.Error("First request ID mismatch")
	}

	if result[1].requestId != 2 {
		t.Error("Second request ID mismatch")
	}
}

func TestUpdateRequestStatus(t *testing.T) {

	h := hotel{}
	s := h.addStaff(1, "Nandini")
	s.addRoom(&h, 101, "Deluxe")

	m := RoomServiceManager{}

	m.createRequest(&h, 1, 101, "food", "HIGH")

	err := m.updateRequestStatus(1, IN_PROGRESS)

	if err != nil {
		t.Error("OPEN request was not moved to IN_PROGRESS")
	}

	if m.request[0].status != IN_PROGRESS {
		t.Error("Request status was not updated to IN_PROGRESS")
	}

	err = m.updateRequestStatus(1, COMPLETED)

	if err != nil {
		t.Error("IN_PROGRESS request was not moved to COMPLETED")
	}

	if m.request[0].status != COMPLETED {
		t.Error("Request status was not updated to COMPLETED")
	}

	err = m.updateRequestStatus(1, CANCELLED)

	if err == nil {
		t.Error("Completed request should not be cancelled")
	}
}

func TestFilterByStatus(t *testing.T) {

	h := hotel{}
	s := h.addStaff(1, "Nandini")
	s.addRoom(&h, 101, "Deluxe")
	s.addRoom(&h, 102, "Suite")

	m := RoomServiceManager{}

	m.createRequest(&h, 1, 101, "food", "HIGH")
	m.createRequest(&h, 2, 102, "towels", "NORMAL")

	m.updateRequestStatus(2, IN_PROGRESS)

	result := m.filterByStatus(OPEN)

	if len(result) != 1 {
		t.Error("OPEN status filter failed")
	}

	if result[0].requestId != 1 {
		t.Error("Incorrect request returned")
	}

	result = m.filterByStatus(IN_PROGRESS)

	if len(result) != 1 {
		t.Error("IN_PROGRESS status filter failed")
	}
}

func TestFilterByRoomNumber(t *testing.T) {

	h := hotel{}
	s := h.addStaff(1, "Nandini")
	s.addRoom(&h, 101, "Deluxe")
	s.addRoom(&h, 102, "Suite")

	m := RoomServiceManager{}

	m.createRequest(&h, 1, 101, "food", "HIGH")
	m.createRequest(&h, 2, 102, "towels", "NORMAL")
	m.createRequest(&h, 3, 101, "housekeeping", "LOW")

	result := m.filterByRoomNumber(101)

	if len(result) != 2 {
		t.Error("Room number filter failed")
	}

	if result[0].roomNumber != 101 {
		t.Error("Incorrect room number returned")
	}
}

func TestFilterByRequestType(t *testing.T) {

	h := hotel{}
	s := h.addStaff(1, "Nandini")
	s.addRoom(&h, 101, "Deluxe")
	s.addRoom(&h, 102, "Suite")

	m := RoomServiceManager{}

	m.createRequest(&h, 1, 101, "food", "HIGH")
	m.createRequest(&h, 2, 102, "towels", "NORMAL")
	m.createRequest(&h, 3, 101, "food", "LOW")

	result := m.filterByRequestType("food")

	if len(result) != 2 {
		t.Error("Request type filter failed")
	}

	if result[0].requestType != "food" {
		t.Error("Incorrect request type returned")
	}
}