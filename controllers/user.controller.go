package controllers

import (
	"net/http"
	"regexp"
	"time"

	"example.com/sarang-apis/models"
	"example.com/sarang-apis/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

// Struct as adapter of link
type UserInfo struct {
	Account string `form:"account"`
}

// Struct as adapter of link
type Transaction struct {
	From   string `form:"from"`
	To     string `form:"to"`
	Amount int    `form:"amount"`
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

// Create user controller
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var username string
	var userInfo UserInfo

	// Get name from link
	if ctx.ShouldBindQuery(&userInfo) == nil {
		username = userInfo.Account
	}

	// Check length (must be 5)
	if len(username) != 5 {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Your username length must be 5!"})
		return
	}
	// Check regex (must have alphanumeric characters)
	reAlphaBeta := regexp.MustCompile("[a-zA-Z]")
	reNum := regexp.MustCompile("[0-9]")
	reSpecial := regexp.MustCompile("[$&+,:;=?@#|'<>.^*()%!- ]")

	//[$&+,:;=?@#|'<>.^*()%!-]
	if (!reAlphaBeta.MatchString(username) && !reNum.MatchString(username)) || reSpecial.MatchString(username) {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Your username must have both alpha and numeric!"})
		return
	}

	// Check already register
	user, err := uc.UserService.GetUser(&username)

	if user != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "This account name has already registered!"})
		return
	}

	// Init new user
	var userCreated = models.User{
		Name:       username,
		Balance:    0,
		IsTransfer: false,
		Sent:       []models.Sent{},
		Received:   []models.Received{},
	}

	// Create user
	err = uc.UserService.CreateUser(&userCreated)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	// Response success
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Transfer controller
// I have created an admin account to transfer to others account without checking the amount.
// This feature help we test the API easily.
func (uc *UserController) Transfer(ctx *gin.Context) {
	var sent string
	var receive string
	var amount int
	var transaction Transaction
	// Get info: from, to and amount from link
	if ctx.ShouldBindQuery(&transaction) == nil {
		sent = transaction.From
		receive = transaction.To
		amount = transaction.Amount
	}

	// Can't transfer to admin
	if receive == "admin" {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Your can't sent money to admin"})
		return
	}

	// Check the sent user
	var checkAdmin bool
	checkAdmin = false
	if sent == "admin" {
		checkAdmin = true
	}

	reAlphaBeta := regexp.MustCompile("[a-zA-Z]")
	reNum := regexp.MustCompile("[0-9]")
	reSpecial := regexp.MustCompile("[$&+,:;=?@#|'<>.^*()%!- ]")

	// Validate name of user in FE
	if len(sent) != 5 || len(receive) != 5 {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The account's name is invalid!"})
		return
	}

	// check self-transfer
	if sent == receive {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "You can't transfer to yourself!"})
		return
	}

	// If admin --> ignore format of name
	if (!reAlphaBeta.MatchString(sent) && !reNum.MatchString(sent)) || reSpecial.MatchString(sent) {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The sent's username is invalid!"})
		return
	}

	// Check received user
	if (!reAlphaBeta.MatchString(receive) && !reNum.MatchString(receive)) || reSpecial.MatchString(receive) {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The receive's username is invalid!"})
		return
	}

	// Check whether user received is exists
	userReceive, err := uc.UserService.GetUser(&receive)
	if userReceive == nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The receive's username does not exist", "receive": receive})
		return
	}

	// Check whether user received is exists
	// (admin is an acount in database)
	userSent, err := uc.UserService.GetUser(&sent)
	if userSent == nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The sent's username does not exist"})
		return
	}

	// Check the sent user balance
	if amount > userSent.Balance && !checkAdmin {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Your balance is not enough!"})
		return
	}

	// Check and reload until user received is free
	for userReceive.IsTransfer {
		userReceive, err = uc.UserService.GetUser(&receive)
		time.Sleep(time.Second / 10) // Estimate the time to resent a request, this will be change to the most value efficient
	}

	// Check and reload until user sent is free
	for userSent.IsTransfer {
		userSent, err = uc.UserService.GetUser(&sent)
		time.Sleep(time.Second / 10) // Estimate the time to resent a request, this will be change to the most value efficient
	}

	err = uc.UserService.UpdateUser(userSent)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	// If both 2 side are free, change value of IsTransfer and start a transfer
	err = uc.UserService.UpdateUser(userReceive)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	var tempSent models.Sent
	var tempReceived models.Received
	if !checkAdmin {
		userSent.Balance -= amount
	}
	userReceive.Balance += amount
	tempSent.To = userReceive.Name
	tempSent.Amount = amount
	tempReceived.From = userSent.Name
	tempReceived.Amount = amount
	userReceive.Received = append(userReceive.Received, tempReceived)
	userSent.Sent = append(userSent.Sent, tempSent)

	userReceive.IsTransfer = false
	userSent.IsTransfer = false

	err = uc.UserService.UpdateUser(userSent)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Failed to update the balance and history!"})
		return
	}

	err = uc.UserService.UpdateUser(userReceive)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Failed to update the balance and history!"})
		return
	}

	// Response success
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Get the details of a user (balance, sent history, received history)
func (uc *UserController) GetDetails(ctx *gin.Context) {
	var username string
	var userInfo UserInfo
	if ctx.ShouldBindQuery(&userInfo) == nil {
		username = userInfo.Account
	}
	// fmt.Print(username)
	if len(username) != 5 {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Your username length must be 5!"})
		return
	}
	// Check regex
	reAlphaBeta := regexp.MustCompile("[a-zA-Z]")
	reNum := regexp.MustCompile("[0-9]")

	if !reAlphaBeta.MatchString(username) || !reNum.MatchString(username) {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Your username must have both alpha and numeric!"})
		return
	}

	// Check already register
	user, err := uc.UserService.GetUser(&username)

	if user != nil && err == nil {
		ctx.JSON(http.StatusOK, gin.H{"balance": user.Balance, "sent": user.Sent, "received": user.Received})
		return
	}
}

// Get info of a user
func (uc *UserController) GetUser(ctx *gin.Context) {
	var username string
	var userInfo UserInfo
	if ctx.ShouldBindQuery(&userInfo) == nil {
		username = userInfo.Account
	}
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Get all of users
func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// Delete a user, I use the method get because i just test this api via url
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var username string
	var userInfo UserInfo
	if ctx.ShouldBindQuery(&userInfo) == nil {
		username = userInfo.Account
	}
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The user isn't exists!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Router: Call controller via parameter
func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.GET("/register", uc.CreateUser)
	userroute.GET("/get", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.GET("/transfer", uc.Transfer)
	userroute.GET("/details", uc.GetDetails)
	userroute.GET("/delete", uc.DeleteUser)
}
