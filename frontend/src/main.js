import "./style.css";

import { LoadAndParseOpenAPI, SelectOpenFile } from "../wailsjs/go/main/App";
import { EventsOn, BrowserOpenURL } from "../wailsjs/runtime/runtime";

const redocContainer = document.getElementById("redoc-container");
const uploadBtn = document.getElementById("upload-btn");
const githubBtn = document.getElementById("github-btn");

const render = async (filePath) => {
  if (filePath !== "") {
    try {
      const specJsonString = await LoadAndParseOpenAPI(filePath);
      redocContainer.textContent = "Rendering documentation...";
      const specObject = JSON.parse(specJsonString);
      if (typeof Redoc !== "undefined") {
        console.log("Available methods in Redoc:", Object.keys(Redoc));
        Redoc.init(
          specObject, // Pass the parsed spec object
          {
            // Redoc options (optional)
            scrollYOffset: 20, // Example: Adjust scroll offset if you have a fixed header
            // theme: { colors: { primary: { main: '#dd5522' } } }
          },
          redocContainer
        );
      }
    } catch (error) {
      console.error("Error:", error);
      // Display the error message from Go (Wails wraps it)
      // Wails error might be an object { message: "..." } or just a string
      const errorMessage =
        typeof error === "object" && error !== null && error.message
          ? error.message
          : String(error);
      redocContainer.textContent = `Error: ${errorMessage}`;
      // Optionally display a more user-friendly error in redocContainer
    }
  }
};

githubBtn.addEventListener("click", function () {
  BrowserOpenURL("https://github.com/carmel/sepcview")
    .then(() => {
      console.log(`Successfully requested to open URL: ${url}`);
    })
    .catch((err) => {
      console.error(`Failed to open URL: ${url}`, err);
      // Optionally display an error message to the user
      alert(`Could not open the link: ${err}`);
    });
});

uploadBtn.addEventListener("click", async function () {
  try {
    const filePath = await SelectOpenFile(); // Assumes Go package 'main', struct 'App'
    render(filePath);
  } catch (error) {
    console.error("Error:", error);
    redocContainer.innerHTML = `<pre style="color: red; padding: 20px;">Error loading spec:\n${errorMessage}</pre>`;
  }
});

document.addEventListener("DOMContentLoaded", function () {
  // 监听后端事件
  EventsOn("fileSelected", async function (filePath) {
    render(filePath);
  });
});
