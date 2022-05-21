package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type BaseClaims struct {
	UserId string `json:"user_id"` //添加自定义的字段内容
}

type MyCustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims //继承
}

//签名
var (
	signingKey  = []byte("12345678")
	contextKey  = "user"
	duration, _ = time.ParseDuration("1m") //过期时间
)

func jwtSuccesss(c *fiber.Ctx) error {
	user := c.Locals(contextKey).(*jwt.Token) //获取通过本地上下文key获取到jwt的请求上下文值
	claims := user.Claims.(*MyCustomClaims)
	userId := claims.UserId
	log.Println(userId)
	//这里可以解析内容将token和redis中的进行对比，前提是获取token的时候要将token存入redis，
	//过期时间和token的jwt过期时间一致

	return c.Next() //进入到hander逻辑处理
}

func jwtError(c *fiber.Ctx, err error) error {
	//没有携带token
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusForbidden).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	//token过期
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func InitJwtMiddleware(app *fiber.App) {
	// JWT 中间件
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:     signingKey,             //签名
		SigningMethod:  "HS256",                //签名方式
		Claims:         &MyCustomClaims{},      //内容
		TokenLookup:    "header:Authorization", //从哪里取
		AuthScheme:     "Bearer",               //前缀
		ContextKey:     contextKey,             //存储当前token信息到请求上下文,key值
		SuccessHandler: jwtSuccesss,            //校验成功
		ErrorHandler:   jwtError,               //当签名校验失败进入的方法
	}))

}

//构建内容
func createClaims(baseClaims BaseClaims) MyCustomClaims {

	claims := MyCustomClaims{
		baseClaims,
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    "签名的发行者",
		},
	}
	return claims
}

//创建token
func CreateToken(baseClaims BaseClaims) (string, error) {
	claims := createClaims(baseClaims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

//解析token
func parse(tokenString string) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		fmt.Printf("%v %v", claims.BaseClaims.UserId, claims.RegisteredClaims.ExpiresAt)
	} else {
		//此处可以判断token是否已经过期了
		fmt.Println(err)
	}
}
