describe("template spec", () => {
  beforeEach(() => {
    // Visit the login page before each test
    cy.visit("http://192.168.0.200:5173/forgot-pass");
  });
  it("Should render the exercise list", () => {
    // Check if the login form or title text is visible
    cy.contains("Forgot Password?").should("be.visible");
  });
});
