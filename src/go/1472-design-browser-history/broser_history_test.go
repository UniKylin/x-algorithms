package leetcode

import "testing"

func TestBrowserHistory(t *testing.T) {
	browserHistory := BrowserHistoryConstructor("leetcode.com")
	browserHistory.Visit("google.com")   // 你原本在浏览 "leetcode.com" 。访问 "google.com"
	browserHistory.Visit("facebook.com") // 你原本在浏览 "google.com" 。访问 "facebook.com"
	browserHistory.Visit("youtube.com")  // 你原本在浏览 "facebook.com" 。访问 "youtube.com"
	browserHistory.Back(1)               // 你原本在浏览 "youtube.com" ，后退到 "facebook.com" 并返回 "facebook.com"
	browserHistory.Back(1)               // 你原本在浏览 "facebook.com" ，后退到 "google.com" 并返回 "google.com"
	browserHistory.Forward(1)            // 你原本在浏览 "google.com" ，前进到 "facebook.com" 并返回 "facebook.com"
	browserHistory.Visit("linkedin.com") // 你原本在浏览 "facebook.com" 。 访问 "linkedin.com"
	browserHistory.Forward(2)            // 你原本在浏览 "linkedin.com" ，你无法前进任何步数。
	browserHistory.Back(2)               // 你原本在浏览 "linkedin.com" ，后退两步依次先到 "facebook.com" ，然后到 "google.com" ，并返回 "google.com"
	browserHistory.Back(7)               // 你原本在浏览 "google.com"， 你只能后退一步到 "leetcode.com" ，并返回 "leetcode.com"

}
