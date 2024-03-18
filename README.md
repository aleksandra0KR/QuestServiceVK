# Бэкенд часть приложения для выполнения пользователем заданий за награду. 
# REST API для взаимодействия c базой данных


Небольшой REST API сервис, засчитывающий задания для пользователя.
Есть пользватель, который может брать задания, которые состоят из подзаданий. После выполенения задания(всех подзаданий в нем) пользователю начисляется награда. Пользователь может выполнить каждое задание только один раз.



# Запросы
## Пользователь / User
## POST
#### ```POST /quest/user``` - создать пользователя
#### ```POST /quest/user/takequest/{quest_id}/{user_id}``` - взять задание
#### ```POST /quest/user/takesubquest/{subquest_id}/{user_id}``` - взять подзадание

## GET
#### ```GET /quest/user/{user_id}``` - получить историю выполненых заданий и подзаданий и баланс

## PUT
#### ```PUT /quest/user/donequest/{quest_id}/{user_id}``` - завершить задание
#### ```PUT /quest/user/donesubquest/{subquest_id}/{user_id}}``` - завершить подзадание
#### ```PUT /quest/user/updatepassword/{user_id}/{new_password}``` - поменять пароль
#### ```PUT /quest/user/updateemail/{user_id}/{new_email}``` - поменять email
#### ```PUT /quest/user/updateusernamel/{user_id}/{new_username}``` - поменять имя пользователя

## DELETE
#### ```DELETE /quest/user/{user_id}``` - удалить пользователя


## Задание / Quest
## POST
#### ```POST /quest/quest``` - создать задание


## PUT
#### ```PUT /quest/quest/updatetitle/{quest_id}/{new_title}``` - поменять название задания
#### ```PUT /quest/user/updatedescription/{quest_id}/{new_description}``` - поменять описание задания 
#### ```PUT /quest/user/updateduedate/{quest_id}/{new_updateduedate}``` - поменять дату завершения задания

## DELETE
#### ```DELETE /quest/quest/{quest_id}``` - удалить задание


## Подзадание / Subuest
## POST
#### ```POST /quest/subquest``` - создать подзадание


## PUT
#### ```PUT /quest/quest/updatetitle/{subquest_id}/{new_title}``` - поменять название подзадания
#### ```PUT /quest/user/updatedescription/{subquest_id}/{new_description}``` - поменять описание подзадания
#### ```PUT /quest/user/updateduedate/{subquest_id}/{new_updateduedate}``` - поменять дату завершения подзадания

## DELETE
#### ```DELETE /quest/subquest/{subquest_id}``` - удалить подзадание

##
- субд - postgreSQL
- миграции - migrate


# Запуск в Docker
```
docker-compose up
```
