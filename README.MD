Пример запроса на создание юзера:

```curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "name=vitya1&password=asdQWE123&dob=11-12-1994&login=vitya@email.com" localhost:80/createUser```

Пример запроса на изменение юзера:

```curl -X PUT "localhost:80/updateUser?login=vitya@email.com1&updatedKey=password&updatedValue=1231231231"```

Пример запроса на получение юзера по логину:

```curl -X GET localhost:80/readUser?login=vitya@email.com```

Пример запроса на получение списка всех юзеров:

```curl -X GET localhost:80/userList```

Пример запроса на удаление юзера по id:

```curl -X DELETE localhost:80/deleteUser?login=vitya@email.com1```