describe("Main page", () => {
    beforeEach(() => {
        cy.cleanDatabase();
    });

    it("creates todo item", () => {
        cy.visit("/");
        cy.get("[data-cy='form-input']").type("test 1");
        cy.get("[data-cy='form-submit']").click();
        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").contains(
            "test 2"
        );

        cy.waitForNetworkIdle(500);

        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").should(
            "have.length",
            1
        );

        cy.get("[data-cy='form-input']").type("test 2");
        cy.get("[data-cy='form-submit']").click();
        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").contains(
            "test 2"
        );

        cy.waitForNetworkIdle(500);

        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").should(
            "have.length",
            2
        );

        cy.get("[data-cy='form-input']").type("test 2");
        cy.get("[data-cy='form-submit']").click();
        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").contains(
            "test 2"
        );

        cy.waitForNetworkIdle(500);

        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").should(
            "have.length",
            2
        );
    });

    it("creates todo and changes it's status", () => {
        cy.visit("/");
        cy.get("[data-cy='form-input']").type("test");
        cy.get("[data-cy='form-submit']").click();

        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").each(($el) => {
            cy.wrap($el).find("button[data-cy='btn-done']").contains("Done");
        });

        cy.get(
            "[data-cy='todo-wrapper'] [data-cy='todo-item'] button[data-cy='btn-done']"
        ).click();

        cy.waitForNetworkIdle(500);

        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").each(($el) => {
            cy.wrap($el)
                .find("button[data-cy='btn-undone']")
                .contains("Undone");
        });

        cy.get(
            "[data-cy='todo-wrapper'] [data-cy='todo-item'] button[data-cy='btn-undone']"
        ).click();

        cy.waitForNetworkIdle(500);

        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").each(($el) => {
            cy.wrap($el).find("button[data-cy='btn-done']").contains("Done");
        });
    });

    it("creates todo and removes them", () => {
        cy.visit("/");
        cy.get("[data-cy='form-input']").type("test 1");
        cy.get("[data-cy='form-submit']").click();

        cy.waitForNetworkIdle(500);

        cy.get("[data-cy='form-input']").type("test 2");
        cy.get("[data-cy='form-submit']").click();

        cy.waitForNetworkIdle(500);

        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").should(
            "have.length",
            2
        );

        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").each(($el) => {
            cy.wrap($el).dblclick();
        });

        cy.waitForNetworkIdle(500);

        cy.get("[data-cy='todo-wrapper'] [data-cy='todo-item']").should(
            "have.length",
            0
        );
    });
});
