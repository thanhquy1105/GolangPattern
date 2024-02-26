package templatemethod

type OTP struct {
	ObjectOTP IOtp
}

func (o *OTP) GenAndSendOTP(otpLenth int) error {
	otp := o.ObjectOTP.genRandomOTP(otpLenth)
	o.ObjectOTP.saveOTPCache(otp)
	message := o.ObjectOTP.getMessage(otp)
	err := o.ObjectOTP.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
