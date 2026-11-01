
package main

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

func (h *hotel) addRoom(r room) room {
	h.roomDetails = append(h.roomDetails, r)
	return r
}

func (h *hotel) addStaff(id int, name string) staff {
	s := staff{id, name}
	h.staffDetails = append(h.staffDetails, s)
	return s
}

func (h *hotel) addGuest(id int, name string) guest {
	g := guest{id, name}
	h.guestDetails = append(h.guestDetails, g)
	return g
}

func (h *hotel) checkOccupancyStatus(number int) bool {
	for _, r := range h.roomDetails {
		if r.number == number {
			return r.occupied
		}
	}
	return false
}

func (h *hotel) viewStaffRecords() []staff {
	return h.staffDetails
}

func (h *hotel) viewGuestRecords() []guest {
	return h.guestDetails
}
