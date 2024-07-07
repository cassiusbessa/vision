package com.github.cassiusbessa.vision.domain.service.dtos;

import com.github.cassiusbessa.vision.domain.core.entities.Profile;

public class LoadProfileResponse {

    private final Profile profile;
    private final String message;

    public LoadProfileResponse(Profile profile, String message) {
        this.profile = profile;
        this.message = message;
    }

    public Profile getProfile() {
        return profile;
    }

    public String getMessage() {
        return message;
    }
}
