FROM chromedp/headless-shell:latest

ENTRYPOINT ["/headless-shell/headless-shell", "--headless", "--disable-gpu", "–-no-first-run", "--disable-dev-shm-usage", "--disable-setuid-sandbox", "--no-sandbox", "--font-render-hinting=none", "--force-color-profile=srgb", "--remote-debugging-address=0.0.0.0", "--remote-debugging-port=9222"]