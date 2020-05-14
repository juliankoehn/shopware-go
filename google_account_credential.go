package shopware

// GoogleAccountCredential see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/GoogleShopping/DataAbstractionLayer/GoogleAccountCredential.php
type GoogleAccountCredential struct {
	AccessToken  string `json:"accessToken"`
	Created      int    `json:"created"`
	Scope        string `json:"scope"`
	IDToken      string `json:"idToken"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}
