const { defineConfig } = require("cypress");

module.exports = defineConfig({
  e2e: {
    baseUrl: "http://192.168.0.200:5173/",
    supportFile: false,
  },

  component: {
    devServer: {
      framework: "react",
      bundler: "webpack",
    },
  },
});
