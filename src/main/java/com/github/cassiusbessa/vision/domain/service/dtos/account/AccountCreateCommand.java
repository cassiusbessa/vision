package com.github.cassiusbessa.vision.domain.service.dtos.account;

public class AccountCreateCommand {

    private final String email;
    private final String password;

    public AccountCreateCommand() {
        this.email = null;
        this.password = null;
    }

    public AccountCreateCommand(String email, String password) {
        this.email = email;
        this.password = password;
    }

    public String getEmail() {
        return email;
    }

    public String getPassword() {
        return password;
    }
}
