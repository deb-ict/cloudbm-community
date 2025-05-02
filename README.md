![Build](https://github.com/deb-ict/cloudbm-community/actions/workflows/build.yml/badge.svg)
[![codecov](https://codecov.io/github/deb-ict/cloudbm-community/branch/main/graph/badge.svg)](https://codecov.io/github/deb-ict/cloudbm-community)
[![codecov](https://codecov.io/github/deb-ict/cloudbm-community/graph/badge.svg?token=ETYWKCYMIB)](https://codecov.io/github/deb-ict/cloudbm-community)
[![godoc](https://godoc.org/github.com/deb-ict/cloudbm-community?status.svg)](https://godoc.org/github.com/deb-ict/cloudbm-community)

# cloudbm-community
Cloud Business Management - Community Edition

# pre
Powershell
`Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned`

## Build the web application
`npm install`
`ng build`

# Add angular module
`ng generate module module/<moduleName>`
`ng generate component module/<moduleName>/view/Add<Model>`
`ng generate component module/<moduleName>/view/Add<Model>`

# TODO

## Contacts
- Company Size
- Company Currency
- Company Language
- Company Invoice Payment Terms

## Products
- Default quantity
- Max quantity
- Tax class
- Dimensions
- Dimenssions class (units)
- Weight
- Weight class (units)
- Status (disable, draft, ...)
- Attributes
- Options
- Discounts
- Tier prices


## Orders
- RecipientType (contact/company)
- RecipientId
- OrderNumber
- OrderDate
- DueDate
- SubTotal
- Vat
- Total
-- ArticleType (product/project)
-- ArticleId
-- Description
-- Quantity
-- UnitPrice


## Orders (opencart based)
- Customer details
- Payments details
- Shipping details
- Order items
- Totals

## Returns
- Order ID
- Order date
- Customer
- Products
    - Id
    - Quantity
    - Reason
    - State
    - Info

/cart/add
{
    product_id: 1
    quantity: 1
}
/cart/update
{
    product_id: 1
    quantity: 2
}
/cart/delete
{
    product_id: 1
}


https://admin-demo.nopcommerce.com/admin/

shop config
- Countries
    - Name
    - Allow billing
    - Allow shipping
    - ISO2 code (BE, NL, ...)
    - ISO3 code (BEL, NED, ...)
    - ISONr (840, 4, ...)
    - SubjectToVat (true/false)
- currencies
    - name
    - code
    - rate