# Wallet Application - Wallet'naka Object Specification

### wallet details:
* #### attributes:
  * cards
  * money
  * state: [OPEN, CLOSED]
  * capacity
* #### rules:
  * you can have as many wallets as possible
  * you cannot do anything a closed wallet

A wallet is some sort of pocket where you can add / take the following items : 
* NIC
* Money
* Bank card
* Drivers license
* Visit card
* ID photo

### For each object, let us analyse the possibilities:
### GLOBAL:
* attributes:
  * size: will decrease the capacity of the wallet on input
* actions:
  * add
  * take
* ### NIC: 
  * #### attributes:
    * first name
    * lastname
    * identifier
    * birthdate
    * creation date
    * work
    * address
    * state
  * #### rules:
    * you can many NICs but the state has to be different
    * the only difference between your NICS are identifier, nationality, state, creation date.
* ### Money:
  * #### attributes:
    * amount
    * currency
  * #### rules:
    * you can take as much money as the available amount
    * you can put as much money as you want (positive value)
* ### Bank card:
  * #### attributes:
    *  expiring date
    * owner name
    * identifier
  * ### rules:
    * you can have as many cards as your number of bank accounts
* ### Drivers license:
  * #### attributes:
      * first name
      * lastname
      * identifier
      * birthdate
      * issue date
      * type
      * NIC id
  * ### rules:
    * you can have as many drivers licenses but the country must be different