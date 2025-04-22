describe("template spec", () => {
  beforeEach(() => {
    // Visit the login page before each test
    cy.visit("http://192.168.0.200:5173/measurements");
  });
  it("Should render the exercise list", () => {
    // Check if the login form or title text is visible
    cy.contains("Enter Measurements").should("be.visible");
  });
  it("Should render the measurements", () => {
    // Check if the login form or title text is visible
  });
  it("Should check if logout works", () => {});
  it("Should check if dropdown menu works", () => {});
  it("Should check if Measurements is added to database", () => {});
  it("Should check if Measurements works", () => {});
});
