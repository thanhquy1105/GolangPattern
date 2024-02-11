package templatemethod

func Main() {
	smsOTP := &SMS{}
	o := &OTP{
		ObjectOTP: smsOTP,
	}
	o.GenAndSendOTP(4)

	emailOTP := &Email{}
	e := &OTP{
		ObjectOTP: emailOTP,
	}
	e.GenAndSendOTP(4)
}
