package com.github.cassiusbessa.vision.domain.service.dtos;

public class AccountCreateCommand {

    private final String email;
    private final String password;

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
