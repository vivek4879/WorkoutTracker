describe("template spec", () => {
  beforeEach(() => {
    // Visit the login page before each test
    cy.visit("http://192.168.0.200:5174/");
  });
  it("Should show the login form", () => {
    // Check if the login form or title text is visible
    cy.contains("Login").should("be.visible");
  });
  it("Should allow a user to log in", () => {
    // Fill in the username
    cy.get('[data-testid="email"]').type("testuser@gmail.com");

    // Fill in the password
    cy.get('[data-testid="password"]').type("mypassword");

    // Submit the form
    cy.get('[data-testid="loginButton"]').click();

    // Check if we are redirected to a dashboard or see a success message
    cy.url().should("include", "/dashboard");
    // or
    // cy.contains('Welcome, testuser').should('be.visible');
  });
});
