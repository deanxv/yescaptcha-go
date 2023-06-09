package req

type CreateTaskRequest struct {
	ClientKey string      `json:"clientKey"`
	Task      interface{} `json:"task"`
	SoftID    string      `json:"softID"`
}

type ImageToTextTaskRequest struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

type NoCaptchaTaskProxylessRequest struct {
	Type       string `json:"type"`
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"`
	PageAction string `json:"pageAction"`
}

type RecaptchaV3TaskProxylessRequest struct {
	Type        string `json:"type"`
	WebsiteURL  string `json:"websiteURL"`
	WebsiteKey  string `json:"websiteKey"`
	IsInvisible bool   `json:"isInvisible"`
}

type RecaptchaV2EnterpriseTaskProxylessRequest struct {
	Type              string `json:"type"`
	WebsiteURL        string `json:"websiteURL"`
	WebsiteKey        string `json:"websiteKey"`
	EnterprisePayload struct {
		S string `json:"s"`
	} `json:"isInvisible"`
}

type ReCaptchaV2ClassificationRequest struct {
	Type       string `json:"type"`
	Image      string `json:"image"`
	Question   string `json:"question"`
	Confidence string `json:"confidence"`
}

type HCaptchaTaskProxylessRequest struct {
	Type        string `json:"type"`
	WebsiteURL  string `json:"websiteURL"`
	WebsiteKey  string `json:"websiteKey"`
	UserAgent   string `json:"userAgent"`
	IsInvisible bool   `json:"isInvisible"`
	Rqdata      string `json:"rqdata"`
}

type HCaptchaClassificationRequest struct {
	Type       string   `json:"type"`
	Queries    []string `json:"queries"`
	Question   string   `json:"question"`
	Coordinate bool     `json:"coordinate"`
}

type FuncaptchaTaskProxylessRequest struct {
	Type                     string `json:"type"`
	WebsiteURL               string `json:"websiteURL"`
	WebsitePublicKey         string `json:"websitePublicKey"`
	FuncaptchaApiJSSubdomain string `json:"funcaptchaApiJSSubdomain"`
	Data                     string `json:"data"`
}

type FunCaptchaClassificationRequest struct {
	Type     string      `json:"type"`
	Image    interface{} `json:"image"`
	Question string      `json:"question"`
}

type TurnstileTaskProxylessRequest struct {
	Type       string `json:"type"`
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"`
}
