# Behaviors

* Hotel should have rooms with room number, type, createdBy
* * Room number must be unique
* * 

* Hotel should have guests with name and id
* * id must be unique

* HotelManager should create hotel and add staffs
* * staffs will have name and id.
* * id is unique

* Staff will add room with number, type
* * staff will add room with existing number
* * staff will add room with type which doesn't exists
* Staff will checkin guest to specfic room
* * Staff will checkin guest to specific room which is already assigned to different guest
* * Staff will checking guest to room which doesn't exists
* * Staff will checking guest which doesn't exist to a specific room
* * Staff will check Occupancy Status of the hotel room

* staff should create service request 
* * request must have unique id --> number, request type, status
* * staff should able to record guest request

* hotel should assgin work(request) to staff
* * request should assigned to a availble staff member
* * request should not be alloted who is assigned to work already



