// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --

// Cypress.Commands.add('login', (email, password) => { ... })
Cypress.Commands.add('login', (username, password) => {
  cy.visit('http://malcorp.test/pages/login')
  cy.get('input[name=username]').type(username)

  // {enter} causes the form to submit
  cy.get('input[name=password]').type(`${password}{enter}`)
})

Cypress.Commands.add('login2', (username, password) => {
  cy.session([username, password], () => {

    cy.request({
      method: 'POST',
      url: 'http://malcorp.test/api/server/login',
      body: {
        username: username,
        password: password,
      }
    })
    .its('body')
    .then(body => {
      cy.setLocalStorage("token", body.data.token);
    })

  })
})

// Cypress.Commands.add('clearLocalStorageCache', { prevSubject: 'element'}, (subject, options) => { ... })


//
//
// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })