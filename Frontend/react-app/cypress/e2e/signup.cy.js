describe("template spec", () => {
  beforeEach(() => {
    // Visit the login page before each test
    cy.visit("http://192.168.0.200:5173/signup");
  });
  it("Should render the exercise list", () => {
    // Check if the login form or title text is visible
    cy.contains("Sign Up").should("be.visible");
  });
  it("Should render the First Name", () => {
    // Check if the login form or title text is visible
  });
  it("Should check if last Name works", () => {});
  it("Should check if email works", () => {});
  it("Should check if password works", () => {});
  it("Should check validation", () => {});
});
