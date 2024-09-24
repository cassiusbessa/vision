package com.github.cassiusbessa.vision.domain.service.dtos.profile;

public class LoadProfileResponse {

    private final ProfileDTO profile;
    private final String message;

    public LoadProfileResponse(ProfileDTO profile, String message) {
        this.profile = profile;
        this.message = message;
    }

    public ProfileDTO getProfile() {
        return profile;
    }

    public String getMessage() {
        return message;
    }
}
