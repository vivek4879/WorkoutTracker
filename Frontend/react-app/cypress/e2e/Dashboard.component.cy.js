describe("template spec", () => {
  beforeEach(() => {
    // Visit the login page before each test
    cy.visit("http://192.168.0.200:5173/dashboard");
  });
  it("Should render the max streak", () => {
    // Check if the login form or title text is visible
    cy.contains("Current Streak").should("be.visible");
  });
  it("Should render the calendar", () => {
    // Check if the login form or title text is visible
    cy.contains("Mon").should("be.visible");
    cy.contains("Tue").should("be.visible");
    cy.contains("Wed").should("be.visible");
    cy.contains("Thurs").should("be.visible");
    cy.contains("Fri").should("be.visible");
    cy.contains("Sat").should("be.visible");
    cy.contains("Sun").should("be.visible");
  });
  it("Should check if logout works", () => {});
  it("Should check if addworkout works", () => {});
  it("Should check if workout cards are rendered", () => {});
  it("Should check if Measurements works", () => {});
});
