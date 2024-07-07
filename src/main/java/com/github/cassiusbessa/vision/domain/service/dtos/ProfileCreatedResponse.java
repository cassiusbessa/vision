package com.github.cassiusbessa.vision.domain.service.dtos;

public class ProfileCreatedResponse {

    private final String message;

    public ProfileCreatedResponse(String message) {
        this.message = message;
    }

    public String getMessage() {
        return message;
    }
}
