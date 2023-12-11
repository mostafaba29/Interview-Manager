# Interview-Manager
-Appointment Booking app typically for managers with a lot of appointments through the day 
-Used fiber framework, gorm for database models (ORM) with sqlite, sessions for authentication
-The user first identify his position as there are different pages for each user:
The employees(receptionist or secretary) have more options as they create the appointments
The managers only can see the appointments assigned to them and then react to them by approving, rejection or updating
-There are also some advanced searching or filteration the user can make like:
showing only the approved appointments, declined appointments
all the appointments from a specific time to another 
all the appointments related to a specific client and more
