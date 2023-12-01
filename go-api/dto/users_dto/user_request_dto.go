package users_dto

type UserRequestDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	DNI      string `json:"dni"`
	Email    string `json:"email"`
	Admin    int    `json:"admin"`
	Token    string `json:"token"`
}

type UsersRequestDto struct {
	Users []UserDto `json:"users"`
}

/*
[{4 2023-07-15 00:00:00 +0000 UTC 2023-07-25 00:00:00 +0000 UTC 1 Cama Matrimonial 1 1 {0      false 0 []} {0  0  [] []}}
 {5 2023-07-15 00:00:00 +0000 UTC 2023-07-25 00:00:00 +0000 UTC 1 Cama Matrimonial 1 1 {0      false 0 []} {0  0  [] []}}
  {6 2023-07-15 00:00:00 +0000 UTC 2023-07-28 00:00:00 +0000 UTC 1 Cama Matrimonial 1 1 {0      false 0 []} {0  0  [] []}}
   {7 2023-07-15 00:00:00 +0000 UTC 2023-07-25 00:00:00 +0000 UTC 1 Cama Matrimonial 1 1 {0      false 0 []} {0  0  [] []}}] */
