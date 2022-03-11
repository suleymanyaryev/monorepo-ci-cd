Cypress.Commands.add("cleanDatabase", () => {
    cy.exec("psql -f ./backend/db.sql", {
        env: {
            PGDATABASE: Cypress.env("PGDATABASE"),
            PGHOST: Cypress.env("PGHOST"),
            PGPORT: Cypress.env("PGPORT"),
            PGUSER: Cypress.env("PGUSER"),
            PGPASSWORD: Cypress.env("PGPASSWORD"),
        },
    });
});
