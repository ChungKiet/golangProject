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

type UserInfo struct {
	Account string `form:"account"`
}

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

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var username string
	var userInfo UserInfo
	if ctx.ShouldBindQuery(&userInfo) == nil {
		username = userInfo.Account
	}

	// Check length
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

	if user != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "This account name has already registered!"})
		return
	}

	var userCreated = models.User{
		Name:        username,
		Balance:     0,
		IsTransform: false,
		Sent:        []models.Sent{},
		Received:    []models.Received{},
	}

	err = uc.UserService.CreateUser(&userCreated)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) Transfer(ctx *gin.Context) {
	var sent string
	var receive string
	var amount int
	var transaction Transaction
	if ctx.ShouldBindQuery(&transaction) == nil {
		sent = transaction.From
		receive = transaction.To
		amount = transaction.Amount
	}

	if receive == "admin" {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Your can't sent money to admin"})
		return
	}

	var checkAdmin bool
	checkAdmin = false
	if sent == "admin" {
		checkAdmin = true
	}

	reAlphaBeta := regexp.MustCompile("[a-zA-Z]")
	reNum := regexp.MustCompile("[0-9]")

	// check dieu kien ten tren fe
	if len(sent) != 5 || len(receive) != 5 {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The account's name is invalid!"})
		return
	}

	if (!reAlphaBeta.MatchString(sent) || !reNum.MatchString(sent)) && !checkAdmin {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The sent's username is invalid!"})
		return
	}

	if !reAlphaBeta.MatchString(receive) || !reNum.MatchString(receive) {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The receive's username is invalid!"})
		return
	}

	userReceive, err := uc.UserService.GetUser(&receive)
	if userReceive == nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The receive's username does not exist", "receive": receive})
		return
	}

	// check admin

	userSent, err := uc.UserService.GetUser(&sent)
	if userSent == nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "The sent's username does not exist"})
		return
	}

	if amount > userSent.Balance && !checkAdmin {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Your balance is not enough!"})
		return
	}

	for userReceive.IsTransform {
		userReceive, err = uc.UserService.GetUser(&receive)
		time.Sleep(time.Second / 10)
	}

	for userSent.IsTransform {
		userSent, err = uc.UserService.GetUser(&sent)
		time.Sleep(time.Second / 10)
	}

	err = uc.UserService.UpdateUser(userSent)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	// set true for user receive
	for userReceive.IsTransform {
		userReceive, err = uc.UserService.GetUser(&receive)
		time.Sleep(time.Second / 10)
	}

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

	userReceive.IsTransform = false
	userSent.IsTransform = false

	err = uc.UserService.UpdateUser(userSent)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": ""})
		return
	}

	err = uc.UserService.UpdateUser(userReceive)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": ""})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

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

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var username string
	var userInfo UserInfo
	if ctx.ShouldBindQuery(&userInfo) == nil {
		username = userInfo.Account
	}
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.GET("/register", uc.CreateUser)
	userroute.GET("/get", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.GET("/transfer", uc.Transfer)
	userroute.GET("/details", uc.GetDetails)
	userroute.GET("/delete", uc.DeleteUser)
}
