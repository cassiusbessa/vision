package com.github.cassiusbessa.vision.domain.service.dtos.auth;

public class AuthCredentials {

    private final String email;
    private final String password;

    public AuthCredentials() {
        this.email = null;
        this.password = null;
    }

    public AuthCredentials(String email, String password) {
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
