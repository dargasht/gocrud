package gocrud

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/kavenegar/kavenegar-go"
	"go.uber.org/zap"
)

// This is a set of helper funcion and structs to handle otp
// We use kavenegar for sending messages
// This is specially created for kavenegar

type OTP struct {
	Code      string
	Timestamp time.Time
}

type OTPStore struct {
	mu    sync.Mutex
	codes map[string]OTP
}

var store = OTPStore{
	codes: make(map[string]OTP),
}

func generateRandomCode() string {
	const digits = "0123456789"
	result := make([]byte, 6)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		result[i] = digits[num.Int64()]
	}
	return string(result)
}
func generateAndStoreOTP(mobile string) string {
	store.mu.Lock()
	defer store.mu.Unlock()
	otp := generateRandomCode()
	store.codes[mobile] = OTP{
		Code:      otp,
		Timestamp: time.Now(),
	}
	return otp
}

// This is for sending otp
func SendOTP(mobile string, logger *zap.Logger) error {

	store.mu.Lock()
	//Removing expired otp
	for _, entry := range store.codes {
		if time.Since(entry.Timestamp) > 2*time.Minute {
			delete(store.codes, mobile) // Remove the expired OTP
		}
	}
	store.mu.Unlock()

	//Check the store for existing otp
	if _, exists := store.codes[mobile]; exists {
		return fmt.Errorf("OTP already sent to %s", mobile)
	}

	otp := generateAndStoreOTP(mobile)
	api := kavenegar.New(GoCRUDConfig.otpApiKey)
	template := "otp"
	params := &kavenegar.VerifyLookupParam{}
	if res, err := api.Verify.Lookup(mobile, template, otp, params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			logger.Info(err.Error())
		case *kavenegar.HTTPError:
			logger.Info(err.Error())
		default:
			logger.Info(err.Error())
		}
	} else {
		logger.Info("MessageId 	= " + strconv.Itoa(res.MessageID))
		logger.Info("Code      	= " + otp)
	}
	return nil
}

// This is for validating otp
func ValidateOTP(mobile, otp string) bool {
	store.mu.Lock()
	defer store.mu.Unlock()
	if entry, exists := store.codes[mobile]; exists {
		if time.Since(entry.Timestamp) > 2*time.Minute {
			delete(store.codes, mobile) // Remove the expired OTP
			return false
		}
		if entry.Code == otp {
			// delete(store.codes, mobile) // Remove the used OTP
			return true
		}
	}
	return false
}
