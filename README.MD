Пример запроса на создание юзера:

```curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d "name=vitya1&password=asdQWE123&dob=11-12-1994&login=vitya@email.com" localhost:80/createUser```

Пример запроса на изменение юзера:

``` ```

Пример запроса на получение юзера по логину:

```curl -X GET localhost:80/readUser?login=vitya@email.com```

Пример запроса на получение списка всех юзеров:

```curl -X GET localhost:80/userList```

Пример запроса на удаление юзера по id:

``` ```