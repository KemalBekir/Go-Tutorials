const puppeteer = require("puppeteer");
const Price = require("../models/price");
const Searches = require("../models/searches");
const mapErrors = require("../utils/mappers");
const userAgent = require("user-agents");
const targetWebsites = {
  "amazon.co.uk": scrapeAmazon,
  // TODO - Add more websites and their corresponding scraping functions
};

async function getAll() {
  return Searches.find({}).populate({
    path: "prices",
    select: ["price", "createdAt"],
  });
}

async function getLatest() {
  return Searches.find({}).sort({ createdAt: "desc" });
}

async function getAllByOwner(owner) {
  return Searches.find({ owner });
}

async function getById(id) {
  return Searches.findById(id).populate({
    path: "prices",
    select: ["price", "createdAt"],
  });
}

async function deleteById(id) {
  await Searches.findByIdAndDelete(id);
}

async function scrape(url, domain) {
  const browser = await puppeteer.launch({ headless: "new" });
  // const browser = await puppeteer.launch({
  //   headless: false,
  //   slowMo: 250, // Adjust the value as needed
  // });
  const page = await browser.newPage();
  await page.setViewport({ width: 1280, height: 800 });

  await page.setUserAgent(userAgent.random().toString());

  await page.goto(url);
  await page.waitForSelector("#sp-cc-accept");
  await page.click("#sp-cc-accept");

  const scrapingFunction = targetWebsites[domain];
  const result = await scrapingFunction(page, url, domain); // Pass the 'page' and 'url' to the scraping function
  await browser.close();

  return result;
}

async function scrapeAmazon(page, url, domain) {
  const productName = await page.evaluate(() => {
    const productTitleElement = document.querySelector("#title");
    if (productTitleElement) {
      return productTitleElement.textContent.trim();
    } else {
      return ""; // or some default value if the element is not found
    }
  });
  const wholePriceElement = await page.$(".a-price-whole");
  const fractionPriceElement = await page.$(".a-price-fraction");

  // Get the text content of the whole and fraction parts
  const wholePrice = await (
    await wholePriceElement.getProperty("textContent")
  ).jsonValue();

  const fractionPrice = await (
    await fractionPriceElement.getProperty("textContent")
  ).jsonValue();

  // Combine the whole and fraction parts into a single price string
  const priceString = wholePrice + fractionPrice;

  // Convert the price string into a floating-point number
  const formattedPrice = parseFloat(priceString.replace(/,/g, ""));

  const imgElement = await page.$(".a-dynamic-image");
  const img = await (await imgElement.getProperty("src")).jsonValue();

  const result = await createOrUpdateSearches(
    url,
    domain,
    productName,
    formattedPrice,
    img
  );

  return result;
}

async function createOrUpdateSearches(
  url,
  domain,
  productName,
  formattedPrice,
  img
) {
  let result = await Searches.findOne({ url });
  if (!result) {
    result = new Searches({
      url,
      domain,
      itemName: productName,
      img,
    });
  }

  const newPrice = new Price({
    price: formattedPrice,
  });
  const savedPrice = await newPrice.save();

  result.prices.push(savedPrice);
  await result.save();

  return result;
}

async function getDailyPrice() {
  try {
    const data = await Searches.find({})
      .populate({
        path: "prices",
        select: ["price", "createdAt"],
      })
      .lean();
    for (const item of data) {
      const browser = await puppeteer.launch({ headless: "false" });
      const page = await browser.newPage();
      await page.goto(item.url);
      await scrapeAmazon(page, item.url, item.domain);
      await browser.close();

      // Generate a random timeout between 5 and 10 seconds (in milliseconds)
      const timeout = Math.floor(Math.random() * (10000 - 5000 + 1)) + 5000;
      console.log(
        `Waiting for ${timeout / 1000} seconds before the next item...`
      );

      // Wait for the specified timeout before processing the next item
      await new Promise((resolve) => setTimeout(resolve, timeout));
    }
    // console.log('Finished');
  } catch (error) {
    console.error(error);
  }
}

module.exports = {
  scrape,
  getById,
  getAll,
  getAllByOwner,
  getLatest,
  deleteById,
  getDailyPrice,
};


//Package json
{
  "name": "server",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "start": "nodemon index.js"
  },
  "dependencies": {
    "bcrypt": "^5.1.0",
    "dotenv": "^16.3.1",
    "express": "^4.18.2",
    "jsonwebtoken": "^9.0.0",
    "mongoose": "^7.2.4",
    "node-cron": "^3.0.2",
    "nodemon": "^2.0.22",
    "puppeteer": "^21.5.2",
    "puppeteer-extra": "^3.3.6",
    "puppeteer-extra-plugin-stealth": "^2.11.2",
    "user-agent": "^1.0.4",
    "user-agents": "^1.1.28"
  }
}