package authorization

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// CreateJwtToken は、提供されたhiddenValueをuserIdとするJWTトークンを生成します。
// を userId とし、"admin" claim を true に設定します。トークンの有効期限は
// 有効期限は 72 時間です。
// JWT トークンは環境変数 JWT_SECRET を使用して作成され、 // HS256 を使用して署名されます。
// HS256署名メソッドを使用して署名される。生成されたトークンは 生成されたトークンは、トークン生成時に発生したエラーとともに // 文字列として返される。
func CreateJwtToken(hiddenValue string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"userId": hiddenValue, // ユーザーIDに変更しても良い
		"admin":  true,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Tokenの期限を72時間で設定
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// IsTokenExpired は、提供されたトークン文字列が期限切れかどうかを検証します。
// トークンをJWT形式でパースし、JWTシグネチャの検証も行います。
// ENV変数 "JWT_SECRET" を使用して署名鍵を取得します。
// トークンが有効で期限切れでない場合、ユーザーのメールアドレスを返します。
// トークンが無効な場合、または期限切れの場合、エラーメッセージを返します。
func IsTokenExpired(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// tokenの期限が切れていないか確認
		expireTime := int64(claims["exp"].(float64))
		if time.Now().Unix() > expireTime {
			return "", fmt.Errorf("is token expired")
		}
		email := claims["userId"].(string)
		return email, nil
	}

	return "", fmt.Errorf("invalid token")
}
