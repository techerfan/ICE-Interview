package controllers

import (
	"fmt"
	"interview/internal/dto"
	"interview/internal/pkg/httpmsg"
	"interview/internal/service/cart"
	"interview/internal/validator/cartvalidator"
	"log"
	"strconv"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CartItemForm struct {
	Product  string `form:"product"   binding:"required"`
	Quantity string `form:"quantity"  binding:"required"`
}

type CartController struct {
	cartService   *cart.Service
	cartValidator cartvalidator.Validator
	ginEngine     *gin.Engine
}

func New(
	cartService *cart.Service,
	cartValidator cartvalidator.Validator,
	ginEngine *gin.Engine,
) *CartController {
	controller := &CartController{
		cartService:   cartService,
		cartValidator: cartValidator,
		ginEngine:     ginEngine,
	}

	controller.setupRoutes(ginEngine)

	return controller
}

func (c *CartController) ShowAddItemForm(ctx *gin.Context) {
	sessionID := ctx.GetString(sessionIDKey)

	data := map[string]interface{}{
		"Error": ctx.Query("error"),
		//"cartItems": cartItems,
	}

	resp, err := c.cartService.GetCartData(ctx.Request.Context(), dto.GetCartDataRequest{SessionID: sessionID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		if code > 500 {
			// TODO: log the error
			ctx.String(code, msg)
			return
		} else if code == 404 {
			// TODO: create the cart here
		}
	}

	var items []map[string]interface{}

	for _, cartItem := range resp.Items {
		item := map[string]interface{}{
			"ID":       cartItem.ID,
			"Quantity": cartItem.Quantity,
			"Price":    cartItem.Price,
			"Product":  cartItem.ProductName,
		}

		items = append(items, item)
	}

	data["CartItems"] = items

	html, err := renderTemplate(data)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.Header("Content-Type", "text/html")
	ctx.String(200, html)
}

func (c *CartController) AddItem(ctx *gin.Context) {
	sessionID := ctx.GetString(sessionIDKey)

	// if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
	// 	ctx.Redirect(302, "/")
	// 	return
	// }

	addItemForm, err := c.getCartItemForm(ctx)
	if err != nil {
		ctx.Redirect(302, "/?error="+err.Error())
		return
	}

	quantity, err := strconv.ParseInt(addItemForm.Quantity, 10, 0)
	if err != nil {
		ctx.Redirect(302, "/?error=invalid quantity")
		return
	}

	req := dto.AddItemToCartRequest{
		SessionID: sessionID,
		Product:   addItemForm.Product,
		Quantity:  int(quantity),
	}

	if _, err := c.cartValidator.ValidateAddItem(req); err != nil {
		ctx.Redirect(302, "/?error="+err.Error())
		return
	}

	err = c.cartService.AddItemToCart(ctx.Request.Context(), req)

	if err != nil {
		msg, _ := httpmsg.Error(err)
		ctx.Redirect(302, "/?error="+msg)
	}

	ctx.Redirect(302, "/")
}

func (c *CartController) DeleteCartItem(ctx *gin.Context) {
	sessionID := ctx.GetString(sessionIDKey)
	// cookie, err := ctx.Request.Cookie("ice_session_id")

	// if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
	// 	ctx.Redirect(302, "/")
	// 	return
	// }

	cartItemIDString := ctx.Query("cart_item_id")
	if cartItemIDString == "" {
		ctx.Redirect(302, "/")
		return
	}

	cartItemID, err := strconv.Atoi(cartItemIDString)
	if err != nil {
		ctx.Redirect(302, "/")
		return
	}

	req := dto.DeleteCartItemRequest{
		SessionID:  sessionID,
		CartItemID: uint(cartItemID),
	}

	if _, err := c.cartValidator.ValidateDeleteItem(req); err != nil {
		ctx.Redirect(302, "/?error="+err.Error())
		return
	}

	err = c.cartService.DeleteCartItem(ctx.Request.Context(), req)
	if err != nil {
		msg, _ := httpmsg.Error(err)
		ctx.Redirect(302, "/?error="+msg)
	}

	ctx.Redirect(302, "/")
}

func (c *CartController) getCartItemForm(ctx *gin.Context) (*CartItemForm, error) {
	if ctx.Request.Body == nil {
		return nil, fmt.Errorf("body cannot be nil")
	}

	form := &CartItemForm{}

	if err := binding.FormPost.Bind(ctx.Request, form); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return form, nil
}

func renderTemplate(pageData interface{}) (string, error) {
	// Read and parse the HTML template file
	tmpl, err := template.ParseFiles("../../static/add_item_form.html")
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	// Create a strings.Builder to store the rendered template
	var renderedTemplate strings.Builder

	err = tmpl.Execute(&renderedTemplate, pageData)
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	// Convert the rendered template to a string
	resultString := renderedTemplate.String()

	return resultString, nil
}
