*** Settings ***

Library    SeleniumLibrary    WITH NAME    SL

*** Variables ***

*** Test Cases ***
This is sample test case
    [Documentation]     Google Test
    [Tags]  test_google

    SL.Open Browser    https://www.google.com/    Chrome
    SL.Input Text    id:APjFqb    Tester
    SL.Click Button    xpath:    //div[@class="FPdoLc l"]//input[]
    
    SL.Close Browser

*** Keywords ***
