import "cypress-localstorage-commands";

function urlJoin(parts) {
  const a = parts.filter((item) => {return item !== ''})
  return a.join('/').trim('/')
}



const css_table_header = '#app table thead tr th'
const css_table_body = '#app table tbody tr td'

const css_input_centre_name = '#app input[name="name"]'
const css_input_centre_description = '#app textarea[name="description"]'
const css_input_button_submit = '#app input.btn'

const username = 'admin.admin'
const password = 'adminadmin'

const url_base = 'http://malcorp.test'
const url_base_api = 'http://malcorp.test/api/server'

const resourse_name = 'centres'
const routes = {
  base: `${resourse_name}`,
  web: {
    list: 'list',
    create_one: 'create_one'
  },
  api: {
    list: '',
    create_one: '',
    getOne: '',
    updateOne: '',
    delete: '',
  }
}

const url_centres_list = urlJoin([url_base, routes.base, routes.web.list])
const url_centres_create_one = urlJoin([url_base, routes.base, routes.web.create_one])
const url_api_centres_create_one = urlJoin([url_base_api, routes.base, routes.api.create_one])
const url_api_centres_delete = urlJoin([url_base_api, routes.base, routes.api.delete])

const payload_create_one = {
  name: 'some new centre name',
  description: 'some description for the centre'
}


describe('Check Centres Page', () => {

  beforeEach(() => {
    cy.login2(username, password)
    cy.log("doing before every step")
  })

  it('Try to Login', () => {

    cy.visit(url_base)

    cy.url().should('include', 'dashboard')

  })

  it('Try to Get Centres List Page', () => {

      cy.visit(url_centres_list)

      cy.get(css_table_header).should('exist')
      cy.get(css_table_body).should('exist')

  })

  it('Try to Find Create One Button', () => {
    cy.visit(url_centres_list)

    cy.get('a').contains("Create One").should('have.attr', 'href').and('include', urlJoin([routes.base, routes.api.create_one]))
  })

  it('Create One New Centre, Check for Autohide of Toast and Delete Centre', () => {
    cy.intercept({
      method: 'POST',
      url: url_api_centres_create_one
    }).as('centreCreate')

    cy.visit(url_centres_create_one)

    cy.get(css_input_centre_name).type(payload_create_one.name)
    cy.get(css_input_centre_description).type(payload_create_one.description)

    cy.get('input.btn').click()
      .then(() => {
        cy.wait('@centreCreate').then(async ({_, response}) => {
          const created_id = response.body.data.id

          cy.get('div.toaster div.toast.bg-success').should('be.visible').should('have.text', 'Created')
          cy.wait(6000)
          cy.get('div.toaster div.toast.bg-success').should('not.exist')

          cy.log(`created_id2: ${created_id}`)

          cy.request({
            method: 'DELETE',
            url: url_api_centres_delete + '/' + created_id,
          })
            .its('body')
            .then(body => {
            expect(body).to.have.property('message', 'success')
          })

        })

      })

  })

  it('Create One New Centre, Check for Autohide of Toast, Close Toast and Delete Centre', () => {
    cy.intercept({
      method: 'POST',
      url: url_api_centres_create_one
    }).as('centreCreate')

    cy.visit(url_centres_create_one)

    cy.get(css_input_centre_name).type(payload_create_one.name)
    cy.get(css_input_centre_description).type(payload_create_one.description)

    cy.get('input.btn').click()
      .then(() => {
        let created_id = cy.wait('@centreCreate').then(async ({_, response}) => {
          const created_id = response.body.data.id

          cy.get('div.toaster div.toast.bg-success').should('be.visible').should('have.text', 'Created')
          cy.get('div.toaster button[type="button"]').should('be.visible').click()
          cy.get('div.toaster div.toast.bg-success').should('not.exist')

          cy.log(`created_id2: ${created_id}`)

          cy.request({
            method: 'DELETE',
            url: url_api_centres_delete + '/' + created_id,
          })
            .its('body')
            .then(body => {
            expect(body).to.have.property('message', 'success')
          })

        })

      })

  })

})

