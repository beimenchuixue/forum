package jwt

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"forum/settings"
	"strings"
	"time"
)

// Header jwt头部数据
type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

func NewHeader(typ string, alg string) *Header {
	return &Header{
		Typ: typ,
		Alg: alg,
	}
}

// Payload 有效负载，存放数据
type Payload struct {
	Exp    int64 `json:"exp"`
	UserId int64 `json:"user_id"`
}

func NewPayload(userId int64) *Payload {
	t := time.Now().Add(time.Duration(settings.Conf.TokenExp) * 60 * time.Second).Unix()
	return &Payload{
		Exp:    t,
		UserId: userId,
	}
}

// Toke 最终token
type Toke struct {
	header  *Header
	payload *Payload
	secret  string
}

func NewToken() *Toke {
	return &Toke{
		header: NewHeader("JWT", "sha256"),
	}
}

// GetToken 获取token
func (t *Toke) GetToken(userId int64) (string, error) {
	// 防止反复生成token
	if t.secret != "" {
		return t.secret, nil
	}
	t.payload = NewPayload(userId)
	// 1. 序列化header 部分
	header, err := json.Marshal(t.header)
	if err != nil {
		return "", err
	}
	headerBase64 := t.encBase64(header)

	// 2. 序列化负载部分
	payload, err := json.Marshal(t.payload)
	if err != nil {
		return "", err
	}
	payloadBase64 := t.encBase64(payload)

	// 3.签名部分
	signature := t.encSha256(headerBase64 + "." + payloadBase64)
	t.secret = fmt.Sprintf("%s.%s.%s", headerBase64, payloadBase64, t.encBase64([]byte(signature)))
	return t.secret, nil
}

// CheckToken 检查token是否合法
func (t *Toke) CheckToken(token string) (p *Payload, err error) {
	tokenData := strings.Split(token, ".")
	// 长度是否为3，不为3则为无效token
	if len(tokenData) != 3 {
		return nil, errors.New("无效token")
	}
	// 获取三部分信息，然后验证签名
	headerBase64, payloadBase64, signatureBase64 := tokenData[0], tokenData[1], tokenData[2]
	signature := t.encSha256(headerBase64 + "." + payloadBase64)
	if signatureBase64 != t.encBase64([]byte(signature)) {
		return nil, errors.New("无效token")
	}

	// 解析 payload
	p = new(Payload)
	payload, err := t.decBase64(payloadBase64)
	if err != nil {
		return nil, errors.New("decBase64: 无效payload")
	}
	err = json.Unmarshal(payload, p)
	if err != nil {
		return nil, errors.New("json.Unmarshal: 无效payload")
	}
	// 返回正确结果
	return p, err
}

// encodeBase64做base64编码
func (t *Toke) encBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// encodeBase64做base64解码
func (t *Toke) decBase64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func (t *Toke) encSha256(data string) string {
	m := sha256.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}
