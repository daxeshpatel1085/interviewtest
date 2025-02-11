package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestLoginOutApplication(t *testing.T) {
	chromedp.Flag("headless", false)
	portalURL := `https://stagingaccess:adminuser@portal-staging.future-doctor.de/public/admin`

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		//schromedp.Flag("start-maximized", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	err := navigateToWebsite(ctx, portalURL)
	if err != nil {
		t.Fatalf("Navigation failed: %v", err)
	}

	emailSelector := `//*[@id="email"]`
	err = enterText(ctx, emailSelector, `daxeshpatel1085@gmail.com`)
	if err != nil {
		t.Fatalf("Failed to enter email: %v", err)
	}
	fmt.Println(`Email entered successfully`)

	passwordSelector := `//*[@id="password"]`
	err = enterText(ctx, passwordSelector, `Dakshesh@2085`)
	if err != nil {
		t.Fatalf("Failed to enter password: %v", err)
	}
	fmt.Println(`Password entered successfully`)

	err = chromedp.Run(ctx, chromedp.KeyEvent("\r"))
	if err != nil {
		t.Fatalf("Failed to press Enter key: %v", err)
	}
	fmt.Println("Enter key pressed")

	OPT := `//*[@id="otp"]`
	err = enterText(ctx, OPT, `111111`)
	if err != nil {
		t.Fatalf("Failed to enter otp: %v", err)
	}
	fmt.Println(`OTP entered successfully`)

	err = chromedp.Run(ctx, chromedp.KeyEvent("\r"))
	if err != nil {
		t.Fatalf("Failed to press Enter key: %v", err)
	}
	fmt.Println("Enter key pressed")

	dashboardSelector := `/html/body/div[3]/div[2]/div/div/div[1]/div`
	err = waitForElement(ctx, dashboardSelector)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	fmt.Println("Login successful!")

}

//// wrong otp entered and testcase failed

func TestLoginToApplicationWrongOTP(t *testing.T) {
	chromedp.Flag("headless", false)
	portalURL := `https://stagingaccess:adminuser@portal-staging.future-doctor.de/public/admin`

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		//chromedp.Flag("start-maximized", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	err := navigateToWebsite(ctx, portalURL)
	if err != nil {
		t.Fatalf("Navigation failed: %v", err)
	}

	emailSelector := `//*[@id="email"]`
	err = enterText(ctx, emailSelector, `daxeshpatel1085@gmail.com`)
	if err != nil {
		t.Fatalf("Failed to enter email: %v", err)
	}
	fmt.Println(`Email entered successfully`)

	passwordSelector := `//*[@id="password"]`
	err = enterText(ctx, passwordSelector, `Dakshesh@2085`)
	if err != nil {
		t.Fatalf("Failed to enter password: %v", err)
	}
	fmt.Println(`Password entered successfully`)

	err = chromedp.Run(ctx, chromedp.KeyEvent("\r"))
	if err != nil {
		t.Fatalf("Failed to press Enter key: %v", err)
	}
	fmt.Println("Enter key pressed")

	OPT := `//*[@id="otp"]`
	err = enterText(ctx, OPT, `22222`)
	if err != nil {
		t.Fatalf("Failed to enter otp: %v", err)
	}
	fmt.Println(`OTP entered successfully`)

	err = chromedp.Run(ctx, chromedp.KeyEvent("\r"))
	if err != nil {
		t.Fatalf("Failed to press Enter key: %v", err)
	}
	fmt.Println("Enter key pressed")

	errorMessage := `/html/body/div[1]/div/form/div[1]`
	err = waitForElement(ctx, errorMessage)
	if err != nil {
		t.Fatalf("error message not shows: %v", err)
	}
	fmt.Println("error message shows successfully!")

}

func navigateToWebsite(ctx context.Context, url string) error {
	var title string
	chromedp.Flag("headless", false)
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Title(&title),
	)
	if err != nil {
		return fmt.Errorf("failed to navigate: %v", err)
	}

	fmt.Printf("Page title: %s\n", title)
	time.Sleep(2 * time.Second)
	return nil
}

func enterText(ctx context.Context, selector, text string) error {
	err := chromedp.Run(ctx,
		chromedp.WaitVisible(selector, chromedp.BySearch),
		chromedp.SendKeys(selector, text),
	)
	if err != nil {
		return fmt.Errorf("failed to enter text: %v", err)
	}

	return nil
}

func clickButton(ctx context.Context, selector string) error {
	err := chromedp.Run(ctx,
		chromedp.WaitVisible(selector, chromedp.BySearch),
		chromedp.Click(selector, chromedp.BySearch),
	)
	if err != nil {
		return fmt.Errorf("failed to click button: %v", err)
	}

	return nil
}

func waitForElement(ctx context.Context, selector string) error {
	err := chromedp.Run(ctx,
		chromedp.WaitVisible(selector, chromedp.BySearch),
	)
	if err != nil {
		return fmt.Errorf("failed to find element: %v", err)
	}

	return nil
}
