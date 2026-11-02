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

* * hotel able to view staff records
* * staff able to view guest records

* hotel should create RoomServiceManager 
* * add request types
* * validate room numbers and request type
* * reject blank request values

* Queue management system for hotel staff
* * all the requests should be maintained in Request Lifecycle 
* * OPEN -> IN_PROGRESS -> COMPLETED (or CANCELLED)
* * staff can  staff to filter and inspect the request queue by status, room number, or request type.

* Add proroity values 
* * LOW, NORMAL AND HIGH
* * Add prority field to RoomServiceManager struct
* * Staff need to assign prority when request is started
* * if both values are same consider frist given request
* * validate the prority 

* Implement prority based queue
* * High priority request first
* * Normal priority request next
* * Low priority request last
* * Assign the availble staff according to the priority
* * staff will open request and mark as complete after work done

* * supervisor will check the prioritized queue
* * supervisor will assign the task for eligible staff member
* * staff starts the works and completes it
* * supervisor views about staff work


