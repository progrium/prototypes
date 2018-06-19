'use strict';

const puppeteer = require('puppeteer');
const timeout = ms => new Promise(res => setTimeout(res, ms));

(async() => {
    await timeout(1000);
  const browser = await puppeteer.launch({
      headless: true,
      userDataDir: "data"
    });
  const pages = await browser.pages();

  await pages[0].goto('https://www.twitch.tv/progrium/manager');

  console.log("waiting for videos...")
  await pages[0].waitForSelector('.video-manager-processed-list .video-card h5');

  console.log("clicking export...")
  const exportSelector = '.video-manager-processed-list .video-card button[data-test-selector="export"]';
  await pages[0].waitForSelector(exportSelector);
  await pages[0].evaluate(selector => {
    document.querySelector(selector).click();
  }, exportSelector);

  console.log("starting export...")  
  const startSelector = '.export-youtube-modal button[data-test-selector="save"]'
  await pages[0].waitForSelector(startSelector);
  await pages[0].evaluate(selector => {
    document.querySelector(selector).click();
  }, startSelector);

  console.log("waiting...")  
  await timeout(5000);

  await browser.close();
})();