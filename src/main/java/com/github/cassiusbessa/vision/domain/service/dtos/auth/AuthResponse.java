package com.github.cassiusbessa.vision.domain.service.dtos.auth;

public class AuthResponse {

    private final String token;
    private final String message;


    public AuthResponse(String token, String message) {
        this.token = token;
        this.message = message;
    }

    public String getToken() {
        return token;
    }

    public String getMessage() {
        return message;
    }
}
