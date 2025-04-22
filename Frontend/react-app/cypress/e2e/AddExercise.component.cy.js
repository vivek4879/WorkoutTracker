describe("template spec", () => {
  beforeEach(() => {
    // Visit the login page before each test
    cy.visit("http://192.168.0.200:5173/AddExercise");
  });
  it("Should render the exercise list", () => {
    // Check if the login form or title text is visible
    cy.contains("Select Exercise").should("be.visible");
  });
  it("Should render the calendar", () => {
    // Check if the login form or title text is visible
    cy.contains("Ab Wheel").should("be.visible");
    cy.contains("Arnold Press").should("be.visible");
    cy.contains("Around The World").should("be.visible");
  });
  it("Should check if logout works", () => {});
  it("Should check if dropdown menu works", () => {});
  it("Should check if workout is added to database", () => {});
  it("Should check if Measurements works", () => {});
});
